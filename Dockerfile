# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/ /app/cmd/
COPY pkg/ /app/pkg/

RUN go build -o /rss-feed-demo cmd/web/*

## Deploy
FROM alpine

WORKDIR /

COPY --from=build /rss-feed-demo /rss-feed-demo
COPY ui/ ui/

EXPOSE 8080

CMD ["/rss-feed-demo"]