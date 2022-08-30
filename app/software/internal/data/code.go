package data

import (
	"context"
	"stb-library/app/software/internal/biz"
)

var _ biz.CodeRepo = (*codeRepo)(nil)

type codeRepo struct {
	data *Data
}

func NewCodeRepo(d *Data) biz.CodeRepo {
	return &codeRepo{
		data: d,
	}
}

func (u *codeRepo) Delete(ctx context.Context, token string) error {
	// rediser.DelCode(u.data.rds, token)
	return nil
}

// GetUser 获取一个user，不存在反馈 err
func (u *codeRepo) GetCodes(key, codeType string) (*biz.Code, error) {
	ur := []*biz.Code{}
	return ur[0], nil
}

func (u *codeRepo) Create(codes []*biz.Code) error {
	return u.data.db.Create(codes).Error
}
