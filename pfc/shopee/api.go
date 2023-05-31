package shopee

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"net/url"
	"strconv"
	"time"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	api := &Api{Setting: setting}
	if api.Setting.ServerUrl == nil || len(*api.Setting.ServerUrl) <= 0 {
		api.Setting.SetServerUrl("https://partner.shopeemobile.com")
	}
	return api
}

func (p *Api) GetAuthUrl(callbackParams string) string {

	c := NewClient(p.Setting)
	c.SetPath(AUTH_PARTNER)
	c.timestamp = strconv.FormatInt(time.Now().Unix(), 10)

	return fmt.Sprintf("%s%s?%s", *p.Setting.ServerUrl, AUTH_PARTNER, model.BodyMap{}.
		Set("partner_id", *p.Setting.Key).
		Set("timestamp", c.timestamp).
		Set("sign", c.sign()).
		Set("redirect", fmt.Sprintf("%s?callbackParams=%s", *p.Setting.AuthCallbackUrl, url.QueryEscape(callbackParams))).EncodeURLParams())
}

/*
	获取Token
	Url : https://open.lazada.com/apps/doc/api?path=%2Fauth%2Ftoken%2Fcreate
	Response: Response
*/

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	partnerId, _ := strconv.ParseInt(*p.Setting.Key, 10, 64)
	Body.Set("partner_id", partnerId)

	c := NewClient(p.Setting)
	c.SetPath(AUTH_ACCESSTOKEN).
		SetMethod(http.POST).
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("code"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.Data = c.HttpReq.Result
	response := GetTokenResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetSeller(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(SHOP_URL).
		SetBody(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetSellerResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetMerchant(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(MERCHAT_URL).
		SetBody(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetMerchantResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {

	partnerId, _ := strconv.ParseInt(*p.Setting.Key, 10, 64)

	if p.Setting.IsMerchant {
		merchantId, _ := strconv.ParseInt(*p.Setting.MerchantId, 10, 64)
		Body.Set("partner_id", partnerId).
			Set("merchant_id", merchantId)
	} else {
		shopId, _ := strconv.ParseInt(*p.Setting.ShopId, 10, 64)
		Body.Set("partner_id", partnerId).
			Set("shop_id", shopId)
	}

	c := NewClient(p.Setting)
	c.SetPath(AUTH_REFRESHTOKEN).
		SetMethod(http.POST).
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("refresh_token"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetTokenResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) StoreRefreshToken(Body model.BodyMap) *model.Client {

	currTime := utils.TimestampSecond()

	c := p.RefreshToken(Body)

	if c.Response.Response.DataTo != nil {
		response := c.Response.Response.DataTo.(GetTokenResponse)
		c.Response.Response.DataTo = model.StoreTokenResponse{
			AccessToken:       response.AccessToken,
			AccessTokenExpire: response.ExpireIn + currTime,
			RefreshToken:      response.RefreshToken,
		}
	}

	return c
}

func (p *Api) GetOrderDetail(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(ORDER_DETAIL).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("order_sn_list"); c.Err != nil {
		return &c.Client
	}

	if len(Body.Get("response_optional_fields")) <= 0 {
		c.Client.Request.Params.Set("response_optional_fields", c.getOrderFields())
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetOrderDetailResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}
