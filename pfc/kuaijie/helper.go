package kuaijie

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/logger"
	ursa "github.com/Epur/ext-sdk/utils/rsa"
)

type Key struct {
	PrivateKey []byte // 私钥
	PublicKey  []byte // 公钥
	EncryptKey []byte // 对称加密密钥明文
	EncryptIV  []byte // 对称加密初始向量明文
}

// new
func KeyNew(privateKey, publicKey string) *Key {
	return &Key{
		PrivateKey: []byte(privateKey),
		PublicKey:  []byte(publicKey),
		//EncryptKey: []byte(encryptKey),
		//EncryptIV:  []byte(encryptIV),
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
		fmt.Println("ParsePKCS8PrivateKey err", err.Error())
		panic(err)
	}

	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err.Error())
		panic(err)
	}

	return signature, nil
}

// 验签
// 验签结果 = SHA256withRSA.verifySign( #{body}, Base64.decode( #{sign} ) )
func (p *Key) Verify(data, sign string) bool {
	//return true

	tmpData, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Printf("Error signData DecodeString: %s\n", err.Error())
		panic(err.Error())
	}
	//fmt.Println("sign:\n", sign)
	//fmt.Println("\n", string(tmpData))
	return p.RsaVerySignWithSha256([]byte(data), tmpData, p.PublicKey)
	//return p.RsaVerySignWithSha256([]byte(data), []byte(sign), p.PublicKey)

}

// 验证
func (p *Key) RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	//fmt.Printf("808080:\n%s\n%s", string(data), string(signData))
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("509 ERROR:", err.Error())
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		fmt.Println("15 ERROR:", err.Error())
		panic(err)
	}
	fmt.Println("RsaVerySignWithSha256 End True")
	return true
}

// 公钥加密
func (p *Key) RsaEncrypt(data []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(p.PublicKey))
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("ERROR:%s", err.Error())
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		fmt.Printf("ERROR:%s", err.Error())
		panic(err)
	}
	return ciphertext, nil
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

// 保护体签名
func (p *Key) SignProtected(data string) (string, error) {
	//logger.KuaijieLoger.Info("PrivateKey:\n", string(p.PrivateKey))
	//logger.KuaijieLoger.Info("PublicKey:\n", string(p.PublicKey))
	//logger.KuaijieLoger.Info("EncryptKey:\n", string(p.EncryptKey))

	b, err := AesEncryptECBPKCS5([]byte(data), p.EncryptKey)
	if err != nil {
		logger.KuaijieLoger.Error("ERROR:", err.Error())
		panic(err)
	}
	//b := uaes.AesEncryptECB([]byte(data), p.EncryptKey)
	sign := base64.StdEncoding.EncodeToString(b)

	return sign, nil
}

//func AesEncryptCBCPKCS5(origData []byte, key []byte) (encrypted []byte) {
//	// 分组秘钥
//	if len(origData) == 0 {
//		logger.KuaijieLoger.Error("ERROR:数据为空")
//		panic("加密数据为空")
//	}
//	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		logger.KuaijieLoger.Error("ERROR:", err)
//		panic(err.Error())
//	}
//	blockSize := block.BlockSize()                              // 获取秘钥块的长度
//	origData = uaes.PKCS5Padding(origData, blockSize)           // 补全码
//	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
//	encrypted = make([]byte, len(origData))                     // 创建数组
//	blockMode.CryptBlocks(encrypted, origData)                  // 加密
//
//	return encrypted
//}

// 公钥加密
func (p *Key) RsaECBEncrypt(data []byte) ([]byte, error) {
	//解密pem格式的公钥
	aaa := ursa.NewRsa(string(p.PublicKey), string(p.PrivateKey))
	b, err := aaa.Encrypt(data)
	if err != nil {
		logger.KuaijieLoger.Error("ERROR:", err.Error())
		panic(err.Error())
	}
	return b, nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// 引入golang官方 https://codereview.appspot.com/7860047/
type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func AesEncryptECBPKCS5(origData []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.KuaijieLoger.Error("ERROR:", err.Error())
		panic(err.Error())
	}
	if len(origData) == 0 {
		logger.KuaijieLoger.Error("ERROR:未检测到源数据")
		fmt.Println("plain content empty")
		return nil, errors.New("未检测到数据")
	}
	ecb1 := NewECBEncrypter(block)
	content := pKCS5Padding(origData, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb1.CryptBlocks(crypted, content)

	return crypted, nil
}
