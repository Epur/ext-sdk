package go1688

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"sort"
	"strings"
)

type Client struct {
	model.Client
	PathUrl string
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{Setting: setting}}
}

func (p *Client) Execute() {

	if p.Request.Method == nil || len(*p.Request.Method) < 0 {
		p.Request.Method = utils.PString(http.GET)
	}
	if p.Request.Params == nil {
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
		http.WithUrl(p.urlParse(p.Request.Path)),
		http.WithMethod(*p.Request.Method),
	)

	if p.Setting.AccessToken != nil && *p.Request.Path != AccessTokenURL {
		p.Request.Params.Set("access_token", *p.Setting.AccessToken)
	}
	if *p.Request.Path != AccessTokenURL {
		p.Request.Params.Set("_aop_signature", p.sign())
	}

	for key, value := range p.Request.Params {
		p.HttpReq.QueryParams.Add(key, value.(string))
	}

	//p.Request.Body.Set("timestamp", fmt.Sprintf("%d", time.Now().UnixMilli())).
	//	Set("sign", p.sign()).
	//	Set("sign_method", p.SignMethod).
	//	Set("app_key", *p.Setting.Key)

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	fmt.Println(string(p.HttpReq.Result))

	result := new(Response)
	_ = json.Unmarshal(p.HttpReq.Result, &result)

	if len(result.Error) > 0 {
		p.Response.Response.Code = result.Error
		p.Response.Response.Message = result.ErrorDescription
		p.Response.Response.RequestId = result.RequestId
	} else if len(result.ErrorCode) > 0 {
		p.Response.Response.Code = result.ErrorCode
		p.Response.Response.Message = result.ErrorMessage
		p.Response.Response.RequestId = result.RequestId
	} else if result.Success != nil && !*result.Success {
		p.Response.Response.Code = result.Message
		p.Response.Response.Message = result.Message
	} else {
		p.Response.Success = true
	}

	p.Response.Response.Data = p.HttpReq.Result

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

func (p *Client) urlParse(apiPath *string) string {

	if apiPath == nil {
		panic("apiPath is void...")
	}

	if *apiPath == AccessTokenURL {
		return fmt.Sprintf("%s%s/%s", *p.Setting.ServerUrl, *apiPath, *p.Setting.Key)
	} else {
		p.handleURI()
		return fmt.Sprintf("%s/%s", *p.Setting.ServerUrl, p.PathUrl)
	}

}

func (p *Client) sign() string {

	/*
		签名
	*/
	var strs []string
	for k := range p.Request.Params {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	var signParamsStr string

	for _, k := range strs {
		signParamsStr += k + p.Request.Params.Get(k)
	}
	//fmt.Println(p.PathUrl + signParamsStr)

	return strings.ToUpper(p.HmacSHA1(*p.Setting.Secret, p.PathUrl+signParamsStr))
}

func (p *Client) handleURI() {
	split := strings.Split(*p.Request.Path, ":")
	spacename := split[0]
	split = strings.Split(split[1], "-")
	apiname := split[0]
	version := split[1]
	p.PathUrl = fmt.Sprintf("param2/%s/%s/%s/%s", version, spacename, apiname, *p.Client.Setting.Key)
}

func (p *Client) HmacSHA1(key string, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
