package lianlianpay_accp

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

type client struct {
	//SignMethod string
	model.Client
	sig   *utils.Sign
	epoch int64
}

// 获取连银通商户的私钥以及连银通的公钥
func NewClient(setting *model.Setting) *client {
	a := &client{Client: model.Client{
		Setting: setting}}

	//通过rsa的方法获取商户私钥以及连连的公钥
	sig, err := utils.NewSign(*setting.RsaPrivateKey, *setting.RsaPublicKey)
	if err != nil {
		panic("公私钥有问题")
	}
	a.sig = sig
	return a
}

/*
 *连连接口调用：含签名
 */

func (p *client) Execute() {

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

	p.epoch = time.Now().Unix()
	//p.t = utils.TimestampSecond()
	//待定-1.开发者使用分配的master_token向连连支付发起HTTPS请求，HTTPS请求头添加Authorization标签，格式为：
	//Authorization: Basic <Base64.encode('developerId: master_token')>
	//p.HttpReq.Header.Set(`Authorization`, fmt.Sprintf("Basic %s", p.Authorization()))
	p.HttpReq.Header.Set(`Signature-Type`, "RSA")
	p.HttpReq.Header.Set("Signature-Data", p.sign())
	//p.HttpReq.Header.Set("LLPAY-Signature", fmt.Sprintf("t=%d, v=%s", p.epoch, p.sign()))
	http.WithRequestType(http.TypeJSON)(p.HttpReq)
	p.HttpReq.Body = p.Request.Body

	if p.Err = p.Client.Execute(); p.Err != nil {
		logger.LianlianLogger.Errorf("执行连银通接口失败:%s", p.Err.Error())
		return
	}

	//验签
	signature := p.HttpReq.Response.Header.Get("Signature-Data")
	if p.Err = p.sig.RsaVerifySignWithMd5(p.HttpReq.Result, signature); p.Err != nil {
		logger.LianlianLogger.Errorf("验签失败:%s", p.Err.Error())
		return
	}

	result := model.BodyMap{}
	_ = json.Unmarshal(p.HttpReq.Result, &result)

	//请求结果代码
	p.Response.Response.Code = result.Get("ret_code")
	//请求结果描述
	p.Response.Response.Message = result.Get("ret_msg")
	//商户系统唯一交易流水号。
	p.Response.Response.RequestId = result.Get("txn_seqno")
	//ACCP系统交易单号:accp_txno 待定
	p.Response.Response.Data = p.HttpReq.Result
	p.Response.Response.Result = p.HttpReq.Result

	if result.Get("ret_code") == "0000" {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

// 计算签名
func (p *client) sign() string {

	body := p.Request.Body.JsonBody()

	ss, err := p.sig.RsaSignWithMd5([]byte(body))
	if err != nil {
		panic(errors.New("rsa 错误"))
	}
	return ss
}

//// 生成鉴权信息
//func (p *client) Authorization() string {
//
//	authStr := fmt.Sprintf("%s:%s", p.Setting.DevelopId, p.Setting.MasterToken)
//
//	ss := base64.StdEncoding.EncodeToString([]byte(authStr))
//	return ss
//}

/*
 *验签：连连公钥验签
 */

func (p *client) VerifySign(oriData, signature string) (bool, error) {
	//var signature string

	//row包含所有返回数据信息，从中可以拿到signature

	if err := p.sig.RsaVerifySignWithMd5([]byte(oriData), signature); err != nil {
		logger.HeliLogger.Error("ERROR:验签失败")
		return false, errors.New("验签失败")
	}

	return true, nil
}

func (p *client) urlParse() string {

	if p.Request.Path == nil {
		panic("apiPath is void...")
	}

	return fmt.Sprintf("%s%s", *p.Setting.ServerUrl, *p.Request.Path)
}
