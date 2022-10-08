package iperf

import (
	"os/exec"
	"stb-library/lib/command"
)

// NetInfo ip 命令反馈信息结构
type NetInfo struct {
	Ifindex   int           `json:"ifindex"`
	Ifname    string        `json:"ifname"`
	Flags     []string      `json:"flags"`
	Mtu       int           `json:"mtu"`
	Qdisc     string        `json:"qdisc"`
	Operstate string        `json:"operstate"`
	Group     string        `json:"group"`
	Txqlen    int           `json:"txqlen"`
	AddrInfo  []NetAddrInfo `json:"addr_info"`
}

type NetAddrInfo struct {
	Family            string `json:"family"`
	Local             string `json:"local"`
	Prefixlen         int    `json:"prefixlen"`
	Scope             string `json:"scope"`
	Label             string `json:"label"`
	ValidLifeTime     int64  `json:"valid_life_time"`
	PreferredLifeTime int64  `json:"preferred_life_time"`
}

// GetLocalAreaIP json 形式反馈 ipv4 网络信息
func GetLocalAreaIP() (string, error) {
	cmd := exec.Command("ip", "-4", "-j", "addr")
	return command.RunCommand(*cmd)
}
