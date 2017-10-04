# Go Cache Store

Small Golang expirement :smile:

### Docker

Only use Docker to build the image if using containerization technologies like:

1. Docker Swarm
2. Kubernetes (k8s)

The image is an alpine linux box, very small so no need to worry about bandwith.

Build times are very fast even from scratch!

### Development

If you are deploying to containerized platforms, build times are fast enough that you should be able to:

`docker-compose up --build` (it just copies and re-compiles)

Working on figuring out an official watch script for dev.

Will need to split up a `Dockerfile.dev` as well as a `docker-compose.dev.yml` for development purposes.
