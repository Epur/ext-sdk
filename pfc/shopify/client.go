package shopify

import (
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/pfc/shopify/constant"
	"github.com/Epur/ext-sdk/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/tangchen2018/go-utils/http"
	http2 "net/http"
	"strings"
)

type client struct {
	model.Client
}

func NewClient(setting *model.Setting) *client {
	return &client{Client: model.Client{Setting: setting}}
}

func (p *client) VerifySessionToken(header map[string]interface{}, payload string, signature string) (bool, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["header"] = header
	atClaims["payload"] = payload
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(*p.Setting.Key))
	if err != nil {
		return false, err
	}
	if strings.Compare(token, signature) != 0 {
		return false, errors.New("令牌验证失败")
	}

	return true, nil
}

// v202309 API调用参考地址：https://partner.tiktokshop.com/docv2/page/64f199679495ef0281851ee5
// 含通用错误码说明
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
	p.HttpReq.Header.Set("Content-Type", "application/json")

	if *p.Request.Path != constant.GETACCESS && *p.Request.Path != constant.AUTH {
		//p.Request.Params.Set("sign", p.sign())
		//设置访问令牌
		if strings.Compare(strings.Trim(p.HttpReq.Header.Get("X-Shopify-Access-Token"), " "), "") == 0 {
			p.HttpReq.Header.Set("X-Shopify-Access-Token", *p.Setting.AccessToken)
		}

	}

	for key, value := range p.Request.Params {
		p.HttpReq.QueryParams.Set(key, value.(string))
	}

	if strings.ToUpper(*p.Request.Method) == http2.MethodPost ||
		strings.ToUpper(*p.Request.Method) == http2.MethodPut {
		http.WithRequestType(http.TypeJSON)(p.HttpReq)
		p.HttpReq.Body = p.Request.Body
	}

	//if p.Err = p.Client.Execute(); p.Err != nil {
	//	return
	//}
	//
	//result := new(Response)
	//_ = json.Unmarshal(p.HttpReq.Result, &result)
	//
	//p.Response.Response.Code = fmt.Sprintf("%d", result.Code)
	//p.Response.Response.Message = result.Message
	//p.Response.Response.RequestId = result.RequestId
	//p.Response.Response.Data = result.Response
	//
	//if result.Code == 0 {
	//	p.Response.Success = true
	//}
	//// 新增错误返回
	//if result.Code != 0 {
	//	p.Err = utils.Err(result.Message)
	//}
	//
	//if p.Response.HttpStatus != 200 {
	//	p.Response.Success = false
	//}
}

/*SERVER_URL:https://%s.myshopify.com
**其中：{shop}是用户的商店名称
 */
func (p *client) urlParse() string {
	return fmt.Sprintf("%s%s", fmt.Sprintf(constant.SERVER_URl, *p.Setting.Shop), *p.Client.Request.Path)
}
