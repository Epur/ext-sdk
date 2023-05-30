package lazada

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

const (
	// APIGatewaySG endpoint
	APIGatewaySG = "https://api.lazada.sg/rest"
	// APIGatewayMY endpoint
	APIGatewayMY = "https://api.lazada.com.my/rest"
	// APIGatewayVN endpoint
	APIGatewayVN = "https://api.lazada.vn/rest"
	// APIGatewayTH endpoint
	APIGatewayTH = "https://api.lazada.co.th/rest"
	// APIGatewayPH endpoint
	APIGatewayPH = "https://api.lazada.com.ph/rest"
	// APIGatewayID endpoint
	APIGatewayID = "https://api.lazada.co.id/rest"

	APICODEURL = "https://api.lazada.com/rest"

	APIREFRESHURL = "https://auth.lazada.com/rest"

	AuthURL = "https://auth.lazada.com/oauth/authorize"
)

const (

	//根据授权码获取访问令牌、刷新令牌
	AccessTokenURL = "/auth/token/create"

	RefreshURL = "/auth/token/refresh"

	//卖家信息
	SELLERURL = "/seller/get"

	UPLOADIMAGE = "/image/upload"
)

type Response struct {
	Code      string          `json:"code"`
	Type      string          `json:"type"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
	Result    json.RawMessage `json:"result"`
}

type ResponseConutryUserInfo struct {
	Country   string `json:"country"`
	SellerId  string `json:"seller_id"`
	UserId    string `json:"user_id"`
	ShortCode string `json:"short_code"`
}

type OrderDetailResponse struct {
	PickUpStore struct {
		StoreName     string `json:"pick_up_store_name"`      // 提货商店名称
		StoreAddr     string `json:"pick_up_store_address"`   //门店地址
		StoreCode     string `json:"pick_up_store_code"`      //门店代码
		StoreOpenHour string `json:"pick_up_store_open_hour"` //提货店开放时间
	} `json:"pick_up_store_info"`
	PurOrderNo         string          `json:"purchase_order_number"`          // 调用SetPackedByMarketPlace时返回
	Name               string          `json:"name"`                           //产品名称
	ProdMainImage      string          `json:"product_main_image"`             //产品主图像URL
	ItemPrice          decimal.Decimal `json:"item_price"`                     //产品价格
	TaxAmt             decimal.Decimal `json:"tax_amount"`                     //税额
	Status             string          `json:"status"`                         //状态
	CancelRetInit      string          `json:"cancel_return_initiator"`        //指示谁发起了取消或退回的订单。可能的值包括取消内部、取消客户、取消失败交付、取消卖家、退货客户和退款内部
	VouPlatform        decimal.Decimal `json:"voucher_platform"`               //Lazada签发的凭证
	VouSeller          decimal.Decimal `json:"voucher_seller"`                 //卖方出具的凭证
	OrderType          string          `json:"order_type"`                     // 订单类型，可能是普通、预售、优惠券、O2O或InStoreO2O
	StagePayStatus     string          `json:"stage_pay_status"`               // 预售阶段预售订单的付款状态。可能的值为空、“未付”或“未付尾款”。（未付：预售定金尚未支付；未付尾款：预售订金已支付但尾款/应付余额未支付）
	WarehoCode         string          `json:"warehouse_code"`                 //多wh卖家的仓库代码
	VouSellerLpi       decimal.Decimal `json:"voucher_seller_lpi"`             //卖方赞助的Lazada奖金
	VouPlatformLpi     decimal.Decimal `json:"voucher_platform_lpi"`           //由Lazada赞助的Lazada奖金
	BuyerId            int64           `json:"buyer_id"`                       //买方ID
	ShipOriFee         decimal.Decimal `json:"shipping_fee_original"`          //运费原价
	ShipDisFeeSeller   decimal.Decimal `json:"shipping_fee_discount_seller"`   //卖方运费折扣
	ShipDisFeePlatform decimal.Decimal `json:"shipping_fee_discount_platform"` //平台运费折扣
	VouCodeSeller      string          `json:"voucher_code_seller"`            // 卖家的凭证代码
	VouCodePlatform    string          `json:"voucher_code_platform"`          //平台的凭证代码
	DeliverOptSof      decimal.Decimal `json:"delivery_option_sof"`            //是否为卖方自有车队的标记，值包括1和0
	IsFbl              decimal.Decimal `json:"is_fbl"`                         //LAZADA是否满足的标记，值包括1和0
	IsReroute          int             `json:"is_reroute"`                     //是否为二次销售的标记，值包括1和0

	Reason         string          `json:"reason"`                 //销售订单原因表中定义的取消、退货或其他原因
	DigDeliverInfo string          `json:"digital_delivery_info"`  //数字送货信息
	PromShipTime   string          `json:"promised_shipping_time"` //承诺发货时间
	OrderId        int64           `json:"order_id"`               // 订单号
	VouAmt         decimal.Decimal `json:"voucher_amount"`         // 凭证金额
	RetStatus      string          `json:"return_status"`          //返回状态
	ShipType       string          `json:"shipping_type"`          //发货类型、直运或仓库
	ShipProvider   string          `json:"shipment_provider"`      //支付费用和退还卖方费用的总和
	Variation      string          `json:"variation"`              //项目收入和其他收入之和
	CreateAt       string          `json:"created_at"`             //以ISO 8601格式创建提要的时间
	InvoiceNo      string          `json:"invoice_number"`         //发票编号
	ShipAmt        decimal.Decimal `json:"shipping_amount"`        //运费
	Currency       string          `json:"currency"`               // ISO 4217兼容货币代码
	OrderFlag      string          `json:"order_flag"`             //订单类型，可能值为GUARANTEE、NORMAL和GLOBAL_COLLECTION。标有“GUARANTEE”或“GLOBAL_COLLECTION”的订单在订单履行中的SLA要求较短。
	ShopId         string          `json:"shop_id"`                //卖方名称
	SlaTimestamp   string          `json:"sla_time_stamp"`         //ISO 8601格式的装运时间SLA（yyyy-MM-dd’T'HH:MM:ssXXX）
	Sku            string          `json:"sku"`                    //产品SKU

	VoucherCode     string          `json:"voucher_code"`           //所有退款总额（如有）
	WalletCredits   decimal.Decimal `json:"wallet_credits"`         //钱包信用卡
	UpdateAt        string          `json:"updated_at"`             //以ISO 8601格式上次更新源的时间
	IsDigital       int             `json:"is_digital"`             // 数字商品与否
	TrackingCodePre string          `json:"tracking_code_pre"`      // 更新时间
	OrderItemId     int64           `json:"order_item_id"`          //订单项目ID
	PackageId       string          `json:"package_id"`             //包源ID
	TrackingCode    string          `json:"tracking_code"`          //从第三方物流发货提供商检索的跟踪代码
	ShipSerCost     decimal.Decimal `json:"shipping_service_cost"`  //运输服务成本
	ExtAttr         string          `json:"extra_attributes"`       //带有额外属性的JSON编码字符串）
	PaidPrice       decimal.Decimal `json:"paid_price"`             //已付价格
	ShipProvidType  string          `json:"shipping_provider_type"` //以下选项之一：EXPRESS、STANDARD、ECONOMY、INSTANT、SELLER_OWN_FLEET、PICKUP_IN_STORE或DIGITAL
	ProdDetailUrl   string          `json:"product_detail_url"`     //产品详细信息URL
	ShopSku         string          `json:"shop_sku"`               //产品外部ID
	ReasonDetail    string          `json:"reason_detail"`          //原因详细信息
	PurOrderId      string          `json:"purchase_order_id"`      //调用SetPackedByMarketPlace时返回
	SkuId           string          `json:"sku_id"`                 //Sku ID
	ProdId          string          `json:"product_id"`             //产品ID
}

type OrderResponse struct {
	BrcNum          string          `json:"branch_number"`                // （仅限泰国）公司客户的税务分支机构代码，由客户在下订单时提供。
	TaxCode         string          `json:"tax_code"`                     //（仅适用于泰国和越南）客户在下订单时提供的增值税代码。
	ExtAttr         string          `json:"extra_attributes"`             //在调用getMarketPlaceOrders时传递给卖方中心的额外属性。
	AddrUpAt        string          `json:"address_updated_at"`           //地址更新于
	ShipFee         decimal.Decimal `json:"shipping_fee"`                 //订单的总运费
	CustFirstName   string          `json:"customer_first_name"`          //客户名字
	PayMethod       string          `json:"payment_method"`               //付款方式。有关详细信息，请参见付款方式选项
	Statuses        []string        `json:"statuses"`                     //订单中项目的唯一状态数组。您可以在响应示例中找到所有不同的状态代码
	Remarks         string          `json:"remarks"`                      //评论
	OrderNum        int64           `json:"order_number"`                 // 订单ID
	OrderId         int64           `json:"order_id"`                     // 订单ID
	Voucher         decimal.Decimal `json:"voucher"`                      //订单的总凭证
	NatRegNum       string          `json:"national_registration_number"` //国家注册号。某些国家需要。
	PromShipTimes   string          `json:"promised_shipping_times"`      //最快订单项目的目标发货时间（如果有）
	ItemsCount      int64           `json:"items_count"`                  //订单中的项目数
	VoucherPlatform decimal.Decimal `json:"voucher_platform"`             //Lazada签发的凭证
	VoucherSeller   decimal.Decimal `json:"voucher_seller"`               //卖方出具的凭证
	CreateAt        string          `json:"created_at"`                   //下订单的日期和时间。
	Price           string          `json:"price"`                        // 此订单的总金额。不是订单的最终交易价格，不包括凭单和发货人
	AddressBilling  struct {
		Addr1     string `json:"address1"`   //已发放退款的累计费用
		Phone2    string `json:"phone2"`     //（小计总额1）-退款
		FirstName string `json:"first_name"` // 创建时间
		Phone     string `json:"phone"`      // 更新时间
		Addr5     string `json:"address5"`   //（小计总额1）-退款
		PostCode  string `json:"post_code"`  // 创建时间

		Addr4    string `json:"address4"`  //（小计总额1）-退款
		LastName string `json:"last_name"` // 创建时间
		Country  string `json:"country"`   // 更新时间
		Addr3    string `json:"address3"`  //运费抵免（如有）
		Addr2    string `json:"address2"`  //其他汇总收入
		City     string `json:"city"`      //支付费用和退还卖方费用的总和
	} `json:"address_billing"` //json格式，地址
	WarehouseCode               string          `json:"warehouse_code"`                 //多wh卖家的仓库代码
	ShippingFeeOriginal         decimal.Decimal `json:"shipping_fee_original"`          //在任何类型的运费促销之前，应向客户收取的原始运费
	ShippingFeeDiscountSeller   decimal.Decimal `json:"shipping_fee_discount_seller"`   //卖方运费折扣
	ShippingFeeDiscountPlatform decimal.Decimal `json:"shipping_fee_discount_platform"` //平台运费折扣
	AddressShipping             struct {
		Addr1     string `json:"address1"`   //地址1
		Phone2    string `json:"phone2"`     //电话2
		FirstName string `json:"first_name"` // 姓名
		Phone     string `json:"phone"`      // 电话
		Addr5     string `json:"address5"`   //地址5
		PostCode  string `json:"post_code"`  // 邮编

		Addr4    string `json:"address4"`  // 地址4
		LastName string `json:"last_name"` // 姓名
		Country  string `json:"country"`   // 国家
		Addr3    string `json:"address3"`  // 地址3
		Addr2    string `json:"address2"`  // 地址2
		City     string `json:"city"`      // 城市
	} `json:"address_shipping"` //json格式，货运地址
	CustomerLastName string `json:"customer_last_name"` //现在是空的。请参见cutomer_first_name
	GiftOption       bool   `json:"gift_option"`        //如果物品是礼物，则为1；如果不是，则为0
	VoucherCode      string `json:"voucher_code"`       // 返回值为凭证id
	UpdateAt         string `json:"updated_at"`         // 上次更改订单的日期和时间。
	DeliverInfo      string `json:"delivery_info"`      //交货信息
	GitMessage       string `json:"gift_message"`       // 客户指定的礼品信息
}
