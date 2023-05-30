package tiktok

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

	p.Request.Params.Set("timestamp", fmt.Sprintf("%d", time.Now().Unix())).
		Set("app_key", *p.Setting.Key)

	if p.Setting.AccessToken != nil && *p.Client.Request.Path != GETACCESS && *p.Client.Request.Path != REFRESHTOKEN {
		p.Request.Params.Set("access_token", *p.Setting.AccessToken)
	}

	if p.Setting.ShopId != nil && *p.Client.Request.Path != GETACCESS && *p.Client.Request.Path != REFRESHTOKEN {
		p.Request.Params.Set("shop_id", *p.Setting.ShopId)
	}

	if *p.Request.Path != GETACCESS {
		p.Request.Params.Set("sign", p.sign())
	}

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)

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

func (p *Client) sign() string {

	keys := []string{}
	union := map[string]string{}

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
	message.WriteString(*p.Client.Setting.Secret)
	msg := message.String()
	hash := hmac.New(sha256.New, []byte(*p.Client.Setting.Secret))
	if _, err := hash.Write([]byte(msg)); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func (p *Client) urlParse() string {
	if *p.Client.Request.Path == GETACCESS || *p.Client.Request.Path == REFRESHTOKEN {
		return fmt.Sprintf("%s%s", AUTHSITE, *p.Client.Request.Path)
	} else {
		return fmt.Sprintf("%s%s", SERVER_URl, *p.Client.Request.Path)
	}
}
