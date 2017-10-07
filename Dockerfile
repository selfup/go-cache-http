
FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/go-cache-http /opt/bin/go-cache-http

CMD ["/opt/bin/go-cache-http"]
