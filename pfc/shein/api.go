package shein

import (
	"encoding/base64"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"strconv"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

/* API授权（返回授权地址）
** 参考地址：https://open.sheincorp.com/documents/system/2169474d-1d4a-41a9-b9fd-427f63f54a63
**/

func (p *Api) GetAuthUrl(callbackParams string) string {

	return fmt.Sprintf("%s%s?%s", SERVER_URL, AUTHSITE, model.BodyMap{}.
		Set("appid", *p.Setting.Key). //appid
		//Set("redirectUrl", url.QueryEscape(callbackParams)).
		Set("redirectUrl", base64.StdEncoding.EncodeToString([]byte(callbackParams))).
		Set("state", "AUTH-SHEIN-"+strconv.FormatInt(utils.TimestampSecond(), 10)).EncodeURLParams())
}

/*
	获取订单列表
	Url :https://open.sheincorp.com/documents/apidoc/detail/2000157-2000001
	Response:OrderListResponse
*/

func (p *Api) GetOrderList(Body model.BodyMap, Params model.BodyMap) *model.Client {

	page := int64(1)
	c := NewClient(p.Setting)
	c.SetPath(ORDER_LIST).
		SetMethod("POST").
		SetParams(Params)

	if c.Err = Params.CheckEmptyError("queryType",
		"startTime",
		"endTime",
		"page_size"); c.Err != nil {
		return &c.Client
	}

	pageSize, _ := strconv.ParseInt(Params.GetString("page_size"), 10, 64)

	result := OrderListResponse{}.Info

	for {

		cResult := OrderListResponse{}
		c.Execute()
		if c.Err != nil || !c.Response.Success {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		//追增加响应订单列表至全局的订单列表
		if cResult.Info.OrderList != nil && len(cResult.Info.OrderList) > 0 {
			result.OrderList = append(result.OrderList, cResult.Info.OrderList...)
		}
		//翻页
		if cResult.Info.Count > (page * pageSize) {
			page++
			c.Request.Params.Set("page", page)
		} else {
			break
		}
	}

	c.Response.Response.DataTo = result

	return &c.Client
}

/*
	获取订单明细
	Url :https://open.sheincorp.com/documents/apidoc/detail/3000257-2000001
	Response: OrderDetailResponse
*/

func (p *Api) GetOrderDetail(Param model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(ORDER_DETAIL).
		SetMethod("POST").
		SetParams(Param) // 改成 Param 传参

	if c.Err = Param.CheckEmptyError("orderNoList"); c.Err != nil {
		return &c.Client
	}

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
获取卖家密钥账号、卖家账号openkeyId、开发者app_id
Url :https://open.sheincorp.com/documents/apidoc/detail/3000051-1000012
Response: OrderDetailResponse
*/
func (p *Api) GetByToken(Param model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(SELLER_SECRET).
		SetMethod("POST").
		SetParams(Param) // 改成 Param 传参

	if c.Err = Param.CheckEmptyError("tempToken"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetByTokenResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
