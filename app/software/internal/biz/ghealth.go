package biz

import "errors"

type GhealthRepo interface {
	SayHello(string) (string, error)
}

type GhealthUseCase struct {
	sLog *SlogUseCase
	repo GhealthRepo
}

func NewGhealthUseCase(repo GhealthRepo, s *SlogUseCase) *GhealthUseCase {
	return &GhealthUseCase{repo: repo, sLog: s}
}

func (c *GhealthUseCase) SayHello(name string) (string, error) {
	return "hello", c.sLog.repo.SendOneLog("test", errors.New("test"))
	// return c.repo.SayHello(name)
}
