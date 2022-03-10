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
	repo           SlogRepo       // 该日志发送第三方 rpc 接受
	logger         *logrus.Logger // 该日志记录本地，默认包含 panic 接受
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
