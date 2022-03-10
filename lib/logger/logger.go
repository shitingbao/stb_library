package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// NewLogrusLog 新建一个 logrus 对象并反馈
// 传入日志生成地址，默认当前目录
// 初始化日志文件，如果已经初始化则跳过,并获取配置参数
// 重定向日志输出的文件
func NewLogrusLog(defaultDirPath string) *logrus.Logger {
	flog, err := os.Create(filepath.Join(defaultDirPath, "log.txt"))
	if err != nil {
		panic(fmt.Errorf("error opening log.txt file: %v", err))
	}

	ferr, err := os.Create(filepath.Join(defaultDirPath, "err.txt"))
	if err != nil {
		panic(fmt.Errorf("error opening err.txt file: %v", err))
	}
	redirectStderr(ferr)
	log.SetFlags(log.LstdFlags | log.Llongfile)

	lg := logrus.New()
	lg.SetOutput(io.MultiWriter(os.Stdout, flog))
	lvl, err := logrus.ParseLevel("debug")
	if err != nil {
		panic(err)
	}
	lg.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	lg.SetLevel(lvl)
	lg.WithFields(logrus.Fields{"set-level": lvl.String()}).Info("initlog")
	return lg
}
