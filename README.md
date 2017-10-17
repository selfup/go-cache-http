# GoCrasHttp

Small Golang expirement :smile:

### Docker

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

### Development

Use a local Go install to develop this software.
