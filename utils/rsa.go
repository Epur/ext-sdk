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
	"fmt"
)

type Sign struct {
	privateKey  *rsa.PrivateKey
	llpublickey *rsa.PublicKey
}

func NewSign(privateKey string, llpublickey string) (*Sign, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return nil, err
	}
	privkey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey", err)
		return nil, err
	}
	llpub, err := base64.StdEncoding.DecodeString(llpublickey)
	if err != nil {
		fmt.Println("DecodeString")
		return nil, err
	}

	llpubkey, err := x509.ParsePKIXPublicKey(llpub)
	if err != nil {
		fmt.Println("ParsePKIXPublicKey", err)
		return nil, err
	}

	return &Sign{
		privateKey:  privkey.(*rsa.PrivateKey),
		llpublickey: llpubkey.(*rsa.PublicKey),
	}, nil
}

// 签名
func (s *Sign) RsaSignWithMd5(content []byte) (sign string, err error) {

	//h := crypto.Hash.New(crypto.SHA256)
	//h.Write(content)
	//hashed := h.Sum(nil)
	//data := hex.EncodeToString(hashed)
	//
	//h2 := crypto.Hash.New(crypto.SHA256)
	//h2.Write([]byte(data))
	//hashed2 := h2.Sum(nil)

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

// 验签
func (s *Sign) RsaVerifySignWithMd5(originalData []byte, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return err
	}

	h := crypto.Hash.New(crypto.MD5)
	h.Write(originalData)
	hashed := h.Sum(nil)
	data := hex.EncodeToString(hashed)

	h2 := crypto.Hash.New(crypto.MD5)
	h2.Write([]byte(data))
	hashed2 := h2.Sum(nil)

	return rsa.VerifyPKCS1v15(&s.privateKey.PublicKey, crypto.MD5, hashed2, sign)
}

func (s *Sign) RsaVerifyLianLianSignWithMd5(originalData []byte, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return err
	}

	hash := md5.New()
	hash.Write(originalData)
	return rsa.VerifyPKCS1v15(s.llpublickey, crypto.MD5, hash.Sum(nil), sign)
}
