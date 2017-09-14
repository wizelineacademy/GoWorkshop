package main

import (
	"log"
	"net"

	pb "github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"github.com/wizelineacademy/GoWorkshop/users/controllers"
	"google.golang.org/grpc"
)

func main() {
	// tcp listener
	lis, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	// grpc server with profiles endpoint
	srv := grpc.NewServer()
	pb.RegisterUsersServer(srv, &controllers.Service{})
	srv.Serve(lis)
}
