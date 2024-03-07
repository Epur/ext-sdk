package cmbchina

import (
	"bytes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/Epur/ext-sdk/utils"
	cryptobin_sm2 "github.com/deatil/go-cryptobin/cryptobin/sm2"
	"github.com/tjfoc/gmsm/sm4"
)

type SM4 struct {
	PrivateKey    string // 客户私钥
	PublicKey     string // 客户公钥
	BankPublicKey string // 银行公钥
	SymKey        string // 对称密钥

	UserId string
}

func SM4New(PrivateKey, PublicKey, BankPublicKey, SymKey string, UserId string) *SM4 {
	s := &SM4{
		PrivateKey:    PrivateKey,
		PublicKey:     PublicKey,
		BankPublicKey: BankPublicKey,
		SymKey:        SymKey,
		UserId:        utils.ZeroFillByStr(UserId, 16, false),
	}
	return s
}

func (p *SM4) Sign(data string) (string, error) {

	sm2keyBytes, _ := base64.StdEncoding.DecodeString(p.PrivateKey)

	sm2Sign := cryptobin_sm2.NewSM2().
		FromPrivateKeyBytes(sm2keyBytes).
		FromString(data).
		WithUID([]byte(p.UserId)).
		SignBytes().
		ToBase64String()
	return sm2Sign, nil
}

func (p *SM4) Verify(sm2signdata string, data string) bool {

	sm2keyBytes, _ := base64.StdEncoding.DecodeString(p.BankPublicKey)

	sssss := cryptobin_sm2.NewSM2().FromBase64String(sm2signdata).
		FromPublicKeyString(hex.EncodeToString(sm2keyBytes)).
		WithUID([]byte(p.UserId)).
		VerifyBytes([]byte(data))

	fmt.Println(sssss.Error())

	return sssss.ToVerify()
}

func (p *SM4) CbcSm4Encrypt(data string) ([]byte, error) {

	ciphertxt, err := p.cbcSm4Encrypt([]byte(p.SymKey), []byte(p.UserId), []byte(data))
	if err != nil {
		return nil, err
	}
	return ciphertxt, nil
}

func (p *SM4) CbcSm4Decrypt(data string) ([]byte, error) {

	ciphertxt, err := p.cbcSm4Decrypt([]byte(p.SymKey), []byte(p.UserId), []byte(data))
	if err != nil {
		return nil, err
	}
	return ciphertxt, nil
}

func (p *SM4) cbcSm4Encrypt(key, iv, plainText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData := p.PKCS7Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

func (p *SM4) cbcSm4Decrypt(key, iv, cipherText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = p.PKCS7UnPadding(origData)
	return origData, nil
}

// pkcs5填充
func (p *SM4) pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func (p *SM4) pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// 使用PKCS7进行填充
func (p *SM4) PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (p *SM4) PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
