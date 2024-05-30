package amazon

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/tangchen2018/go-utils/http"
	"net/url"
	"time"
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
	return p.token(Body, "refresh_token")
}

func (p *Api) NoToken() *model.Client {
	return p.token(model.BodyMap{}, "client_credentials")
}

func (p *Api) token(Body model.BodyMap, grantType string) *model.Client {
	c := NewClient(p.Setting)
	Body.Set("grant_type", grantType)
	Body.Set("client_id", *p.Setting.Key)
	Body.Set("client_secret", *p.Setting.Secret)
	if grantType == "refresh_token" {
		if c.Err = Body.CheckEmptyError("refresh_token"); c.Err != nil {
			return &c.Client
		}
	} else if grantType == "client_credentials" {
		Body.Set("scope", "sellingpartnerapi::notifications")
	}
	c.SetPath(TokenURL).
		SetMethod(http.POST).
		SetParams(Body)

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
	c := p.RefreshToken(Body)
	if c.Response.Response.DataTo != nil {
		response := c.Response.Response.DataTo.(GetTokenResponse)
		c.Response.Response.DataTo = model.StoreTokenResponse{
			AccessToken:        response.AccessToken,
			AccessTokenExpire:  time.Now().Unix() + response.ExpiresIn,
			RefreshToken:       response.RefreshToken,
			RefreshTokenExpire: time.Now().Unix() + 365*24*60*60,
		}
	}
	return c
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

func (p *Api) GetOrders(Body model.BodyMap) *model.Client {
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
	response := GetOrdersResponse{}
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

func (p *Api) GetOrderItems(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf("/orders/v0/orders/%s/orderItems", Body.Get("orderId"))).
		SetMethod(http.GET)

	if c.Err = Body.CheckEmptyError("orderId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetOrderItemsResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetDestinations() *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf("/notifications/v1/destinations")).
		SetMethod(http.GET)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetDestinationsResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) CreateDestination(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf("/notifications/v1/destinations")).
		SetMethod(http.POST).
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("resourceSpecification", "name"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := CreateDestinationResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) CreateSubscription(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf("/notifications/v1/subscriptions/%s", Body.Get("notificationType"))).
		SetMethod(http.POST).
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("notificationType", "payloadVersion", "destinationId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := CreateSubscriptionResponse{}
	if c.Err = json.Unmarshal(c.HttpReq.Result, &response); c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.DataTo = response
	return &c.Client
}
