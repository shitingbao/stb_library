package biz

import (
	"stb-library/lib/logger"

	"github.com/sirupsen/logrus"
)

type SlogRepo interface {
	SendOneLog(topic string, err error) error
	SendOneLogMes(topic string, content interface{}) error
}

type SlogUseCase struct {
	repo           SlogRepo       // 该日志发送第三方 rpc 接收
	logger         *logrus.Logger // 该日志记录本地，默认包含 panic 拦截
	defaultFileDir DefaultFileDir
}

func NewSlogUseCase(defaultDir DefaultFileDir, repo SlogRepo) *SlogUseCase {
	lg := logger.NewLogrusLog(defaultDir.DefaultDirPath)
	return &SlogUseCase{
		repo:           repo,
		logger:         lg,
		defaultFileDir: defaultDir,
	}
}

func (c *SlogUseCase) SendOneLog(topic string, err error) error {
	return c.repo.SendOneLog(topic, err)
}

func (c *SlogUseCase) SendOneLogMes(topic string, content interface{}) error {
	return c.repo.SendOneLogMes(topic, content)
}

func (c *SlogUseCase) Info(args ...interface{}) {
	c.logger.Info(args...)
}

func (c *SlogUseCase) Warning(args ...interface{}) {
	c.logger.Warning(args...)
}

func (c *SlogUseCase) Error(args ...interface{}) {
	c.logger.Error(args...)
}

func (c *SlogUseCase) Panic(args ...interface{}) {
	c.logger.Panic(args...)
}

func (c *SlogUseCase) WithFields(fields logrus.Fields) *logrus.Entry {
	return c.logger.WithFields(fields)
}
