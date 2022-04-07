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
	DeleteMenu(userId, menuId, parentId int, NewSort string) error
	UpdateMenuSort(userId, parentId int, newSort string) error
	UpdateAscription(userId, menuId, parentId, flagParentId int, newSort string) error
	UpdateMenuName(userId, parentId int, name string) error
}

func NewXiaojiCase(defaultDir DefaultFileDir, s *SlogUseCase) *XiaojiUseCase {
	return &XiaojiUseCase{defaultFileDir: defaultDir, sLog: s}
}

func (x *XiaojiUseCase) GetMenuList(userId, parentId int) (model.MenuResult, error) {
	return x.xiaoji.GetMenuList(userId, parentId)
}

func (x *XiaojiUseCase) CreateMenu(userId, parentId int, name string) error {
	return x.xiaoji.CreateMenu(userId, parentId, name)
}

func (x *XiaojiUseCase) DeleteMenu(userId, menuId, parentId int, NewSort string) error {
	return x.xiaoji.DeleteMenu(userId, menuId, parentId, NewSort)
}

func (x *XiaojiUseCase) UpdateMenuSort(userId int, parentId int, newSort string) error {
	return x.xiaoji.UpdateMenuSort(userId, parentId, newSort)
}

func (x *XiaojiUseCase) UpdateAscription(userId int, menuId int, parentId int, flagParentId int, newSort string) error {
	if parentId == 0 || flagParentId == 0 {
		return errors.New("parentId or flagParentId can not nil")
	}
	return x.xiaoji.UpdateAscription(userId, menuId, parentId, flagParentId, newSort)
}

func (x *XiaojiUseCase) UpdateMenuName(userId int, parentId int, name string) error {
	return x.xiaoji.UpdateMenuName(userId, parentId, name)
}
