FROM golang:1.9-alpine

ENV PORT 8080
EXPOSE 8080

COPY . /go/src/github.com/selfup/go-cache-http

RUN go install github.com/selfup/go-cache-http

ENTRYPOINT /go/bin/go-cache-http
