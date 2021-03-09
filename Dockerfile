FROM golang:1-alpine as builder

WORKDIR /app

RUN set -x \
    && apk add --no-cache \
        git \
    && go get github.com/markbates/pkger/cmd/pkger

COPY go.mod .
RUN go mod download

ARG GOOS=linux
ARG GOARCH=amd64
COPY . .
RUN set -x \
    && go generate \
    && go build -ldflags="-w -s"

FROM composer:2

WORKDIR /app

RUN apk add --no-cache npm

ENV PATH="/app:$PATH"

COPY --from=builder /app/scaffold .
CMD ["scaffold"]
