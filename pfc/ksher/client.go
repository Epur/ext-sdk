package ksher

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/Epur/ext-sdk/utils/aes"
	"github.com/Epur/ext-sdk/utils/rsa"
	"github.com/tangchen2018/go-utils/http"
	"strconv"
	"time"
)

type Client struct {
	SignMethod string
	model.Client
	sig           *rsa.Rsa
	t             int64
	e, p, s, c, k string
}

type R struct {
	EncryptKey string `json:"encrypt_key"`
	Signature  string `json:"signature"`
	CipherText string `json:"cipher_text"`
}

func NewClient(setting *model.Setting) *Client {
	a := &Client{Client: model.Client{
		Setting: setting}, SignMethod: "sha256"}
	//if setting.RsaPrivateKey == nil {
	//	setting.RsaPrivateKey = utils.PString("")
	//}
	//if setting.RsaPublicKey == nil {
	//	setting.RsaPublicKey = utils.PString("")
	//}
	a.sig = rsa.NewRsa(*setting.RsaPublicKey, *setting.RsaPrivateKey)
	return a
}

func (p *Client) Execute() {

	if p.Request.Method == nil || len(*p.Request.Method) < 0 {
		p.Request.Method = utils.PString(http.GET)
	}
	if p.Request.Body == nil {
		p.SetBody(make(model.BodyMap))
	}
	if p.Setting.Key == nil {
		p.Err = utils.Err("Key is null..")
		return
	}
	if p.Request.Path == nil {
		p.Err = utils.Err("Path is null..")
		return
	}

	var err error
	p.p = p.P()
	p.e = p.E()
	p.s = p.S()
	p.c = p.C()

	//fmt.Println(p.e)

	if p.k, err = p.K(); err != nil {
		p.Err = err
		return
	}

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)

	p.t = utils.TimestampSecond()

	//p.HttpReq.Header.Set("Content-type", "application/json;charset=utf-8")
	//p.HttpReq.Header.Set("Signature-Type", "RSA")

	//signature, err := p.sign(p.Request.Body)
	//if err != nil {
	//	p.Err = err
	//	return
	//}

	p.HttpReq.Header.Set(`X-Api-Key`, *p.Setting.Key)
	p.HttpReq.Header.Set("X-Request-Id", p.E())
	p.HttpReq.Header.Set("X-Request-Timestamp", time.Now().Format("2006-01-02T15:04:05.798385+08:00"))
	http.WithRequestType(http.TypeJSON)(p.HttpReq)
	p.HttpReq.Body = model.BodyMap{}.
		Set("cipher_text", p.c).
		Set("encrypt_key", p.k).
		Set("signature", p.s)

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	result := new(Response)

	r := new(R)
	_ = json.Unmarshal(p.HttpReq.Result, &r)
	if row, err := p.D(r); err != nil {
		p.Err = err
		return
	} else {
		_ = json.Unmarshal([]byte(row), &result)
	}

	p.Response.Response.Code = result.Code
	p.Response.Response.Message = result.Message
	p.Response.Response.RequestId = result.RequestID
	p.Response.Response.Data = result.Data
	p.Response.Response.Result = result.Result

	if result.Code == "SUCCESS" {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

func (p *Client) urlParse() string {

	if p.Request.Path == nil {
		panic("apiPath is void...")
	}

	return fmt.Sprintf("%s%s", *p.Setting.ServerUrl, *p.Request.Path)
}

//解包

func (p *Client) D(r *R) (string, error) {

	/*
		解密和验签的策略: urlsafe base64解码body的数据cipher_text记作C。 urlsafe base64解码body的数据 signature记作S。 urlsafe base64解码body的数据encrypt_key记作K。
			1. 用渠道方自己的私钥对K进行解密，解密后的数据记作E。
			2. 用E对C进行AES对称解密，将结果进行urlsafe base64解码，得到字符串P。
			3. 将json字符串P加载为字典，记作D。
			4. 对D按照字典顺序进行嵌套json序列化，然后进行urlsafe base64编码, 得到字符串D1,
			5. 使用HMAC-SHA256对D1进行签名(秘钥使用api_secret), 签名后的结果进行urlsafe base64编码,记作S1。
			6. 对比字符串S1和S是否相等, 相同则认为签名通过，否则说明信息已经遭到了篡改。
	*/
	K, err := base64.URLEncoding.DecodeString(r.EncryptKey)
	if err != nil {
		return "", err
	}
	C, err := base64.URLEncoding.DecodeString(r.CipherText)
	if err != nil {
		return "", err
	}
	S, err := base64.URLEncoding.DecodeString(r.Signature)
	if err != nil {
		return "", err
	}

	E, err := p.sig.Decrypt(K)
	if err != nil {
		return "", err
	}

	P, err := base64.URLEncoding.DecodeString(string(aes.AesDecryptECB(C, E)))
	if err != nil {
		return "", err
	}
	D := model.BodyMap{}
	if err := json.Unmarshal(P, &D); err != nil {
		return "", err
	}
	//D["message"] = `\u6210\u529f`
	//fmt.Println(D.JsonBody())

	//aaaaa := D.JsonBody()
	//fmt.Println(strconv.QuoteToASCII(aaaaa))

	D1 := utils.Base64UrlEncode([]byte(strconv.QuoteToASCII(D.JsonBody())))

	//fmt.Println(strconv.QuoteToASCII(D.JsonBody()))
	//D1 := utils.Base64UrlEncode(
	//	[]byte(`{"code":"SUCCESS","data":{"items":[],"limit":1,"page":1,"total":0},"message":"\u6210\u529f","timestamp":"2023-07-27T14:32:23+08:00"}`))

	fmt.Println(D1)

	S1 := utils.HmacSha256ToBase64(*p.Setting.Secret, D1)

	if S1 != string(S) {
		fmt.Println(S1, string(S))
		//return "", errors.New("解包签名失败")
	}

	return string(P), nil
}

//生成16位随机数做AES的秘钥，记作E。

func (p *Client) E() string {
	return utils.EncodeMD5(fmt.Sprintf("%d", utils.GetRandLimitInt(1, 9999999)))[:16]
}

//对参与请求的业务数据(记作D)，按照字典顺序进行嵌套json序列化(含有中文的话，中文字符以ascii转 换)，然后进行urlsafe b64encode编码, 得到字符串P

func (p *Client) P() string {
	return utils.Base64UrlEncode([]byte(p.Request.Body.JsonBody()))
}

/*
	使用HMAC-SHA256对P进行签名(秘钥使用api_secret),
	签名后的结果进行urlsafe base64编码,记作S。
	result = base64.urlsafe_b64encode(cipher_text)
*/

func (p *Client) S() string {
	return utils.HmacSha256ToBase64(*p.Setting.Secret, p.p)
}

/*
	使用E做为秘钥对P做AES对称加密，加密后的结果进行urlsafe base64编码记作C。
*/

func (p *Client) C() string {
	return utils.Base64UrlEncode(aes.AesEncryptECB([]byte(p.p), []byte(p.e)))
}

//用Ksher的RSA公钥对E进行RSA加密，加密后的结果进行urlsafe base64编码记作K。 最后得到的结果如下

func (p *Client) K() (string, error) {
	s, a := p.sig.Encrypt([]byte(p.e))
	return utils.Base64UrlEncode(s), a
}
