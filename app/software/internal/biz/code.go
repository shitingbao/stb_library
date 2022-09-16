package biz

import (
	"context"
	"encoding/json"
	"errors"
	"path"
	"stb-library/app/software/internal/conf"
	"stb-library/app/software/internal/model"
	"stb-library/lib/office"

	"gopkg.in/mgo.v2/bson"
)

type CodeRepo interface {
	Delete(context.Context, string) error
	GetCodes(int, string, []string, []string) ([]bson.M, error)
	Create([]model.Code) error
	GetHeaderCode(int, string, []string) ([]bson.M, error)
	CreateHeaders([]model.Code) error
}

type CodeUseCase struct {
	office *conf.Office
	repo   CodeRepo
}

func NewCodeCase(repo CodeRepo, o *conf.Office) *CodeUseCase {
	return &CodeUseCase{repo: repo, office: o}
}

func (c *CodeUseCase) Create(arg string) error {
	codes := []model.Code{}
	if err := json.Unmarshal([]byte(arg), &codes); err != nil {
		return err
	}
	return c.repo.Create(codes)
}

func (c *CodeUseCase) GetCodes(num int, lan string, key []string, filters []string) ([]bson.M, error) {
	return c.repo.GetCodes(num, lan, key, filters)
}

func (c *CodeUseCase) CreateHeaders(arg string) error {
	codes := []model.Code{}
	if err := json.Unmarshal([]byte(arg), &codes); err != nil {
		return err
	}
	return c.repo.CreateHeaders(codes)
}

func (c *CodeUseCase) GetHeaderCode(num int, key string, filters []string) ([]bson.M, error) {
	return c.repo.GetHeaderCode(num, key, filters)
}

func (c *CodeUseCase) CreateDocx(param, assetsPath string) ([]bson.M, error) {
	arg := model.ArgDocx{}
	if err := json.Unmarshal([]byte(param), &arg); err != nil {
		return nil, err
	}

	head, err := c.repo.GetHeaderCode(1, arg.Language, arg.HeaderFilters)
	if err != nil || len(head) < 1 {
		return nil, err
	}
	hd, ok := head[0]["content"].(string)
	if !ok {
		return nil, errors.New("codes error")
	}
	codes, err := c.repo.GetCodes(arg.ContentsNum, arg.Language, arg.ContentsKeys, arg.ContentFilters)
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
	doc, err := office.NewOfficeDocx(c.office.Docx.Secret)
	if err != nil {
		return nil, err
	}
	doc.CreateDocx(path.Join(assetsPath, "software.docx"), arg.HeaderContent, contentList)
	return nil, nil
}
