package biz

type SlogRepo interface {
	SendOneLog(topic string, err error) error
	SendOneLogMes(topic string, content interface{}) error
}

type SlogUseCase struct {
	repo SlogRepo
}

func NewSlogUseCase(repo SlogRepo) *SlogUseCase {
	return &SlogUseCase{repo: repo}
}

func (c *SlogUseCase) SendOneLog(topic string, err error) error {
	return c.repo.SendOneLog(topic, err)
}

func (c *SlogUseCase) SendOneLogMes(topic string, content interface{}) error {
	return c.repo.SendOneLogMes(topic, content)
}
