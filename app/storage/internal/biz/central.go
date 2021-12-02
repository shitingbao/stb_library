package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type CentralRepo interface {
	SayHello(string) (string, error)
}

type CentralUseCase struct {
	repo CentralRepo

	log *log.Helper
}

func NewCentralUseCase(repo CentralRepo, logger log.Logger) *CentralUseCase {
	return &CentralUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *CentralUseCase) SayHello(name string) (string, error) {
	return c.repo.SayHello(name)
}
