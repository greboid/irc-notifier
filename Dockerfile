FROM golang:1.24 as builder

WORKDIR /app
COPY . /app
RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/notifier; \
     go run github.com/google/go-licenses@latest save ./... --save_path=/notices;

FROM ghcr.io/greboid/dockerbase/nonroot:1.20250803.0

COPY --from=builder /app/main /irc-notifier
COPY --from=builder /notices /notices
CMD ["/irc-notifier"]
