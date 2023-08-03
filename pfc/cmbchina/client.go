package cmbchina

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"time"
)

type Client struct {
	SignMethod string
	model.Client
	Userid string
	t      int64
	sig    *SM4
}

func NewClient(setting *model.Setting) *Client {
	a := &Client{Client: model.Client{
		Setting: setting}, SignMethod: "sha256"}

	a.Userid = *setting.UserId

	a.sig = SM4New(
		setting.SM4.PrivateKey,
		setting.SM4.PublicKey,
		setting.SM4.BankPublicKey,
		setting.SM4.SymKey, a.Userid)

	return a
}

func (p *Client) Execute() {
	if p.Request.Method == nil || len(*p.Request.Method) < 0 {
		p.Request.Method = utils.PString(http.GET)
	}
	if p.Request.Body == nil {
		p.SetBody(make(model.BodyMap))
	}
	if p.Request.Path == nil {
		p.Err = utils.Err("Path is null..")
		return
	}

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)

	row, err := p.requestParams()
	if err != nil {
		p.Err = err
		return
	}

	for key := range row {
		p.HttpReq.QueryParams.Add(key, row.Get(key))
	}

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	if p.Response.HttpStatus != 200 {
		p.Err = errors.New(string(p.HttpReq.Result))
		return
	}

	response, err := p.responseParams()
	if err != nil {
		p.Err = err
		return
	}
	////
	//result := new(Response)

	fmt.Println(p.HttpReq.Url)

	data, _ := json.Marshal(response["response"].(map[string]interface{})["body"])

	p.Response.Response.Code = response["response"].(map[string]interface{})["head"].(map[string]interface{})["resultcode"].(string)
	p.Response.Response.Message = response["response"].(map[string]interface{})["head"].(map[string]interface{})["resultmsg"].(string)
	p.Response.Response.RequestId = response["response"].(map[string]interface{})["head"].(map[string]interface{})["reqid"].(string)
	p.Response.Response.Data = data
	p.Response.Response.Result = data

	if p.Response.Response.Code == "SUC0000" {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

func (p *Client) responseParams() (model.BodyMap, error) {

	result, _ := base64.StdEncoding.DecodeString(string(p.HttpReq.Result))

	//解密
	body, err := p.sig.CbcSm4Decrypt(string(result))
	if err != nil {
		return nil, err
	}

	row := model.BodyMap{}
	_ = json.Unmarshal(body, &row)

	//拿到signature
	signature := row["signature"].(map[string]interface{})["sigdat"].(string)

	row["signature"].(map[string]interface{})["sigdat"] = "__signature_sigdat__"

	fmt.Println(body)
	fmt.Println(p.sig.Verify(signature, row.JsonBody()), string(row.JsonBody()))

	return row, nil
}

func (p *Client) requestParams() (model.BodyMap, error) {
	head := model.BodyMap{}
	body := model.BodyMap{}
	signature := model.BodyMap{}

	head.Set("funcode", *p.Request.Path).
		Set("userid", p.Userid).
		Set("reqid", time.Now().Format("20060102150405798385")+utils.EncodeMD5(fmt.Sprintf("%d", utils.GetRandLimitInt(1, 9999999)))[:6])
	body = p.Request.Body

	signature.Set("sigtim", time.Now().Format("20060102150405")).
		Set("sigdat", "__signature_sigdat__")

	request := model.BodyMap{}.Set("request", model.BodyMap{}.Set("head", head).Set("body", body)).Set("signature", signature)

	data := request.JsonBody() //待签名字符串

	//签名
	t1, err := p.sig.Sign(data)
	if err != nil {
		return nil, err
	}
	signature.Set("sigdat", t1)

	request.Set("signature", signature)

	fmt.Println(request.JsonBody())
	fmt.Println(t1)

	//加密
	t2, err := p.sig.CbcSm4Encrypt(request.JsonBody())
	if err != nil {
		return nil, err
	}

	r1 := model.BodyMap{}.Set("UID", p.Userid).Set("FUNCODE", *p.Request.Path).Set("ALG", "SM").
		Set("DATA", utils.Base64Encode(t2))

	fmt.Println(r1["DATA"])

	return r1, nil
}

//		p.HttpReq = http.New(
//			http.WithUrl(p.urlParse()),
//			http.WithMethod(*p.Request.Method),
//		)
//
//		p.t = utils.TimestampSecond()
//
//		p.requestParams()
//
//		p.HttpReq.Header.Set(`X-Api-Key`, *p.Setting.Key)
//		p.HttpReq.Header.Set("X-Request-Id", p.E())
//		p.HttpReq.Header.Set("X-Request-Timestamp", time.Now().Format("2006-01-02T15:04:05.798385+08:00"))
//		http.WithRequestType(http.TypeJSON)(p.HttpReq)
//		p.HttpReq.Body = model.BodyMap{}.
//			Set("cipher_text", p.c).
//			Set("encrypt_key", p.k).
//			Set("signature", p.s)
//
//		if p.Err = p.Client.Execute(); p.Err != nil {
//			return
//		}
//
//		result := new(Response)
//
//		p.Response.Response.Code = result.Code
//		p.Response.Response.Message = result.Message
//		p.Response.Response.RequestId = result.RequestID
//		p.Response.Response.Data = result.Data
//		p.Response.Response.Result = result.Result
//
//		if result.Code == "SUCCESS" {
//			p.Response.Success = true
//		}
//
//		if p.Response.HttpStatus != 200 {
//			p.Response.Success = false
//		}
//	}
func (p *Client) urlParse() string {

	if p.Request.Path == nil {
		panic("apiPath is void...")
	}

	return *p.Setting.ServerUrl
}

//
//func (p *Client) requestParams() error {
//
//	head := model.BodyMap{}
//	body := model.BodyMap{}
//	signature := model.BodyMap{}
//
//	head.Set("funcode", *p.Request.Path).
//		Set("userid", p.Userid).
//		Set("reqid", time.Now().Format("20060102150405")+utils.EncodeMD5(fmt.Sprintf("%d", utils.GetRandLimitInt(1, 9999999)))[:6])
//	body = p.Request.Body
//
//	signature.Set("sigtim", time.Now().Format("20060102150405")).
//		Set("sigdat", "__signature_sigdat__")
//
//	request := model.BodyMap{}.Set("request", model.BodyMap{}.Set("head", head).Set("body", body).Set("signature", signature))
//
//	s := request.JsonBody() //待签名字符串
//
//	signStr, err := sm2.Sign(HexToPri("HAyo83HmtOVVqOYgoJdTNhnXuYyVcWUJ4d3p/26blNk="), []byte(p.Userid), []byte(s))
//	if err != nil {
//		return err
//	}
//	sign := base64.StdEncoding.EncodeToString(signStr)
//	//sm2.Sm2Sign()
//}
//
//func HexToPri(priStr string) *sm2.PrivateKey {
//	privateKeyByte, _ := hex.DecodeString(priStr)
//	pri, err := sm2.RawBytesToPrivateKey(privateKeyByte)
//	if err != nil {
//		panic("私钥加载异常")
//	}
//	return pri
//}
