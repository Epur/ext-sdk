package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

var key = []byte("AB195F10A2B6413DBDE3747F18A3B948")

var iv = []byte("space-station-de")

type AES_CBC struct {
}

func main() {
	ds, err := Decrypt("rgDgjyfcT6pNEW+A6LB23hvZGZ5EpE8DdF2jXgQjxln1CSzU8916ftvBn4wF1bnurJFZA1Qt2oGHAIOMqSKAng==")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(ds))
	}
}

func Encrypt(origData []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func Decrypt(crypted string) ([]byte, error) {
	decodeData, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originByte := make([]byte, len(decodeData))
	blockMode := cipher.NewCBCDecrypter(block, iv[:block.BlockSize()])
	blockMode.CryptBlocks(originByte, decodeData)
	if len(originByte) == 0 {
		return nil, errors.New("blockMode.CryptBlocks error")
	}
	return PKCS7UnPadding(originByte), nil
}

func PKCS7UnPadding(origData []byte) (bs []byte) {
	length := len(origData)
	unPaddingNumber := int(origData[length-1]) // 找到Byte数组最后的填充byte 数字
	if unPaddingNumber <= 16 {
		bs = origData[:(length - unPaddingNumber)] // 只截取返回有效数字内的byte数组
	} else {
		bs = origData
	}
	return
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//import (
//	"bytes"
//	"crypto/aes"
//	"crypto/cipher"
//	"crypto/rand"
//	"encoding/base64"
//	"encoding/hex"
//	"io"
//	"log"
//)
//
//func main() {
//	origData := []byte("Hello World")                 // 待加密的数据
//	key := []byte("AB195F10A2B6413DBDE3747F18A3B948") // 加密的密钥
//	log.Println("原文：", string(origData))
//
//	log.Println("------------------ CBC模式 --------------------")
//	encrypted := AesEncryptCBC(origData, key)
//	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
//	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
//	encrypted = []byte("rgDgjyfcT6pNEW+A6LB23hvZGZ5EpE8DdF2jXgQjxln1CSzU8916ftvBn4wF1bnurJFZA1Qt2oGHAIOMqSKAng==")
//	decrypted := AesDecryptCBC(encrypted, key)
//	log.Println("解密结果：", string(decrypted))
//
//	log.Println("------------------ ECB模式 --------------------")
//	encrypted = AesEncryptECB(origData, key)
//	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
//	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
//	decrypted = AesDecryptECB(encrypted, key)
//	log.Println("解密结果：", string(decrypted))
//
//	log.Println("------------------ CFB模式 --------------------")
//	encrypted = AesEncryptCFB(origData, key)
//	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
//	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
//	decrypted = AesDecryptCFB(encrypted, key)
//	log.Println("解密结果：", string(decrypted))
//}
//
//// =================== CBC ======================
//func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
//	// 分组秘钥
//	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
//	block, _ := aes.NewCipher(key)
//	blockSize := block.BlockSize()                              // 获取秘钥块的长度
//	origData = pkcs5Padding(origData, blockSize)                // 补全码
//	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
//	encrypted = make([]byte, len(origData))                     // 创建数组
//	blockMode.CryptBlocks(encrypted, origData)                  // 加密
//	return encrypted
//}
//func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
//	block, _ := aes.NewCipher(key) // 分组秘钥
//	//blockSize := block.BlockSize() // 获取秘钥块的长度
//	iv := "space-station-de"
//	blockMode := cipher.NewCBCDecrypter(block, []byte(iv)) // 加密模式
//	decrypted = make([]byte, len(encrypted))               // 创建数组
//	blockMode.CryptBlocks(decrypted, encrypted)            // 解密
//	decrypted = pkcs5UnPadding(decrypted)                  // 去除补全码
//	return decrypted
//}
//func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext, padtext...)
//}
//func pkcs5UnPadding(origData []byte) []byte {
//	length := len(origData)
//	unpadding := int(origData[length-1])
//	return origData[:(length - unpadding)]
//}
//
//// =================== ECB ======================
//func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
//	cipher, _ := aes.NewCipher(generateKey(key))
//	length := (len(origData) + aes.BlockSize) / aes.BlockSize
//	plain := make([]byte, length*aes.BlockSize)
//	copy(plain, origData)
//	pad := byte(len(plain) - len(origData))
//	for i := len(origData); i < len(plain); i++ {
//		plain[i] = pad
//	}
//	encrypted = make([]byte, len(plain))
//	// 分组分块加密
//	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
//	}
//
//	return encrypted
//}
//func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
//	cipher, _ := aes.NewCipher(generateKey(key))
//	decrypted = make([]byte, len(encrypted))
//	//
//	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
//	}
//
//	trim := 0
//	if len(decrypted) > 0 {
//		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
//	}
//
//	return decrypted[:trim]
//}
//func generateKey(key []byte) (genKey []byte) {
//	genKey = make([]byte, 16)
//	copy(genKey, key)
//	for i := 16; i < len(key); {
//		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
//			genKey[j] ^= key[i]
//		}
//	}
//	return genKey
//}
//
//// =================== CFB ======================
//func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//	encrypted = make([]byte, aes.BlockSize+len(origData))
//	iv := encrypted[:aes.BlockSize]
//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//		panic(err)
//	}
//	stream := cipher.NewCFBEncrypter(block, iv)
//	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
//	return encrypted
//}
//func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
//	block, _ := aes.NewCipher(key)
//	if len(encrypted) < aes.BlockSize {
//		panic("ciphertext too short")
//	}
//	iv := encrypted[:aes.BlockSize]
//	encrypted = encrypted[aes.BlockSize:]
//
//	stream := cipher.NewCFBDecrypter(block, iv)
//	stream.XORKeyStream(encrypted, encrypted)
//	return encrypted
//}

//
//import (
//	"fmt"
//	"github.com/Epur/ext-sdk/utils/aes"
//	"log"
//)
//
//func main() {
//	// 示例密钥和初始化向量
//	//key := []byte("AB195F10A2B6413DBDE3747F18A3B948")                                   // 16字节长度
//	//iv := []byte("space-station-de")                                                    // 16字节长度
//	//encryptedText := "KsojcDmWm9MS99kEECoL3E5d5dPYRklFs1jkBvBXdvqiPq6Rk19d5vToehuDPsa2" // 加密文本的Base64编码
//	//
//	//// 解密
//	//decryptedText, err := aes.DecryptAESWithIV(key, iv, encryptedText)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	//decryptedText := aes.AesDecryptCFB([]byte("rgDgjyfcT6pNEW+A6LB23hvZGZ5EpE8DdF2jXgQjxln1CSzU8916ftvBn4wF1bnurJFZA1Qt2oGHAIOMqSKAng=="), []byte("AB195F10A2B6413DBDE3747F18A3B948"))
//	//
//	//// 输出解密结果
//	//log.Println("Decrypted text:", string(decryptedText))
//
//	decryptedText, err := aes.DecryptAES("rgDgjyfcT6pNEW+A6LB23hvZGZ5EpE8DdF2jXgQjxln1CSzU8916ftvBn4wF1bnurJFZA1Qt2oGHAIOMqSKAng==", "AB195F10A2B6413DBDE3747F18A3B948")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	log.Println("Decrypted text:", string(decryptedText))
//}

//	func main() {
//		orderDetail := `{
//	   "code": "0",
//	   "msg": "OK",
//	   "info": [
//	       {
//	           "orderNo": "GSONK315800M5DC",
//	           "orderType": 1,
//	           "performanceType": 1,
//	           "orderStatus": 2,
//	           "isCod": 2,
//	           "isOverLimitOrder": 1,
//	           "unpackingStatus": 3,
//	           "orderTag": 0,
//	           "deliveryType": 1,
//	           "printOrderStatus": 1,
//	           "invoiceStatus": 2,
//	           "orderGoodsInfoList": [
//	               {
//	                   "goodsId": 3000000000232252,
//	                   "skuCode": "I5dtvb4nk1dx",
//	                   "skc": "sM23040300586693",
//	                   "goodsSn": "SPMP发布自定义保存草稿2",
//	                   "sellerSku": "",
//	                   "goodsStatus": 1,
//	                   "newGoodsStatus": 2,
//	                   "skuAttribute": [
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "CN"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "US"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7葡文-次规格值自定义2",
//	                           "language": "PT"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7泰语-次规格值自定义2",
//	                           "language": "TH"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "ES"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "IT"
//	                       }
//	                   ],
//	                   "goodsTitle": "GoodsName111111111111",
//	                   "spuPicURL": "http://imgdeal-test01.shein.com/pi_img/2021/11/03/16359230894287290938_thumbnail_220x293.jpg",
//	                   "goodsWeight": 500.00,
//	                   "storageTag": 1,
//	                   "performanceTag": 2,
//	                   "goodsExchangeTag": 1,
//	                   "unpackingGroupNo": "FG24022932245297156",
//	                   "unpackingGroupInvoiceStatus": 2,
//	                   "beExchangeEntityId": 0,
//	                   "orderCurrency": "MXN",
//	                   "sellerCurrencyPrice": 280.45,
//	                   "orderCurrencyStoreCouponPrice": 0.00,
//	                   "orderCurrencyPromotionPrice": 0.00,
//	                   "commission": 0.00,
//	                   "commissionRate": 0.0000,
//	                   "serviceCharge": 0.00,
//	                   "performanceServiceCharge": 0.00,
//	                   "estimatedIncome": 280.45,
//	                   "spuName": "M2304030058",
//	                   "saleTax": 0.00,
//	                   "warehouseCode": null,
//	                   "warehouseName": null,
//	                   "sellerCurrencyDiscountPrice": 801.06
//	               },
//	               {
//	                   "goodsId": 3000000000232253,
//	                   "skuCode": "I5dtvb4nk1dx",
//	                   "skc": "sM23040300586693",
//	                   "goodsSn": "SPMP发布自定义保存草稿2",
//	                   "sellerSku": "",
//	                   "goodsStatus": 1,
//	                   "newGoodsStatus": 2,
//	                   "skuAttribute": [
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "CN"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "US"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7葡文-次规格值自定义2",
//	                           "language": "PT"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7泰语-次规格值自定义2",
//	                           "language": "TH"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "ES"
//	                       },
//	                       {
//	                           "attrValueId": "1007750,39,1007752",
//	                           "attrName": "主规格值自定义2-7-次规格值自定义2",
//	                           "language": "IT"
//	                       }
//	                   ],
//	                   "goodsTitle": "GoodsName111111111111",
//	                   "spuPicURL": "http://imgdeal-test01.shein.com/pi_img/2021/11/03/16359230894287290938_thumbnail_220x293.jpg",
//	                   "goodsWeight": 500.00,
//	                   "storageTag": 1,
//	                   "performanceTag": 2,
//	                   "goodsExchangeTag": 1,
//	                   "unpackingGroupNo": "FG24022932245297155",
//	                   "unpackingGroupInvoiceStatus": 2,
//	                   "beExchangeEntityId": 0,
//	                   "orderCurrency": "MXN",
//	                   "sellerCurrencyPrice": 280.45,
//	                   "orderCurrencyStoreCouponPrice": 0.00,
//	                   "orderCurrencyPromotionPrice": 0.00,
//	                   "commission": 0.00,
//	                   "commissionRate": 0.0000,
//	                   "serviceCharge": 0.00,
//	                   "performanceServiceCharge": 0.00,
//	                   "estimatedIncome": 280.45,
//	                   "spuName": "M2304030058",
//	                   "saleTax": 0.00,
//	                   "warehouseCode": null,
//	                   "warehouseName": null,
//	                   "sellerCurrencyDiscountPrice": 801.06
//	               }
//	           ],
//	           "packageWaybillList": [
//	               {
//	                   "packageNo": "GC24030727444473856",
//	                   "waybillNo": "UAT301009158929",
//	                   "carrier": "222222",
//	                   "carrierCode": "TYDHL-BA1",
//	                   "productInventoryList": [
//	                       {
//	                           "productId": "3000000000232253"
//	                       }
//	                   ],
//	                   "packageLabel": "",
//	                   "sortingCode": "H09",
//	                   "expressSortingCode": "NYN",
//	                   "isCutOffSeller": 2
//	               },
//	               {
//	                   "packageNo": "GC24030727444473857",
//	                   "waybillNo": "UAT301008583535",
//	                   "carrier": "222222",
//	                   "carrierCode": "TYDHL-BA1",
//	                   "productInventoryList": [
//	                       {
//	                           "productId": "3000000000232252"
//	                       }
//	                   ],
//	                   "packageLabel": "",
//	                   "sortingCode": "H09",
//	                   "expressSortingCode": "NYN",
//	                   "isCutOffSeller": 2
//	               }
//	           ],
//	           "orderCurrency": "MXN",
//	           "productTotalPrice": 560.90,
//	           "storeDiscountTotalPrice": 0.00,
//	           "promotionDiscountTotalPrice": 0.00,
//	           "totalCommission": 0.00,
//	           "totalServiceCharge": 0.00,
//	           "totalPerformanceServiceCharge": 0.00,
//	           "estimatedGrossIncome": 560.90,
//	           "totalSaleTax": 0.00,
//	           "orderTime": "2024-02-29T17:03:31.000+0800",
//	           "paymentTime": "2024-02-29T17:03:32.000+0800",
//	           "orderAllocateTime": "2024-02-29T17:05:05.000+0800",
//	           "requestDeliveryTime": "2024-03-01T17:05:05.000+0800",
//	           "sellerDeliveryTime": "",
//	           "warehouseDeliveryTime": "",
//	           "printingTime": "2024-03-07T17:00:55.000+0800",
//	           "scheduleDeliveryTime": "",
//	           "pickUpTime": "",
//	           "orderReceiptTime": "",
//	           "orderRejectionTime": "",
//	           "orderReportedLossTime": "",
//	           "orderReturnTime": "",
//	           "orderMsgUpdateTime": "2024-03-07T17:00:58.360+0800",
//	           "billNo": "GSONK315800M5DC",
//	           "salesArea": 1,
//	           "stockMode": 3,
//	           "salesSite": "shein-mx",
//	           "storeCode": 3342233928,
//	           "settleActuallyPrice": null,
//	           "unProcessReason": [],
//	           "packageInvoiceProblems": [
//	               {
//	                   "problemCode": null,
//	                   "problemDescEnglish": "The order package requires a CTE logistics invoice, and the platform logistics provider is currently invoicing. Please be patient and wait",
//	                   "problemField": null,
//	                   "proposalEnglish": null,
//	                   "packageNo": "GC24030727444473856"
//	               },
//	               {
//	                   "problemCode": null,
//	                   "problemDescEnglish": "The order package requires a CTE logistics invoice, and the platform logistics provider is currently invoicing. Please be patient and wait",
//	                   "problemField": null,
//	                   "proposalEnglish": null,
//	                   "packageNo": "GC24030727444473857"
//	               }
//	           ],
//	           "expectedCollectTime": ""
//	       }
//	   ],
//	   "bbl": {}
//	}`
//
//		response := shein.OrderDetailResponse{}
//
//		err := json.Unmarshal([]byte(orderDetail), &response)
//		if err != nil {
//			fmt.Printf("%s\n", err.Error())
//		} else {
//			fmt.Printf("%#v\n", response)
//			//fmt.Println(response)
//		}
//		fmt.Println("===================")
//	}
//package main
//
//import (
//	"bytes"
//	"crypto/aes"
//	"crypto/cipher"
//	"encoding/base64"
//	"fmt"
//)
//
//var (
//	initialVector = "space-station-de"
//	passphrase    = "AB195F10A2B6413DBDE3747F18A3B948"
//)
//
//func main() {
//	var plainText = "hello world"
//
//	encryptedData := AESEncrypt(plainText, []byte(passphrase))
//	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
//	fmt.Println(encryptedString)
//	encryptedString = "rgDgjyfcT6pNEW+A6LB23hvZGZ5EpE8DdF2jXgQjxln1CSzU8916ftvBn4wF1bnurJFZA1Qt2oGHAIOMqSKAng=="
//
//	encryptedData, _ = base64.StdEncoding.DecodeString(encryptedString)
//	decryptedText := AESDecrypt(encryptedData, []byte(passphrase))
//	fmt.Println(string(decryptedText))
//}
//
//func AESEncrypt(src string, key []byte) []byte {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		fmt.Println("key error1", err)
//	}
//	if src == "" {
//		fmt.Println("plain content empty")
//	}
//	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
//	content := []byte(src)
//	content = PKCS5Padding(content, block.BlockSize())
//	crypted := make([]byte, len(content))
//	ecb.CryptBlocks(crypted, content)
//
//	return crypted
//}
//
//func AESDecrypt(crypt []byte, key []byte) []byte {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		fmt.Println("key error1", err)
//	}
//	if len(crypt) == 0 {
//		fmt.Println("plain content empty")
//	}
//	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
//	decrypted := make([]byte, len(crypt))
//	ecb.CryptBlocks(decrypted, crypt)
//
//	return PKCS5Trimming(decrypted)
//}
//
//func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext, padtext...)
//}
//
//func PKCS5Trimming(encrypt []byte) []byte {
//	padding := encrypt[len(encrypt)-1]
//	return encrypt[:len(encrypt)-int(padding)]
//}
//
//import (
//	"bytes"
//	"crypto/aes"
//	"crypto/cipher"
//	"crypto/md5"
//	"encoding/base64"
//	"encoding/hex"
//	"fmt"
//)
//
//func getKeyBytes(key string) []byte {
//	keyBytes := []byte(key)
//	switch l := len(keyBytes); {
//	case l < 16:
//		keyBytes = append(keyBytes, make([]byte, 16-l)...)
//	case l > 16:
//		keyBytes = keyBytes[:16]
//	}
//	return keyBytes
//}
//
//func encrypt(key string, origData []byte) ([]byte, error) {
//	keyBytes := getKeyBytes(key)
//	block, err := aes.NewCipher(keyBytes)
//	if err != nil {
//		return nil, err
//	}
//	blockSize := block.BlockSize()
//	origData = PKCS5Padding(origData, blockSize)
//	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
//	crypted := make([]byte, len(origData))
//	blockMode.CryptBlocks(crypted, origData)
//	return crypted, nil
//}
//
//func decrpt(key string, crypted []byte) ([]byte, error) {
//	keyBytes := getKeyBytes(key)
//	block, err := aes.NewCipher(keyBytes)
//	if err != nil {
//		return nil, err
//	}
//	blockSize := block.BlockSize()
//	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
//	origData := make([]byte, len(crypted))
//	blockMode.CryptBlocks(origData, crypted)
//	origData = PKCS5UnPadding(origData)
//	return origData, nil
//}
//
//func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext, padtext...)
//}
//
//func PKCS5UnPadding(origData []byte) []byte {
//	length := len(origData)
//	unpadding := int(origData[length-1])
//	return origData[:(length - unpadding)]
//}
//
//func Encrypt(key string, val string) (string, error) {
//	origData := []byte(val)
//	crypted, err := encrypt(key, origData)
//	if err != nil {
//		return "", err
//	}
//	return base64.URLEncoding.EncodeToString(crypted), nil
//}
//
//func Decrypt(key string, val []byte) (string, error) {
//	origData, err := decrpt(key, val)
//	if err != nil {
//		return "", err
//	}
//	return string(origData), nil
//}
//
//func main() {
//	m := md5.Sum([]byte("cvrhsdftredhghgfjhgwsfresdsfhjk"))
//	key := hex.EncodeToString(m[:])[0:16]
//	cal := "VYGThW0MXPf4v88IKP/o4g=="
//	// url 解码
//	//unescape, _ := url.QueryUnescape(cal)
//	// base64 解码
//	decodeString, _ := base64.StdEncoding.DecodeString(cal)
//
//	// 揭秘
//	decrypt, _ := Decrypt(key, decodeString)
//	fmt.Println(decrypt)
//
//}
