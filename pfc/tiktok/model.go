package tiktok

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

const (
	SERVER_URl   = `https://open-api.tiktokglobalshop.com`
	AUTHSITE     = "https://auth.tiktok-shops.com"
	GETACCESS    = "/api/v2/token/get" // 获取访问令牌
	AUTH         = `/oauth/authorize`
	REFRESHTOKEN = "/api/v2/token/refresh"
)

type Response struct {
	RequestId string          `json:"request_id"` // 请求Id
	Code      int             `json:"code"`       // 返回码
	Message   string          `json:"message"`    // 返回信息
	Response  json.RawMessage `json:"data"`       // 响应信息
}

type getOrderListResponse struct {
	OrderList  []OrderListRow `json:"order_list"` //订单数据
	Total      int            `json:"total"`
	NextCursor string         `json:"next_cursor"`
	More       bool           `json:"more"`
}

type OrderListRow struct {
	OrderId     string `json:"order_id"`     //订单id
	OrderStatus int    `json:"order_status"` //订单状态
	UpdateTime  int64  `json:"update_time"`  //更新时间
}

type Order struct {
	OrderId     string `json:"order_id"`     //订单id
	OrderStatus int    `json:"order_status"` //订单状态
	UpdateTime  int64  `json:"update_time"`  //更新时间
}

//订单数据
type OrderDetailResponse struct {
	OrderId      string `json:"order_id"`             //订单id
	OrderStatus  int    `json:"order_status"`         //订单状态
	PayMethod    string `json:"payment_method"`       //支付方法
	DeliverOpt   string `json:"delivery_option"`      //快递方式
	ShipProvider string `json:"shipping_provider"`    //配送提供商
	ShipprovId   string `json:"shipping_provider_id"` //配送提供商id
	CreateTime   string `json:"create_time"`          //时间戳
	PayTime      int64  `json:"paid_time"`            //支付时间
	BuyerMessage string `json:"buyer_message"`        //买家信息

	IsCod bool `json:"is_cod"` //是否COD

	PayInfo                *PaymentInfo   `json:"payment_info"`      //支付信息
	RecipAddr              *RecipientAddr `json:"recipient_address"` //票据地址
	ItemList               []*Item        `json:"item_list"`         //订单条目
	CancelReason           string         `json:"cancel_reason"`     //取消原因
	CancelUser             string         `json:"cancel_user"`       //取消用户
	ExtStatus              int            `json:"ext_status"`        //订单子状态
	OrderStatusOld         string         `json:"order_status_old"`  //订单原状态
	TrackingNumber         string         `json:"tracking_number"`   //流水号
	RtsTime                int64          `json:"rts_time"`          //商家发货时间
	RtsSla                 int64          `json:"rts_sla"`           //最新的装运时间
	TtsSla                 int64          `json:"tts_sla"`           //最新的采集时间
	CancelOrderSla         int64          `json:"cancel_order_sla"`  //订单自动取消时间
	UpdateTime             int64          `json:"update_time"`       //更新时间
	PackageList            []*Package     `json:"package_list"`
	ReceiverAddressUpdated int            `json:"receiver_address_updated"` //收件地址更新标识，0-没有更新，1更新
	BuyerUid               string         `json:"buyer_uid"`                //运费
	SplitOrCombineTag      string         `json:"split_or_combine_tag"`     //拆分、合并标签
	FulfilmentType         int            `json:"fulfillment_type"`         //完成类型
	SellerNote             string         `json:"seller_note"`              //卖家note
	WarehouseId            string         `json:"warehouse_id"`             //仓库Id
	PaymethodMethodType    int            `json:"payment_method_type"`      //支付方法类型
	PaymethodMethodName    string         `json:"payment_method_name"`      //支付方法名
	OrderLineList          []*OrderLine   `json:"order_line_list"`          //订单行列表
}

//订单明细-条目
type Item struct {
	SkuId                string          `json:"sku_id"`                      //sku Id
	ProdId               string          `json:"product_id"`                  //产品id
	SkuName              string          `json:"sku_name"`                    //sku 名称
	Quantity             int             `json:"quantity"`                    //数量
	SellerSku            string          `json:"seller_sku"`                  //卖家sku
	ProdName             string          `json:"product_name"`                //产品名称
	SkuImage             string          `json:"sku_image"`                   //sku图片
	SkuOriPrice          decimal.Decimal `json:"sku_original_price"`          //sku初始价格
	SkuSalePrice         decimal.Decimal `json:"sku_sale_price"`              //卖家note
	SkuPDiscount         decimal.Decimal `json:"sku_platform_discount"`       //sku销售价格
	SkuSellerDiscount    decimal.Decimal `json:"sku_seller_discount"`         //sku卖家折扣
	SkuExtStatus         int             `json:"sku_ext_status"`              //sku 额外状态
	SkuDisplayStatus     int             `json:"sku_display_status"`          //sku显示状态
	SkuCancelReason      string          `json:"sku_cancel_reason"`           //sku取消原因
	SkuCancelUser        string          `json:"sku_cancel_user"`             //sku取消用户
	SkuRtsTime           int             `json:"sku_rts_time"`                //sku rts时间
	SkuType              int             `json:"sku_type"`                    //sku类型
	SkuPDiscountSumTotal decimal.Decimal `json:"sku_platform_discount_total"` //sku平台折扣总额
	SkuSmallOrderFee     decimal.Decimal `json:"sku_small_order_fee"`         //sku小订单费
}

