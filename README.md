### Golang [Clean Architecture]() gRPC Auth microservice example with Prometheus, Grafana monitoring and Jaeger opentracing ⚡️

#### 👨‍💻 Full list what has been used:
* [Docker](https://www.docker.com/) - Docker
* [Prometheus](https://prometheus.io/) - Prometheus
* [Grafana](https://grafana.com/) - Grafana
* [Jaeger](https://www.jaegertracing.io/) - Jaeger tracing

#### Recommendation for local development most comfortable usage:
    make local // run all containers
    make run // run the application

#### 🙌👨‍💻🚀 Docker-compose files:
    docker-compose.local.yml - run postgresql, redis, aws, prometheus, grafana containers
    docker-compose.dev.yml - run all in docker

### Docker development usage:
    make docker

### Local development usage:
    make local
    make run

### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090

### Grafana UI:

http://localhost:3000