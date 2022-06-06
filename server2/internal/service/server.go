package service

import (
	"context"

	pb "server2/api/server"
)

type ServerService struct {
	pb.UnimplementedServerServer
}

func NewServerService() *ServerService {
	return &ServerService{}
}

func (s *ServerService) Createserver(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	return &pb.CreateServerReply{}, nil
}
func (s *ServerService) Updateserver(ctx context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
	return &pb.UpdateServerReply{}, nil
}
func (s *ServerService) Deleteserver(ctx context.Context, req *pb.DeleteServerRequest) (*pb.DeleteServerReply, error) {
	return &pb.DeleteServerReply{}, nil
}
func (s *ServerService) Getserver(ctx context.Context, req *pb.GetServerRequest) (*pb.GetServerReply, error) {
	return &pb.GetServerReply{}, nil
}
func (s *ServerService) Listserver(ctx context.Context, req *pb.ListServerRequest) (*pb.ListServerReply, error) {
	return &pb.ListServerReply{}, nil
}
