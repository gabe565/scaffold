ARG GO_VERSION=1.18
ARG PHP_VERSION=8.1

FROM golang:$GO_VERSION-alpine as builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

ARG GOOS=linux
ARG GOARCH=amd64
COPY . .
RUN set -x \
    && go generate \
    && go build -ldflags="-w -s"


FROM php:$PHP_VERSION-alpine
WORKDIR /data
RUN apk add --no-cache npm
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer
COPY --from=builder /app/scaffold /usr/local/bin
CMD ["scaffold"]
