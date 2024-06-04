package shopify

import (
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/pfc/shopify/constant"
	"github.com/Epur/ext-sdk/utils"
	"net/http"
	"net/url"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

//交互令牌即授权参考地址：https://shopify.dev/docs/apps/build/authentication-authorization
/*
**参考地址：https://shopify.dev/docs/apps/auth/get-access-tokens/token-exchange
**交互访问令牌
 */

func (p *Api) GetExchgToken(Body model.BodyMap) *model.Client {

	Body.Set("subject_token_type", "urn:ietf:params:oauth:token-type:id_token")
	Body.Set("requested_token_type", "urn:shopify:params:oauth:token-type:online-access-token")
	Body.Set("client_id", p.Setting.UserId)
	Body.Set("client_secret", p.Setting.Secret)
	Body.Set("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange")

	c := NewClient(p.Setting)
	c.SetPath(constant.EXCHGACCESS).
		SetParams(Body)

	//校验必输项字段
	if c.Err = Body.CheckEmptyError("subject_token_type",
		"requested_token_type",
		"client_id",
		"client_secret",
		"grant_type"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := ExchangeTokenResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
**参考地址：https://shopify.dev/docs/apps/auth/get-access-tokens/authorization-code-grant/getting-started
**步骤 4：获取访问令牌
 */

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	Body.Set("client_id", p.Setting.UserId)
	Body.Set("client_secret", p.Setting.Secret)

	c := NewClient(p.Setting)
	c.SetPath(constant.GETACCESS).
		SetParams(Body)

	//校验必输项字段
	if c.Err = Body.CheckEmptyError(
		"client_id",
		"client_secret",
		"code"); c.Err != nil {
		return &c.Client
	}
	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := AccessTokenResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
获取授权链接
shop:店铺名称
callbackParams : 同步回调的自定义参数
参考地址：https://shopify.dev/docs/apps/auth/get-access-tokens/authorization-code-grant/getting-started
*/

func (p *Api) GetAuthUrl(shop, callbackParams string) string {

	return fmt.Sprintf("%s%s?%s", fmt.Sprintf(constant.AUTHSITE, shop), constant.AUTH, model.BodyMap{}.
		Set("client_id", *p.Setting.UserId).
		Set("redirect_uri", url.QueryEscape(callbackParams)).
		Set("scope", constant.AUTH_SCOPE).
		Set("state", utils.GetRandLimitInt(1, 9999999)).
		Set("grant_options[]", "per-user").EncodeURLParams())
}

/*
	获取订单列表
	Url :https://shopify.dev/docs/api/admin-rest/2024-04/resources/order#get-orders?status=any
	Response:OrderListResponse
	todo:分页还未找到？
*/

func (p *Api) GetOrderList(Body model.BodyMap, Params model.BodyMap) *model.Client {

	var cursor *string
	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf(constant.ORDER_LIST, constant.API_V1)).
		SetMethod("GET").
		SetParams(Params)

	//if c.Err = Params.CheckEmptyError("page_size"); c.Err != nil {
	//	return &c.Client
	//}
	//报文头设置访问令牌
	c.HttpReq.Header.Set("X-Shopify-Access-Token", *p.Setting.AccessToken)
	c.Request.Params.Set("fields", constant.ORDER_LIST_FIELDS)
	result := OrderListResponse{}

	for {

		if cursor != nil && len(*cursor) > 0 {
			// 将page_token 参数 传给Query  page_token只能存在于params中
			// 去掉地址符
			c.Request.Params.Set("page_token", *cursor)
		}

		cResult := OrderListResponse{}

		c.Execute()
		if c.Err != nil || !c.Response.Success {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		//追增加响应订单列表至全局的订单列表
		if cResult.Orders != nil && len(cResult.Orders) > 0 {
			result.Orders = append(result.Orders, cResult.Orders...)
		}
		//todo:如何在响应中获取条数或者翻页标识？
		//if cResult.NextPageToken == "" {
		//	result.Total = cResult.TotalCount
		//	break
		//} else {
		//	cursor = &cResult.NextPageToken
		//}
	}

	c.Response.Response.DataTo = result

	return &c.Client
}

/*
	获取订单明细
	Url : https://shopify.dev/docs/api/admin-rest/2024-04/resources/order#get-orders-order-id?fields=id,line-items,name,total-price
	Response:
*/

func (p *Api) GetOrderDetail(Param model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	if c.Err = Param.CheckEmptyError("order_id"); c.Err != nil {
		return &c.Client
	}
	c.SetPath(fmt.Sprintf(constant.ORDER_DETAIL, constant.API_V1, Param.GetString("order_id"))).
		SetMethod("GET").
		SetParams(Param) // 改成 Param 传参
	c.HttpReq.Header.Set("X-Shopify-Access-Token", *p.Setting.AccessToken)
	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := OrderDetailResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	2024.5.23 16:12
	获取与授权店铺关联的所有商店
	Url : https://shopify.dev/docs/api/admin-rest/2024-04/resources/shop#get-shop
	Response: ShopConfigResponse
*/

func (p *Api) GetActiveShop(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(fmt.Sprintf(constant.ORDER_DETAIL, constant.SHOP_CONFIG)).
		SetMethod(http.MethodGet).
		SetParams(Body)
	c.HttpReq.Header.Set("X-Shopify-Access-Token", *p.Setting.AccessToken)
	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	var response ShopConfigResponse
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
