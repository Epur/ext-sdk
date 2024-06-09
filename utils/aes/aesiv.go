package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// aes解密 AES/IV/PKCS5
func DecryptAESWithIV(key []byte, iv []byte, encryptedText string) (string, error) {
	// 解码Base64加密文本
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// 创建密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建一个CBC模式的cipher.BlockMode
	if len(encryptedBytes)%aes.BlockSize != 0 {
		panic("blocksize must be multiple of decrypted bytes length")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密
	mode.CryptBlocks(encryptedBytes, encryptedBytes)

	// 去除填充
	encryptedBytes = pkcs5Unpad(encryptedBytes, aes.BlockSize)

	return string(encryptedBytes), nil
}

// pkcs5Unpad 移除PKCS7填充
func pkcs5Unpad(buf []byte, blockSize int) []byte {
	if blockSize <= 0 {
		panic("block size is not positive")
	}
	n := len(buf)
	if n%blockSize != 0 {
		panic("data not multiple of block size")
	}
	c := buf[n-1]
	if c > byte(blockSize) {
		panic("padding is not valid")
	}
	for i := 0; i < int(c); i++ {
		if buf[n-1-i] != c {
			panic("padding is not valid")
		}
	}
	return buf[:n-int(c)]
}
