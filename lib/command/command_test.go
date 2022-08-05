package command

import (
	"log"
	"testing"
)

func CommandTest(t *testing.T) {
	out, err := RunCommand("sc", "create", "your server")
	log.Println(out, err)
}
