package rtc

import (
	"bytes"
	cr "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// doc address https://help.aliyun.com/document_detail/159037.html

var (
	appID  = "appID"
	appKey = "appKey"

	defaltTimestamp = time.Now().Add(time.Hour).Unix() // 默认 48 小时过期时间
)

// timestamp 过期时间，时间戳
func CreateRtcToken(channelID, user string, timestamp ...int64) (token string, err error) {
	tokenTimestamp := defaltTimestamp
	if len(timestamp) > 0 {
		tokenTimestamp = timestamp[0]
	}
	var b bytes.Buffer
	b.WriteString(appID)
	b.WriteString(appKey)
	b.WriteString(channelID)
	b.WriteString(createUserID(channelID, user))
	b.WriteString(getNonce())
	b.WriteString(fmt.Sprint(tokenTimestamp))

	h := sha256.New()
	if _, err = h.Write(b.Bytes()); err != nil {
		return "", err
	}

	s := h.Sum(nil)
	token = hex.EncodeToString(s)
	return
}

// 以 AK 开头，详情看文档
func getNonce() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("AK-%v", uuid)
}

func createUserID(channelID, user string) string {
	var b bytes.Buffer
	b.WriteString(channelID)
	b.WriteString("/")
	b.WriteString(user)

	h := sha256.New()
	// if _, err := h.Write([]byte(b.String())); err != nil {
	if _, err := h.Write(b.Bytes()); err != nil {
		return buildRandom(16)
	}

	s := h.Sum(nil)
	uid := hex.EncodeToString(s)
	return uid[:16]
}

func buildRandom(length int) string {
	if length <= 0 {
		return ""
	}

	b := make([]byte, length/2+1)
	_, _ = cr.Read(b)
	s := hex.EncodeToString(b)
	return s[:length]
}
