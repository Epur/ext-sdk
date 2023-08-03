package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
)

type Rsa struct {
	privateKey    string
	publicKey     string
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
}

func NewRsa(publicKey, privateKey string) *Rsa {
	rsaObj := &Rsa{
		privateKey: privateKey,
		publicKey:  publicKey,
	}

	rsaObj.init()

	return rsaObj
}

func (this *Rsa) init() {
	keyBytes, err := base64.StdEncoding.DecodeString(this.privateKey)
	if err != nil {
		panic(err.Error())
		//fmt.Println("DecodeString:", err)
		//return nil, err
	}
	privkey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err != nil {
		panic(err.Error())
		//fmt.Println("ParsePKCS8PrivateKey", err)
		//return nil, err
	}

	this.rsaPrivateKey = privkey
	llpub, err := base64.StdEncoding.DecodeString(this.publicKey)
	if err != nil {
		panic(err.Error())
		//fmt.Println("DecodeString")
		//return nil, err
	}

	llpubkey, err := x509.ParsePKIXPublicKey(llpub)
	if err != nil {
		panic(err.Error())
	}
	this.rsaPublicKey = llpubkey.(*rsa.PublicKey)
}

/**
 * 加密
 */
func (this *Rsa) Encrypt(data []byte) ([]byte, error) {
	blockLength := this.rsaPublicKey.N.BitLen()/8 - 11
	if len(data) <= blockLength {
		return rsa.EncryptPKCS1v15(rand.Reader, this.rsaPublicKey, []byte(data))
	}

	buffer := bytes.NewBufferString("")

	pages := len(data) / blockLength

	for index := 0; index <= pages; index++ {
		start := index * blockLength
		end := (index + 1) * blockLength
		if index == pages {
			if start == len(data) {
				continue
			}
			end = len(data)
		}

		chunk, err := rsa.EncryptPKCS1v15(rand.Reader, this.rsaPublicKey, data[start:end])
		if err != nil {
			return nil, err
		}
		buffer.Write(chunk)
	}
	return buffer.Bytes(), nil
}

/**
 * 解密
 */
func (this *Rsa) Decrypt(secretData []byte) ([]byte, error) {

	blockLength := this.rsaPublicKey.N.BitLen() / 8
	if len(secretData) <= blockLength {
		return rsa.DecryptPKCS1v15(rand.Reader, this.rsaPrivateKey, secretData)
	}

	buffer := bytes.NewBufferString("")

	pages := len(secretData) / blockLength
	for index := 0; index <= pages; index++ {
		start := index * blockLength
		end := (index + 1) * blockLength
		if index == pages {
			if start == len(secretData) {
				continue
			}
			end = len(secretData)
		}

		chunk, err := rsa.DecryptPKCS1v15(rand.Reader, this.rsaPrivateKey, secretData[start:end])
		if err != nil {
			return nil, err
		}
		buffer.Write(chunk)
	}
	return buffer.Bytes(), nil
}

/**
 * 签名
 */
func (this *Rsa) Sign(data []byte, algorithmSign crypto.Hash) ([]byte, error) {
	hash := algorithmSign.New()
	hash.Write(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, this.rsaPrivateKey, algorithmSign, hash.Sum(nil))
	if err != nil {
		return nil, err
	}
	return sign, err
}

/**
 * 验签
 */
func (this *Rsa) Verify(data []byte, sign []byte, algorithmSign crypto.Hash) bool {
	h := algorithmSign.New()
	h.Write(data)
	return rsa.VerifyPKCS1v15(this.rsaPublicKey, algorithmSign, h.Sum(nil), sign) == nil
}

/**
 * 生成pkcs1格式公钥私钥
 */
func (this *Rsa) CreateKeys(keyLength int) (privateKey, publicKey string) {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return
	}

	privateKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsaPrivateKey),
	}))

	derPkix, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		return
	}

	publicKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}))
	return
}

/**
 * 生成pkcs8格式公钥私钥
 */
func (this *Rsa) CreatePkcs8Keys(keyLength int) (privateKey, publicKey string) {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return
	}

	privateKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: this.MarshalPKCS8PrivateKey(rsaPrivateKey),
	}))

	derPkix, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		return
	}

	publicKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}))
	return
}

func (this *Rsa) MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)
	k, _ := asn1.Marshal(info)
	return k
}
