package tiktokv2

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
		Set("service_id", *p.Setting.Id).
		Set("state", url.QueryEscape(callbackParams)).EncodeURLParams())
}

/*
	获取Token
	Url : https://partner.tiktokshop.com/docv2/page/64f199619495ef0281851e1c
	Response: GetTokenResponse
*/

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	Body.Set("grant_type", "authorized_code")

	c := NewClient(p.Setting)
	c.SetPath(GETACCESS).
		SetParams(Body)

	//auth_code实为访问令牌
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
	Url : https://partner.tiktokshop.com/docv2/page/64f199619495ef0281851e1c
	Response: GetTokenResponse
*/

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(REFRESHTOKEN).
		SetParams(Body)

	//刷新令牌由外部调用者提供
	if c.Err = Body.CheckEmptyError("refresh_token"); c.Err != nil {
		return &c.Client
	}

	c.Request.Params.Set("grant_type", "refresh_token").
		Set("app_secret", *p.Setting.Secret).Set("app_key", *p.Setting.Key)

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
	Url : https://partner.tiktokshop.com/docv2/page/650aa8094a0bb702c06df242?external_id=650aa8094a0bb702c06df242
	Response: GetOrderListResponse
*/

func (p *Api) GetOrderList(Body model.BodyMap) *model.Client {

	var cursor *string
	c := NewClient(p.Setting)
	c.SetPath(`/order/202309/orders/search`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("page_size"); c.Err != nil {
		return &c.Client
	}

	result := GetOrderListResponse{}

	for {

		if cursor != nil && len(*cursor) > 0 {
			c.Request.Body.Set("page_token", cursor)
		}

		cResult := getOrderListResponse{}

		c.Execute()
		if c.Err != nil || !c.Response.Success {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		if cResult.Orders != nil && len(cResult.Orders) > 0 {
			for index := range cResult.Orders {
				result.List = append(result.List, cResult.Orders[index])
			}
		}

		if cResult.NextPageToken == "" {
			result.Total = cResult.TotalCount
			break
		} else {
			cursor = &cResult.NextPageToken
		}
	}

	c.Response.Response.DataTo = result

	return &c.Client
}

/*
	获取订单明细
	Url : https://partner.tiktokshop.com/docv2/page/650aa8ccc16ffe02b8f167a0?external_id=650aa8ccc16ffe02b8f167a0#Back%20To%20Top
	Response: GetOrderDetailResponse
*/

func (p *Api) GetOrderDetail(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/order/202309/orders`).
		SetMethod("GET").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("ids"); c.Err != nil {
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
Url : https://partner.tiktokshop.com/docv2/page/65854ffb8f559302d8a6acda?external_id=6503081a56e2bb0289dd6d7d
Response: GetProductListResponse
*/
func (p *Api) GetProductList(Body model.BodyMap) *model.Client {

	/*
		"1": "Published",
		"2": "Created",
		"3": "Draft",
		"4": "Deleted"
	*/

	c := NewClient(p.Setting)
	c.SetPath(`/product/202309/products/search`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("page_size"); c.Err != nil {
		return &c.Client
	}

	result := GetProductListResponse{}
	pageSize := Body.Get("page_size")

	for {

		c.Request.Params.Set("page_size", fmt.Sprintf("%s", pageSize))

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

		//page++

		//fmt.Println(page, len(cResult.List), cResult.Total)

		if len(result.List) >= cResult.Total {
			break
		}
	}

	c.Response.Response.DataTo = result

	return &c.Client
}

/*
	获取产品详情
	Url : https://partner.tiktokshop.com/docv2/page/6509d85b4a0bb702c057fdda?external_id=6509d85b4a0bb702c057fdda#Back%20To%20Top
	Response: GetProductDetailResponse
*/

func (p *Api) GetProductDetail(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("product_id"); c.Err != nil {
		return &c.Client
	}
	c.SetPath(fmt.Sprintf("/product/202309/products/%s", Body.Get("product_id")))

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetProductDetailResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}