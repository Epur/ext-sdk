package shein

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	http2 "net/http"
	"strings"
	"time"
)

type client struct {
	model.Client
}

func NewClient(setting *model.Setting) *client {
	return &client{Client: model.Client{Setting: setting}}
}

func (p *client) Execute() {

	if p.Request.Method == nil || len(*p.Request.Method) < 0 {
		p.Request.Method = utils.PString(http.GET)
	}
	if p.Request.Params == nil {
		p.SetParams(make(model.BodyMap))
	}
	if p.Setting.Key == nil {
		p.Err = utils.Err("Key is null..")
		return
	}
	if p.Request.Path == nil {
		p.Err = utils.Err("Path is null..")
		return
	}

	// 先进行httpRep初始化
	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)
	// 目前只用到 application/json 类型 后续根据业务处理
	p.HttpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//卖家账号ID
	if p.Request.Params.Get("openKeyId") != "" {
		p.HttpReq.Header.Set("x-lt-openKeyId", p.Request.Params.Get("openKeyId")) // 访问用户隐私数据时的唯一权限标识 openKeyId
	}
	if *p.Request.Path == SELLER_SECRET {
		p.HttpReq.Header.Set("x-lt-appid", *p.Setting.Key) //开发者平台appid
	}
	//当前时间戳（毫秒）
	p.HttpReq.Header.Set("x-lt-timestamp", fmt.Sprintf("%d", time.Now().UnixMilli()))
	p.HttpReq.Header.Set("x-lt-signature", p.sign())

	for key, value := range p.Request.Params {
		//滤掉卖家id
		if key == "openKeyId" {
			continue
		}
		p.HttpReq.QueryParams.Set(key, value.(string))
	}

	if strings.ToUpper(*p.Request.Method) == http2.MethodPost ||
		strings.ToUpper(*p.Request.Method) == http2.MethodPut {
		http.WithRequestType(http.TypeJSON)(p.HttpReq)
		p.HttpReq.Body = p.Request.Body
	}

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	result := new(Response)
	_ = json.Unmarshal(p.HttpReq.Result, &result)

	p.Response.Response.Code = fmt.Sprintf("%d", result.Code)
	p.Response.Response.Message = result.Msg
	//p.Response.Response.RequestId = result.RequestId
	p.Response.Response.Data = result.Info

	if result.Code == SUCCESS {
		p.Response.Success = true
	}
	// 新增错误返回
	if result.Code != SUCCESS {
		p.Err = utils.Err(result.Msg)
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

/*签名算法：hmacSha256
**参考地址：https://open.sheincorp.com/documents/system/passwdrule
 */
func (p *client) sign() string {

	//获取contentType类型
	var message bytes.Buffer
	for idx := range SignFactors {
		v := p.HttpReq.Header.Get(SignFactors[idx])
		if strings.Compare(strings.Trim(v, " "), "") == 0 {
			continue
		}
		message.WriteString(fmt.Sprintf("%s&", v))
	}
	//增加url:从域名com后开始截取到结尾，urlPath=/open-api/order/purchase-order-infos
	message.WriteString(fmt.Sprintf("%s", *p.Request.Path))
	msg := message.String()
	//msg = msg[:len(msg)-1]

	randomKey := utils.GenerateRandomString(5)
	//return  randomKey + hex.EncodeToString(hash.Sum(nil))
	signedStr, _ := p.generateSHA256(msg, randomKey)

	//return  randomKey + hex.EncodeToString(hash.Sum(nil))
	return signedStr
}

// 参考地址：https://open.sheincorp.com/documents/system/passwdrule
func (p *client) generateSHA256(input string, randomKey string) (string, error) {
	hash := hmac.New(sha256.New, []byte(*p.Client.Setting.Secret+randomKey))
	if _, err := hash.Write([]byte(input)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", randomKey, base64.StdEncoding.EncodeToString(hash.Sum(nil))), nil
}

func (p *client) urlParse() string {
	return fmt.Sprintf("%s%s", SERVER_URL, *p.Client.Request.Path)

}
