package kuaijie

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/logger"
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
	if p.Request.Protected == nil {
		p.SetProtected(make(model.BodyMap))
	}
	if p.Request.Path == nil {
		logger.KuaijieLoger.Error("ERROR:Path is null..")
		p.Err = utils.Err("Path is null..")
		return
	}

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)

	row, err := p.requestParams()
	if err != nil {
		logger.KuaijieLoger.Error("ERROR:", err.Error())
		p.Err = err
		return
	}

	for key := range row {
		p.HttpReq.QueryParams.Add(key, row.Get(key))
	}
	p.HttpReq.Header.Set("Content-Type", "application/json")
	http.WithRequestType(http.TypeJSON)(p.HttpReq)
	p.HttpReq.Body = row

	if p.Err = p.Client.Execute(); p.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", p.Err.Error())
		return
	}
	//fmt.Println("HttpStatus:", p.Response.HttpStatus)

	if p.Response.HttpStatus != 200 {
		p.Err = errors.New(string(p.HttpReq.Result))
		return
	}

	response, err := p.responseParams()
	if err != nil {
		logger.KuaijieLoger.Error("ERROR:", err.Error())
		p.Err = err
		return
	}

	result := new(Response)
	//data, _ := json.Marshal(response["body"])
	_ = json.Unmarshal([]byte(response["body"].(string)), &result)
	//fmt.Println("data:", string(data))
	//fmt.Printf("\nresult:%+v\n", result)

	p.Response.Response.Code = result.MsgPublic.RspCode
	p.Response.Response.Message = result.MsgPublic.RspMsg
	p.Response.Response.RequestId = result.MsgPublic.CusTraceNo
	p.Response.Response.Data = result.MsgPrivate
	p.Response.Response.Result = json.RawMessage(response["body"].(string))
	//fmt.Printf("\np.Response.Response:%+v\n", p.Response.Response)

	if p.Response.Response.Code == "0000" {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

func (p *Client) responseParams() (model.BodyMap, error) {
	var signature, bodys string
	//result, _ := base64.StdEncoding.DecodeString(string(p.HttpReq.Result))

	////解密
	//body, err := p.key.CbcSm4Decrypt(string(result))
	//if err != nil {
	//	return nil, err
	//}

	//fmt.Println("Result:", string(p.HttpReq.Result))

	row := model.BodyMap{}
	_ = json.Unmarshal(p.HttpReq.Result, &row)
	//fmt.Printf("response:\n%+v\n", row)
	logger.KuaijieLoger.Infof("响应报文:%s", row.JsonBody())

	//拿到signature
	if row["sign"] != nil {
		signature = row["sign"].(string)
		//fmt.Printf("签名:%v\n", signature)
	} else {
		fmt.Println("没有签名")
		return row, nil
	}
	if row["body"] != nil {
		bodys = row["body"].(string)
		//fmt.Printf("response body:%+v\n", bodys)
	}

	//row["signature"].(map[string]interface{})["sigdat"] = "__signature_sigdat__"
	//fmt.Printf("row.JsonBody():%+v\n", row.JsonBody())
	if !p.key.Verify(bodys, signature) {
		logger.KuaijieLoger.Error("ERROR:验签失败")
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

	bodyMsg := model.BodyMap{}.
		Set("msgPublic", msgPublic).
		//Set("msgProtected", msgProtected).
		//Set("secretKey", secretKey).
		Set("msgPrivate", msgPrivate)
	if p.Request.Protected == nil {
		logger.KuaijieLoger.Info("protected:", p.Request.Protected.JsonBody())
		protected, _ := p.key.SignProtected(p.Request.Protected.JsonBody())
		logger.KuaijieLoger.Info("protected:", protected)
		bodyMsg = bodyMsg.Set("msgProtected", protected)
	}

	data := bodyMsg.JsonBody() //待签名字符串
	//fmt.Println("body:", data)

	//签名
	sign, err := p.key.Sign(data)
	if err != nil {
		logger.KuaijieLoger.Error("ERROR:", err.Error())
		return nil, err
	}
	//signature.Set("", t1)

	request := model.BodyMap{}.
		Set("body", data).
		Set("sign", sign)

	//fmt.Println(request.JsonBody())
	logger.KuaijieLoger.Infof("请求报文:%s", request.JsonBody())
	//fmt.Println(sign)

	////加密
	//t2, err := p.sig(request.JsonBody())
	//if err != nil {
	//	return nil, err
	//}

	//r1 := model.BodyMap{}.Set("UID", p.CusCode).
	//	Set("FUNCODE", *p.Request.Path).
	//	Set("ALG", "SM").
	//	Set("DATA", utils.Base64Encode([]byte("t2")))
	//
	//fmt.Println(r1["DATA"])

	return request, nil
}

func (p *Client) urlParse() string {

	if p.Request.Path == nil {
		panic("apiPath is void...")
	}

	return fmt.Sprintf("%s%s", *p.Setting.ServerUrl, *p.Request.Path)
}
