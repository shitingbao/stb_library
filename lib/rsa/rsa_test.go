package rsa

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestRsa(t *testing.T) { //rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey := GenRsaKey()
	log.Println("这是私钥：", string(prvKey))
	log.Println("这是公钥：", string(pubKey))

	log.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	log.Println("-----对消息进行签名操作-----")
	signData := RsaSignWithSha256([]byte(data), prvKey)
	log.Println("消息的签名信息： ", hex.EncodeToString(signData))
	log.Println("\n对签名信息进行验证...")
	if RsaVerySignWithSha256([]byte(data), signData, pubKey) {
		log.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	log.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext := RsaEncrypt([]byte(data), pubKey)
	log.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := RsaDecrypt(ciphertext, prvKey)
	log.Println("私钥解密后的数据：", string(sourceData))
}
