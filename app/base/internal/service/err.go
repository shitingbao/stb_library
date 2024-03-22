package service

import (
	"context"

	pb "stb-library/api/base/v1"
)

type ErrService struct {
	pb.UnimplementedErrServer
}

func NewErrService() *ErrService {
	return &ErrService{}
}

func (s *ErrService) CreateErr(ctx context.Context, req *pb.CreateErrRequest) (*pb.CreateErrReply, error) {
	return &pb.CreateErrReply{}, nil
}
func (s *ErrService) UpdateErr(ctx context.Context, req *pb.UpdateErrRequest) (*pb.UpdateErrReply, error) {
	return &pb.UpdateErrReply{}, nil
}
func (s *ErrService) DeleteErr(ctx context.Context, req *pb.DeleteErrRequest) (*pb.DeleteErrReply, error) {
	return &pb.DeleteErrReply{}, nil
}
func (s *ErrService) GetErr(ctx context.Context, req *pb.GetErrRequest) (*pb.GetErrReply, error) {
	return &pb.GetErrReply{}, nil
}
func (s *ErrService) ListErr(ctx context.Context, req *pb.ListErrRequest) (*pb.ListErrReply, error) {
	return &pb.ListErrReply{}, nil
}
