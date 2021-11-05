package data

import (
	"context"
	"encoding/json"
	"errors"
	"stb-library/app/storage/internal/biz"
	"stb-library/lib/rediser"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(d *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: d,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) DelUser(ctx context.Context, token string) {
	rediser.DelUser(u.data.rds, token)
}

//GetUser 获取一个user，不存在反馈 err
func (u *userRepo) GetUser(username string) (*biz.User, error) {
	ur := []*biz.User{}
	if err := u.data.db.Table("user").Where("user_name = ?", username).Scan(&ur).Error; err != nil {
		// logrus.WithFields(logrus.Fields{"get user": err}).Error("user")
		return nil, err
	}
	if len(ur) == 0 {
		return nil, errors.New("no exists this user")
	}
	return ur[0], nil
}

//IsExistUser 判断用户是否存在，存在为true
func (u *userRepo) IsExistUser(username string) (bool, error) {
	var num int64
	if err := u.data.db.Table("user").Where("user_name = ?", username).Count(&num).Error; err != nil {
		// logrus.WithFields(logrus.Fields{"get user": err}).Error("user")
		return false, err
	}
	return num > 0, nil
}

// SaveUser 注册保存用户信息状态
func (u *userRepo) SaveUser(token string, val biz.UserResult) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return rediser.SaveUser(u.data.rds, token, string(b))
}

func (u *userRepo) InsertUser(user *biz.User) error {
	return u.data.db.Create(user).Error
}

func (u *userRepo) GetRoles(id int) ([]biz.Role, error) {
	roles := []biz.Role{}
	if err := u.data.db.Table("role").Where("user_id = ?", id).Scan(&roles).Error; err != nil {
		return roles, err
	}
	return roles, nil
}

func (u *userRepo) GetUserInfo(token string) (*biz.UserResult, error) {
	model := &biz.UserResult{}
	uInfo, err := rediser.GetUserInfo(u.data.rds, token)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(uInfo), model); err != nil {
		return model, err
	}
	return model, nil
}
