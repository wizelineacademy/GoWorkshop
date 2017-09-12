package main

import (
	"log"
	"net"

	"github.com/wizelineacademy/GoWorkshop/list/controllers"
	pb "github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/shared/config"
	"google.golang.org/grpc"
)

func main() {
	// tcp listener
	lis, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	config.Init()

	// grpc server with profiles endpoint
	srv := grpc.NewServer()
	pb.RegisterListServer(srv, &controllers.Service{})
	srv.Serve(lis)
}
