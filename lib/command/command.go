package command

import (
	"errors"
	"os/exec"
)

func RunCommand(commands ...string) (string, error) {
	if len(commands) < 1 {
		return "", errors.New("commands is nil")
	}
	cmd := exec.Command(commands[0], commands[:1]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	outresult, errResult := "", ""

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)
			if n > 0 {
				errResult += string(buf[:n])
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
				outresult += string(buf[:n])
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
