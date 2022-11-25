package base

import (
	"log"
	"testing"
)

func TestOpenFileLine(t *testing.T) {
	OpenLastOk()
}

func TestCloseFileLine(t *testing.T) {
	CloseLastOk()
}

func TestGetStatus(t *testing.T) {
	log.Println(GetLastOkStatus("./hw.ini", "lastOKBoot"))
}
