package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 对数据进行加密操作
func EncyptogRSA(src []byte, path string) (res []byte, err error) {
	//1.获取秘钥（从本地磁盘读取）
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	// 2、将得到的字符串解码
	block, _ := pem.Decode(b)

	// 使用X509将解码之后的数据 解析出来
	//x509.MarshalPKCS1PublicKey(block):解析之后无法用，所以采用以下方法：ParsePKIXPublicKey
	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes) //对应于生成秘钥的x509.MarshalPKIXPublicKey(&publicKey)
	//keyInit1,err:=x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return
	}
	//4.使用公钥加密数据
	pubKey := keyInit.(*rsa.PublicKey)
	res, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	return
}

// 对数据进行解密操作
func DecrptogRSA(src []byte, path string) (res []byte, err error) {
	//1.获取秘钥（从本地磁盘读取）
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	block, _ := pem.Decode(b)                                 //解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes) //还原数据
	res, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	return
}
