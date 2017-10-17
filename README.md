# GoCrasHttp

Small Golang expirement :smile:

Currently up as an [alpine container on heroku](https://gocrashttp.herokuapp.com/)

## Docker

Only use Docker to build the image if using containerization technologies like:

1. Docker Swarm
2. Kubernetes (k8s)
3. Google Container Platform
4. Heroku (ex: `heroku container:push web --app <app_name>`)
5. AWS/Azure Container Solutions
6. A VPS with only Docker as a dependency (Digital Ocean/Linode box)

The image for production is using alpine linux.

Very small so no need to worry about bandwith.

Build times are very fast even from scratch!

## Development

Use a local Go install to develop this software.

### Get **Gin** for dev:

**Linux**:
```
go get github.com/codegangsta/gin && source ~/.bashrc
```

**macOS**:
```
go get github.com/codegangsta/gin && source ~/.bash_profile
```

### Use **Gin** for dev:

```
./scripts/dev.sh
```

OR

```
gin run main.go
```

_Refer to the `./scripts` dir for other tasks that can be run_
