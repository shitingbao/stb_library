package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// type User struct {
// 	UserName string `form:"user_name" json:"user_name"`
// 	Password string `form:"password" json:"password"`
// }

//User 用户对象
type User struct {
	Name       string
	Pwd        []byte //对应mysql的varbinary,末尾不会填充，不能使用binary，因为不足会使用ox0填充导致取出的时候多18位的0
	Avatar     string
	Email      string
	Phone      string
	Salt       string
	CreateTime string
}

func (User) TableName() string {
	return "user"
}

type UserRepo interface {
	Login(context.Context, *User) error
	Logout(context.Context, *User) error
}

type Usercase struct {
	repo UserRepo
	log  *log.Helper
}
