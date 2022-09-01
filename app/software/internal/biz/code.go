package biz

import (
	"context"
	"errors"
	"log"
	"stb-library/app/software/internal/model"
	"stb-library/lib/office"

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
	GetCodes(int, []string, []string) ([]bson.M, error)
	Create([]Code) error
	GetHeaderCode(int, string, []string) ([]bson.M, error)
	CreateHeaders([]Code) error
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

func (c *CodeUseCase) GetCodes(num int, key []string, filters []string) ([]bson.M, error) {
	return c.repo.GetCodes(num, key, filters)
}

func (c *CodeUseCase) CreateHeaders(codes []Code) error {
	return c.repo.CreateHeaders(codes)
}

func (c *CodeUseCase) GetHeaderCode(num int, key string, filters []string) ([]bson.M, error) {
	return c.repo.GetHeaderCode(num, key, filters)
}

func (c *CodeUseCase) CreateDocx(arg model.ArgDocx) ([]bson.M, error) {
	head, err := c.repo.GetHeaderCode(1, arg.TitleKey, arg.TitleFilters)
	if err != nil || len(head) < 1 {
		return nil, err
	}
	codes, err := c.repo.GetCodes(arg.ContentsNum, arg.ContentsKey, arg.ContentFilters)
	if err != nil {
		return nil, err
	}
	log.Println(head, codes[0][""])
	titilecon, ok := codes[0]["content"].(string)
	if !ok {
		return nil, errors.New("codes error")
	}
	contentList := []string{}
	for _, c := range codes {
		con, ok := c["content"].(string)
		if !ok {
			return nil, errors.New("codes error")
		}
		contentList = append(contentList, con)
	}
	office.CreateDocx("./test.docx", titilecon, contentList)
	return nil, nil
}
