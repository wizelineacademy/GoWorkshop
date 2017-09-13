![diagram](https://github.com/wizelineacademy/GoWorkshop/raw/master/diagram.png)

### "ToDo List" experiment project

This repository is an example project which demonstrates the use of microservices for a fictional ToDo list application. The ToDo backend is powered by 3 microservices, all of which happen to be written in Go, using MongoDB for manage the database and Docker to isolate and deploy the ecosystem.

### Services organization

The application consists of the following application services:

| Service  | Port  | Description                   | Methods                              |
|----------|-------|-------------------------------|--------------------------------------|
| users    | 50000 | Provides users information    | CreateUser                           |
| list     | 50001 | Manages items in todo lists   | CreateItem, GetUserItems, DeleteItem |
| notifier | 50002 | Send email notifications      | Will be implemented during Workshop  |

Client web application is working on [http://127.0.0.1:3000](http://127.0.0.1:3000).

### Workshop task

 - Create `notifier` service.
 - Call `notifier` from `users` service to send simple notification to user's email address after account is created by `CreateUser` procedure.

#### Tips

New service needs separate folder, can be copied from `users` or `list` and modified later. The best place to start is to define a service using [proto3 language](https://developers.google.com/protocol-buffers/docs/proto3), in `proto/notifier/service.proto`. Then implement handlers in `notifier/controllers/handlers.go`.

![diagram2](https://github.com/wizelineacademy/GoWorkshop/raw/master/diagram2.png)

Service should be also added to `docker-compose.yml` and work on port 50002 (`ports` option in this file). Use `docker-compose build notifier` to build service image.

### Requirements

 - Install [Docker](https://www.docker.com/get-docker)
 - Install [Docker Compose](https://docs.docker.com/compose/install)

### Run

```
docker-compose up -d
```

Go to [http://127.0.0.1:3000](http://127.0.0.1:3000) to test gRPC from webapp.

### Client

Client is built with Go and located in `web`, working on [http://127.0.0.1:3000](http://127.0.0.1:3000).

### Generate source code for the gRPC client from .proto files

Service Docker containers generate these files for you, so it isn't necessary to run it locally, however here is instruction for Golang:

 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)
 - Install go plugin: `go get github.com/golang/protobuf/protoc-gen-go`

```
protoc --go_out=plugins=grpc:. proto/users/service.proto
protoc --go_out=plugins=grpc:. proto/list/service.proto
```
