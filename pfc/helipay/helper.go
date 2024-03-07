package helipay

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/Epur/ext-sdk/logger"
	"github.com/deatil/go-cryptobin/gm/sm2"
	"github.com/deatil/go-cryptobin/pkcs12"
	"github.com/jinzhu/copier"
	//"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"io/ioutil"
	"os"
	//pkcs12 "software.sslmate.com/src/go-pkcs12"
	"sync"
)

type SM2 struct {
	//PrivateKey string
	//PublicKey  string
	PrivateKey *sm2.PrivateKey // 商户私钥，商户加签
	PublicKey  *sm2.PublicKey  // 合利宝公钥，验签

	UserId string
}

func NewSM2(privateKeypath, password, publicKey string) *SM2 {
	//根据私钥存放路径以及口令，导出私钥
	priKey, err := LoadPrivateKey(privateKeypath, password)
	if err != nil {
		logger.HeliLogger.Errorf("加载私钥失败:[%s]\n", err.Error())
		return nil
	}
	//根据合利宝公钥字符串导出合利宝公钥
	pubKey := LoadPublicKey(publicKey)
	if pubKey == nil {
		return nil
	}

	return &SM2{
		PrivateKey: priKey,
		PublicKey:  pubKey,
		UserId:     "",
	}
}

var certCAs []*x509.Certificate

var initonce sync.Once

func LoadPublicKey(publicKey string) *sm2.PublicKey {
	prefix := `-----BEGIN CERTIFICATE-----`
	postfix := `-----END CERTIFICATE-----`

	apublicKey := fmt.Sprintf("%s\n%s%s", prefix, publicKey, postfix)
	initonce.Do(func() {
		block, _ := pem.Decode([]byte(apublicKey))
		ca, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			panic(err)
		}
		certCAs = append(certCAs, ca)
	})
	if len(certCAs) > 0 {
		v := certCAs[0].PublicKey.(*ecdsa.PublicKey)
		key := sm2.PublicKey{}
		err := copier.Copy(&key, v)
		if err != nil {
			return nil
		}
		return &key
	}
	return nil
}

/*
 * 商户私钥加签
 */
func (p *SM2) Sign(data string) (string, error) {
	sign, err := p.PrivateKey.Sign(nil, []byte(data), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

//func (p *SM2) Sign(data string) (string, error) {
//
//	sm2keyBytes, _ := base64.StdEncoding.DecodeString(p.PrivateKey)
//
//	sm2Sign := cryptobin_sm2.NewSM2().
//		FromPrivateKeyBytes(sm2keyBytes).
//		FromString(data).
//		SignHex([]byte(p.UserId)).
//		ToBase64String()
//	return sm2Sign, nil
//}

/*
 * 合利宝公钥验签
 */
func (p *SM2) Verify(data string, sign string) bool {
	signbytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		logger.HeliLogger.Errorf("解码失败:[%s]\n", err.Error())
		return false
	}
	logger.HeliLogger.Infof("验签成功，data=[%s]，sign=[%s]\n", data, signbytes)
	boo := p.PublicKey.Verify([]byte(data), signbytes, nil)
	if boo {
		logger.HeliLogger.Infof("验签成功，data=[%s]，sign=[%s]\n", data, sign)
	} else {
		logger.HeliLogger.Infof("验签失败，data=[%s]，sign=[%s]\n", data, sign)
	}
	return true
}

//func (p *SM2) Verify(sm2signdata string, data string) bool {
//
//	sm2keyBytes, _ := base64.StdEncoding.DecodeString(p.PublicKey)
//
//	sssss := cryptobin_sm2.NewSM2().FromBase64String(sm2signdata).
//		FromPublicKeyString(hex.EncodeToString(sm2keyBytes)).
//		VerifyHex([]byte(data), []byte(p.UserId))
//
//	fmt.Println(sssss.Error())
//
//	return sssss.ToVerify()
//}

func LoadPrivateKey(privateKeyName, privatePassword string) (*sm2.PrivateKey, error) {
	f, err := os.Open(privateKeyName)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	// 因为pfx证书公钥和密钥是成对的，所以要先转成pem.Block
	//blocks, err := pkcs12.ToPEM(bytes, privatePassword)
	//prikey, certs, err := pkcs12.DecodeAll(bytes, privatePassword)
	//if err != nil {
	//	return nil, err
	//}
	pkcs121, err := pkcs12.LoadPKCS12FromBytes(bytes, privatePassword)
	fmt.Println(pkcs121)
	if err != nil {
		return nil, err
	}
	prikey, _, err := pkcs121.GetPrivateKey()
	if err != nil {
		return nil, err
	}
	fmt.Println(prikey)
	//fmt.Println(certs)
	v := prikey.(*sm2.PrivateKey)
	//key := sm2.PrivateKey{}
	//err = copier.Copy(&key, v)
	//if err != nil {
	//	fmt.Println("copier failed :" + err.Error())
	//	return nil, err
	//}
	return v, nil
}
