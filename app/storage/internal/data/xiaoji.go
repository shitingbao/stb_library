package data

import (
	"stb-library/app/storage/internal/biz"
	"strconv"

	"gorm.io/gorm"
)

var _ biz.XiaojiRepo = (*xiaojiRepo)(nil)

type xiaojiRepo struct {
	data *Data
}

func NewXiaojiRepo(d *Data) biz.XiaojiRepo {
	return &xiaojiRepo{
		data: d,
	}
}

func (x *xiaojiRepo) GetMenuList(userId, parentId int) (biz.MenuResult, error) {
	res := biz.MenuResult{}
	menuList, err := x.menuList(userId, parentId)
	if err != nil {
		return res, err
	}

	sort, err := x.menuSort(userId, parentId)
	if err != nil {
		return res, err
	}
	res.MenuList = menuList
	res.Sort = sort
	return res, nil
}

// parentId 为 0 代表第一级目录，不为 0 代表第二级目录
func (x *xiaojiRepo) menuList(userId, parentId int) ([]biz.Menu, error) {
	list := []biz.Menu{}
	if err := x.data.db.Table("menu").
		Where("user_id = ? and parent_id = ?", userId, parentId).Scan(&list).Error; err != nil {
		return nil, err
	}
	return []biz.Menu{}, nil
}

// 反馈目录 id 顺序集合
// 一级目录没有根目录id menu_id
func (x *xiaojiRepo) menuSort(userId, parentId int) (string, error) {
	menuSort := ""
	if err := x.data.db.Table("menu_user").Select("menu_sort").
		Where("user_id = ? and parent_id = ?", userId, parentId).
		Scan(&menuSort).Error; err != nil {
		return "", err
	}
	return menuSort, nil
}

// CreateMenu 创建菜单，注意一级菜单和二级的菜单排序
func (x *xiaojiRepo) CreateMenu(userId, parentId int, name string) error {
	db := x.data.db.Begin()
	defer db.Commit()
	menu, err := x.createMenu(userId, parentId, name, db)
	if err != nil {
		db.Rollback()
		return err
	}
	// 二级目录添加更新排序列表即可
	if parentId != 0 {
		if err := x.updateMenuSort(userId, menu.Id, parentId, db); err != nil {
			db.Rollback()
			return err
		}
		return nil
	}

	if err := x.createFirstMenuSort(userId, menu.Id, parentId, db); err != nil {
		db.Rollback()
		return err
	}
	return nil
}

// 创建菜单
func (x *xiaojiRepo) createMenu(userId, parentId int, name string, db *gorm.DB) (biz.Menu, error) {
	menu := biz.Menu{
		UserId:   userId,
		ParentId: parentId,
		Name:     name,
	}
	err := db.Table("menu").Create(&menu).Error
	return menu, err
}

// 生成 或者 更新 第一级目录排序列表
// 更新本级目录排序，并且生成下一级排序
func (x *xiaojiRepo) createFirstMenuSort(userId, menuId, parentId int, db *gorm.DB) error {
	var num int64 = 0
	if err := db.Table("menu").
		Where("user_id = ? and parent_id = ?", userId, parentId).
		Count(&num).Error; err != nil {
		return err
	}
	//说明还没目录 create
	if num == 0 {
		if err := x.createMenuSort(userId, parentId, strconv.Itoa(menuId), db); err != nil {
			return err
		}
	} else {
		if err := x.updateMenuSort(userId, menuId, parentId, db); err != nil {
			return err
		}
	}
	// 第一级目录新增时，为下级目录添加空的排序目录
	return x.createMenuSort(userId, menuId, "", db)
}

func (x *xiaojiRepo) createMenuSort(userId, menuId int, sort string, db *gorm.DB) error {
	menuUser := &biz.MenuUser{
		UserId:   userId,
		ParentId: menuId,
		MenuSort: sort,
	}
	return db.Table("menu_user").Create(menuUser).Error
}

// 更新目录排序
func (x *xiaojiRepo) updateMenuSort(userId, menuId, parentId int, db *gorm.DB) error {
	sort, err := x.menuSort(userId, parentId)
	if err != nil {
		return err
	}
	if sort == "" {
		sort = strconv.Itoa(menuId)
	} else {
		sort += "," + strconv.Itoa(menuId)
	}
	if err := db.Table("menu_user").
		Where("user_id = ? and parent_id = ?", userId, parentId).
		Update("sort", sort).Error; err != nil {
		return err
	}
	return nil
}

// 删除直接传新生成的排序列表
func (x *xiaojiRepo) DeleteMenu(userId, menuId, parentId int, NewSort string) error {
	db := x.data.db.Begin()
	defer db.Commit()
	if err := db.Table("menu").Where("id = ?", menuId).
		Delete(&biz.Menu{}).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Table("menu_user").Where("user_id = ? and parent_id = ?", userId, parentId).
		Update("sort", NewSort).Error; err != nil {
		db.Rollback()
		return err
	}
	if parentId != 0 {
		return nil
	}
	// 一级目录需要删除以他为父节点的所有子节点
	if err := db.Table("menu_user").Where("user_id = ? and parent_id = ?", userId, menuId).
		Delete(&biz.MenuUser{}).Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
