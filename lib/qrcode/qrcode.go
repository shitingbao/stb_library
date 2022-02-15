// qrcode 二维码生成 和 解码
package qrcode

import (
	"errors"
	"os"
	"path"

	"github.com/pborman/uuid"
	qrcode "github.com/skip2/go-qrcode"
	qrcodedecode "github.com/tuotoo/qrcode"
)

type qrOptions struct {
	filePath    string // 输出目录
	fileName    string // 文件名称
	size        int    // 图片的尺码，默认都是长宽一致
	completeURL string // 完整的图片路径
}

type Option func(*qrOptions) error

func WithFilePath(path string) Option {
	return func(o *qrOptions) error {
		f, err := os.Stat(path)
		if err != nil {
			return err
		}
		if !f.IsDir() {
			return errors.New("WithFilePath path is not a Dir")
		}
		o.filePath = path
		return nil
	}
}

func WithSize(size int) Option {
	return func(o *qrOptions) error {
		if size < 0 || size > 500 {
			return errors.New("WithWidth size should between 0 and 500")
		}
		o.size = size
		return nil
	}
}

// GenerateQR 生成二维码，以图片的形式输出
// 可输入需要生成的路径，默认放入默认地址
// 反馈图片生成完整 URL
func GenerateQR(message string, opts ...Option) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	opt := &qrOptions{
		filePath: path.Join(dir, "assets"),
		fileName: uuid.New(),
		size:     250,
	}
	for _, o := range opts {
		if err := o(opt); err != nil {
			return "", err
		}
	}
	opt.completeURL = path.Join(opt.filePath, opt.fileName+".png")
	return opt.completeURL, qrcode.WriteFile(message, qrcode.Medium, opt.size, opt.completeURL)
}

// 将该路径下的二维码图片解析，反馈 Matrix 文件对象
func Decode(path string) (*qrcodedecode.Matrix, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return qrcodedecode.Decode(f)
}
