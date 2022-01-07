package biz

import (
	"context"
	v1 "stb-library/api/central/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type CentralRepo interface {
	SayHelloData(string) (string, error)
	HealthyData(string) (string, error)
}

type CentralUsecase struct {
	repo CentralRepo
	log  *log.Helper
}

func NewCentralUsecase(repo CentralRepo, logger log.Logger) *CentralUsecase {
	return &CentralUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *CentralUsecase) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	mes, err := s.repo.SayHelloData(req.Name)
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{
		Message: mes,
	}, nil
}

func (s *CentralUsecase) Healthy(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	mes, err := s.repo.HealthyData(req.Name)
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{
		Message: mes,
	}, nil
}
