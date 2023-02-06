package snsq

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 接受 nsq 请求过来的参数
type AuthReq struct {
	RemoteIP   string `form:"remote_ip"`
	TLS        bool   `form:"tls"`
	Secret     string `form:"secret"`
	CommonName string `form:"common_name"`
}

type Authorization struct {
	Topic       string   `json:"topic"`       // 内容需要是符合正则表达式的形势，别问为啥，官方定的
	Channels    []string `json:"channels"`    // 内容需要是符合正则表达式的形势
	Permissions []string `json:"permissions"` // 订阅或者发布 subscribe ｜ publish

}

// 反馈给 nsq 的结构，参照官方的例子结构，https://nsq.io/components/nsqd.html#auth
type AuthResp struct {
	TTL            int             `json:"ttl"`            // 过期时间
	Authorizations []Authorization `json:"authorizations"` // 允许（或者不允许）哪些主体和通道
	Identity       string          `json:"identity"`       // 身份
	IdentityURL    string          `json:"identity_url"`   // 身份地址，就是提供验证的接口地址，例如这里本地的就是就是 127.0.0.1:1325
	Expires        time.Time
}

func Example() {
	g := aPIRoute()
	g.Run(":1325")
}

func aPIRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/auth", auth) // 鉴权的接口，nsq 使用了 auth-http-address 参数后，连接到 nsq 会自动请求这个接口
	r.GET("/ping", ping) // 还有其他接口自定义，详情看官方文档，不需要就不用
	return r
}

func ping(c *gin.Context) { c.JSON(http.StatusOK, "pong") }

// 请求到该接口后，自定义权限逻辑
func auth(c *gin.Context) {
	req := &AuthReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		log.Println(err)
		return
	}
	if req.Secret == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "NOT_AUTHORIZED",
		})
		return
	}
	da := Authorization{
		Topic:    ".*",           // 这里设置全部可以，如果指定哪些 Topic 可以，在这里反馈即可，比如指定 “test”，那就该连接只能操作 test 的 topic
		Channels: []string{".*"}, // 这个同上
		Permissions: []string{
			"subscribe",
			"publish"},
	}
	res := AuthResp{
		TTL:            60,
		Identity:       "your name",
		IdentityURL:    "http://127.0.0.1:1325",
		Authorizations: []Authorization{da},
	}
	c.JSON(http.StatusOK, res)
}
