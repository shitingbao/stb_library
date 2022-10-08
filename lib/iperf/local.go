package iperf

import (
	"os/exec"
	"stb-library/lib/command"
)

// GetLocalAreaIP json 形式反馈 ipv4 网络信息
func GetLocalAreaIP() (string, error) {
	cmd := exec.Command("ip", "-4", "-j", "addr")
	return command.RunCommand(*cmd)
}
