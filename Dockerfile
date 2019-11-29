FROM golang:1.13-alpine

ADD . /go/src/github.com/sad-unicorn/gray-goose-bar

RUN go install github.com/sad-unicorn/gray-goose-bar

ENTRYPOINT /go/bin/gray-goose-bar
