package go1688

import (
	"encoding/json"
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"strconv"
	"strings"
	"time"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	api := &Api{Setting: setting}
	if api.Setting.ServerUrl == nil || len(*api.Setting.ServerUrl) <= 0 {
		api.Setting.SetServerUrl("https://gw.open.1688.com/openapi")
	}
	return api
}

func (p *Api) GetAuthUrl(callbackParams string) string {

	return fmt.Sprintf("%s?%s", AuthURL, model.BodyMap{}.
		Set("client_id", *p.Setting.Key).
		Set("redirect_uri", fmt.Sprintf("%s?callbackParams=%s", *p.Setting.AuthCallbackUrl, callbackParams)).
		Set("site", "1688").
		Set("state", "state").EncodeURLParams(),
	)
}

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	Body.Set("grant_type", "authorization_code").
		Set("need_refresh_token", "true").
		Set("client_id", *p.Setting.Key).
		Set("client_secret", *p.Setting.Secret)

	c := NewClient(p.Setting)
	c.SetPath(AccessTokenURL).
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("code"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	if c.Response.Success {
		c.Client.Response.Response.Data = c.HttpReq.Result

		tmpResponse := map[string]interface{}{}
		if c.Err = json.Unmarshal(c.HttpReq.Result, &tmpResponse); c.Err != nil {
			return &c.Client
		}

		response := GetTokenResponse{
			AliId:         tmpResponse["aliId"].(string),
			AccessToken:   tmpResponse["access_token"].(string),
			RefreshToken:  tmpResponse["refresh_token"].(string),
			ResourceOwner: tmpResponse["resource_owner"].(string),
			MemberId:      tmpResponse["memberId"].(string),
		}
		response.ExpiresIn, _ = strconv.ParseInt(tmpResponse["expires_in"].(string), 10, 64)

		aTmp := strings.Split(tmpResponse["refresh_token_timeout"].(string), "+")
		location, _ := time.LoadLocation("Asia/Shanghai")
		s, _ := time.ParseInLocation(fmt.Sprintf("20060102150405000+%s", aTmp[1]),
			tmpResponse["refresh_token_timeout"].(string), location)
		response.RefreshExpiresIn = s.UnixMilli() / 1000

		c.Response.Response.DataTo = response
	}

	return &c.Client
}

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {

	Body.Set("grant_type", "refresh_token").
		Set("client_id", *p.Setting.Key).
		Set("client_secret", *p.Setting.Secret)

	c := NewClient(p.Setting)
	c.SetPath(AccessTokenURL).
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("refresh_token"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	fmt.Println(c.Response.Success)

	if c.Response.Success {

		c.Client.Response.Response.Data = c.HttpReq.Result

		tmpResponse := map[string]interface{}{}
		if c.Err = json.Unmarshal(c.HttpReq.Result, &tmpResponse); c.Err != nil {
			return &c.Client
		}

		response := GetTokenResponse{
			AliId:         tmpResponse["aliId"].(string),
			AccessToken:   tmpResponse["access_token"].(string),
			ResourceOwner: tmpResponse["resource_owner"].(string),
			MemberId:      tmpResponse["memberId"].(string),
		}
		response.ExpiresIn, _ = strconv.ParseInt(tmpResponse["expires_in"].(string), 10, 64)
		c.Response.Response.DataTo = response
	}

	return &c.Client
}

func (p *Api) StoreRefreshToken(Body model.BodyMap) *model.Client {

	currTime := utils.TimestampSecond()

	c := p.RefreshToken(Body)

	if c.Response.Response.DataTo != nil {
		response := c.Response.Response.DataTo.(GetTokenResponse)
		c.Response.Response.DataTo = model.StoreTokenResponse{
			AccessToken:       response.AccessToken,
			AccessTokenExpire: response.ExpiresIn + currTime,
			//RefreshToken:      response.RefreshToken,
			//RefreshTokenExpire: response.RefreshExpiresIn + currTime,
		}
	}

	return c
}

func (p *Api) PushedProductList(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.product.push:alibaba.cross.syncProductListPushed-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("productIdList"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

func (p *Api) GetProductList(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.product:alibaba.cross.productList-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("productIdList"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}
