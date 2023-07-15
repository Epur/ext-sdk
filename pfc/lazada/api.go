package lazada

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
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
		Set("response_type", "code").
		Set("client_id", *p.Setting.Key).
		Set("redirect_uri", fmt.Sprintf("%s?callbackParams=%s", *p.Setting.AuthCallbackUrl, url.QueryEscape(callbackParams))).EncodeURLParams())
}

/*
	获取Token
	Url : https://open.lazada.com/apps/doc/api?path=%2Fauth%2Ftoken%2Fcreate
	Response: Response
*/

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(AccessTokenURL).
		//SetMethod(http.POST).
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

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(RefreshURL).
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
			AccessToken:        response.AccessToken,
			AccessTokenExpire:  response.ExpiresIn + currTime,
			RefreshToken:       response.RefreshToken,
			RefreshTokenExpire: response.RefreshExpiresIn + currTime,
		}
	}

	return c
}

func (p *Api) GetSeller(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(SELLERURL).
		SetBody(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetSellerResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	获取订单信息
	Url : https://open.lazada.com/apps/doc/api?path=%2Forder%2Fget
	Response: GetOrderDetailResponse
*/

func (p *Api) GetOrder(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/order/get`).
		//SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("order_id"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetOrderResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	获取订单列表
	Url : https://open.lazada.com/apps/doc/api?path=%2Forders%2Fget
	Response: GetOrderDetailResponse
*/

// 获取订单列表

func (p *Api) GetOrderList(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/orders/get`).
		SetBody(Body)
	result := GetOrderListResponse{}

	offset := 0

	for {
		//Body.Set("offset", fmt.Sprintf("%d", offset))

		c.Request.Body.Set("offset", fmt.Sprintf("%d", offset))
		cResult := getOrderListResponse{}

		//fmt.Println("请求->", Body)

		c.Execute()
		if c.Err != nil {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		if cResult.OrderList != nil && len(cResult.OrderList) > 0 {
			for index := range cResult.OrderList {
				result.List = append(result.List, cResult.OrderList[index])
			}
		}

		//fmt.Println("cResult.OrderList->", cResult.OrderList)

		offset += cResult.Count

		if cResult.CountTotal <= offset {
			break
		}

		if cResult.Count <= 0 || len(cResult.OrderList) <= 0 {
			result.Total = cResult.CountTotal
			break
		}
	}

	c.Response.Response.DataTo = result

	return &c.Client
}

/*
	获取订单明细
	Url : https://open.lazada.com/apps/doc/api?path=%2Forder%2Fitems%2Fget
	Response: GetOrderDetailResponse
*/

func (p *Api) GetOrderDetail(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/order/items/get`).
		//SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("order_id"); c.Err != nil {
		return &c.Client
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
