package services

import (
	"context"
	"github.com/Sakenzhassulan/salt-service/config"
	"github.com/Sakenzhassulan/salt-service/internal/pb"
	"github.com/Sakenzhassulan/salt-service/internal/utils"
	"net"
)

type Server struct {
	Lis net.Listener
}

func NewServer(config config.Config) (*Server, error) {
	lis, err := net.Listen("tcp", config.ServerPort)
	if err != nil {
		return nil, err
	}
	return &Server{
		Lis: lis,
	}, err
}

func (s *Server) GenerateSalt(ctx context.Context, request *pb.GenerateSaltRequest) (*pb.GenerateSaltResponse, error) {
	randomString := utils.GenerateRandomString()
	return &pb.GenerateSaltResponse{
		Salt: randomString,
	}, nil
}
