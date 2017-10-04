FROM golang:1.9-alpine

EXPOSE 8080

# Copy into a valid GOPATH
COPY . /go/src/github.com/golang/example/go_cache_server

RUN go install github.com/golang/example/go_cache_server

# Run using the go binary
ENTRYPOINT /go/bin/go_cache_server
