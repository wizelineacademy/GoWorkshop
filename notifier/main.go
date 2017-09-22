package main

import (
	"log"
	"net"

	"github.com/wizelineacademy/GoWorkshop/notifier/handlers"
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"google.golang.org/grpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	srv := grpc.NewServer()
	notifier.RegisterNotifierServer(srv, &handlers.Service{})
	srv.Serve(listener)
}
