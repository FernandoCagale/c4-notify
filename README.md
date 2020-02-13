**c4-notify - docker**

`Docker Mongodb`

```sh
$ docker run --network host --name mongo -d mongo
```

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

`Docker build c4-notify`

```sh
$   docker build -t c4-notify .
```

`Docker c4-notify`

```sh
$   docker run -d --name c4-notify -p 8080:8080 c4-notify
```

**c4-notify - local**


```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download wire "dependency injection"`

```sh
$   go get -u github.com/google/wire/cmd/wire
```

`generate wire_gen.go`

```sh
$   wire
```

`generate build`

```sh
$   go build -o bin/application
```


```sh
$   ./bin/application
```