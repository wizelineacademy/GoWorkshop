package main

import (
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/wizelineacademy/GoWorkshop/notifier/server"
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.WithError(err).Error("could not start service")
		return
	}

	log.Info("starting service on :8080")

	srv := grpc.NewServer()
	notifier.RegisterNotifierServer(srv, &server.Server{})
	srv.Serve(listener)
}
