package shopee

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

const (
	PATH_TYPE_MERCHANT = "0"
	PATH_TYPE_SHOP     = "1"
	PATH_TYPE_PUBLIC   = "2"

	REFRESH_TOKEN_EXPRISE = 30 * 60 * 60 * 24

	//虾皮主站点
	//SHOPSITE = "https://partner.test-stable.shopeemobile.com"

	SHOPSITE = "https://partner.shopeemobile.com"

	//获取授权码
	AUTH_PARTNER = "/api/v2/shop/auth_partner"

	//获取访问令牌(根据授权码获取访问令牌及刷新令牌)
	AUTH_ACCESSTOKEN = "/api/v2/auth/token/get"

	//获取访问令牌(根据刷新令牌获取访问令牌)
	AUTH_REFRESHTOKEN = "/api/v2/auth/access_token/get"

	//AUTH_CALLBACK = "https://api.epur1688.cn/admin/v1/business/cbec/auth/2/callback"
)

const (
	//获取店铺信息
	SHOP_URL = "/api/v2/shop/get_shop_info"

	MERCHAT_URL = "/api/v2/merchant/get_merchant_info"

	SHOP_WAREHOUSE_DETAIL = "/api/v2/shop/get_warehouse_detail"
)

const (
	//获取订单信息
	ORDER_LIST = "/api/v2/order/get_order_list"
	//获取订单明细
	ORDER_DETAIL = "/api/v2/order/get_order_detail"
	//获取商品类目列表
	CATEGORY        = "/api/v2/product/get_category"
	GLOBAL_CATEGORY = "/api/v2/global_product/get_category"
)

type Response struct {
	RequestId string          `json:"request_id"` //请求Id
	Error     string          `json:"error"`      //返回码
	Message   string          `json:"message"`    //返回信息
	Response  json.RawMessage `json:"response"`   //响应信息
}

type OrderItem struct {
	OrderSn              string          `json:"order_sn"`               //默认情况下返回。Shopee订单的唯一标识符
	Region               string          `json:"region"`                 //默认情况下返回。表示订单所在地区的两位数代码
	Currency             string          `json:"currency"`               //币种
	Cod                  bool            `json:"cod"`                    //货到付款标识
	TotalAmount          decimal.Decimal `json:"total_amount"`           //总金额
	OrderStatus          string          `json:"order_status"`           //订单状态
	ShippingCarrier      string          `json:"shipping_carrier"`       //承运人
	PaymentMethod        string          `json:"payment_method"`         //支付方法
	EstimatedShippingFee decimal.Decimal `json:"estimated_shipping_fee"` //估计运费是Shopee根据特定物流快递员标准计算的估计费用

	MessageToSeller string `json:"message_to_seller"` //默认情况下返回。给卖家的消息
	CreateTime      int64  `json:"create_time"`       //默认情况下返回。指示订单创建日期和时间的时间戳
	UpdateTime      int64  `json:"update_time"`       //默认情况下返回。时间戳，表示订单值上次发生变化的时间，例如订单状态从“已付款”更改为“已完成”。
	DayToShip       int    `json:"days_to_ship"`      //默认情况下返回。卖家在Shopee上列出物品时设置的发货准备时间
	ShipByDate      int    `json:"ship_by_date"`      //默认情况下返回。寄出包裹的最后期限
	BuyerUserId     int64  `json:"buyer_user_id"`     //此订单买家的用户id
	BuyerUsername   string `json:"buyer_username"`    //买家姓名

	RecipientAddress struct {
		Name        string `json:"name"`         //收件人的地址名称
		Phone       string `json:"phone"`        //电话
		Town        string `json:"town"`         //镇
		District    string `json:"district"`     //区
		City        string `json:"city"`         //市
		State       string `json:"state"`        //州
		Region      string `json:"region"`       //地区
		Zipcode     string `json:"zipcode"`      //邮编
		FullAddress string `json:"full_address"` //地址
	} `json:"recipient_address"` //此对象包含收件人地址的详细信息。
	ActualShippingFee decimal.Decimal `json:"actual_shipping_fee"` //实际货运费用
	GoodsToDeclare    bool            `json:"goods_to_declare"`    //待申报货物标识

	Note           string `json:"note"`             //备注
	NoteUpdateTime int64  `json:"note_update_time"` //备注更新时间

	ItemList    []*Item `json:"item_list"`   //此对象包含此API调用结果的详细细分
	PayTime     int64   `json:"pay_time"`    //支付时间
	Dropshipper string  `json:"dropshipper"` //托运人

	DropshipperPhone  string `json:"dropshipper_phone"`   //托运人电话
	SplitUp           bool   `json:"split_up"`            //拆分标识
	BuyerCancelReason string `json:"buyer_cancel_reason"` //买家取消原因

	CancelBy                   string `json:"cancel_by"`                     //取消人
	CancelReason               string `json:"cancel_reason"`                 //取消原因
	ActualShippingFeeConfirmed bool   `json:"actual_shipping_fee_confirmed"` //实际发货去人
	BuyerCpfId                 string `json:"buyer_cpf_id"`                  //买家cpf Id
	FulfillmentFlag            string `json:"fulfillment_flag"`              //支付方法类型
	PickupDoneTime             int64  `json:"pickup_done_time"`              //拾取完成时间
	PackageList                []*struct {
		PackageNumber   string `json:"package_number"`   //包裹id
		LogisticsStatus string `json:"logistics_status"` //包裹状态
		ShippingCarrier string `json:"shipping_carrier"` //承运人
		ItemList        []*struct {
			ItemId        int64 `json:"item_id"`        //条目id
			ModelId       int64 `json:"model_id"`       //模型Id
			ModelQuantity int64 `json:"model_quantity"` //数量
		} `json:"item_list"` //条目
		ParcelChargableWeightGram int `json:"parcel_chargeable_weight_gram"` //包裹可计费重量
	} `json:"package_list"`
	InvoiceData *struct {
		Number             string `json:"number"`               //编号
		SeriesNumber       string `json:"series_number"`        //流水号
		AccessKey          string `json:"access_key"`           //
		IssueDate          string `json:"issue_date"`           //发布时间
		TotalValue         string `json:"total_value"`          //总价值
		ProductsTotalValue string `json:"products_total_value"` //产品总价格
		TaxCode            string `json:"tax_code"`             //税号
	} `json:"invoice_data"`
	CheckoutShippingCarrier string `json:"checkout_shipping_carrier"` //签出发货承运人

	ReverseShippingFee        decimal.Decimal `json:"reverse_shipping_fee"`         //反向装运费用
	OrderChargeableWeightGram int             `json:"order_chargeable_weight_gram"` //订单可收费重量克
	EdtFrom                   int64           `json:"edt_from"`                     //订单可收费重量克
	EdtTo                     int64           `json:"edt_to"`                       //订单可收费重量克
	PrescriptionImages        []string        `json:"prescription_images"`          //订单可收费重量克
	PrescriptionCheckStauts   int             `json:"prescription_check_status"`    //订单可收费重量
}

