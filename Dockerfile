FROM golang:1.12-alpine
RUN apk add git

ENV GO111MODULE=on