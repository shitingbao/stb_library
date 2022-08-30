package biz

import (
	"context"
)


type Code struct {
	UserName string `form:"username" json:"username" gorm:"column:username"`
	RealName string `form:"realName" json:"realName" gorm:"column:real_name"`
	Avatar   string `form:"avatar" json:"avatar"`
	Email    string `form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
	HomePath string `form:"homePath" json:"homePath" gorm:"column:home_path"`
}


func (Code) TableName() string {
	return "code"
}


type ArgCode struct {
	Password string `form:"password" json:"password"`
}

type CodeRepo interface {
	Delete( context.Context,  string) error
	GetCodes(string, string) (*Code, error)
	Create([]*Code) error 
}

type CodeUseCase struct {
	sLog *SlogUseCase

	repo CodeRepo
}

func NewCodeCase(repo CodeRepo, s *SlogUseCase) *CodeUseCase {
	return &CodeUseCase{repo: repo, sLog: s}
}