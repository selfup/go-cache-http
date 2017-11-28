FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/gocrashttp /opt/bin/gocrashttp

CMD ["/opt/bin/gocrashttp"]
