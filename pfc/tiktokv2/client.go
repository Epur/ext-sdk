package tiktokv2

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	http2 "net/http"
	"sort"
	"strings"
	"time"
)

type Client struct {
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{Setting: setting}}
}

// v202309 API调用参考地址：https://partner.tiktokshop.com/docv2/page/64f199679495ef0281851ee5
// 含通用错误码说明
func (p *Client) Execute() {

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

	//query 参数
	p.Request.Params.Set("timestamp", fmt.Sprintf("%d", time.Now().Unix())).
		Set("app_key", *p.Setting.Key)

	if *p.Request.Path != GETACCESS && *p.Request.Path != REFRESHTOKEN {
		p.Request.Params.Set("sign", p.sign())
	}

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)

	//请求头增加访问令牌
	if p.Setting.AccessToken != nil && *p.Client.Request.Path != GETACCESS && *p.Client.Request.Path != REFRESHTOKEN {
		p.HttpReq.Header.Set("x-tts-access-token", *p.Setting.AccessToken)
	}

	for key, value := range p.Request.Params {
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
	p.Response.Response.Message = result.Message
	p.Response.Response.RequestId = result.RequestId
	p.Response.Response.Data = result.Response

	if result.Code == 0 {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

/*
**参考地址：https://partner.tiktokshop.com/docv2/page/64f199709495ef0281851fd0
**增加body，参与加签
 */
func (p *Client) sign() string {

	keys := []string{}
	union := map[string]string{}
	//获取contentType类型
	mediaType := p.HttpReq.Header.Get("Content-type")
	for key, val := range p.Client.Request.Params {
		if key == "access_token" || key == "sign" {
			continue
		}
		union[key] = val.(string)
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("%s%s", *p.Client.Setting.Secret, *p.Client.Request.Path))
	for _, key := range keys {
		message.WriteString(fmt.Sprintf("%s%s", key, union[key]))
	}
	//增加body，在不为form-data情况下
	if mediaType != "multipart/form-data" {
		if p.Request.Body != nil {
			body, _ := json.Marshal(p.Request.Body)
			message.WriteString(string(body))
		}
	}
	message.WriteString(*p.Client.Setting.Secret)
	msg := message.String()

	signedStr, _ := p.generateSHA256(msg)
	//return hex.EncodeToString(hash.Sum(nil))
	return signedStr
}

func (p *Client) generateSHA256(input string) (string, error) {
	hash := hmac.New(sha256.New, []byte(*p.Client.Setting.Secret))
	if _, err := hash.Write([]byte(input)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (p *Client) urlParse() string {
	if *p.Client.Request.Path == GETACCESS || *p.Client.Request.Path == REFRESHTOKEN {
		return fmt.Sprintf("%s%s", AUTHSITE, *p.Client.Request.Path)
	} else {
		return fmt.Sprintf("%s%s", SERVER_URl, *p.Client.Request.Path)
	}
}
