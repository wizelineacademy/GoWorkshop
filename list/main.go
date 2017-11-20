package main

import (
	"log"
	"net"

	"github.com/wizelineacademy/GoWorkshop/list/server"
	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"google.golang.org/grpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	srv := grpc.NewServer()
	list.RegisterListServer(srv, &server.Server{})
	srv.Serve(listener)
}
