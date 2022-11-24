package base

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"stb-library/lib/command"
	"strconv"
	"strings"
)

// GetAllDirFile 便利所有文件内文件，反馈所有文件路径,isAbsolute代表是否反馈完整路径，或者只反馈文件名称
// isAbsolute为true反馈当前开始的完整路径，[file/aa/aa.txt file/aa/bb/bb.txt]，为false只反馈文件名，[aa.txt bb.txt]
func GetAllDirFile(url string, isAbsolute bool) ([]string, error) {
	fList := []string{}
	ft, err := os.ReadDir(url)
	if err != nil {
		return fList, err
	}
	for _, v := range ft {
		if v.IsDir() {
			ft, err := GetAllDirFile(path.Join(url, v.Name()), isAbsolute)
			if err != nil {
				return fList, err
			}
			fList = append(fList, ft...)
			continue
		}
		if isAbsolute {
			fList = append(fList, path.Join(url, v.Name()))
		} else {
			fList = append(fList, v.Name())
		}
	}
	return fList, nil
}

// 文件后缀操作
func fileNameOpera() {
	fullFilename := "/Users/itfanr/Documents/test.txt"
	log.Println(path.Dir(fullFilename)) //获取当前目录，"/Users/itfanr/Documents"
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀(test.txt)
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)

	var fileSuffix string
	fileSuffix = path.Ext(fullFilename) //获取文件后缀(.txt)
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名(test)
	fmt.Println("filenameOnly =", filenameOnly)
}

// GetFileDiskSize 获取实际文件磁盘占用
func GetFileDiskSize(url string) int64 {
	cmd := exec.Command("du", "-Lb", url)
	out, err := command.RunCommand(*cmd)
	if err != nil {
		return 0
	}
	log.Println("out:", out)
	status := "[0-9]*"
	matchMe := regexp.MustCompile(status)
	log.Println("find:", matchMe.FindString(out))
	size, _ := strconv.ParseInt(matchMe.FindString(out), 10, 64)
	return size
}

// 更新某一行数据
func updateFileLine(fileName string) error {
	fi, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var pos int64 = 0
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if strings.HasPrefix(line, "lastOKBoot") {
			fi.WriteAt([]byte("lastOKBoot=bbbbbbbb\n"), pos)
			log.Println("ok:", line, pos)
			break
		}
		pos += int64(len(line))
		log.Println("line:", line, len(line))
	}
	return nil
}
