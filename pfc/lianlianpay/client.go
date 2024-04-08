package lianlianpay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/logger"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"net/url"
	"strings"
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

	p.HttpReq.Header.Set("Content-type", " application/json")
	p.HttpReq.Header.Set(`LLPAY-Signature`, p.sign())
	p.HttpReq.Header.Set("Authorization", p.authorization())
	http.WithRequestType(http.TypeJSON)(p.HttpReq)
	p.HttpReq.Body = p.Request.Body

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	//验签
	if _, p.Err = p.responseParams(); p.Err != nil {
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

	//获取请求路径
	path := p.urlParse()
	s, _ := url.Parse(path)

	body := p.Request.Body.JsonBody()
	if body == "{}" {
		body = ""
	}
	//去掉query_params参数
	payload := fmt.Sprintf("%s&%s&%d&%s",
		*p.Request.Method,
		s.Path,
		p.t,
		body,
		//query_params,
	)

	fmt.Println(payload)

	ss, err := p.sig.RsaSignWithSHA256([]byte(payload))
	if err != nil {
		panic(errors.New("rsa 错误"))
	}
	return fmt.Sprintf("t=%d,v=%s", p.t, ss)

}

//转码开发者id及mastertoken

func (p *Client) authorization() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", *p.Setting.DevelopId, *p.Setting.MasterToken))))
}

//验签及解密敏感字段

func (p *Client) responseParams() (model.BodyMap, error) {

	//body := model.BodyMap{}
	//获取body
	body := p.Request.Body
	//获取签名信息
	signature := p.HttpReq.Header.Get("Signature-Data")
	//验证签名信息
	if strings.Compare(signature, "") == 0 {
		logger.LianlianLogger.Errorf("%s", "签名串不存在，存在伪造可能")
		return nil, errors.New("签名串不存在，存在伪造可能")
	}
	signData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	err = p.sig.RsaVerifySignWithMd5(signData, signature)

	return nil, err
}
