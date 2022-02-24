package ffmpeg

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

var (
	defaultOrder      = "ffmpeg"   // 默认执行目标工具
	defaultInPath     = "./"       // 默认读取地址
	defaultOutPath    = "outFiles" // 输出文件地址
	defaultTargetType = "mp4"      // 目标文件后缀
	defaultSourceType = "wav"      // 源文件后缀

)

type Ffmpeg interface {
	MovToMp4() error
	// DefaultTransform 默认转化，传入需要转化的文件路径，输出转化后的文名称
	DefaultTransform(string) (string, error)
}

type (
	Option func(*options)
)

type options struct {
	Order      string
	InPath     string
	OutPath    string
	TargetType string
	SourceType string
}

type ffmpeg struct {
	order      string
	inPath     string
	outPath    string
	targetType string
	sourceType string
}

func NewFfmpeg(opts ...Option) Ffmpeg {
	f := &options{
		Order:      defaultOrder,
		InPath:     defaultInPath,
		OutPath:    defaultOutPath,
		TargetType: defaultTargetType,
		SourceType: defaultSourceType,
	}
	for _, o := range opts {
		o(f)
	}
	return &ffmpeg{
		order:      f.Order,
		inPath:     f.InPath,
		outPath:    f.OutPath,
		targetType: f.TargetType,
		sourceType: f.SourceType,
	}
}

// MovToMp4
func (f *ffmpeg) MovToMp4() error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	fds, err := os.ReadDir(f.inPath)
	if err != nil {
		return err
	}
	wholePath := path.Join(rootPath, f.outPath+strconv.Itoa(int(time.Now().Unix())))
	if err := os.MkdirAll(wholePath, os.ModePerm); err != nil {
		return err
	}
	for _, v := range fds {
		l := strings.Split(v.Name(), ".") // 防止名字中有多个英文 “ . ”
		if l[len(l)-1] != f.sourceType {
			continue
		}
		fileName := strings.Join(l[:len(l)-1], ".")
		fileName += "." + f.targetType
		wholeOutPath := path.Join(wholePath, fileName)

		ecPath := path.Join(rootPath, f.order)
		//不要写整条命令！！！
		//不要写整条命令！！！
		//不要写整条命令！！！
		cmd := exec.Command(ecPath, "-i", path.Join(rootPath, v.Name()), "-qscale", "0", wholeOutPath) //ffmpeg -i input.mov -qscale 0 output.mp4
		if err := cmd.Run(); err != nil {
			return err
		} else {
			log.Println("success change ", v.Name())
		}
	}
	return nil
}

func (f *ffmpeg) DefaultTransform(filePath string) (string, error) {
	fullPath := fileExtension(filePath)
	fullFileName := fullPath + f.targetType

	cmd := exec.Command(f.order, "-i", filePath, "-qscale", "10", fullFileName) //ffmpeg -i input.mov -qscale 0 output.mp4
	if err := cmd.Run(); err != nil {
		return "", errors.New("转化文件出错-" + err.Error())
	}
	filename := path.Base(fullFileName) //获取文件名带后缀(test.txt)
	return filename, nil
}

// 去除文件后缀
func fileExtension(fullFilename string) string {
	basePath := path.Dir(fullFilename)            //获取当前目录，"/Users/itfanr/Documents"
	filenameWithSuffix := path.Base(fullFilename) //获取文件名带后缀(test.txt)
	fileSuffix := path.Ext(fullFilename)          //获取文件后缀(.txt)

	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名(test)
	return path.Join(basePath, filenameOnly)
}

// WithFfmpegOrder ffmpeg所在的目录
func WithFfmpegOrder(od string) Option {
	return func(o *options) {
		o.Order = path.Join(od, "ffmpeg")
	}
}

func WithFfmpegInpath(inpath string) Option {
	return func(o *options) {
		o.InPath = inpath
	}
}

func WithFfmpegOutpath(outpath string) Option {
	return func(o *options) {
		o.OutPath = outpath
	}
}

func WithFfmpegTargetType(targetType string) Option {
	return func(o *options) {
		o.TargetType = targetType
	}
}

func WithFfmpegSourceType(sourceType string) Option {
	return func(o *options) {
		o.SourceType = sourceType
	}
}

// 使用例子
func FfmpegTestLoad() error {
	fmg := NewFfmpeg(
		WithFfmpegSourceType("wav"),
		WithFfmpegTargetType("mp4"),
	)
	return fmg.MovToMp4()
}
