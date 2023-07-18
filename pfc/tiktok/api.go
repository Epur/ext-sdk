package tiktok

import (
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"net/url"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

/*
	获取授权链接
	callbackParams : 同步回调的自定义参数
*/
func (p *Api) GetAuthUrl(callbackParams string) string {

	return fmt.Sprintf("%s%s?%s", AUTHSITE, AUTH, model.BodyMap{}.
		Set("app_key", *p.Setting.Key).
		Set("state", url.QueryEscape(callbackParams)).EncodeURLParams())
}

/*
	获取Token
	Url : https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg
	Response: GetTokenResponse
*/

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	Body.Set("grant_type", "authorized_code")

	c := NewClient(p.Setting)
	c.SetPath(GETACCESS).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("auth_code"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetTokenResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	刷新令牌
	Url : https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg
	Response: GetTokenResponse
*/

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(REFRESHTOKEN).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("refresh_token"); c.Err != nil {
		return &c.Client
	}

	c.Request.Params.Set("grant_type", "refresh_token").
		Set("app_secret", *p.Setting.Secret)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetTokenResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) StoreRefreshToken(Body model.BodyMap) *model.Client {

	c := p.RefreshToken(Body)

	if c.Response.Response.DataTo != nil {
		response := c.Response.Response.DataTo.(GetTokenResponse)
		c.Response.Response.DataTo = model.StoreTokenResponse{
			AccessToken:        response.AccessToken,
			AccessTokenExpire:  response.AccessTokenExpireIn,
			RefreshToken:       response.RefreshToken,
			RefreshTokenExpire: response.RefreshTokenExpireIn,
		}
	}

	return c
}

/*
	获取订单列表
	Url : https://partner.tiktokshop.com/doc/page/262815?external_id=262815
	Response: GetOrderListResponse
*/

func (p *Api) GetOrderList(Body model.BodyMap) *model.Client {

	var cursor *string
	c := NewClient(p.Setting)
	c.SetPath(`/api/orders/search`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("page_size"); c.Err != nil {
		return &c.Client
	}

	result := GetOrderListResponse{}

	for {

		if cursor != nil && len(*cursor) > 0 {
			c.Request.Body.Set("cursor", cursor)
		}

		cResult := getOrderListResponse{}

		c.Execute()
		if c.Err != nil || !c.Response.Success {
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

		if cResult.More == false {
			result.Total = cResult.Total
			break
		} else {
			cursor = &cResult.NextCursor
		}
	}

	c.Response.Response.DataTo = result

	return &c.Client
}

/*
	获取订单明细
	Url : https://partner.tiktokshop.com/doc/page/262814?external_id=262814
	Response: GetOrderDetailResponse
*/

func (p *Api) GetOrderDetail(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/api/orders/detail/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("order_id_list"); c.Err != nil {
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

/*
	获取站点产品列表
	Url : https://partner.tiktokshop.com/doc/page/262788?external_id=262788
	Response: GetOrderDetailResponse
*/
func (p *Api) GetProductList(Body model.BodyMap) *model.Client {

	/*
		"1": "Published",
		"2": "Created",
		"3": "Draft",
		"4": "Deleted"
	*/

	c := NewClient(p.Setting)
	c.SetPath(`/api/products/search`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("page_number"); c.Err != nil {
		return &c.Client
	}

	result := GetProductListResponse{}
	page := 1

	for {

		c.Request.Body.Set("page_size", fmt.Sprintf("%d", page))

		cResult := GetProductListResponse{}

		c.Execute()
		if c.Err != nil || !c.Response.Success {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		if len(cResult.List) > 0 {
			for index := range cResult.List {
				result.List = append(result.List, cResult.List[index])
			}
		}

		page++

		fmt.Println(page, len(cResult.List), cResult.Total)

		if len(result.List) >= cResult.Total {
			break
		}
	}

	c.Response.Response.DataTo = result

	return &c.Client
}
