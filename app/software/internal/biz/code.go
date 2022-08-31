package biz

import (
	"context"
)

type Code struct {
	Key        string `form:"key" json:"key" gorm:"column:key"`
	Content    string `form:"content" json:"content" gorm:"column:content"`
	CodeType   int    `form:"code_type" json:"code_type" gorm:"column:code_type"`
	CodeLength int    `form:"code_length" json:"code_length" gorm:"column:code_length"`
}

func (Code) TableName() string {
	return "code"
}

type ArgCode struct {
	Password string `form:"password" json:"password"`
}

type CodeRepo interface {
	Delete(context.Context, string) error
	GetCodes([]int) ([]Code, error)
	Create([]Code) error
	GetCodesMAx() (int64, error)
}

type CodeUseCase struct {
	sLog *SlogUseCase
	repo CodeRepo
}

func NewCodeCase(repo CodeRepo, s *SlogUseCase) *CodeUseCase {
	return &CodeUseCase{repo: repo, sLog: s}
}
