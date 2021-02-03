package RSA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"testing"
)

//var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIICXgIBAAKBgQC0ZRaS/mQCq5LzD5tsNnwO83eXq7PVa99aIuB8MEUseP1pGm28hAAIcRushzEqPl4Lnvg8ZzJAxTTo8bBk8WxinsAXx/142tVXQYcTRJbHcuZx8ukfXn1Rhfo8ZN9wUzoAlnCgQ3Gsba8zpomKpHLF7APBE+KnwUcknxepYBCI7wIDAQABAoGBAITWB3h4kSaSNyR6sqVNva64w7DRBBy9UXwrQIjSdq2X7mrtxi7SOL/+ojU6XA7SXChMiFjaWNvvG9YI2y6JxJfhT1+WQw/4wPE44wTLZ19I4ko/gQPxaZqr5YI2RhmGF5NycgYXohKmD9+GuZb/HDkTlxwEy9s4d3L6tIO3Fd0BAkEA5Sskp4NU65qunWWNGWHeNaNsh9dY3y4mx06ETkXtHAsSDqPxXIELVihMwKXqaFWcCRJR4l0p6TQ3whfrq3RsgQJBAMmEDgS0GTAoxB7VOmeYOONe1UNfdeJ4+8pIJiw7CWhlBAZo7N+hErJlK+YED3XT996PeBLzeVrDNCWrv9Ra/W8CQQDlJcWqBmVUjMALPG7hMX4azkWIcyk3SJOX+QfMqJkV0HAG2aFPKO2oAfX1MDH6j5fNblcYbKWqvP8AAiExuOGBAkA6UUMVTWF0tHd9TMvoKv9bnZguNTQSZFzJv3N8nWEtmv49NKRIW37jra+0kzw+Jye2euKO6XMXNyKiPD/5npN1AkEAz/Os5CEHaPVKW6myQyVPnxZJSVLub3vN5kMrg1ffpIj3QHMY0M4HNJEK6wMJmwLEkSxWjSbZiXzYXik3yY+Vbg==
//-----END RSA PRIVATE KEY-----
//`)
//
//// 公钥: 根据私钥生成
////openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
//var publicKey = []byte(`
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0ZRaS/mQCq5LzD5tsNnwO83eXq7PVa99aIuB8MEUseP1pGm28hAAIcRushzEqPl4Lnvg8ZzJAxTTo8bBk8WxinsAXx/142tVXQYcTRJbHcuZx8ukfXn1Rhfo8ZN9wUzoAlnCgQ3Gsba8zpomKpHLF7APBE+KnwUcknxepYBCI7wIDAQAB
//-----END PUBLIC KEY-----
//`)


func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	publicKey,_ := ioutil.ReadFile("./RsaPublicKey.txt")
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	privateKey,_ := ioutil.ReadFile("./RsaPrivateKey.txt")
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func TestRsa(t *testing.T){
	secret_key,_ := RsaEncrypt([]byte("周婉秋"))
	t.Log(secret_key)
	txt,_ := RsaDecrypt(secret_key)
	t.Log(string(txt))
}