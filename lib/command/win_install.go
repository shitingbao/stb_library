package command

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	defaultDir = "C:\\win-agent"
)

func Install() {
	checkDir()
	log.Println(runCommand("sc", "stop", "win-agent"))
	log.Println(runCommand("sc", "delete", "win-agent"))
	log.Println(runCommand("sc", "create", "win-agent", `binpath="C:\win-agent\win-agent.exe"`, "start=auto"))
	log.Println(runCommand("sc", "start", "win-agent"))
}

func runCommand(commands ...string) (string, error) {
	if len(commands) < 1 {
		return "", errors.New("commands is nil")
	}
	log.Println(commands[0], ":", commands[1:])
	cmd := exec.Command(commands[0], commands[1:]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	outresult, errResult := "", ""

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)
			if err != nil {
				errResult = err.Error()
				return
			}

			if n > 0 {
				errResult += string(gbkToUtf8(buf[:n]))
			}
			if n == 0 {
				break
			}
			if err != nil {
				errResult = err.Error()
				return
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stdout.Read(buf)

			if n == 0 {
				break
			}

			if n > 0 {
				outresult += string(gbkToUtf8(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				return
			}
		}
	}()

	cmd.Wait()
	if len(errResult) > 0 {
		return "", errors.New(errResult)
	}
	return outresult, nil
}

func gbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, _ := ioutil.ReadAll(reader)
	return d
}

func checkDir() {
	// Install()
	dir, _ := os.Getwd()
	log.Println(dir)
	if dir != defaultDir {
		log.Println("Mkdir:", os.Mkdir(defaultDir, 0777))
		copyFile(dir)
	}
}

func copyFile(wdDir string) {
	// Open original file
	original, err := os.Open(path.Join(wdDir, "win-agent.exe"))
	if err != nil {
		log.Println("os.Open:", err)
		return
	}
	defer original.Close()

	// Create new file
	newFile, err := os.Create("C:\\win-agent\\win-agent.exe")
	if err != nil {
		log.Println("os.Create:", err)
		return
	}
	defer newFile.Close()

	//This will copy
	bytesWritten, err := io.Copy(newFile, original)
	if err != nil {
		log.Println("os.Copy:", err)
		return
	}
	fmt.Printf("Bytes Written: %d\n", bytesWritten)
}
