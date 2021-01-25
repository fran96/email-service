# Build
FROM golang:1.15.3-alpine as builder

RUN apk --no-cache --update add gcc libc-dev git

WORKDIR /build

COPY src .

RUN go build -a -ldflags="-w -s" -o output cmd/main.go

CMD ["./output"]
