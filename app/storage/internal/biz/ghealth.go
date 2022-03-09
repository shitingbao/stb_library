package biz

import "errors"

type GhealthRepo interface {
	SayHello(string) (string, error)
}

type GhealthUseCase struct {
	slog *SlogUseCase
	repo GhealthRepo
}

func NewGhealthUseCase(repo GhealthRepo, s *SlogUseCase) *GhealthUseCase {
	return &GhealthUseCase{repo: repo, slog: s}
}

func (c *GhealthUseCase) SayHello(name string) (string, error) {
	return "hello", c.slog.repo.SendOneLog("test", errors.New("test"))
	// return c.repo.SayHello(name)
}
