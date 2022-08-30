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

func (u *codeRepo) GetCodes(codeType int, key []string, values []int) ([]biz.Code, error) {
	codes := []biz.Code{}
	if err := u.data.db.Table("code").Where("key in (?) and code_type = ? and id in (values)", key, codeType).Scan(&codes).Error; err != nil {
		return nil, err
	}
	return codes, nil
}

func (u *codeRepo) GetCodesMAx() (int64, error) {
	var num int64
	if err := u.data.db.Table("code").Count(&num).Error; err != nil {
		return 0, err
	}
	return num, nil
}

func (u *codeRepo) Create(codes []biz.Code) error {
	return u.data.db.Create(&codes).Error
}
