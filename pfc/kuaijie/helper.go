package kuaijie

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

type Key struct {
	PrivateKey []byte // 私钥
	PublicKey  []byte // 公钥
}

// new
func KeyNew(privateKey, publicKey string) *Key {
	return &Key{
		PrivateKey: []byte(privateKey),
		PublicKey:  []byte(publicKey),
	}
}

// 报文发送方签名
// 1、 将报文体“body”的值（不含“body”这个 key）作为待签名字符串，使用报文发送方的
// 私钥对其进行 SHA256withRSA 签名运算，得到签名字节流
// 2、 对签名字节流进行 Base64 编码，得到数字签名字符串
// 3、 将数字签名字符串写入数字签名域“sign”
// 数字签名(sign) = Base64.encode( SHA256withRSA.sign( #{body} ) )
func (p *Key) Sign(data string) (string, error) {

	b, err := p.RsaWithSHA256([]byte(data), p.PrivateKey)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(b)

	return sign, nil
}

// SHA256withRSA
func (p *Key) RsaWithSHA256(origData, prvKey []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(origData)
	block, _ := pem.Decode(prvKey)
	if block == nil {
		fmt.Printf("prvKey\n%s", prvKey)
		//return nil, errors.New("private key error")
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature, nil
}

// 验签
func (p *Key) Verify(data, signData string) bool {

	return p.RsaVerySignWithSha256([]byte(data), []byte(signData), []byte(p.PublicKey))
}

// 验证
func (p *Key) RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 公钥加密
func (p *Key) RsaEncrypt(data []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(p.PublicKey))
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func (p *Key) RsaDecrypt(ciphertext []byte) []byte {
	//获取私钥
	block, _ := pem.Decode([]byte(p.PrivateKey))
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}
