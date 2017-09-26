![diagram](https://github.com/wizelineacademy/GoWorkshop/raw/master/diagram.png)

### "TODO List" experiment project

This repository is an example project which demonstrates the use of microservices for a fictional TODO list application. The TODO backend is powered by 3 microservices, all of which happen to be written in Go, using MongoDB for manage the database and Docker to isolate and deploy the ecosystem.

In real world each service should live in a separate repository, so teams can work separately and don't overlap each other, however in this demo project they just located in separate folders for easy use.

### Services organization

The application consists of the following application services:

| Service  | Port  | Description                   | Methods                              |
|----------|-------|-------------------------------|--------------------------------------|
| users    | 50000 | Provides users information    | CreateUser                           |
| list     | 50001 | Manages items in todo lists   | CreateItem, GetUserItems, DeleteItem |
| notifier | 50002 | Send email notifications      | NewUser                              |

Client web application is working on [http://127.0.0.1:3030](http://127.0.0.1:3030).

![diagram2](https://github.com/wizelineacademy/GoWorkshop/raw/master/diagram2.png)

### Workshop task

 - Create `notifier` service, can be copied from `users` or `list` and modified.
 - Add `notifier` do `docker-compose.yml`, also add it to `users`'s options: `depends_on`, `links`.
 - Implement server in `notifier/server/server.go`. Use `shared` package, it's already has a function `SendEmail` to send email by SMTP.
 - Call `notifier` from `users` service to send simple notification to user's email address after account is created by `CreateUser` procedure.
 - Run `docker-compose build notifier users` to build new images.

### Requirements

 - Install [Docker](https://www.docker.com/get-docker)
 - Install [Docker Compose](https://docs.docker.com/compose/install)

### Run

```
docker-compose pull
docker-compose up
```

Go to [http://127.0.0.1:3030](http://127.0.0.1:3030) to test gRPC from webapp.

### Generate source code for the gRPC client from .proto files

 - Install [Go](https://golang.org/dl/)
 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)
 - Install protoc plugin: `go get github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go`

```
protoc --go_out=plugins=grpc:. proto/users/service.proto
protoc --go_out=plugins=grpc:. proto/list/service.proto
protoc --go_out=plugins=grpc:. proto/notifier/service.proto
```
