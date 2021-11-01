package data

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"stb-library/app/storage/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) Login(ctx context.Context, b *biz.User) error {
	return errors.New("")
}

func (u *userRepo) Logout(ctx context.Context, b *biz.User) error {
	return errors.New("")
}

//GetUser 获取一个user
func (u *userRepo) GetUser(username string) (*biz.User, error) {
	ur := biz.User{}
	if err := u.data.db.Table("user").Where("name = ?", username).Scan(&u).Error; err != nil {
		// logrus.WithFields(logrus.Fields{"get user": err}).Error("user")
		return nil, err
	}
	return &ur, nil
}

//IsExistUser 判断用户是否存在，存在为true
func (u *userRepo) IsExistUser(username string) (bool, error) {
	var num int64
	if err := u.data.db.Table("user").Where("name = ?", username).Count(&num).Error; err != nil {
		// logrus.WithFields(logrus.Fields{"get user": err}).Error("user")
		return false, err
	}
	return num > 0, nil
}

//前端的hex字符串
func (u *userRepo) huexEncode(md5Pwd string) []byte {
	decoded, err := hex.DecodeString(md5Pwd)
	if err != nil {
		logrus.WithFields(logrus.Fields{"decode": err}).Error("hex")
	}
	return decoded
}

//BuildIserSalt 随机获取用户中一段+uuid生成随机盐，防止代码泄密密码生成过程被破解
func (u *userRepo) BuildIserSalt(user string) string {
	rand.Seed(time.Now().UnixNano())
	sl := rand.Intn(len(user))
	return user[sl:] + base64.RawURLEncoding.EncodeToString(uuid.NewUUID())
}

//buildUserPassword 根据密码文本和盐生成密文
func (u *userRepo) buildUserPassword(pwdMd5, salt []byte) ([]byte, error) {
	return scrypt.Key(pwdMd5, salt, 16384, 8, 1, 32)
}

//Equal 密文验证
func (u *userRepo) Equal(pwd, salt string, uPwd []byte) bool {
	bPwd := u.BuildPas(pwd, salt)
	return bytes.Equal(bPwd, uPwd)
}

//BuildPas 解析前端的hex密码文本，并调用密文生成函数
func (u *userRepo) BuildPas(pwd, salt string) []byte {
	bPwd, err := u.buildUserPassword(u.huexEncode(pwd), []byte(salt))
	if err != nil {
		// logrus.WithFields(logrus.Fields{"pwd": err}).Error("validPwdMd5")
	}
	return bPwd
}
