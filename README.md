# c4-notify [Flow](https://github.com/FernandoCagale/c4-kustomize)

[![Build Status](https://travis-ci.org/FernandoCagale/c4-notify.svg?branch=master)](https://travis-ci.org/FernandoCagale/c4-notify)

### Docker

`running docker multi-stage builds and publish c4-notify`

```sh
$   ./scripts/publish.sh
```

### Kubernetes and Istio - [YAML](https://github.com/FernandoCagale/c4-kustomize/tree/master/c4-notify/base)

    *   deployment.yaml
    *   service.yaml

# Running local

### Dependencies [docker-compose](https://github.com/FernandoCagale/c4-kustomize/blob/master/docker-compose.yml)

```sh
$   docker-compose up -d
```

### Standard Go Project [Layout](https://github.com/golang-standards/project-layout)

```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download "dependency injection"` [wire](https://github.com/google/wire)

```sh
$   go get -u github.com/google/wire/cmd/wire
```

```sh
$   ./scripts/start.sh
```