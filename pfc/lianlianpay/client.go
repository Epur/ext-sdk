package lianlianpay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"net/url"
)

type Client struct {
	SignMethod string
	model.Client
	sig *utils.Sign
	t   int64
}

func NewClient(setting *model.Setting) *Client {
	a := &Client{Client: model.Client{
		Setting: setting}, SignMethod: "sha256"}

	sig, err := utils.NewSign(*setting.RsaPrivateKey, *setting.RsaPublicKey)
	if err != nil {
		panic("公私钥有问题")
	}
	a.sig = sig
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

	p.HttpReq.Header.Set(`LLPAY-Signature`, p.sign())
	p.HttpReq.Header.Set("Authorization", p.authorization())
	http.WithRequestType(http.TypeJSON)(p.HttpReq)
	p.HttpReq.Body = p.Request.Body
	//
	//if *p.Request.Method == http.GET {
	//	for key, value := range p.Request.Body {
	//		p.HttpReq.QueryParams.Add(key, value.(string))
	//	}
	//} else {
	//	http.WithRequestType(http.TypeJSON)(p.HttpReq)
	//	p.HttpReq.Body = p.Request.Body
	//}

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	result := new(Response)
	_ = json.Unmarshal(p.HttpReq.Result, &result)

	p.Response.Response.Code = result.Code
	p.Response.Response.Message = result.Message
	p.Response.Response.RequestId = result.RequestID
	p.Response.Response.Data = result.Data
	p.Response.Response.Result = result.Result

	if result.Code == "000000" {
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

func (p *Client) sign() string {

	/*
		签名
	*/

	query_params := ""
	if p.Request.Params != nil {
		query_params = "&" + p.Request.Params.EncodeURLParams()
	}

	path := p.urlParse()
	s, _ := url.Parse(path)

	body := p.Request.Body.JsonBody()
	if body == "{}" {
		body = ""
	}

	payload := fmt.Sprintf("%s&%s&%d&%s%s",
		*p.Request.Method,
		s.Path,
		p.t,
		body,
		query_params,
	)

	fmt.Println(payload)

	ss, err := p.sig.RsaSignWithMd5([]byte(payload))
	if err != nil {
		panic(errors.New("rsa 错误"))
	}
	return fmt.Sprintf("t=%d,v=%s", p.t, ss)

	//body := fmt.Sprintf("%s")
	//dd, err := json.Marshal(data)
	//if err != nil {
	//	return "", errors.New("签名数据解析失败")
	//}
	//return p.sig.RsaSignWithMd5(dd)
}

func (p *Client) authorization() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", *p.Setting.Key, *p.Setting.Secret))))
}
