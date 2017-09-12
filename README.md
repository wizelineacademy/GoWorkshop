![diagram](https://github.com/wizelineacademy/GoWorkshop/raw/master/diagram.png)

### "ToDo List" experiment project

This repository is an example project which demonstrates the use of microservices for a fictional ToDo list application. The ToDo backend is powered by 3 microservices, all of which happen to be written in Go, using MongoDB for manage the database and Docker to isolate and deploy the ecosystem.

### Services organization

The application consists of the following application services:

| Service  | Port  | Description                   | Methods                             |
|----------|-------|-------------------------------|-------------------------------------|
| users    | 50000 | Provides users information    | CreateUser, GetAllUsers             |
| list     | 50001 | Manages items in todo lists   | CreateItem, GetAllItems, DeleteItem |
| notifier | 50002 | Send email notifications      | Will be implemented during Workshop |

Client web application is working on [http://127.0.0.1:3000](http://127.0.0.1:3000).

### Workshop task

 - Create `notifier` service.
 - Call `notifier` from `users` service to send email notification.

#### Tips

New service needs separate folder, can be copied from `list` and modified later. The best place to start is to define a service using [proto3 language](https://developers.google.com/protocol-buffers/docs/proto3), in `proto/notifier/service.proto`. Then implement handlers in `notifier/controllers/handlers.go`.

Service should be also added to `docker-compose.yml` and work on port 50002.

### Requirements

 - Install [Docker](https://www.docker.com/get-docker)
 - Install [Docker Compose](https://docs.docker.com/compose/install)

### Run

```
docker-compose up -d
```

### Client

Client is built with Go and located in `web`, working on [http://127.0.0.1:3000](http://127.0.0.1:3000).

### How to test gRPC

[omgRPC](https://github.com/troylelandshields/omgrpc) aims to be a GUI client for interacting with gRPC services, similar to what Postman is for REST APIs.

 - Install [omgRPC](https://github.com/troylelandshields/omgrpc/releases)
 - Open Protofile (Select `proto/users/service.proto` file to test `users` service)
 - Use `127.0.0.1:50000` address for `users`

### Generate source code for the gRPC client from .proto files

 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)

```
protoc --go_out=plugins=grpc:. proto/users/service.proto
protoc --go_out=plugins=grpc:. proto/list/service.proto
```
