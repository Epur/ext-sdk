package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

type Sign struct {
	privateKey  *rsa.PrivateKey
	llpublickey *rsa.PublicKey
}

// 获取密钥（连连方式):20240402，直接从pem文件粘贴到privateKey及llpublickey中，无需去掉换行

func NewSign(privateKey string, llpublickey string) (*Sign, error) {

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, fmt.Errorf("failed to decode private key")
	}
	privkey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey", err)
		return nil, err
	}

	blocks, _ := pem.Decode([]byte(llpublickey))
	if block == nil {
		return nil, fmt.Errorf("failed to decode public key")
	}

	llpubkey, err := x509.ParsePKIXPublicKey(blocks.Bytes)
	if err != nil {
		fmt.Println("ParsePKIXPublicKey", err)
		return nil, err
	}

	return &Sign{
		privateKey:  privkey.(*rsa.PrivateKey),
		llpublickey: llpubkey.(*rsa.PublicKey),
	}, nil
}

// 签名(修改成RSA(MD5(完整JSON报文))

func (s *Sign) RsaSignWithMd5(content []byte) (sign string, err error) {

	h := sha256.New()
	h.Write(content)
	d := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.MD5, d)
	if err != nil {
		fmt.Println("SignPKCS1v15:", err)
		return "", err
	}

	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

// 验签(连连银通验签)

func (s *Sign) RsaVerifySignWithMd5(originalData []byte, signData string) error {
	signatureByte, _ := hex.DecodeString(signData)

	h2 := md5.New()
	h2.Write(originalData)
	hashed2 := h2.Sum(nil)

	return rsa.VerifyPKCS1v15(s.llpublickey, crypto.MD5, hashed2, signatureByte)
}

// 连连国际 签名
func (s *Sign) RsaSignWithSHA256(content []byte) (sign string, err error) {

	h := sha256.New()
	h.Write(content)
	d := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, d)
	if err != nil {
		fmt.Println("SignPKCS1v15:", err)
		return "", err
	}

	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

//连连国际使用验签

func (s *Sign) RsaVerifySignWithSHA256(originalData []byte, signature string) error {
	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return err
	}

	hashed := sha256.Sum256(originalData)

	return rsa.VerifyPKCS1v15(s.llpublickey, crypto.SHA256, hashed[:], sign)
}

//连连国际私钥解密敏感字段

func (s *Sign) RsaDecrypt(encrypted string) (string, error) {
	sign, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return "", err
	}

	ss, err := rsa.DecryptPKCS1v15(rand.Reader, s.privateKey, sign)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return "", err
	}
	return string(ss), nil
}

//连连国际公钥加密敏感字段

func (s *Sign) RsaEncrypt(source string) (string, error) {

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, s.llpublickey, []byte(source))
	if err != nil {
		return "", fmt.Errorf("encryption error: %w", err)
	}

	encrypted := base64.StdEncoding.EncodeToString(ciphertext)
	return encrypted, nil
}

// 连连国际 证件号码hash
func (s *Sign) Sha256Hex(content []byte) (sign string) {

	h := sha256.New()
	h.Write(content)
	out := hex.EncodeToString(h.Sum(nil))
	return out
}
