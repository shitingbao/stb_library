package biz

import "errors"

type CentralRepo interface {
	SayHello(string) (string, error)
}

type CentralUseCase struct {
	slog *SlogUseCase
	repo CentralRepo
}

func NewCentralUseCase(repo CentralRepo, s *SlogUseCase) *CentralUseCase {
	return &CentralUseCase{repo: repo, slog: s}
}

func (c *CentralUseCase) SayHello(name string) (string, error) {
	return "hello", c.slog.repo.SendOneLog("test", errors.New("test"))
	// return c.repo.SayHello(name)
}
