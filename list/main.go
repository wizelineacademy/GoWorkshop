package main

import (
	"log"
	"net"

	"github.com/wizelineacademy/GoWorkshop/list/controllers"
	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	srv := grpc.NewServer()
	list.RegisterListServer(srv, &controllers.Service{})
	srv.Serve(lis)
}
