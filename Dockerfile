FROM golang:1.9-alpine

ENV PORT 8080
EXPOSE 8080

COPY . /go/src/github.com/selfup/go-cache-server

RUN go install github.com/selfup/go-cache-server

ENTRYPOINT /go/bin/go-cache-server
