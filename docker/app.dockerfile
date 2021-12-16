FROM golang:alpine

WORKDIR /golang-live-coding-challenge

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./golang-live-coding-challenge

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./golang-live-coding-challenge"