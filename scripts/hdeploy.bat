setx CGO_ENABLED 0
setx GOOS linux
setx GOARCH amd64

rmdir .docker_build
mkdir .docker_build

go build -v -o .docker_build\gocrashttp .
docker-compose build --no-cache
heroku container:push web --app gocrashttp
