package lazada

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
	"sort"
	"strings"
	"time"
)

type Client struct {
	SignMethod string
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{
		Setting: setting}, SignMethod: "sha256"}
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

	if p.Setting.AccessToken != nil && *p.Request.Path != AccessTokenURL && *p.Request.Path != RefreshURL {
		p.Request.Body.Set("access_token", *p.Setting.AccessToken)
	}
	p.Request.Body.Set("timestamp", fmt.Sprintf("%d", time.Now().UnixMilli())).
		Set("sign", p.sign()).
		Set("sign_method", p.SignMethod).
		Set("app_key", *p.Setting.Key)

	if *p.Request.Method == http.GET {
		for key, value := range p.Request.Body {
			p.HttpReq.QueryParams.Add(key, value.(string))
		}
	} else {
		http.WithRequestType(http.TypeJSON)(p.HttpReq)
		p.HttpReq.Body = p.Request.Body
	}
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

	if result.Code == "0" {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

func (p *Client) getServerURL() string {
	switch *p.Setting.SiteNo {
	case "sg":
		return APIGatewaySG
	case "my":
		return APIGatewayMY
	case "vn":
		return APIGatewayVN
	case "th":
		return APIGatewayTH
	case "ph":
		return APIGatewayPH
	case "id":
		return APIGatewayID
	case "refresh":
		return APIREFRESHURL
	}
	return APICODEURL
}

func (p *Client) urlParse() string {

	var apiServerURL string

	if p.Request.Path == nil {
		panic("apiPath is void...")
	}

	if *p.Request.Path == RefreshURL || *p.Request.Path == AccessTokenURL {
		apiServerURL = APIREFRESHURL
	} else {
		apiServerURL = p.getServerURL()
	}
	return fmt.Sprintf("%s%s", apiServerURL, *p.Request.Path)
}

func (p *Client) sign() string {

	/*
		签名
	*/

	keys := []string{"sign_method", "app_key"}
	union := map[string]string{
		"sign_method": p.SignMethod,
		"app_key":     *p.Setting.Key,
	}

	for key, val := range p.Request.Body {
		if key == "image" {
			continue
		}
		union[key] = val.(string)
		keys = append(keys, key)
	}

	// sort sys params and api params by key
	sort.Strings(keys)

	var message bytes.Buffer
	message.WriteString(fmt.Sprintf("%s", *p.Request.Path))
	for _, key := range keys {
		message.WriteString(fmt.Sprintf("%s%s", key, union[key]))
	}

	//fmt.Println(message.String())

	hash := hmac.New(sha256.New, []byte(*p.Setting.Secret))
	hash.Write(message.Bytes())
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}
