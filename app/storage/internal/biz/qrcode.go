package biz

import (
	"context"
	"stb-library/lib/qrcode"

	"github.com/go-kratos/kratos/v2/log"
)

type QrcodeUseCase struct {
	log *log.Helper
}

func NewQrcodeCase(repo UserRepo, logger log.Logger) *QrcodeUseCase {
	return &QrcodeUseCase{log: log.NewHelper(logger)}
}

func (u *QrcodeUseCase) QrcodeEncoder(ctx context.Context, mes string) (string, error) {
	return qrcode.GenerateQR(mes)
}

func (u *QrcodeUseCase) QrcodeDecoder(ctx context.Context, imageURL string) (string, error) {
	q, err := qrcode.Decode(imageURL)
	if err != nil {
		return "", err
	}
	return q.Content, nil
}
