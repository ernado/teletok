FROM gcr.io/distroless/static

ADD teletok /usr/local/bin/teletok

ENTRYPOINT ["teletok"]