type Item struct {
	ItemId                 int64           `json:"item_id"`                  //商店物品的唯一标识符
	ItemName               string          `json:"item_name"`                //项目的名称
	ItemSku                string          `json:"item_sku"`                 //商品SKU（库存单位）是卖家定义的标识符，有时称为父SKU。项目SKU可以分配给Shopee Listings中的项目。
	ModelId                int64           `json:"model_id"`                 //属于同一项的模型的ID
	ModelName              string          `json:"model_name"`               //属于同一项的模型的名称。卖家可以提供相同商品的型号。例如，卖家可以为t恤设计创建一个固定价格的清单，并提供不同颜色和尺寸的t恤。在这种情况下，每个颜色和尺寸组合都是一个单独的模型。每个型号可以有不同的数量和价格。
	ModelSku               string          `json:"model_sku"`                //型号SKU（库存单位）是卖家定义的标识符。仅供卖方使用。许多卖家将SKU分配给特定类型、尺寸和颜色的商品，这些商品是Shopee Listings中一个商品的模型。
	ModelQuantityPurchased int             `json:"model_quantity_purchased"` //同一买家同时从一个物品中购买的相同物品的数量
	ModelOriginalPrice     decimal.Decimal `json:"model_original_price"`     //物品以挂牌货币表示的原始价格
	ModelDiscountedPrice   decimal.Decimal `json:"model_discounted_price"`   //以上市货币表示的项目的折扣后价格。如果没有折扣，该值将与model_original_price的值相同。对于捆绑交易项目，该值将返回0，因为根据设计，捆绑交易折扣不会细分到项目/模型级别。由于技术限制，如果我们不将其配置为0，则该值将返回捆绑交易前的价格。如果要计算捆绑交易项目的项目级别折扣价格，请调用GetEscrowDetails。
	Wholesale              bool            `json:"wholesale"`                //该值指示买家是否以批发价购买订单项目
	Weight                 decimal.Decimal `json:"weight"`                   //物品的重量
	AddOnDeal              bool            `json:"add_on_deal"`              //表示此项目是否属于附加交易
	MainItem               bool            `json:"main_item"`                //指示该项目是主项目还是子项目。True表示主项，false表示子项
	AddOnDealId            int64           `json:"add_on_deal_id"`           //用于区分购物车和订单中的项目组的唯一ID
	PromotionType          string          `json:"promotion_type"`           //促销类型
	PromotionId            int64           `json:"promotion_id"`             //促销id
	OrderItemId            int64           `json:"order_item_id"`            //订单条目id
	PromotionGroupId       int64           `json:"promotion_group_id"`       //促销组id
	ImageInfo              *struct {
		ImageUrl string `json:"image_url"`
	} `json:"image_info"` //图片
	ProductLocationId  []string `json:"product_location_id"`  //产品位置id
	IsPrescriptionItem bool     `json:"is_prescription_item"` //产品位置id
	IsB2cOwnedItem     bool     `json:"is_b2c_owned_item"`    //产品位置id
}
