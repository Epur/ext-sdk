/*
Copyright 2023 Steven Zhong

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package desede

import (
	"crypto/des"
)

// PKCS5Padding 填充明文分组到8字节的整数倍
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(ciphertext, padtext...)
}

// PKCS5Unpadding 去除填充数据
func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// TripleEcbDesEncrypt 使用desede/ECB/PKCS5Padding加密数据
func TripleEcbDesEncrypt(origData, key []byte) ([]byte, error) {
	// 复制24字节的密钥
	tkey := make([]byte, 24)
	copy(tkey, key)
	// 分割成三个子密钥
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]
	// 创建DES加密器
	block1, err := des.NewCipher(k1)
	if err != nil {
		return nil, err
	}
	block2, err := des.NewCipher(k2)
	if err != nil {
		return nil, err
	}
	block3, err := des.NewCipher(k3)
	if err != nil {
		return nil, err
	}
	// 填充明文分组
	bs := block1.BlockSize()
	origData = PKCS5Padding(origData, bs)
	// 分配加密后的数据空间
	crypted := make([]byte, len(origData))
	tmpData := make([]byte, len(origData))
	// 进行三重DES加密
	for i := 0; i < len(origData); i += bs {
		block1.Encrypt(tmpData[i:i+bs], origData[i:i+bs])
		block2.Decrypt(tmpData[i:i+bs], tmpData[i:i+bs])
		block3.Encrypt(crypted[i:i+bs], tmpData[i:i+bs])
	}
	return crypted, nil
}

// TripleEcbDesDecrypt 使用desede/ECB/PKCS5Padding解密数据
func TripleEcbDesDecrypt(crypted, key []byte) ([]byte, error) {
	// 复制24字节的密钥
	tkey := make([]byte, 24)
	copy(tkey, key)
	// 分割成三个子密钥
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]
	// 创建DES解密器
	block1, err := des.NewCipher(k1)
	if err != nil {
		return nil, err
	}
	block2, err := des.NewCipher(k2)
	if err != nil {
		return nil, err
	}
	block3, err := des.NewCipher(k3)
	if err != nil {
		return nil, err
	}
	// 分配解密后的数据空间
	origData := make([]byte, len(crypted))
	tmpData := make([]byte, len(crypted))
	// 进行三重DES解密
	bs := block1.BlockSize()
	for i := 0; i < len(crypted); i += bs {
		block3.Decrypt(tmpData[i:i+bs], crypted[i:i+bs])
		block2.Encrypt(tmpData[i:i+bs], tmpData[i:i+bs])
		block1.Decrypt(origData[i:i+bs], tmpData[i:i+bs])
	}
	// 去除填充数据
	origData = PKCS5Unpadding(origData)
	return origData, nil
}

//func main() {
//	// 定义密钥，长度必须是24字节
//	key := []byte("J68yys6TYv5YNHTMtZrZnIa5")
//	// 定义明文
//	plaintext := `{"merchantNo":"E1808319765","productType":"APPPAY","type":"FeeCollection","value":"OWN_CASHACCOUNT"}`
//	//plaintext := []byte(`{"merchantNo":"E1808319765","productType":"APPPAY","type":"FeeCollection","value":"OWN_CASHACCOUNT"}`)
//	fmt.Println("明文：", string(plaintext))
//	// 加密明文
//	ciphertext, err := TripleEcbDesEncrypt([]byte(plaintext), key)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("密文（十六进制）：", base64.StdEncoding.EncodeToString(ciphertext))
//
//	// 定义一个16进制的字符串
//	//hexString := ""
//	// 调用hex.DecodeString函数，返回一个字节切片和一个错误
//	//ciphertext, err = hex.DecodeString(hexString)
//	// 检查是否有错误
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 解密密文
//	newtext, err := TripleEcbDesDecrypt(ciphertext, key)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("解密后的明文：", string(newtext))
//}
