FROM golang:1.15.3-alpine3.12

ENV GO111MODULE=on

RUN apk update && \
    apk add emacs curl git && \
    go get github.com/pilu/fresh