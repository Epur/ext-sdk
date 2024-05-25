package tiktokv2

import (
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/logger"
	"github.com/Epur/ext-sdk/model"
	"net/http"
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
	2024.5.22 11:36
	获取授权店铺信息
	Url : https://partner.tiktokshop.com/docv2/page/6507ead7b99d5302be949ba9?external_id=6507ead7b99d5302be949ba9
	Response: ShopListResponse
*/

func (p *Api) GetSellerShop(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(GET_AUTHORIZED_SHOP).
		SetMethod(http.MethodGet).
		SetParams(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	var response ShopListResponse
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	2024.5.23 16:12
	获取与授权店铺关联的所有商店
	Url : https://partner.tiktokshop.com/docv2/page/650a69e24a0bb702c067291c?external_id=650a69e24a0bb702c067291c
	Response: ShopListResponse
*/

func (p *Api) GetActiveShop(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(GET_SELLER_SHOP).
		SetMethod(http.MethodGet).
		SetParams(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	var response ActiveShopListResponse
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

	// c.Execute()已经存在 app_key的字段 所以这里需要删除 2024.5.22 11:36
	c.Request.Params.Set("grant_type", "refresh_token").
		Set("app_secret", *p.Setting.Secret)
	//.Set("app_key", *p.Setting.Key)

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

func (p *Api) GetOrderList(Body model.BodyMap, Params model.BodyMap) *model.Client {

	var cursor *string
	c := NewClient(p.Setting)
	c.SetPath(`/order/202309/orders/search`).
		SetMethod("POST").
		SetBody(Body).
		SetParams(Params) // 新增 Query 参数赋值  page_size只能存在于params中

	if c.Err = Params.CheckEmptyError("page_size"); c.Err != nil {
		return &c.Client
	}

	result := GetOrderListResponse{}

	for {

		if cursor != nil && len(*cursor) > 0 {
			// 将page_token 参数 传给Query  page_token只能存在于params中
			// 去掉地址符
			c.Request.Params.Set("page_token", *cursor)
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

func (p *Api) GetOrderDetail(Param model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/order/202309/orders`).
		SetMethod("GET").
		SetParams(Param) // 改成 Param 传参

	if c.Err = Param.CheckEmptyError("ids"); c.Err != nil {
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
	获取包裹面单
	Url : https://partner.tiktokshop.com/docv2/page/650aa5fac16ffe02b8f112ca
	Response: GetOrderDetailResponse
*/

func (p *Api) PrintAwb(Param model.BodyMap, PackageId string) *model.Client {
	c := NewClient(p.Setting)
	if len(PackageId) <= 0 || PackageId == "" {
		c.Client.Err = errors.New("PackageId is empty")
		return &c.Client
	}

	c.SetPath(fmt.Sprintf(GET_SHIPPING_DOCUMENTS, PackageId)).
		SetMethod("GET").
		SetParams(Param) // 改成 Param 传参

	if c.Err = Param.CheckEmptyError("document_type"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetOrderPrintAwbResponse{}
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

// 订单发货（运单号回填）
func (p *Api) OrderShipPackage202309(body model.BodyMap) error {
	logger.SdkLogger.Info("OrderShipPackage202309...")

	c := NewClient(p.Setting)
	c.SetMethod("POST").
		SetBody(body).
		SetParams(model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher)).
		SetPath(SHIP_PACKAGE)

	c.Execute()
	if c.Err != nil {
		logger.SdkLogger.Error(c.Err.Error())
		return c.Err
	}
	logger.SdkLogger.Info("OrderShipPackage202309 END")
	return nil
}

// 获取店铺关联仓库
func (p *Api) LogisticsWarehouses202309() *model.Client {
	logger.SdkLogger.Info("LogisticsWarehouses202309...")

	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(LOGISTICS_WAREHOUSES).
		SetParams(model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher))

	c.Execute()
	if c.Err != nil {
		logger.SdkLogger.Error(c.Err.Error())
		return &c.Client
	}
	response := LogisticsWarehousesResult{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response

	logger.SdkLogger.Info("LogisticsWarehouses202309 END")
	return &c.Client
}

// 获取卖家指定仓库订阅的配送选项列表
func (p *Api) LogisticsWarehousesDeliveryOptions202309(WarehousesId *string) *model.Client {
	logger.SdkLogger.Info("LogisticsWarehousesDeliveryOptions202309...")

	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(fmt.Sprintf(LOGISTICS_WAREHOUSES_DELIVERY_OPTIONS, *WarehousesId)).
		SetParams(model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher))

	c.Execute()
	if c.Err != nil {
		logger.SdkLogger.Error(c.Err.Error())
		return &c.Client
	}
	response := LogisticsWarehousesDeliveryOptionsResult{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response

	logger.SdkLogger.Info("LogisticsWarehousesDeliveryOptions202309 END")
	return &c.Client
}

// 获取指定配送选项对应的配送商
func (p *Api) LogisticsWarehousesDeliveryOptionsShip202309(DeliveryOptionsId *string) *model.Client {
	logger.SdkLogger.Info("LogisticsWarehousesDeliveryOptionsShip202309...")

	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(fmt.Sprintf(LOGISTICS_DELIVERY_OPTIONS_SHIP, *DeliveryOptionsId)).
		SetParams(model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher))

	c.Execute()
	if c.Err != nil {
		logger.SdkLogger.Error(c.Err.Error())
		return &c.Client
	}
	response := LogisticsWarehousesDeliveryOptionsShipResult{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response

	logger.SdkLogger.Info("LogisticsWarehousesDeliveryOptionsShip202309 END")
	return &c.Client
}

// 获取包裹详情
func (p *Api) GetPackageDetail202309(packageId *string) *model.Client {
	logger.SdkLogger.Info("GetPackageDetail202309...")

	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(fmt.Sprintf(GET_PACKAGE_DETAIL_BY_PACKAGEID, *packageId)).
		SetParams(model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher))

	c.Execute()
	if c.Err != nil {
		logger.SdkLogger.Error(c.Err.Error())
		return &c.Client
	}
	response := PackageDetail{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response

	logger.SdkLogger.Info("GetPackageDetail202309 END")
	return &c.Client
}

// 获取合格的运输服务
func (p *Api) GetShippingServices202309(orderId *string) *model.Client {
	logger.SdkLogger.Info("GetShippingServices202309...")

	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(fmt.Sprintf(GET_SHIPPING_SERVICES, *orderId)).
		SetParams(model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher))

	c.Execute()
	if c.Err != nil {
		logger.SdkLogger.Error(c.Err.Error())
		return &c.Client
	}
	response := ShippingServices{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response

	logger.SdkLogger.Info("GetShippingServices202309 END")
	return &c.Client
}

func (p *Api) GetOrderStatementTransaction202309(orderId *string) *model.Client {
	logger.SdkLogger.Infof("GetOrderStatementTransaction...%s", *orderId)

	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(fmt.Sprintf(STATEMENT_ORDER_TRANSACTIONS_GET_URL, *orderId)).
		SetParams(model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher))

	c.Execute()
	if c.Err != nil {
		logger.SdkLogger.Error(c.Err.Error())
		return &c.Client
	}
	response := OrderStatementTransactionsResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

func (p *Api) GetStatementTransaction202309(data StatementTransactionRequest) *model.Client {
	logger.SdkLogger.Infof("GetStatementTransaction... %v", *data.StatementId)

	param := model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher).
		Set("sort_field", "order_create_time")

	if data.PageSize != nil {
		param.Set("page_size", *data.PageSize)
	}
	if data.PageToken != nil {
		param.Set("page_token", *data.PageToken)
	}
	if data.SortOrder != nil {
		param.Set("sort_order", *data.SortOrder)
	}

	result := StatementTransactionResponse{}
	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(fmt.Sprintf(STATEMENT_TRANSACTIONS_GET_URL, *data.StatementId)).
		SetParams(param)

	for {

		c.Request.Params.Set("page_size", "100")

		cResult := StatementTransactionResponse{}

		c.Execute()
		if c.Err != nil || !c.Response.Success {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		if len(cResult.StatementTransactions) > 0 {
			result.StatementTransactions = append(result.StatementTransactions, cResult.StatementTransactions...)
		}

		if len(result.StatementTransactions) >= cResult.TotalCount {
			break
		}
	}
	c.Response.Response.DataTo = result

	logger.SdkLogger.Info("GetStatementTransaction202309....END")

	return &c.Client
}

func (p *Api) GetStatement202309(data StatementRequest) *model.Client {
	logger.SdkLogger.Info("GetStatement202309...")

	param := model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher).
		Set("sort_field", "statement_time").Set("page_size", "100")

	if data.PageToken != nil {
		param.Set("page_token", *data.PageToken)
	}
	if data.SortOrder != nil {
		param.Set("sort_order", *data.SortOrder)
	}
	if data.StatementTimeGe != nil && data.StatementTimeIt != nil {
		param.Set("statement_time_ge", fmt.Sprintf("%d", *data.StatementTimeGe)).
			Set("statement_time_it", fmt.Sprintf("%d", *data.StatementTimeIt))
	}

	if data.PaymentStatus != nil {
		param.Set("payment_status", *data.PaymentStatus)
	}

	result := StatementResponse{}
	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(STATEMENT_GET_URL).
		SetParams(param)

	for {

		c.Request.Params.Set("page_size", "100")

		cResult := StatementResponse{}

		c.Execute()
		if c.Err != nil || !c.Response.Success {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		if len(cResult.Statements) > 0 {
			result.Statements = append(result.Statements, cResult.Statements...)
		}

		if len(result.NextPageToken) <= 0 {
			break
		}
	}
	c.Response.Response.DataTo = result

	logger.SdkLogger.Info("GetStatement202309....END")

	return &c.Client
}

func (p *Api) GetPayments202309(beginTime, endTime *int64) *model.Client {
	logger.SdkLogger.Info("GetPayments202309...")

	param := model.BodyMap{}.Set("shop_cipher", *p.Setting.ShopCipher).
		Set("sort_field", "create_time").Set("page_size", "100")

	if beginTime != nil && endTime != nil {
		param.Set("create_time_lt", fmt.Sprintf("%d", *beginTime)).
			Set("create_time_ge", fmt.Sprintf("%d", *endTime))
	}

	result := PaymentsResponse{}
	c := NewClient(p.Setting)
	c.SetMethod("GET").
		SetPath(STATEMENT_PAYMENTS).
		SetParams(param)

	for {

		c.Request.Params.Set("page_size", "100")

		cResult := PaymentsResponse{}

		c.Execute()
		if c.Err != nil || !c.Response.Success {
			return &c.Client
		}

		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return &c.Client
		}

		if len(cResult.Payments) > 0 {
			result.Payments = append(result.Payments, cResult.Payments...)
		}

		if len(result.NextPageToken) <= 0 {
			break
		}
	}
	c.Response.Response.DataTo = result

	logger.SdkLogger.Info("GetPayments202309....END")

	return &c.Client
}
