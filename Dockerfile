FROM golang:1.11-alpine

ENV CGO_ENABLED 1

RUN apk update \
  && apk add --no-cache sqlite \
  && apk add --no-cache sqlite-dev \
  && apk add --no-cache build-base \
  && apk add --no-cache git

COPY ./bot /go/src/pollrbot
WORKDIR /go/src/pollrbot

RUN go get ./
RUN go build

RUN mkdir -p /db && chmod 777 /db

WORKDIR /go/src/pollrbot
CMD DB=/db/database.db APITOKEN=$TOKEN pollrbot

EXPOSE 8443
