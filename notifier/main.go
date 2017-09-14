package main

import (
	"log"
	"net"

	"github.com/wizelineacademy/GoWorkshop/notifier/controllers"
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	srv := grpc.NewServer()
	notifier.RegisterNotifierServer(srv, &controllers.Service{})
	srv.Serve(lis)
}
