package biz

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

type CodeHeader struct {
	Key        string `form:"key" json:"key" gorm:"column:key"`
	Content    string `form:"content" json:"content" gorm:"column:content"`
	CodeLength int    `form:"code_length" json:"code_length" gorm:"column:code_length"`
}

func (CodeHeader) TableName() string {
	return "code"
}

type CodeHeaderRepo interface {
	Delete(context.Context, string) error
	GetCodes(int, string, []string) ([]bson.M, error)
	Create([]CodeHeader) error
}

type CodeHeaderUseCase struct {
	sLog *SlogUseCase
	repo CodeHeaderRepo
}

func NewCodeHeaderCase(repo CodeHeaderRepo, s *SlogUseCase) *CodeHeaderUseCase {
	return &CodeHeaderUseCase{repo: repo, sLog: s}
}

func (c *CodeHeaderUseCase) Create(codes []CodeHeader) error {
	return c.repo.Create(codes)
}

func (c *CodeHeaderUseCase) GetCodes(num int, key string, filters []string) ([]bson.M, error) {
	return c.repo.GetCodes(num, key, filters)
}
