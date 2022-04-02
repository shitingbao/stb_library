package data

import "stb-library/app/storage/internal/biz"

var _ biz.XiaojiRepo = (*xiaojiRepo)(nil)

type xiaojiRepo struct {
	data *Data
}

func NewXiaojiRepo(d *Data) biz.XiaojiRepo {
	return &xiaojiRepo{
		data: d,
	}
}

func (x *xiaojiRepo) GetMenuList(userId int) ([]biz.Menu, error) {
	return []biz.Menu{}, nil
}
