FROM golang:1.9-alpine

EXPOSE 8080

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/golang/example/go_cache_server

RUN go install github.com/golang/example/go_cache_server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/go_cache_server
