package model

// menu 表 parent_id 为 0 代表一级目录，否则代表父节点目录 id
type Menu struct {
	Id       int    `form:"id" json:"id" gorm:"column:id"`
	UserId   int    `form:"user_id" json:"user_id" gorm:"column:user_id"`
	ParentId int    `form:"parent_id" json:"parent_id" gorm:"column:parent_id"`
	Name     string `form:"name" json:"name" gorm:"column:name"`
}

// menu_user 表 parent_id 为 0 代表一级目录，否则代表父节点目录 id
type MenuUser struct {
	Id       int    `form:"id" json:"id" gorm:"column:id"`
	UserId   int    `form:"user_id" json:"user_id" gorm:"column:user_id"`
	ParentId int    `form:"parent_id" json:"parent_id" gorm:"column:parent_id"`
	MenuSort string `form:"menu_sort" json:"menu_sort" gorm:"column:menu_sort"`
}

type MenuResult struct {
	MenuList []Menu
	Sort     string
}

type ArgMenu struct {
	Id           int    `form:"id" json:"id" gorm:"column:id"`
	UserId       int    `form:"user_id" json:"user_id" gorm:"column:user_id"`
	ParentId     int    `form:"parent_id" json:"parent_id" gorm:"column:parent_id"`
	FlagParentId int    `form:"flag_parent_id" json:"flag_parent_id"`
	Name         string `form:"name" json:"name" gorm:"column:name"`
	NewSort      string `form:"new_sort" json:"new_sort" gorm:"column:new_sort"`
}
