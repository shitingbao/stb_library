package biz

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/pborman/uuid"

	"golang.org/x/crypto/scrypt"
)

var UserVisitor = "visitor"

type UserBase struct {
	UserName string `form:"username" json:"username"`
	RealName string `form:"realName" json:"realName"`
	Avatar   string `form:"avatar" json:"avatar"`
	Email    string `form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
	HomePath string `form:"homePath" json:"homePath"`
}

type Role struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

//User 用户对象
type User struct {
	ID int `json:"userId"`
	UserBase
	Password []byte //对应mysql的varbinary,末尾不会填充，不能使用binary，因为不足会使用ox0填充导致取出的时候多18位的0
	Salt     string `form:"salt" json:"salt"`
	// CreateTime string `form:"create_time" json:"create_time"`
}

func (User) TableName() string {
	return "user"
}

type UserResult struct {
	User
	Token string `form:"token" json:"token"`
	Roles []Role `json:"roles"`
}

var Visitor = UserResult{
	User: User{
		UserBase: UserBase{
			UserName: "visitor",
			RealName: "visitor",
			HomePath: "/tools",
		},
	},
	Token: "visitor",
	Roles: []Role{
		{
			RoleName: "visitor",
			Value:    "visitor",
		},
	},
}

type ArgUser struct {
	UserBase
	Password string `form:"password" json:"password"`
}

type UserRepo interface {
	GetUser(username string) (*User, error)
	IsExistUser(username string) (bool, error)
	SaveUser(token string, val UserResult) error
	DelUser(ctx context.Context, token string)
	InsertUser(user *User) error
	GetRoles(id int) ([]Role, error)

	GetUserInfo(token string) (*UserResult, error)
	// Login(context.Context, *ArgUser) error
	// Logout(context.Context, *User) error
}

type UserUseCase struct {
	slog *SlogUseCase

	repo UserRepo
}

func NewUserCase(repo UserRepo, s *SlogUseCase) *UserUseCase {
	return &UserUseCase{repo: repo, slog: s}
}

func (u *UserUseCase) LoginOut(ctx context.Context, token string) error {
	u.repo.DelUser(ctx, token)
	return nil
}

//equal 密文验证
func (u *UserUseCase) equal(pwd, salt string, uPwd []byte) (bool, error) {
	bPwd, err := u.buildPas(pwd, salt)
	if err != nil {
		return false, err
	}
	return bytes.Equal(bPwd, uPwd), nil
}

func (u *UserUseCase) Login(ctx context.Context, pa *ArgUser) (UserResult, error) {
	userModel := UserResult{}
	if pa.UserName == "" || pa.Password == "" {
		return userModel, errors.New("UserName or Password cant not nil")
	}
	usr, err := u.repo.GetUser(pa.UserName)
	if err != nil {
		return userModel, err
	}
	isExists, err := u.equal(pa.Password, usr.Salt, usr.Password)
	if err != nil {
		return userModel, err
	}
	if !isExists {
		return userModel, errors.New("password have error")
	}

	token := uuid.NewUUID().String()

	roles, err := u.repo.GetRoles(usr.ID)
	if err != nil {
		return userModel, err
	}
	userModel.UserBase = usr.UserBase
	userModel.Token = token
	userModel.Roles = roles

	if err := u.repo.SaveUser(token, userModel); err != nil {
		return userModel, err
	}

	return userModel, nil
}

//前端的hex字符串
func (u *UserUseCase) huexEncode(md5Pwd string) ([]byte, error) {
	decoded, err := hex.DecodeString(md5Pwd)
	if err != nil {
		// logrus.WithFields(logrus.Fields{"decode": err}).Error("hex")
		return nil, err
	}
	return decoded, nil
}

//buildUserPassword 根据密码文本和盐生成密文
func (u *UserUseCase) buildUserPassword(pwdMd5, salt []byte) ([]byte, error) {
	return scrypt.Key(pwdMd5, salt, 16384, 8, 1, 32)
}

//buildPas 解析前端的hex密码文本，并调用密文生成函数
func (u *UserUseCase) buildPas(pwd, salt string) ([]byte, error) {
	h, err := u.huexEncode(pwd)
	if err != nil {
		return nil, err
	}
	bPwd, err := u.buildUserPassword(h, []byte(salt))
	if err != nil {
		// logrus.WithFields(logrus.Fields{"pwd": err}).Error("validPwdMd5")
		return nil, err
	}
	return bPwd, nil
}

//BuildIserSalt 随机获取用户中一段+uuid生成随机盐，防止代码泄密密码生成过程被破解
func (u *UserUseCase) BuildIserSalt(user string) string {
	rand.Seed(time.Now().UnixNano())
	sl := rand.Intn(len(user))
	return user[sl:] + base64.RawURLEncoding.EncodeToString(uuid.NewUUID())
}

func (u *UserUseCase) UserRegister(ctx context.Context, pa *ArgUser) error {
	if pa.UserName == "" || pa.Password == "" {
		return errors.New("必填内容不能为空")
	}
	isExist, err := u.repo.IsExistUser(pa.UserName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("this user have exists")
	}
	//两次加密一次解密，双向加单向
	salt := u.BuildIserSalt(pa.UserName)
	bPwd, err := u.buildPas(pa.Password, salt)
	if err != nil {
		return err
	}
	user := &User{}
	user.UserBase = pa.UserBase
	user.Salt = salt
	user.Password = bPwd

	return u.repo.InsertUser(user)
}

func (u *UserUseCase) GetUserInfo(token string) (*UserResult, error) {
	return u.repo.GetUserInfo(token)
}
