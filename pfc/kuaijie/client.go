package kuaijie

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"time"
)

type Client struct {
	model.Client
	CusCode string
	t       int64
	key     *Key
}

func NewClient(setting *model.Setting) *Client {
	a := &Client{Client: model.Client{
		Setting: setting}}

	a.CusCode = *setting.UserId
	a.key = KeyNew(*setting.RsaPrivateKey, *setting.RsaPublicKey)

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

	fmt.Println(p.Client.Request.Path)
	if p.Err = p.Client.Execute(); p.Err != nil {
		fmt.Println("ERROR:", p.Err.Error())
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

	//fmt.Println("url:", p.HttpReq.Url)

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

	//result, _ := base64.StdEncoding.DecodeString(string(p.HttpReq.Result))

	////解密
	//body, err := p.key.CbcSm4Decrypt(string(result))
	//if err != nil {
	//	return nil, err
	//}

	fmt.Println(string(p.HttpReq.Result))

	row := model.BodyMap{}
	_ = json.Unmarshal(p.HttpReq.Result, &row)

	//拿到signature
	signature := row["sign"].(string)

	//row["signature"].(map[string]interface{})["sigdat"] = "__signature_sigdat__"

	if !p.key.Verify(signature, row.JsonBody()) {
		return nil, errors.New("验签失败")
	}

	return row, nil
}

func (p *Client) requestParams() (model.BodyMap, error) {
	msgPublic := model.BodyMap{}
	msgPrivate := model.BodyMap{}
	//msgProtected := model.BodyMap{}
	//secretKey := model.BodyMap{}

	//head.Set("funcode", *p.Request.Path)
	//Set("userid", p.Userid).
	//	Set("reqid", time.Now().Format("20060102150405798385")+utils.EncodeMD5(fmt.Sprintf("%d", utils.GetRandLimitInt(1, 9999999)))[:6])
	msgPublic.Set("version", "2.0").
		Set("cusReqTime", time.Now().Format("20060102150405")).
		Set("cusTraceNo", time.Now().Format("20060102150405798385")+utils.EncodeMD5(fmt.Sprintf("%d", utils.GetRandLimitInt(1, 9999999)))[:6]).
		Set("cusCode", p.CusCode)

	msgPrivate = p.Request.Body

	//signature.Set("sigtim", time.Now().Format("20060102150405")).
	//	Set("sigdat", "__signature_sigdat__")

	badyMsg := model.BodyMap{}.
		Set("msgPublic", msgPublic).
		//Set("msgProtected", msgProtected).
		//Set("secretKey", secretKey).
		Set("msgPrivate", msgPrivate)

	data := badyMsg.JsonBody() //待签名字符串
	fmt.Println("body:", data)

	//签名
	sign, err := p.key.Sign(data)
	if err != nil {
		return nil, err
	}
	//signature.Set("", t1)

	request := model.BodyMap{}.
		Set("body", badyMsg).
		Set("sign", sign)

	fmt.Println(request.JsonBody())
	fmt.Println(sign)

	////加密
	//t2, err := p.sig(request.JsonBody())
	//if err != nil {
	//	return nil, err
	//}

	r1 := model.BodyMap{}.Set("UID", p.CusCode).
		Set("FUNCODE", *p.Request.Path).
		Set("ALG", "SM").
		Set("DATA", utils.Base64Encode([]byte("t2")))

	fmt.Println(r1["DATA"])

	return r1, nil
}

func (p *Client) urlParse() string {

	if p.Request.Path == nil {
		panic("apiPath is void...")
	}

	return *p.Setting.ServerUrl
}
