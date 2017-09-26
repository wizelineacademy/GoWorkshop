package main

import (
	"log"
	"net"

	"github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"github.com/wizelineacademy/GoWorkshop/users/server"
	"google.golang.org/grpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	srv := grpc.NewServer()
	users.RegisterUsersServer(srv, &server.Server{})
	srv.Serve(listener)
}
