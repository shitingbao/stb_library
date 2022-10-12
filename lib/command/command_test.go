package command

import (
	"log"
	"os/exec"
	"testing"
)

func CommandTest(t *testing.T) {
	cmd := exec.Command("sc", "create", "your server")
	out, err := RunCommand(*cmd)
	log.Println(out, err)
}
