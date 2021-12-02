package data

import (
	"context"
	v1 "stb-library/api/central/v1"
	"stb-library/app/storage/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CentralRepo = (*centralRepo)(nil)

type centralRepo struct {
	data *Data
	log  *log.Helper
}

func NewCentralRepo(da *Data, lg log.Logger) biz.CentralRepo {
	return &centralRepo{
		data: da,
		log:  log.NewHelper(log.With(lg, "module", "data/user")),
	}
}

func (c *centralRepo) SayHello(name string) (string, error) {
	res, err := c.data.ce.SayHello(context.TODO(), &v1.HelloRequest{Name: name})
	return res.Message, err
}
