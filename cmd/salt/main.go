package main

import (
	"github.com/Sakenzhassulan/salt-service/config"
	"github.com/Sakenzhassulan/salt-service/internal/pb"
	"github.com/Sakenzhassulan/salt-service/internal/services"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	server, err := services.NewServer(conf)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSaltServiceServer(grpcServer, server)
	if err = grpcServer.Serve(server.Lis); err != nil {
		log.Fatal(err)
	}
}
