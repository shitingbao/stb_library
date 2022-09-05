package biz

import (
	"context"
	"errors"
	"stb-library/app/software/internal/model"
	"stb-library/lib/office"

	"gopkg.in/mgo.v2/bson"
)

type CodeRepo interface {
	Delete(context.Context, string) error
	GetCodes(int, string, []string, []string) ([]bson.M, error)
	Create(model.ArgCode) error
	GetHeaderCode(int, string, []string) ([]bson.M, error)
	CreateHeaders(model.ArgCode) error
}

type CodeUseCase struct {
	sLog *SlogUseCase
	repo CodeRepo
}

func NewCodeCase(repo CodeRepo, s *SlogUseCase) *CodeUseCase {
	return &CodeUseCase{repo: repo, sLog: s}
}

func (c *CodeUseCase) Create(arg model.ArgCode) error {
	return c.repo.Create(arg)
}

func (c *CodeUseCase) GetCodes(num int, lan string, key []string, filters []string) ([]bson.M, error) {
	return c.repo.GetCodes(num, lan, key, filters)
}

func (c *CodeUseCase) CreateHeaders(arg model.ArgCode) error {
	return c.repo.CreateHeaders(arg)
}

func (c *CodeUseCase) GetHeaderCode(num int, key string, filters []string) ([]bson.M, error) {
	return c.repo.GetHeaderCode(num, key, filters)
}

func (c *CodeUseCase) CreateDocx(arg model.ArgDocx) ([]bson.M, error) {
	head, err := c.repo.GetHeaderCode(1, arg.Language, arg.HeaderFilters)
	if err != nil || len(head) < 1 {
		return nil, err
	}
	hd, ok := head[0]["content"].(string)
	if !ok {
		return nil, errors.New("codes error")
	}
	codes, err := c.repo.GetCodes(arg.ContentsNum, arg.Language, arg.ContentsKey, arg.ContentFilters)
	if err != nil || len(codes) < 1 {

		return nil, errors.New("codes error")
	}

	contentList := []string{}
	contentList = append(contentList, hd) // 语言类型
	for _, c := range codes {
		con, ok := c["content"].(string)
		if !ok {
			return nil, errors.New("codes error")
		}
		contentList = append(contentList, arg.ContentTitle)
		contentList = append(contentList, con)
	}
	// log.Println(hd, contentList)
	office.CreateDocx("./test.docx", arg.HeaderContent, contentList)
	return nil, nil
}
