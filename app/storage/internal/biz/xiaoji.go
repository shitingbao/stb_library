package biz

import (
	"errors"
	"stb-library/app/storage/internal/model"
)

type XiaojiUseCase struct {
	sLog           *SlogUseCase
	defaultFileDir DefaultFileDir
	xiaoji         XiaojiRepo
}

type XiaojiRepo interface {
	GetMenuList(userId, parentId int) (model.MenuResult, error)
	CreateMenu(userId, parentId int, name string) error
	DeleteMenu(menuId, userId, parentId int) error
	UpdateMenuSort(userId, parentId int, newSort string) error
	UpdateAscription(menuId, userId, parentId, flagParentId int) error
	UpdateMenuName(Id int, name string) error
}

func NewXiaojiCase(defaultDir DefaultFileDir, s *SlogUseCase, x XiaojiRepo) *XiaojiUseCase {
	return &XiaojiUseCase{defaultFileDir: defaultDir, sLog: s, xiaoji: x}
}

func (x *XiaojiUseCase) GetMenuList(userId, parentId int) (model.MenuResult, error) {
	return x.xiaoji.GetMenuList(userId, parentId)
}

func (x *XiaojiUseCase) CreateMenu(userId, parentId int, name string) error {
	return x.xiaoji.CreateMenu(userId, parentId, name)
}

func (x *XiaojiUseCase) DeleteMenu(menuId, userId, parentId int) error {
	return x.xiaoji.DeleteMenu(menuId, userId, parentId)
}

func (x *XiaojiUseCase) UpdateMenuSort(userId, parentId int, newSort string) error {
	return x.xiaoji.UpdateMenuSort(userId, parentId, newSort)
}

func (x *XiaojiUseCase) UpdateAscription(menuId int, userId int, parentId int, flagParentId int) error {
	if parentId == 0 || flagParentId == 0 {
		return errors.New("parentId or flagParentId can not nil")
	}
	return x.xiaoji.UpdateAscription(menuId, userId, parentId, flagParentId)
}

func (x *XiaojiUseCase) UpdateMenuName(Id int, name string) error {
	return x.xiaoji.UpdateMenuName(Id, name)
}
