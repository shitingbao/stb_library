package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "stb-library/api/central/v1"
	"stb-library/app/central/internal/biz"
)

type CentralService struct {
	pb.UnimplementedCentralServer

	central *biz.CentralUsecase
	log     *log.Helper
}

// 唯一的服务，注册入 http 、 grpc
// 如果有其他逻辑，统一放在 CentralService 下，因为一个系统就是一个服务
// 与其他模块不同的是，其他模块可根据业务进行细化，多对象的划分
func NewCentralService(cen *biz.CentralUsecase, lg log.Logger) *CentralService {
	return &CentralService{
		central: cen,
		log:     log.NewHelper(lg),
	}
}

func (s *CentralService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return s.central.SayHello(ctx, req)
}

func (s *CentralService) Healthy(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return s.central.Healthy(ctx, req)
}
