package base

import (
	"archive/zip"
	"context"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/pborman/uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 提取同级目录的 zip 文件，解压并提取所有内容
func Load() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	fileDirs, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, fls := range fileDirs {
		if fls.IsDir() {
			continue
		}
		if strings.HasSuffix(fls.Name(), "zip") {
			load(fls.Name())
		}
	}
}

func load(basePath string) {
	baseName := getBaseFileName(basePath)
	if err := UnZip("./", basePath); err != nil {
		return
	}
	h, err := NewHub(baseName)
	if err != nil {
		log.Println(err)
		return
	}
	defer h.Close()
	go func() {
		h.ReadAllFile(baseName)
		close(h.FileChan)
	}()
	h.Write()
}

type hub struct {
	Ctx      context.Context
	FileChan chan string
	File     *os.File
}

func NewHub(fPath string) (*hub, error) {
	w, err := os.OpenFile((path.Join(fPath, uuid.New()+".txt")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return &hub{}, err
	}
	return &hub{
		Ctx:      context.Background(),
		FileChan: make(chan string, 5),
		File:     w,
	}, nil
}

func (h *hub) Close() {
	if h.File != nil {
		h.File.Close()
	}
}

func (h *hub) Write() {
	num := 1
	for {
		select {
		case f, ok := <-h.FileChan:
			if !ok {
				log.Println("解析完毕")
				return
			}
			b, err := os.ReadFile(f)
			if err != nil {
				log.Println("ReadFile:", err)
				continue
			}
			h.File.Write(b)
			log.Println("wirte:", num)
			num++
		case <-h.Ctx.Done():
			return
		}
	}
}

// 解压 src 目录的 zip 文件到对应 dst 地址
func UnZip(dst, src string) (err error) {
	if dst != "" {
		os.MkdirAll(dst, os.ModePerm)
	}
	zr, err := zip.OpenReader(src)
	if err != nil {
		return
	}
	defer zr.Close()
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}

		sourceFile, err := file.Open()
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		// 创建要写出的文件对应的 Write
		dataFile, err := os.Create(path)
		if err != nil {
			return err
		}
		defer dataFile.Close()

		typeFile, err := file.Open() // 注意每个流读取完就关闭了，不能复用，得用两个流，判断文件编码格式
		if err != nil {
			return err
		}
		b, err := io.ReadAll(typeFile)
		if err != nil {
			log.Println("ReadAll:", err)
			return err
		}
		if utf8.Valid(b) {
			if _, err := io.Copy(dataFile, sourceFile); err != nil {
				log.Println("Copy:", err)
				return err
			}
			continue
		}
		fr := transform.NewReader(sourceFile, simplifiedchinese.GBK.NewDecoder())
		if _, err := io.Copy(dataFile, fr); err != nil {
			log.Println("NewReader Copy:", err)
			return err
		}
	}
	return nil
}

// 获取到除文件后缀的完整路径
func getBaseFileName(fName string) string {
	return path.Join(path.Dir(fName), strings.TrimSuffix(path.Base(fName), path.Ext(fName)))
}

func (h *hub) ReadAllFile(basePath string) {
	fdir, err := os.ReadDir(basePath)
	if err != nil {
		log.Println("ReadDir:", err)
		return
	}
	for _, fl := range fdir {
		if fl.IsDir() {
			h.ReadAllFile(path.Join(basePath, fl.Name()))
			continue
		}
		if strings.HasSuffix(fl.Name(), "java") || strings.HasSuffix(fl.Name(), "dll") || strings.HasSuffix(fl.Name(), "php") {
			h.FileChan <- path.Join(basePath, fl.Name())
		}
	}
}
