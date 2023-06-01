package go1688

type GetTokenResponse struct {
	*Response

	AccessToken      string `json:"access_token"`
	AliId            string `json:"aliId"`
	RefreshToken     string `json:"refresh_token"`
	ResourceOwner    string `json:"resource_owner"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	MemberId         string `json:"member_id"`
}

// 跨境场景下将商品加入铺货列表
type PushedProductRequest struct {
	ProductIdList []int64 `json:"productIdList"` // 1688的商品ID列表,列表长度不能超过20个
}

// 跨境场景获取商品详情
type ProductInfoRequest struct {
	ProductId int64 `json:"productId"` // 1688的商品ID列表,列表长度不能超过20个
}

// 创建订单
type CreateOrderRequest struct {
	Flow            string                  `json:"flow"`            // 流程 general（创建大市场订单），fenxiao（创建分销订单）
	Message         string                  `json:"message"`         // 买家留言
	IsvBizType      string                  `json:"isvBizType"`      // 开放平台业务码,默认为cross。cross(跨境业务),cross_daigou（跨境代购业务）
	AddressParam    CreateOrderAddressParam `json:"addressParam"`    // 收货地址信息
	CargoParamList  []CreateOrderCargoParam `json:"cargoParamList"`  // 商品信息
	TradeType       string                  `json:"tradeType"`       // 由于不同的商品支持的交易方式不同，没有一种交易方式是全局通用的，所以当前下单可使用的交易方式必须通过下单预览接口的tradeModeNameList获取。交易方式类型说明：fxassure（交易4.0通用担保交易），alipay（大市场通用的支付宝担保交易（目前在做切流，后续会下掉）），period（普通账期交易）, assure（大买家企业采购询报价下单时需要使用的担保交易流程）, creditBuy（诚E赊），bank（银行转账），631staged（631分阶段付款），37staged（37分阶段）；此字段不传则系统默认会选取一个可用的交易方式下单，如果开通了诚E赊默认是creditBuy（诚E赊），未开通诚E赊默认使用的方式是支付宝担宝交易。
	ShopPromotionId string                  `json:"shopPromotionId"` // 店铺优惠ID，通过“创建订单前预览数据接口”获得。为空默认使用默认优惠
	AnonymousBuyer  bool                    `json:"anonymousBuyer"`  // 是否匿名下单
}

// 创建订单地址信息
type CreateOrderAddressParam struct {
	AddressId    int64  `json:"addressId"`    // 收货地址id
	FullName     string `json:"fullName"`     // 收货人姓名
	Mobile       string `json:"mobile"`       // 手机
	Phone        string `json:"phone"`        // 电话
	PostCode     string `json:"postCode"`     // 邮编
	CityText     string `json:"cityText"`     // 市文本
	ProvinceText string `json:"provinceText"` // 省份文本
	AreaText     string `json:"areaText"`     // 区文本
	TownText     string `json:"townText"`     // 镇文本
	Address      string `json:"address"`      // 街道地址
	DistrictCode string `json:"districtCode"` // 地址编码
}

// 订单商品信息
type CreateOrderCargoParam struct {
	SpecId   string `json:"specId"`
	Quantity int    `json:"quantity"`
	OfferId  int64  `json:"offerId"`
}

// 解析地址返回信息
type ParseAddressCodeResponse struct {
	*Response
	Result struct {
		Address         string `json:"address"`         // 街道地址，不包括省市编码
		AddressCode     string `json:"addressCode"`     // 地址区域编码
		AddressCodeText string `json:"addressCodeText"` // 地址区域编码对应的文本（包括国家，省，城市）
		AddressId       int64  `json:"addressId"`       // addressId
		BizType         string `json:"bizType"`         // 记录收货地址的业务类型
		IsDefault       bool   `json:"isDefault"`       // 是否为默认
		FullName        string `json:"fullName"`        // 收货人姓名
		Latest          bool   `json:"latest"`          // 是否是最后选择的收货地址
		Mobile          string `json:"mobile"`          // 手机号
		Phone           string `json:"phone"`           // 电话
		PostCode        string `json:"postCode"`        // 邮编
	} `json:"result"`
}

// 创建订单返回信息
type CreateOrderResponse struct {
	*Response
	Result struct {
		TotalSuccessAmount int64  `json:"totalSuccessAmount"` // 订单总金额（单位分），一次创建多个订单时，该字段为空
		OrderId            string `json:"orderId"`            // 订单ID，一次创建多个订单时，该字段为空
		Success            bool   `json:"success"`            // 是否成功
		Code               string `json:"code"`               // 错误码
		Message            string `json:"message"`            // 错误信息
		PostFee            int64  `json:"postFee"`            // 运费，单位：分，一次创建多个订单时，该字段为空
		OrderList          []struct {
			PostFee                        int    `json:"postFee"`      // 运费
			OrderAmmount                   int    `json:"orderAmmount"` // 订单实付款金额，单位为分
			Discount                       int    `json:"discount"`     // 描述信息
			SumPaymentNoCarriageFromClient int    `json:"sumPaymentNoCarriageFromClient"`
			MergePay                       bool   `json:"mergePay"`
			OrderId                        string `json:"orderId"` // 订单号
			ChooseFreeFreight              bool   `json:"chooseFreeFreight"`
		} `json:"orderList"` // 一次创建多个订单
	} `json:"result"`
}

// 发起免密支付
type OrderPreparePayRequest struct {
	Trade struct {
		OrderIds []string `json:"orderIds"` // 订单ID
	} `json:"tradeWithholdPreparePayParam"`
}
