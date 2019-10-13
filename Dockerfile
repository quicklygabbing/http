FROM golang:1.12-alpine

WORKDIR $GOPATH/src/github.com/quicklygabbing/http

COPY . .

RUN apk update && apk add bash make git gcc libc-dev
RUN make http
