package amazon

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/tangchen2018/go-utils/http"
	"net/url"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

func (p *Api) GetAuthUrl(callbackParams string) string {
	return fmt.Sprintf("%s?%s", AuthURL, model.BodyMap{}.
		Set("application_id", *p.Setting.DevelopId).
		Set("version", "beta").
		Set("state", url.QueryEscape(callbackParams)).
		Set("redirect_uri", *p.Setting.AuthCallbackUrl).EncodeURLParams())
}

func (p *Api) GetToken(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	Body.Set("grant_type", "authorization_code")
	Body.Set("client_id", *p.Setting.Key)
	Body.Set("client_secret", *p.Setting.Secret)
	Body.Set("redirect_uri", *p.Setting.ServerUrl)
	c.SetPath(TokenURL).
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("code", "redirect_uri"); c.Err != nil {
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
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	Body.Set("grant_type", "refresh_token")
	Body.Set("client_id", *p.Setting.Key)
	Body.Set("client_secret", *p.Setting.Secret)
	c.SetPath(TokenURL).
		SetMethod(http.POST).
		SetParams(Body)

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
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) StoreRefreshToken(Body model.BodyMap) *model.Client {
	//TODO implement me
	panic("implement me")
}

func (p *Api) GetSeller() *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("/sellers/v1/marketplaceParticipations").
		SetMethod(http.GET)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetSellerResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetOrder(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf("/orders/v0/orders/%s", Body.Get("orderId"))).
		SetMethod(http.GET)

	if c.Err = Body.CheckEmptyError("orderId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetOrderResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetOrderList(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("/orders/v0/orders").
		SetMethod(http.GET).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("MarketplaceIds"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetOrderListResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetOrderDetail(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf("/orders/v0/orders/%s", Body["orderId"])).
		SetMethod(http.GET).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("orderId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetOrderDetailResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}
