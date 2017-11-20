package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/users/server"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err == nil {
		log.Info("service started on :8080")
	} else {
		log.WithError(err).Fatal("could not start service")
	}

	srv := grpc.NewServer()
	users.RegisterUsersServer(srv, &server.Server{})
	srv.Serve(listener)
}
