package data

import (
	"context"
	v1 "stb-library/api/central/v1"
	"stb-library/app/storage/internal/biz"
)

var _ biz.CentralRepo = (*centralRepo)(nil)

type centralRepo struct {
	data *Data
}

func NewCentralRepo(da *Data) biz.CentralRepo {
	return &centralRepo{
		data: da,
	}
}

func (c *centralRepo) SayHello(name string) (string, error) {
	res, err := c.data.ce.SayHello(context.Background(), &v1.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return res.Message, err
}
