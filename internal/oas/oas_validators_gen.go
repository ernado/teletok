// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s Data) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.VideoAuthorDiggCount.Set {
			if err := func() error {
				if err := s.VideoAuthorDiggCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_author_diggCount",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoAuthorFollowerCount.Set {
			if err := func() error {
				if err := s.VideoAuthorFollowerCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_author_followerCount",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoAuthorFollowingCount.Set {
			if err := func() error {
				if err := s.VideoAuthorFollowingCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_author_followingCount",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoAuthorHeartCount.Set {
			if err := func() error {
				if err := s.VideoAuthorHeartCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_author_heartCount",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoAuthorVideoCount.Set {
			if err := func() error {
				if err := s.VideoAuthorVideoCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_author_videoCount",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoCommentCount.Set {
			if err := func() error {
				if err := s.VideoCommentCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_comment_count",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoDiggCount.Set {
			if err := func() error {
				if err := s.VideoDiggCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_digg_count",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoDownloadCount.Set {
			if err := func() error {
				if err := s.VideoDownloadCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_download_count",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoPlayCount.Set {
			if err := func() error {
				if err := s.VideoPlayCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_play_count",
			Error: err,
		})
	}
	if err := func() error {
		if s.VideoShareCount.Set {
			if err := func() error {
				if err := s.VideoShareCount.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_share_count",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s MaybeNumber) Validate() error {
	switch s.Type {
	case Int64MaybeNumber:
		return nil // no validation needed
	case NoneMaybeNumber:
		if err := s.None.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s None) Validate() error {
	switch s {
	case "None":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
