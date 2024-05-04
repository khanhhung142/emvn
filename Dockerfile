FROM golang:1.21-alpine as builder

ENV GO111MODULE=on \
  CGO_ENABLED=1

RUN apk update && apk add gcc make git libc-dev binutils-gold 

WORKDIR /build
COPY . .

RUN go build -tags musl -ldflags '-extldflags "-static"' -o ./main main.go

# Create a non-root user
RUN adduser -D -g '' app

USER app

FROM alpine:latest

WORKDIR /home
COPY --from=builder /build/main .
COPY --from=builder /build/entrypoint.sh .
COPY --from=builder /build/config/yaml/config.yaml.template ./config/yaml/config.yaml.template

RUN apk update && apk add --no-cache bash jq gettext util-linux curl && \
  chmod 500 /home/entrypoint.sh && \
  rm -rf /var/lib/{apt,dpkg,cache,log}/ 

ENTRYPOINT [ "./entrypoint.sh" ]

CMD ["./main"]