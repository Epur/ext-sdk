package amazon

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	http2 "net/http"
	"runtime"
	"strings"
)

type Client struct {
	SignMethod string
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{
		Setting: setting}, SignMethod: "HmacSHA256"}
}

func (c *Client) Execute() {
	if c.Request.Method == nil || len(*c.Request.Method) < 0 {
		c.Request.Method = utils.PString(http.GET)
	}
	if c.Request.Body == nil {
		c.SetBody(make(model.BodyMap))
	}
	if c.Setting.Key == nil {
		c.Err = utils.Err("Key is null..")
		return
	}
	if c.Request.Path == nil {
		c.Err = utils.Err("Path is null..")
		return
	}

	c.HttpReq = http.New(
		http.WithUrl(c.urlParse()),
		http.WithMethod(*c.Request.Method),
	)
	c.HttpReq.Header.Set("User-Agent", fmt.Sprintf(
		"ext-sdk-amz/v1.0.0 (Language=%s; Platform=%s-%s)",
		strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH))

	if c.Setting.AccessToken != nil && *c.Request.Path != TokenURL {
		c.HttpReq.Header.Set("X-Amz-Access-Token", *c.Setting.AccessToken)
	}

	for key, value := range c.Request.Params {
		// 使用类型 switch
		switch v := value.(type) {
		//数组类型
		case []string:
			for _, vv := range v {
				c.HttpReq.QueryParams.Add(key, vv)
			}
		default:
			c.HttpReq.QueryParams.Set(key, value.(string))
		}
	}

	if strings.ToUpper(*c.Request.Method) == http2.MethodPost ||
		strings.ToUpper(*c.Request.Method) == http2.MethodPut {
		http.WithRequestType(http.TypeJSON)(c.HttpReq)
		c.HttpReq.Body = c.Request.Body
	}

	if c.Err = c.Client.Execute(); c.Err != nil {
		return
	}

	result := new(Response)
	_ = json.Unmarshal(c.HttpReq.Result, &result)

	code := "Success"
	msg := "success"
	if result.Errors != nil {
		code = result.Errors[0].Code
		msg = result.Errors[0].Message
	}

	c.Response.Response.Code = code
	c.Response.Response.Message = msg
	c.Response.Response.RequestId = ""
	c.Response.Response.Data = result.Payload

	if result.Errors == nil {
		c.Response.Success = true
	}

	if c.Response.HttpStatus != 200 {
		c.Response.Success = false
	}
}

func (c *Client) urlParse() string {
	urlPath := ""
	if strings.HasPrefix(*c.Request.Path, "http") {
		urlPath = *c.Request.Path
	} else {
		urlPath = c.getServerURL() + *c.Request.Path
	}
	return urlPath
}

func (c *Client) getServerURL() string {
	siteNo := c.Setting.SiteNo
	api := APIGatewayUS
	if siteNo == nil {
		return api
	}
	switch strings.ToUpper(*siteNo) {
	case "US", "MX", "CA", "BR":
		api = APIGatewayUS
	case "ES", "UK", "GB", "FR", "BE", "NL", "DE", "IT", "SE", "ZA", "PL", "EG", "TR", "SA", "AE", "IN":
		api = APIGatewayEU
	case "SG", "AU", "JP":
		api = APIGatewayFE
	case "refresh":
		api = TokenURL
	}
	return api
}
