package data

import (
	"stb-library/app/base/internal/biz"
)

var _ biz.GhealthRepo = (*healthRepo)(nil)

type healthRepo struct {
	data *Data
}

func NewHealthRepo(da *Data) biz.GhealthRepo {
	return &healthRepo{
		data: da,
	}
}

func (c *healthRepo) SayHello(name string) (string, error) {
	return "hello " + name, nil
}
