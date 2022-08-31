package biz

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

type Code struct {
	Key        string `form:"key" json:"key" gorm:"column:key"`
	Content    string `form:"content" json:"content" gorm:"column:content"`
	CodeLength int    `form:"code_length" json:"code_length" gorm:"column:code_length"`
}

func (Code) TableName() string {
	return "code"
}

type ArgCode struct {
	Num     int      `form:"num" json:"num"`
	Key     string   `form:"key" json:"key"`
	Filters []string `form:"filters" json:"filters"`
}

type CodeRepo interface {
	Delete(context.Context, string) error
	GetCodes(int, string, []string) ([]bson.M, error)
	Create([]Code) error
}

type CodeUseCase struct {
	sLog *SlogUseCase
	repo CodeRepo
}

func NewCodeCase(repo CodeRepo, s *SlogUseCase) *CodeUseCase {
	return &CodeUseCase{repo: repo, sLog: s}
}

func (c *CodeUseCase) Create(codes []Code) error {
	return c.repo.Create(codes)
}

func (c *CodeUseCase) GetCodes(num int, key string, filters []string) ([]bson.M, error) {
	return c.repo.GetCodes(num, key, filters)
}
