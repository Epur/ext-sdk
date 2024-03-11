package helipay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/logger"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"sort"
	"strings"
)

type client struct {
	model.Client
	UserId string //商户编号，合利宝统一分配
	key    SM2
}

/*
 * 生成合利宝的客户端，用于后续的操作
 */
func NewClient(setting *model.Setting) *client {
	a := &client{Client: model.Client{
		Setting: setting}}
	if setting.UserId != nil {
		a.UserId = *setting.UserId
	}
	a.key.PrivateKey = setting.SM2.PrivateKey
	a.key.PublicKey = setting.SM2.PublicKey
	return a
}

/*
 * 执行一系列流程：
 * 1、构造请求参数（含header)
 * 2、发送Http请求
 * 3、解析请求报文
 */
func (p *client) Execute() {
	if p.Request.Method == nil || len(*p.Request.Method) < 0 {
		p.Request.Method = utils.PString(http.GET)
	}
	if p.Request.Body == nil {
		p.SetBody(make(model.BodyMap))
	}
	logger.HeliLogger.Infof("Protected:%+v", p.Request.Protected)
	if p.Request.Protected == nil {
		p.SetProtected(make(model.BodyMap))
	}
	if p.Request.Path == nil {
		logger.HeliLogger.Error("ERROR:Path is null..")
		p.Err = utils.Err("Path is null..")
		return
	}

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)

	row, err := p.requestParams()
	if err != nil {
		logger.HeliLogger.Error("ERROR:", err.Error())
		p.Err = err
		return
	}

	// 将报文添加到请求url中
	//for key := range row {
	//	p.HttpReq.QueryParams.Add(key, row.Get(key))
	//}
	p.HttpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	http.WithRequestType(http.TypeForm)(p.HttpReq)
	p.HttpReq.Body = row

	if p.Err = p.Client.Execute(); p.Err != nil {
		logger.HeliLogger.Error("ERROR:", p.Err.Error())
		return
	}

	if p.Response.HttpStatus != 200 {
		p.Err = errors.New(string(p.HttpReq.Result))
		return
	}

	response, err := p.responseParams()
	if err != nil {
		logger.HeliLogger.Error("ERROR:", err.Error())
		p.Err = err
		return
	}
	logger.HeliLogger.Info(response)
	result := new(Response)
	//data, _ := json.Marshal(response["body"])
	_ = json.Unmarshal(p.HttpReq.Result, &result)
	//fmt.Println("data:", string(data))
	//fmt.Printf("\nresult:%+v\n", result)

	p.Response.Response.Code = result.Rt5RetCode           //返回码
	p.Response.Response.Message = result.Rt6RetMsg         //响应信息
	p.Response.Response.RequestId = result.Rt9SerialNumber //平台流水号
	p.Response.Response.Data = p.HttpReq.Result
	//p.Response.Response.Data = result.MsgPrivate
	//p.Response.Response.Result = json.RawMessage(response["body"].(string))
	////fmt.Printf("\np.Response.Response:%+v\n", p.Response.Response)
	//
	//if p.Response.Response.Code == "0000" {
	//	p.Response.Success = true
	//}
	//
	//if p.Response.HttpStatus != 200 {
	//	p.Response.Success = false
	//}
	p.Response.Response.Result = p.HttpReq.Result

	if p.Response.Response.Code == "0000" {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

/*
 *返回body数据,并且包含签名及签名算法信息
 */
func (p *client) requestParams() (model.BodyMap, error) {
	body := model.BodyMap{}
	body = p.Request.Body

	var keys []string

	//获取Key,并重排序
	for k, _ := range body {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//获取交易类型P1_bizType，如MerchantSettlement的P6_notifyUrl不参与签名
	bizType := body.Get("P1_bizType")

	data := bytes.Buffer{}
	for _, v := range keys {
		if bizType == BIZ_TYPE_MS && v == "P6_notifyUrl" {
			continue
		}
		vv := body.Get(v)
		data.WriteString(fmt.Sprintf("%s%s", "&", vv))
	}

	signData := data.String()

	//签名
	if bizType != BIZ_TYPE_MAQ {
		t1, err := p.key.Sign(signData)
		if err != nil {
			logger.CmbcLogger.Error("ERROR:", err.Error())
			return nil, err
		}
		fmt.Println(t1)
		body.Set("sign", t1)
	} else {
		//商户余额查询用md5算法签名(增加商户密钥)
		signData = fmt.Sprintf("%s&%s", signData, p.key.MerchantKey)
		t1, err := p.key.SignWithMD5(signData)
		if err != nil {
			logger.CmbcLogger.Error("ERROR:", err.Error())
			return nil, err
		}
		fmt.Println(t1)
		body.Set("sign", t1)
	}

	fmt.Println(body)

	return body, nil
}

func (p *client) responseParams() (model.BodyMap, error) {
	var signature string

	row := model.BodyMap{}
	_ = json.Unmarshal(p.HttpReq.Result, &row)

	//row包含所有返回数据信息，从中可以拿到signature
	if row["sign"] != nil {
		signature = row["sign"].(string)
		//fmt.Printf("签名:%v\n", signature)
	} else {
		fmt.Println("无需验签")
		return row, nil
	}

	var keys []string

	//获取Key,并重排序
	for k, _ := range row {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//调整前顺序为(rt10,rt11, ..., rt1, rt2..., rt9)调整后顺序(rt1,rt2,...,rt10,rt11...)
	idx := 0
	for _, v := range keys {
		if vv := strings.Contains(v, "rt1_"); vv {
			break
		}
		idx++
	}
	if idx > 0 {
		keys = append(keys[idx:], keys[:idx]...)
	}
	//获取交易类型P1_bizType，如MerchantSettlement的P6_notifyUrl不参与签名
	bizType := row.Get("P1_bizType")

	data := bytes.Buffer{}
	for _, v := range keys {
		if v == "sign" {
			continue
		}
		//交易类型为MerchantSettlement的rt3_retMsg字段不参与验签
		if bizType == BIZ_TYPE_MS && v == "rt3_retMsg" {
			continue
		}
		vv := row.Get(v)
		data.WriteString(fmt.Sprintf("%s%s", "&", vv))
	}
	if bizType != BIZ_TYPE_MAQ {
		if !p.key.Verify(data.String(), signature) {
			logger.HeliLogger.Error("ERROR:验签失败")
			return nil, errors.New("验签失败")
		}
	} else {
		signData := fmt.Sprintf("%s&%s", data.String(), p.key.MerchantKey)
		if !p.key.VerifyWithMD5(signData, signature) {
			logger.HeliLogger.Error("ERROR:验签（md5)失败")
			return nil, errors.New("验签失败")
		}
	}

	return row, nil
}

func (p *client) urlParse() string {

	if p.Request.Path == nil {
		panic("apiPath is void...")
	}

	return fmt.Sprintf("%s%s", *p.Setting.ServerUrl, *p.Request.Path)
}

//func GetPrivateKey(privateKeyName, privatePassword string) (*rsa.PrivateKey, error) {
//	f, err := os.Open(privateKeyName)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := ioutil.ReadAll(f)
//	if err != nil {
//		return nil, err
//	}
//
//	prikey, certs, err := pkcs12.DecodeAll(bytes, privatePassword)
//	if err != nil {
//		return nil, err
//	}
//
//	fmt.Println(prikey)
//	fmt.Println(certs)
//	return nil, nil
//}

/*
 *直接根据body数据验证签名有效性
 */

func (p *client) VerifySign(row model.BodyMap) (bool, error) {
	var signature string

	//row包含所有返回数据信息，从中可以拿到signature
	if row["sign"] != nil {
		signature = row["sign"].(string)
		//fmt.Printf("签名:%v\n", signature)
	} else {
		fmt.Println("无需验签")
		return true, nil
	}

	var keys []string

	//获取Key,并重排序
	for k, _ := range row {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//调整前顺序为(rt10,rt11, ..., rt1, rt2..., rt9)调整后顺序(rt1,rt2,...,rt10,rt11...)
	idx := 0
	for _, v := range keys {
		if vv := strings.Contains(v, "rt1_"); vv {
			break
		}
		idx++
	}
	if idx > 0 {
		keys = append(keys[idx:], keys[:idx]...)
	}

	data := bytes.Buffer{}
	for _, v := range keys {
		if v == "sign" {
			continue
		}
		vv := row.Get(v)
		data.WriteString(fmt.Sprintf("%s%s", "&", vv))
	}

	if !p.key.Verify(data.String(), signature) {
		logger.HeliLogger.Error("ERROR:验签失败")
		return false, errors.New("验签失败")
	}

	return true, nil
}
