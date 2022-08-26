package biz

import (
	"errors"
)

type CentralRepo interface {
	SayHello(string) (string, error)
}

type CentralUseCase struct {
	sLog *SlogUseCase
	repo CentralRepo
}

func NewCentralUseCase(repo CentralRepo, s *SlogUseCase) *CentralUseCase {
	return &CentralUseCase{repo: repo, sLog: s}
}

func (c *CentralUseCase) SayHello(name string) (string, error) {
	c.sLog.logger.Info("this is info")
	c.sLog.logger.Error("this is error")
	// c.slog.logger.WithFields(logrus.Fields{"hello": "panic"}).Panic("logger panic")
	return "hello", c.sLog.repo.SendOneLog("test", errors.New("test"))
}
