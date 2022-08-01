package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/go-faster/errors"
	"github.com/gotd/contrib/oteltg"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"teletok/internal/app"
	"teletok/internal/oas"
)

func main() {
	app.Run(func(ctx context.Context, logger *zap.Logger) error {
		metrics, err := app.NewMetrics(logger, app.Config{
			Name: "teletok",
		})
		if err != nil {
			return errors.Wrap(err, "failed to create metrics")
		}

		mw, err := oteltg.New(metrics.MeterProvider(), metrics.TracerProvider())
		if err != nil {
			return errors.Wrap(err, "failed to create metrics middleware")
		}

		dispatcher := tg.NewUpdateDispatcher()
		opt := telegram.Options{
			UpdateHandler: dispatcher,
			Logger:        logger.Named("gotd"),
			Middlewares:   []telegram.Middleware{mw},
		}

		tiktok, err := oas.NewClient(os.Getenv("API_URL"),
			oas.WithMeterProvider(metrics.MeterProvider()),
			oas.WithTracerProvider(metrics.TracerProvider()),
		)
		if err != nil {
			return errors.Wrap(err, "failed to create tiktok client")
		}

		tracer := metrics.TracerProvider().Tracer("")

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			if err := telegram.BotFromEnvironment(ctx, opt, func(ctx context.Context, client *telegram.Client) error {
				var (
					api    = tg.NewClient(client)
					sender = message.NewSender(api)
				)
				dispatcher.OnNewMessage(func(ctx context.Context, e tg.Entities, u *tg.UpdateNewMessage) (err error) {
					m, ok := u.Message.(*tg.Message)
					if !ok || m.Out {
						return nil
					}

					ctx, span := tracer.Start(ctx, "OnNewMessage")
					defer func() {
						if err != nil {
							span.RecordError(err)
							span.SetStatus(codes.Error, err.Error())
						} else {
							span.SetStatus(codes.Ok, "OK")
						}
						span.End()
					}()

					var (
						reply  = sender.Reply(e, u)
						lg     = logger.With(zap.Int("msg_id", m.ID))
						answer = sender.Answer(e, u)
						action = answer.TypingAction()
					)

					peerUser, ok := m.PeerID.(*tg.PeerUser)
					if !ok {
						if _, err := reply.Text(ctx, "Invalid"); err != nil {
							return err
						}
						return nil
					}

					user := e.Users[peerUser.UserID]
					if user == nil {
						return nil
					}

					span.SetAttributes(
						attribute.Int64("telegram.user.id", peerUser.UserID),
						attribute.String("telegram.user.username", user.Username),
					)

					lg = lg.With(zap.String("user", user.Username))

					if m.Message == "/start" {
						if _, err := reply.Text(ctx, "Hi, I'm teletok bot. I'm here to help you to download tiktok videos. Send me some links."); err != nil {
							return errors.Wrap(err, "failed to send start message")
						}
						return nil
					}

					uri, err := url.Parse(m.Message)
					if err != nil {
						if _, err := reply.Text(ctx, "Invalid URL"); err != nil {
							return errors.Wrap(err, "failed to send invalid url message")
						}
						return nil
					}

					if err := action.RecordVideo(ctx); err != nil {
						return err
					}

					stage := atomic.NewString("fetching data")
					done := make(chan struct{})

					g, gCtx := errgroup.WithContext(ctx)

					g.Go(func() error {
						defer close(done)
						data, err := tiktok.GetData(gCtx, oas.GetDataParams{
							URL: uri.String(),
						})
						if err != nil {
							return err
						}
						if data.Status != "success" {
							return fmt.Errorf("status: %s", data.Status)
						}

						stage.Store("uploading")

						upload, err := uploader.NewUploader(api).WithThreads(2).FromURL(ctx, data.NwmVideoURL.Value)
						if err != nil {
							return fmt.Errorf("upload: %w", err)
						}
						lg.Info("Uploaded")

						videoName := data.VideoTitle.Value
						if videoName == "" {
							videoName = "video"
						}

						var caption []message.StyledTextOption
						caption = append(caption,
							styling.Bold(data.VideoAuthorNickname.Value),
							styling.TextURL(videoName, data.OriginalURL.Value),
						)
						if _, err := reply.Media(gCtx,
							message.UploadedDocument(upload, caption...).
								Filename(fmt.Sprintf("%s.mp4", data.VideoAwemeID.Or("video"))).
								Video().
								SupportsStreaming(),
						); err != nil {
							return err
						}

						return nil
					})
					g.Go(func() error {
						ticker := time.NewTicker(time.Second * 3)
						defer ticker.Stop()
						defer lg.Info("Progress done")

						reportProgress := func() error {
							switch stage.Load() {
							case "uploading":
								return action.UploadVideo(gCtx, 50)
							default:
								return action.RecordVideo(gCtx)
							}
						}

						if err := reportProgress(); err != nil {
							return err
						}

						tick := func() error {
							if err := reportProgress(); err != nil {
								return fmt.Errorf("report: %w", err)
							}

							return nil
						}

						if err := tick(); err != nil {
							return err
						}

						for {
							select {
							case <-done:
								_ = action.Cancel(gCtx)

								return nil
							case <-ticker.C:
								if err := tick(); err != nil {
									return err
								}
							}
						}
					})

					if err := g.Wait(); err != nil {
						lg.Error("Failed", zap.Error(err))
						_, err := reply.Text(ctx, fmt.Sprintf("Upload failed: %v", err))
						return err
					}

					return nil
				})

				return nil
			}, telegram.RunUntilCanceled); err != nil {
				return errors.Wrap(err, "run")
			}

			return nil
		})
		g.Go(func() error {
			return metrics.Run(ctx)
		})

		return g.Wait()
	})
}