//订单明细-订单line
type OrderLine struct {
	OrderLineId      string          `json:"order_line_id"`          //orderline Id
	SkuId            string          `json:"sku_id"`                 //skuid
	ExtStatus        int             `json:"ext_status"`             //额外状态
	DisplayStatus    int             `json:"display_status"`         //显示状态
	ProdId           string          `json:"product_id"`             //产品id
	ProdName         string          `json:"product_name"`           //产品名称
	SkuName          string          `json:"sku_name"`               //sku名称
	SellerSku        string          `json:"seller_sku"`             //卖家sku
	SkuImage         string          `json:"sku_image"`              //sku图片
	SkuType          int             `json:"sku_type"`               //sku类型
	OrgPrice         decimal.Decimal `json:"original_price"`         //初始价格
	SalePrice        decimal.Decimal `json:"sale_price"`             //销售价格
	PlatformDiscount decimal.Decimal `json:"platform_discount"`      //平台折扣
	SellerDiscount   decimal.Decimal `json:"seller_discount"`        //卖家折扣
	RtsTime          int             `json:"rts_time"`               //rts时间
	CancelReason     string          `json:"cancel_reason"`          //取消原因
	CancelUser       string          `json:"cancel_user"`            //取消用户
	PackageId        int             `json:"package_id"`             //包裹id
	PackStatus       int             `json:"package_status"`         //包裹状态
	PackFrzStatus    int             `json:"package_freeze_status"`  //包裹冻结状态
	ShipProvideId    string          `json:"shipping_provider_id"`   //货运提供商Id
	ShipProvideName  string          `json:"shipping_provider_name"` //货运提供商名称
	TrackingNum      string          `json:"tracking_number"`        //跟踪号
}

//订单明细-包裹
type Package struct {
	PackageId string `json:"package_id"` //包裹id
}

//订单明细-支付信息
type PaymentInfo struct {
	Currency                string          `json:"currency"`                       //币种
	SubTotal                decimal.Decimal `json:"sub_total"`                      //小计
	ShipFee                 decimal.Decimal `json:"shipping_fee"`                   //运费
	SellerDiscount          decimal.Decimal `json:"seller_discount"`                //卖家折扣
	PlatformDiscount        decimal.Decimal `json:"platform_discount"`              //平台折扣
	TotalAmt                decimal.Decimal `json:"total_amount"`                   //总金额
	OrgTotalPrdPrice        decimal.Decimal `json:"original_total_product_price"`   //初始合计总额
	OrgShipFee              decimal.Decimal `json:"original_shipping_fee"`          //初试运费
	ShipFeeSellerDiscount   decimal.Decimal `json:"shipping_fee_seller_discount"`   //卖家发货折扣
	ShipFeePlatFormDiscount decimal.Decimal `json:"shipping_fee_platform_discount"` //平台运费折扣
	Tax                     decimal.Decimal `json:"taxes"`                          //税费
	SmallOrderFee           decimal.Decimal `json:"small_order_fee"`                //小订单费
}

//订单明细-收件人地址
type RecipientAddr struct {
	FullAddr     string   `json:"full_address"`      //地址
	Region       string   `json:"region"`            //地区
	State        string   `json:"state"`             //州
	City         string   `json:"city"`              //城市
	District     string   `json:"district"`          //区县
	Town         string   `json:"town"`              //镇
	Phone        string   `json:"phone"`             //电话
	Name         string   `json:"name"`              //姓名
	Zipcode      string   `json:"zipcode"`           //邮编
	AddrDetail   string   `json:"address_detail"`    //地址明细
	AddrLineList []string `json:"address_line_list"` //地址行列表
	RegionCode   string   `json:"region_code"`       //地区代码
}

type ProductListResponse struct {

	/*
		{
		"1": "Published",
		"2": "Created",
		"3": "Draft",
		"4": "Deleted"
		}
	*/

	CreateTime int64  `json:"create_time"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Status     int64  `json:"status"`
	UpdTime    int64  `json:"upd_time"`
	Skus       []struct {
		Id        string `json:"id"`
		SellerSku string `json:"seller_sku"`
	}
}
