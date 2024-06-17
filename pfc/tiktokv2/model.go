package tiktokv2

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

const (
	SERVER_URl   = `https://open-api.tiktokglobalshop.com`
	AUTH_URL     = "http://auth.tiktok-shops.com"
	AUTHSITE     = "https://services.tiktokshop.com"
	AUTHSITE_US  = "https://services.us.tiktokshop.com"
	GETACCESS    = "/api/v2/token/get" // 获取访问令牌
	AUTH         = `/open/authorize`
	REFRESHTOKEN = "/api/v2/token/refresh"

	GET_SHIPPING_DOCUMENTS = "/fulfillment/202309/packages/%v/shipping_documents" // 获取面单

	GET_AUTHORIZED_SHOP = "/authorization/202309/shops" // 获取店铺信息
	GET_SELLER_SHOP     = "/seller/202309/shops"        // 获取授权店铺与卖家关联的所有商店

	// 获取包裹详情
	GET_SHIPPING_SERVICES = "/fulfillment/202309/orders/%s/shipping_services/query"

	// 获取包裹详情
	GET_PACKAGE_DETAIL_BY_PACKAGEID = "/fulfillment/202309/packages/%s"

	// 运单号回填（原发货功能）
	SHIP_PACKAGE = "/fulfillment/202309/packages/ship"

	// 获取卖家关联的所有仓库信息
	LOGISTICS_WAREHOUSES = "/logistics/202309/warehouses"

	// 获取卖家指定仓库订阅的配送选项列表  %s == warehouse_id
	LOGISTICS_WAREHOUSES_DELIVERY_OPTIONS = "/logistics/202309/warehouses/%s/delivery_options"

	// 获取指定配送选项对应的配送商  %s = delivery_option_id
	LOGISTICS_DELIVERY_OPTIONS_SHIP = "/logistics/202309/delivery_options/%s/shipping_providers"

	// 订单级结算
	STATEMENT_ORDER_TRANSACTIONS_GET_URL = "/finance/202309/orders/%s/statement_transactions"

	// 结算交易明细
	STATEMENT_TRANSACTIONS_GET_URL = "/finance/202309/statements/%s/statement_transactions"

	// 获取结算交易记录
	STATEMENT_GET_URL = "/finance/202309/statements"

	// 获取付款交易记录
	STATEMENT_PAYMENTS = "/finance/202309/payments"

	UPLOAD_PRODUCT_IMAGES = "/product/202309/images/upload"
)

type Response struct {
	RequestId string          `json:"request_id"` // 请求Id
	Code      int             `json:"code"`       // 返回码
	Message   string          `json:"message"`    // 返回信息
	Response  json.RawMessage `json:"data"`       // 响应信息
}

// 订单列表信息
type getOrderListResponse struct {
	TotalCount    int            `json:"total_count"`
	NextPageToken string         `json:"next_page_token"`
	Orders        []OrderListRow `json:"orders"` //订单数据
}

// 订单数据
type OrderListRow struct {
	Id                                  string            `json:"id"`                                     //Tiktok shop order id
	BuyerMessage                        string            `json:"buyer_message"`                          //The note from buyer.
	CancellationInitiator               string            `json:"cancellation_initiator"`                 //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	ShippingProviderId                  string            `json:"shipping_provider_id"`                   //Tiktok shop order id
	CreateTime                          int64             `json:"create_time"`                            //The note from buyer.
	ShippingProvider                    string            `json:"shipping_provider"`                      //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	Packages                            []OrderPackage    `json:"packages"`                               //Tiktok shop order id
	Payment                             *Payment          `json:"payment"`                                //The note from buyer.
	RecipientAddress                    *RecipientAddress `json:"recipient_address"`                      //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	Status                              string            `json:"status"`                                 //Tiktok shop order id
	FulfillmentType                     string            `json:"fulfillment_type"`                       //The note from buyer.
	TrackingNumber                      string            `json:"tracking_number"`                        //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	PaidTime                            int64             `json:"paid_time"`                              //Tiktok shop order id
	RtsSlaTime                          int64             `json:"rts_sla_time"`                           //The note from buyer.
	TtsSlaTime                          int64             `json:"tts_sla_time"`                           //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	CancelReason                        string            `json:"cancel_reason"`                          //Tiktok shop order id
	UpdateTime                          int64             `json:"update_time"`                            //The note from buyer.
	PaymentMethodName                   string            `json:"payment_method_name"`                    //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	RtsTime                             int64             `json:"rts_time"`                               //Tiktok shop order id
	UserId                              string            `json:"user_id"`                                //The note from buyer.
	SplitOrCombineTag                   string            `json:"split_or_combine_tag"`                   //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	HasUpdatedrRecipientAddress         bool              `json:"has_updated_recipient_address"`          //Tiktok shop order id
	CancelOrderSlaTime                  int64             `json:"cancel_order_sla_time"`                  //The note from buyer.
	WarehouseId                         string            `json:"warehouse_id"`                           //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	RequestCancelTime                   int64             `json:"request_cancel_time"`                    //Tiktok shop order id
	ShippingType                        string            `json:"shipping_type"`                          //The note from buyer.
	deliveryOptionName                  string            `json:"delivery_option_name"`                   //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	SellerNote                          string            `json:"seller_note"`                            //Tiktok shop order id
	DeliverySlaTime                     int64             `json:"delivery_sla_time"`                      //The note from buyer.
	IsCod                               bool              `json:"is_cod"`                                 //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	DeliveryOptionId                    string            `json:"delivery_option_id"`                     //The note from buyer.
	CancelTime                          int64             `json:"cancel_time"`                            //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	NeedUploadInvoice                   string            `json:"need_upload_invoice"`                    //Tiktok shop order id
	IsSampleOrder                       bool              `json:"is_sample_order"`                        //The note from buyer.
	Cpf                                 string            `json:"cpf"`                                    //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	LineItems                           []LineItem        `json:"line_items"`                             //Tiktok shop order id
	BuyerEmail                          string            `json:"buyer_email"`                            //The note from buyer.
	DeliveryDueTime                     int64             `json:"delivery_due_time"`                      //The note from buyer.
	IsOnHoldOrder                       bool              `json:"is_on_hold_order"`                       //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	ShippingDueTime                     int64             `json:"shipping_due_time"`                      //Tiktok shop order id
	CollectionDueTime                   int64             `json:"collection_due_time"`                    //The note from buyer.
	DeliveryOptionRequireddDeliveryTime int64             `json:"delivery_option_required_delivery_time"` //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	CollectionTime                      int64             `json:"collection_time"`                        //Tiktok shop order id
	DeliveryTime                        int64             `json:"delivery_time"`                          //The note from buyer.
	IsBuyerRequestCancel                bool              `json:"is_buyer_request_cancel"`                //The note from buyer.
	IsReplacementOrder                  bool              `json:"is_replacement_order"`                   //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	ReplacedOrderId                     string            `json:"replaced_order_id"`                      //Tiktok shop order id
}

type OrderPackage struct {
	Id string `json:"replaced_order_id"`
}

type Payment struct {
	Currency                    string `json:"currency"`
	SubTotal                    string `json:"sub_total"`
	ShippingFee                 string `json:"shipping_fee"`
	SellerDiscount              string `json:"seller_discount"`
	PlatformDiscount            string `json:"platform_discount"`
	TotalAmount                 string `json:"total_amount"`
	OriginalTotalProductPrice   string `json:"original_total_product_price"`
	OriginalShippingFee         string `json:"original_shipping_fee"`
	ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
	ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
	Tax                         string `json:"tax"`
	SmallOrderFee               string `json:"small_order_fee"`
	ShippingFeeTax              string `json:"shipping_fee_tax"`
	ProductTax                  string `json:"product_tax"`
	RetailDeliveryFee           string `json:"retail_delivery_fee"`
}

type RecipientAddress struct {
	FullAddress         string               `json:"full_address"`
	PhoneNumber         string               `json:"phone_number"`
	Name                string               `json:"name"`
	PostalCode          string               `json:"postal_code"`
	AddressDetail       string               `json:"address_detail"`
	RegionCode          string               `json:"region_code"`
	AddressLine1        string               `json:"address_line1"`
	AddressLine2        string               `json:"address_line2"`
	AddressLine3        string               `json:"address_line3"`
	AddressLine4        string               `json:"address_line4"`
	DistrictInfo        []DistrictInfo       `json:"district_info"`
	DeliveryPreferences *DeliveryPreferences `json:"delivery_preferences"`
}

type DistrictInfo struct {
	AddressLevelName string `json:"address_level_name"`
	AddressName      string `json:"address_name"`
	AddressLevel     string `json:"address_level"`
}

type DeliveryPreferences struct {
	DropOffLocation string `json:"drop_off_location"`
}

type LineItem struct {
	Id                   string    `json:"id"`
	SkuId                string    `json:"sku_id"`
	DisplayStatus        string    `json:"display_status"`
	ProductId            string    `json:"product_id"`
	ProductName          string    `json:"product_name"`
	SellerSku            string    `json:"seller_sku"`
	SkuImage             string    `json:"sku_image"`
	SkuName              string    `json:"sku_name"`
	OriginalPrice        string    `json:"original_price"`
	SalePrice            string    `json:"sale_price"`
	PlatformDiscount     string    `json:"platform_discount"`
	SellerDiscount       string    `json:"seller_discount"`
	SkuType              string    `json:"sku_type"`
	CancelReason         string    `json:"cancel_reason"`
	CancelUser           string    `json:"cancel_user"`
	RtsTime              int       `json:"rts_time"`
	PackageStatus        string    `json:"package_status"`
	Currency             string    `json:"currency"`
	ShippingProviderName string    `json:"shipping_provider_name"`
	TrackingNumber       string    `json:"tracking_number"`
	ShippingProviderId   string    `json:"shipping_provider_id"`
	IsGift               bool      `json:"is_gift"`
	ItemTax              []ItemTax `json:"item_tax"`
	SmallOrderFee        string    `json:"small_order_fee"`
	PackageId            string    `json:"package_id"`
	RetailDeliveryFee    string    `json:"retail_delivery_fee"`
}

type ItemTax struct {
	TaxType   string `json:"tax_type"`
	TaxAmount string `json:"tax_amount"`
	TaxRate   string `json:"tax_rate"`
}
type Order struct {
	OrderId     string `json:"order_id"`     //订单id
	OrderStatus int    `json:"order_status"` //订单状态
	UpdateTime  int64  `json:"update_time"`  //更新时间
}

// 订单明细数据
type OrderDetailResponse struct {
	Id                                  string             `json:"id"`                     //Tiktok shop order id
	CancellationInitiator               string             `json:"cancellation_initiator"` //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	ShippingProvider                    string             `json:"shipping_provider"`
	ShippingProviderId                  string             `json:"shipping_provider_id"` //Tiktok shop order id
	CreateTime                          int64              `json:"create_time"`          //The note from buyer.
	Status                              string             `json:"status"`               //支付方法
	RtsTime                             int64              `json:"rts_time"`             //商家发货时间
	Payment                             *PaymentM          `json:"payment"`              //商家发货时间
	RecipientAddress                    *RecipientAddressM `json:"recipient_address"`
	BuyerMessage                        string             `json:"buyer_message"`
	UserId                              string             `json:"user_id"`
	TrackingNumber                      string             `json:"tracking_number"` //流水号
	CancelReason                        string             `json:"cancel_reason"`   //取消原因
	RtsSlaTime                          int64              `json:"rts_sla_time"`
	PayTime                             int64              `json:"paid_time"` //支付时间
	CancelOrderSlaTime                  int64              `json:"cancel_order_sla_time"`
	SellerNote                          string             `json:"seller_note"` //卖家note
	TtsSlaTime                          int64              `json:"tts_sla_time"`
	HasUpdatedrRecipientAddress         bool               `json:"has_updated_recipient_address"` //Tiktok shop order id
	UpdateTime                          int64              `json:"update_time"`                   //更新时间
	Packages                            []Package          `json:"packages"`
	FulfilmentType                      int                `json:"fulfillment_type"` //完成类型
	IsSampleOrder                       bool               `json:"is_sample_order"`
	WarehouseId                         string             `json:"warehouse_id"`         //仓库Id
	SplitOrCombineTag                   string             `json:"split_or_combine_tag"` //拆分、合并标签
	ShippingType                        string             `json:"shipping_type"`
	Cpf                                 string             `json:"cpf"`
	DeliveryOptionId                    string             `json:"delivery_option_id"`
	DeliverySlaTime                     int64              `json:"delivery_sla_time"`
	PaymethodMethodName                 string             `json:"payment_method_name"` //支付方法名
	ShippingDueTime                     int64              `json:"shipping_due_time"`
	LineItems                           []LineItemM        `json:"line_items"`
	deliveryOptionName                  string             `json:"delivery_option_name"`
	BuyerEmail                          string             `json:"buyer_email"`
	DeliveryTime                        int64              `json:"delivery_time"`
	NeedUploadInvoice                   string             `json:"need_upload_invoice"`
	IsCod                               bool               `json:"is_cod"` //是否COD
	RequestCancelTime                   int64              `json:"request_cancel_time"`
	DeliveryOptionRequireddDeliveryTime int64              `json:"delivery_option_required_delivery_time"` //Cancel request initiator. Avaliable value: SELLER/ BUYER/ SYSTEM .
	CollectionDueTime                   int64              `json:"collection_due_time"`
	IsBuyerRequestCancel                bool               `json:"is_buyer_request_cancel"`
	DeliveryDueTime                     int64              `json:"delivery_due_time"`
	CollectionTime                      int64              `json:"collection_time"`
	IsOnHoldOrder                       bool               `json:"is_on_hold_order"`
	CancelTime                          int64              `json:"cancel_time"`
	IsReplacementOrder                  bool               `json:"is_replacement_order"`
	ReplacedOrderId                     string             `json:"replaced_order_id"`
}

type PaymentM struct {
	Currency                    string `json:"currency"`
	SubTotal                    string `json:"sub_total"`
	ShippingFee                 string `json:"shipping_fee"`
	SellerDiscount              string `json:"seller_discount"`
	PlatformDiscount            string `json:"platform_discount"`
	TotalAmount                 string `json:"total_amount"`
	OriginalTotalProductPrice   string `json:"original_total_product_price"`
	OriginalShippingFee         string `json:"original_shipping_fee"`
	ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
	ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
	Tax                         string `json:"tax"`
	SmallOrderFee               string `json:"small_order_fee"`
	ShippingFeeTax              string `json:"shipping_fee_tax"`
	ProductTax                  string `json:"product_tax"`
	RetailDeliveryFee           string `json:"retail_delivery_fee"`
}

type RecipientAddressM struct {
	FullAddress         string               `json:"full_address"`
	PhoneNumber         string               `json:"phone_number"`
	Name                string               `json:"name"`
	PostalCode          string               `json:"postal_code"`
	AddressDetail       string               `json:"address_detail"`
	RegionCode          string               `json:"region_code"`
	AddressLine1        string               `json:"address_line1"`
	AddressLine2        string               `json:"address_line2"`
	AddressLine3        string               `json:"address_line3"`
	AddressLine4        string               `json:"address_line4"`
	DistrictInfo        []DistrictInfo       `json:"district_info"`
	DeliveryPreferences *DeliveryPreferences `json:"delivery_preferences"`
}

type LineItemM struct {
	Id                   string    `json:"id"`
	SkuId                string    `json:"sku_id"`
	CombinedListingSkus  []CLS     `json:"combined_listing_skus"`
	ProductId            string    `json:"product_id"`
	ProductName          string    `json:"product_name"`
	SkuName              string    `json:"sku_name"`
	SkuImage             string    `json:"sku_image"`
	OriginalPrice        string    `json:"original_price"`
	SalePrice            string    `json:"sale_price"`
	PlatformDiscount     string    `json:"platform_discount"`
	DisplayStatus        string    `json:"display_status"`
	CancelUser           string    `json:"cancel_user"`
	SkuType              string    `json:"sku_type"`
	SellerSku            string    `json:"seller_sku"`
	ShippingProviderId   string    `json:"shipping_provider_id"`
	SellerDiscount       string    `json:"seller_discount"`
	Currency             string    `json:"currency"`
	PackageId            string    `json:"package_id"`
	RtsTime              int       `json:"rts_time"`
	PackageStatus        string    `json:"package_status"`
	ShippingProviderName string    `json:"shipping_provider_name"`
	IsGift               bool      `json:"is_gift"`
	CancelReason         string    `json:"cancel_reason"`
	SmallOrderFee        string    `json:"small_order_fee"`
	RetailDeliveryFee    string    `json:"retail_delivery_fee"`
	TrackingNumber       string    `json:"tracking_number"`
	DeliveryOptionName   string    `json:"delivery_option_name"`
	ItemTax              []ItemTax `json:"item_tax"`
}

type CLS struct {
	SkuId     string `json:"sku_id"`     //sku Id
	SkuCount  int    `json:"sku_count"`  //sku Id
	ProductId string `json:"product_id"` //产品id
}

// 订单明细-条目
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

// 订单明细-订单line
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

// 订单明细-包裹
type Package struct {
	PackageId string `json:"package_id"` //包裹id
}

// 订单明细-支付信息
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

// 订单明细-收件人地址
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

type ProductDetailResponse struct {
	AuditFailedReasons []struct {
		Position    string   `json:"position"`
		Reasons     []string `json:"reasons"`
		Suggestions []string `json:"suggestions"`
	} `json:"audit_failed_reasons"`
	Brand struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"brand"`
	CategoryChains []struct {
		Id        string `json:"id"`
		IsLeaf    bool   `json:"is_leaf"`
		LocalName string `json:"local_name"`
		ParentId  string `json:"parent_id"`
	} `json:"category_chains"`
	Certifications []struct {
		Files []struct {
			Format string   `json:"format"`
			Id     string   `json:"id"`
			Name   string   `json:"name"`
			Urls   []string `json:"urls"`
		} `json:"files"`
		Id     string `json:"id"`
		Images []struct {
			Height    int      `json:"height"`
			ThumbUrls []string `json:"thumb_urls"`
			Uri       string   `json:"uri"`
			Urls      []string `json:"urls"`
			Width     int      `json:"width"`
		} `json:"images"`
		Title string `json:"title"`
	} `json:"certifications"`
	CreateTime      int `json:"create_time"`
	DeliveryOptions []struct {
		Id          string `json:"id"`
		IsAvailable bool   `json:"is_available"`
		Name        string `json:"name"`
	} `json:"delivery_options"`
	Description       string `json:"description"`
	ExternalProductId string `json:"external_product_id"`
	Id                string `json:"id"`
	IsCodAllowed      bool   `json:"is_cod_allowed"`
	IsNotForSale      bool   `json:"is_not_for_sale"`
	MainImages        []struct {
		Height    int      `json:"height"`
		ThumbUrls []string `json:"thumb_urls"`
		Uri       string   `json:"uri"`
		Urls      []string `json:"urls"`
		Width     int      `json:"width"`
	} `json:"main_images"`
	Manufacturer struct {
		Address     string `json:"address"`
		Email       string `json:"email"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"manufacturer"`
	PackageDimensions struct {
		Height string `json:"height"`
		Length string `json:"length"`
		Unit   string `json:"unit"`
		Width  string `json:"width"`
	} `json:"package_dimensions"`
	PackageWeight struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"package_weight"`
	ProductAttributes []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Values []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"product_attributes"`
	ProductTypes string `json:"product_types"`
	SizeChart    struct {
		Image struct {
			Height    int      `json:"height"`
			ThumbUrls []string `json:"thumb_urls"`
			Uri       string   `json:"uri"`
			Urls      []string `json:"urls"`
			Width     int      `json:"width"`
		} `json:"image"`
		Template struct {
			Id string `json:"id"`
		} `json:"template"`
	} `json:"size_chart"`
	Skus []struct {
		CombinedSkus []struct {
			ProductId string `json:"product_id"`
			SkuCount  int    `json:"sku_count"`
			SkuId     string `json:"sku_id"`
		} `json:"combined_skus"`
		ExternalSkuId       string `json:"external_sku_id"`
		GlobalListingPolicy struct {
			InventoryType   string `json:"inventory_type"`
			PriceSync       bool   `json:"price_sync"`
			ReplicateSource struct {
				ProductId string `json:"product_id"`
				ShopId    string `json:"shop_id"`
				SkuId     string `json:"sku_id"`
			} `json:"replicate_source"`
		} `json:"global_listing_policy"`
		Id             string `json:"id"`
		IdentifierCode struct {
			Code string `json:"code"`
			Type string `json:"type"`
		} `json:"identifier_code"`
		Inventory []struct {
			Quantity    int    `json:"quantity"`
			WarehouseId string `json:"warehouse_id"`
		} `json:"inventory"`
		Price struct {
			Currency          string `json:"currency"`
			SalePrice         string `json:"sale_price"`
			TaxExclusivePrice string `json:"tax_exclusive_price"`
			UnitPrice         string `json:"unit_price"`
		} `json:"price"`
		SalesAttributes []struct {
			Id     string `json:"id"`
			Name   string `json:"name"`
			SkuImg struct {
				Height    int      `json:"height"`
				ThumbUrls []string `json:"thumb_urls"`
				Uri       string   `json:"uri"`
				Urls      []string `json:"urls"`
				Width     int      `json:"width"`
			} `json:"sku_img"`
			ValueId   string `json:"value_id"`
			ValueName string `json:"value_name"`
		} `json:"sales_attributes"`
		SellerSku    string `json:"seller_sku"`
		SkuUnitCount string `json:"sku_unit_count"`
	} `json:"skus"`
	Status     string `json:"status"`
	Title      string `json:"title"`
	UpdateTime int    `json:"update_time"`
	Video      struct {
		CoverUrl string `json:"cover_url"`
		Format   string `json:"format"`
		Height   int    `json:"height"`
		Id       string `json:"id"`
		Size     int    `json:"size"`
		Url      string `json:"url"`
		Width    int    `json:"width"`
	} `json:"video"`
}

type ImageEx struct {
	Height       int64    `json:"height"`
	Width        int64    `json:"width"`
	ThumbUrlList []string `json:"thumb_url_list"`
	Id           string   `json:"id"`
	UrlList      []string `json:"url_List"`
}

// 订单级交易记录（TikTok订单级结算）
type OrderStatementTransactionsResponse struct {
	OrderCreateTime            int64                      `json:"order_create_time"`      // 订单创建时间
	OrderId                    string                     `json:"order_id"`               // 订单编号
	OrderStatementTransactions OrderStatementTransactions `json:"statement_transactions"` // 对账单交易列表
}

type OrderStatementTransactions []OrderStatementTransaction

// 对账单交易列表
type OrderStatementTransaction struct {
	OrderStatementId                   string                   `json:"id"`                                     // 事务唯一键
	StatementId                        string                   `json:"statement_id"`                           // 结算ID
	Currency                           string                   `json:"currency"`                               // 币种
	ActualReturnShippingFeeAmount      decimal.Decimal          `json:"actual_return_shipping_fee_amount"`      // 实际退货运费金额
	ActualShippingFeeAmount            decimal.Decimal          `json:"actual_shipping_fee_amount"`             // 实际运费金额
	AdjustmentAmount                   decimal.Decimal          `json:"adjustment_amount"`                      // 调整金额
	AffiliateCommissionAmount          decimal.Decimal          `json:"affiliate_commission_amount"`            // 会员佣金金额
	AffiliatePartnerCommissionAmount   decimal.Decimal          `json:"affiliate_partner_commission_amount"`    // 会员合作伙伴佣金金额
	AfterSellerDiscountsSubtotalAmount decimal.Decimal          `json:"after_seller_discounts_subtotal_amount"` // 卖家折扣后小计金额
	CustomerOrderRefundAmount          decimal.Decimal          `json:"customer_order_refund_amount"`           // 客户订单退款金额
	CustomerPaymentAmount              decimal.Decimal          `json:"customer_payment_amount"`                // 客户付款金额
	CustomerRefundAmount               decimal.Decimal          `json:"customer_refund_amount"`                 // 客户退款金额
	CustomerShippingFeeAmount          decimal.Decimal          `json:"customer_shipping_fee_amount"`           // 自定义运费金额
	FeeAmount                          decimal.Decimal          `json:"fee_amount"`                             // 费用金额
	PlatformCommissionAmount           decimal.Decimal          `json:"platform_commission_amount"`             // 平台佣金金额
	PlatformDiscountAmount             decimal.Decimal          `json:"platform_discount_amount"`               // 平台折扣金额
	PlatformDiscountRefundAmount       decimal.Decimal          `json:"platform_discount_refund_amount"`        // 平台折扣退款金额
	PlatformRefundSubsidyAmount        decimal.Decimal          `json:"platform_refund_subsidy_amount"`         // 平台退款补贴金额
	PlatformShippingFeeDiscountAmount  decimal.Decimal          `json:"platform_shipping_fee_discount_amount"`  // 平台运费折扣金额
	ReferralFeeAmount                  decimal.Decimal          `json:"referral_fee_amount"`                    // 销售佣金金额
	RetailDeliveryFeeAmount            decimal.Decimal          `json:"retail_delivery_fee_amount"`             // 零售配送费金额
	RetailDeliveryFeePaymentAmount     decimal.Decimal          `json:"retail_delivery_fee_payment_amount"`     // 零售配送费支付金额
	RetailDeliveryFeeRefundAmount      decimal.Decimal          `json:"retail_delivery_fee_refund_amount"`      // 零售配送费退款金额
	RevenueAmount                      decimal.Decimal          `json:"revenue_amount"`                         // 收入金额
	SalesTaxAmount                     decimal.Decimal          `json:"sales_tax_amount"`                       // 销售税金额
	SalesTaxPaymentAmount              decimal.Decimal          `json:"sales_tax_payment_amount"`               // 销售税缴纳金额
	SalesTaxRefundAmount               decimal.Decimal          `json:"sales_tax_refund_amount"`                // 销售税退税金额
	SettlementAmount                   decimal.Decimal          `json:"settlement_amount"`                      // 结算金额
	ShippingFeeAmount                  decimal.Decimal          `json:"shipping_fee_amount"`                    // 运费金额
	ShippingFeeSubsidyAmount           decimal.Decimal          `json:"shipping_fee_subsidy_amount"`            // 运费补贴金额
	TransactionFeeAmount               decimal.Decimal          `json:"transaction_fee_amount"`                 // 交易费用金额
	StatementTime                      int64                    `json:"statement_time"`                         // 结算时间
	Status                             string                   `json:"status"`                                 // 状态 事务的状态。在正常情况下，响应中仅包含已结算的订单。
	SkuStatementTransactions           SkuStatementTransactions `json:"sku_statement_transactions"`             // SKU 对帐单交易记录
}

type SkuStatementTransactions []SkuStatementTransaction

// SKU 对帐单交易记录
type SkuStatementTransaction struct {
	SkuId                              string          `json:"sku_id"`                                 // skuId
	SkuName                            string          `json:"sku_name"`                               // sku名称
	ProductName                        string          `json:"product_name"`                           // 商品名称
	Currency                           string          `json:"currency"`                               // 币种
	Quantity                           int64           `json:"quantity"`                               // 数量
	ActualReturnShippingFeeAmount      decimal.Decimal `json:"actual_return_shipping_fee_amount"`      // 实际退货运费金额
	ActualShippingFeeAmount            decimal.Decimal `json:"actual_shipping_fee_amount"`             // 实际运费金额
	AdjustmentAmount                   decimal.Decimal `json:"adjustment_amount"`                      // 调整金额
	AffiliateCommissionAmount          decimal.Decimal `json:"affiliate_commission_amount"`            // 会员佣金金额
	AffiliatePartnerCommissionAmount   decimal.Decimal `json:"affiliate_partner_commission_amount"`    // 会员合作伙伴佣金金额
	AfterSellerDiscountsSubtotalAmount decimal.Decimal `json:"after_seller_discounts_subtotal_amount"` // 卖家折扣后小计金额
	CustomerOrderRefundAmount          decimal.Decimal `json:"customer_order_refund_amount"`           // 客户订单退款金额
	CustomerPaymentAmount              decimal.Decimal `json:"customer_payment_amount"`                // 客户付款金额
	CustomerRefundAmount               decimal.Decimal `json:"customer_refund_amount"`                 // 客户退款金额
	CustomerShippingFeeAmount          decimal.Decimal `json:"customer_shipping_fee_amount"`           // 自定义运费金额
	FeeAmount                          decimal.Decimal `json:"fee_amount"`                             // 费用金额
	PlatformCommissionAmount           decimal.Decimal `json:"platform_commission_amount"`             // 平台佣金金额
	PlatformDiscountAmount             decimal.Decimal `json:"platform_discount_amount"`               // 平台折扣金额
	PlatformDiscountRefundAmount       decimal.Decimal `json:"platform_discount_refund_amount"`        // 平台折扣退款金额
	PlatformRefundSubsidyAmount        decimal.Decimal `json:"platform_refund_subsidy_amount"`         // 平台退款补贴金额
	PlatformShippingFeeDiscountAmount  decimal.Decimal `json:"platform_shipping_fee_discount_amount"`  // 平台运费折扣金额
	ReferralFeeAmount                  decimal.Decimal `json:"referral_fee_amount"`                    // 销售佣金金额
	RetailDeliveryFeeAmount            decimal.Decimal `json:"retail_delivery_fee_amount"`             // 零售配送费金额
	RetailDeliveryFeePaymentAmount     decimal.Decimal `json:"retail_delivery_fee_payment_amount"`     // 零售配送费支付金额
	RetailDeliveryFeeRefundAmount      decimal.Decimal `json:"retail_delivery_fee_refund_amount"`      // 零售配送费退款金额
	RevenueAmount                      decimal.Decimal `json:"revenue_amount"`                         // 收入金额
	SalesTaxAmount                     decimal.Decimal `json:"sales_tax_amount"`                       // 销售税金额
	SalesTaxPaymentAmount              decimal.Decimal `json:"sales_tax_payment_amount"`               // 销售税缴纳金额
	SalesTaxRefundAmount               decimal.Decimal `json:"sales_tax_refund_amount"`                // 销售税退税金额
	SettlementAmount                   decimal.Decimal `json:"settlement_amount"`                      // 结算金额
	ShippingFeeAmount                  decimal.Decimal `json:"shipping_fee_amount"`                    // 运费金额
	ShippingFeeSubsidyAmount           decimal.Decimal `json:"shipping_fee_subsidy_amount"`            // 运费补贴金额
	TransactionFeeAmount               decimal.Decimal `json:"transaction_fee_amount"`                 // 交易费用金额
}

type StatementTransactionRequest struct {
	StatementId *string `json:"statement_id"` // 结算ID
	PageToken   *string `json:"page_token"`   // 本页访问token (非必要)
	SortOrder   *string `json:"sort_order"`   // 正序ASC 倒叙DESC  (非必要)
	PageSize    *int64  `json:"page_size"`    // 页码 (非必要)
}

// 结算交易记录
type StatementTransactionResponse struct {
	AdjustmentAmount      decimal.Decimal       `json:"adjustment_amount"` // 调整金额
	Currency              string                `json:"currency"`          // 币种
	FeeAmount             decimal.Decimal       `json:"fee_amount"`        // 费用金额
	NextPageToken         string                `json:"next_page_token"`   // 下一页访问token
	RevenueAmount         decimal.Decimal       `json:"revenue_amount"`    // 收入金额
	SettlementAmount      decimal.Decimal       `json:"settlement_amount"` // 结算金额
	StatementId           string                `json:"statement_id"`      // 结算ID
	StatementTime         int64                 `json:"statement_time"`    // 结算时间
	StatementTransactions StatementTransactions `json:"statement_transactions"`
	TotalCount            int                   `json:"total_count"`
}

type StatementTransactions []StatementTransaction

type StatementTransaction struct {
	StatementTransactionId string `json:"id"`                  // 事务唯一键
	OrderId                string `json:"order_id"`            // 订单ID
	AdjustmentId           string `json:"adjustment_id"`       // 调整ID
	AdjustmentOrderId      string `json:"adjustment_order_id"` // 调整订单ID
	Currency               string `json:"currency"`            // 币种
	/*
		调整类型 100 -->SHIPPING_FEE_ADJUSTMENT 		 						(当卖家支付的运费存在差异或错误时进行调整)
				110 -->SHIPPING_FEE_COMPENSATION 	 						(因实际运费与预付运费之间的差额而给予卖家的补偿)
				120 -->CHARGE_BACK 					 						(在客户成功对账户对账单或交易报告中的项目提出异议后，退还到支付卡的扣款)
				130 -->CUSTOMER_SERVICE_COMPENSATION 						(这是客服在售后期后支付给客户的额外补偿或补偿)
				140 -->PROMOTION_ADJUSTMENT  								(当卖家参与平台促销活动时，促销价格与卖家实际支付的金额之间存在差异时的调整)
	*/
	Type                               string          `json:"type"`
	ActualReturnShippingFeeAmount      decimal.Decimal `json:"actual_return_shipping_fee_amount"`      // 实际退货运费金额
	ActualShippingFeeAmount            decimal.Decimal `json:"actual_shipping_fee_amount"`             // 实际运费金额
	AdjustmentAmount                   decimal.Decimal `json:"adjustment_amount"`                      // 调整金额
	AffiliateCommissionAmount          decimal.Decimal `json:"affiliate_commission_amount"`            // 会员佣金金额
	AffiliatePartnerCommissionAmount   decimal.Decimal `json:"affiliate_partner_commission_amount"`    // 会员合作伙伴佣金金额
	AfterSellerDiscountsSubtotalAmount decimal.Decimal `json:"after_seller_discounts_subtotal_amount"` // 卖家折扣后小计金额
	CustomerOrderRefundAmount          decimal.Decimal `json:"customer_order_refund_amount"`           // 客户订单退款金额
	CustomerPaymentAmount              decimal.Decimal `json:"customer_payment_amount"`                // 客户付款金额
	CustomerRefundAmount               decimal.Decimal `json:"customer_refund_amount"`                 // 客户退款金额
	CustomerShippingFeeAmount          decimal.Decimal `json:"customer_shipping_fee_amount"`           // 自定义运费金额
	FeeAmount                          decimal.Decimal `json:"fee_amount"`                             // 费用金额
	PlatformCommissionAmount           decimal.Decimal `json:"platform_commission_amount"`             // 平台佣金金额
	PlatformDiscountAmount             decimal.Decimal `json:"platform_discount_amount"`               // 平台折扣金额
	PlatformDiscountRefundAmount       decimal.Decimal `json:"platform_discount_refund_amount"`        // 平台折扣退款金额
	PlatformRefundSubsidyAmount        decimal.Decimal `json:"platform_refund_subsidy_amount"`         // 平台退款补贴金额
	PlatformShippingFeeDiscountAmount  decimal.Decimal `json:"platform_shipping_fee_discount_amount"`  // 平台运费折扣金额
	ReferralFeeAmount                  decimal.Decimal `json:"referral_fee_amount"`                    // 销售佣金金额
	RetailDeliveryFeeAmount            decimal.Decimal `json:"retail_delivery_fee_amount"`             // 零售配送费金额
	RetailDeliveryFeePaymentAmount     decimal.Decimal `json:"retail_delivery_fee_payment_amount"`     // 零售配送费支付金额
	RetailDeliveryFeeRefundAmount      decimal.Decimal `json:"retail_delivery_fee_refund_amount"`      // 零售配送费退款金额
	RevenueAmount                      decimal.Decimal `json:"revenue_amount"`                         // 收入金额
	SalesTaxAmount                     decimal.Decimal `json:"sales_tax_amount"`                       // 销售税金额
	SalesTaxPaymentAmount              decimal.Decimal `json:"sales_tax_payment_amount"`               // 销售税缴纳金额
	SalesTaxRefundAmount               decimal.Decimal `json:"sales_tax_refund_amount"`                // 销售税退税金额
	SettlementAmount                   decimal.Decimal `json:"settlement_amount"`                      // 结算金额
	ShippingFeeAmount                  decimal.Decimal `json:"shipping_fee_amount"`                    // 运费金额
	ShippingFeeSubsidyAmount           decimal.Decimal `json:"shipping_fee_subsidy_amount"`            // 运费补贴金额
	TransactionFeeAmount               decimal.Decimal `json:"transaction_fee_amount"`                 // 交易费用金额
	OrderCreateTime                    int             `json:"order_create_time"`                      // 订单创建时间
}

type StatementRequest struct {
	/*
		- PAID：the payment has been transferred to the Seller
		- FAILED：the payment failed
		- PROCESSING：the payment is currently processing. If the payment is successful, the status will transition to PAID. If not, it will be FAILED.
	*/
	PaymentStatus   *string `json:"payment_status"`    // 支付状态 (非必要)
	StatementTimeIt *int64  `json:"statement_time_it"` // 结算开始时间 (非必要)
	StatementTimeGe *int64  `json:"statement_time_ge"` // 结算结束时间 (非必要)
	PageToken       *string `json:"page_token"`        // 本页访问token (非必要)
	SortOrder       *string `json:"sort_order"`        // 正序ASC 倒叙DESC  (非必要)
	PageSize        *int64  `json:"page_size"`         // 页码 (非必要)
}

// 结算
type StatementResponse struct {
	NextPageToken string     `json:"next_page_token"`
	Statements    Statements `json:"statements"`
}

type Statements []Statement

type Statement struct {
	AdjustmentAmount decimal.Decimal `json:"adjustment_amount"`
	Currency         string          `json:"currency"`
	FeeAmount        decimal.Decimal `json:"fee_amount"`
	Id               string          `json:"id"`
	PaymentId        string          `json:"payment_id"`
	PaymentStatus    string          `json:"payment_status"`
	RevenueAmount    decimal.Decimal `json:"revenue_amount"`
	SettlementAmount decimal.Decimal `json:"settlement_amount"`
	StatementTime    int64           `json:"statement_time"`
}

type PaymentsResponse struct {
	NextPageToken string `json:"next_page_token"`
	Payments      []struct {
		Amount struct {
			Currency string          `json:"currency"`
			Value    decimal.Decimal `json:"value"`
		} `json:"amount"`
		BankAccount                 string          `json:"bank_account"`
		CreateTime                  int64           `json:"create_time"`
		ExchangeRate                decimal.Decimal `json:"exchange_rate"`
		Id                          string          `json:"id"`
		PaidTime                    int             `json:"paid_time"`
		PaymentAmountBeforeExchange struct {
			Currency string          `json:"currency"`
			Value    decimal.Decimal `json:"value"`
		} `json:"payment_amount_before_exchange"`
		ReserveAmount struct {
			Currency string          `json:"currency"`
			Value    decimal.Decimal `json:"value"`
		} `json:"reserve_amount"`
		SettlementAmount struct {
			Currency string          `json:"currency"`
			Value    decimal.Decimal `json:"value"`
		} `json:"settlement_amount"`
		Status string `json:"status"`
	} `json:"payments"`
}

type OrderShipPackageParams []OrderShipPackage

type OrderShipPackage struct {
	PackageId      string       `json:"id"`              // 包裹ID
	HandoverMethod string       `json:"handover_method"` // PICKUP 运输提供商将从卖家的取件地址取件 DROP_OFF 卖家需要将包裹投递到指定地点
	PickupSlot     PickupSlot   `json:"pickup_slot"`     // 包裹取件时间段
	SelfShipment   SelfShipment `json:"self_shipment"`   // 仅卖家运输包裹需要
}
type PickupSlot struct {
	StartTime int64 `json:"start_time"` // 包裹取件时间段的开始日期
	EndTime   int64 `json:"end_time"`   // 包裹取件时间段的结束日期
}
type SelfShipment struct {
	TrackingNumber     string `json:"tracking_number"`
	ShippingProviderId string `json:"shipping_provider_id"`
}

// 获取店铺关联仓库列表
type LogisticsWarehousesResult struct {
	Warehouses []LogisticsWarehouses `json:"warehouses"`
}
type LogisticsWarehouses struct {
	WarehousesId string                     `json:"id"`         // 物流仓库ID
	IsDefault    bool                       `json:"is_default"` // 默认仓库。如果商品发布时没有指定仓库，则将使用默认仓库
	Name         string                     `json:"name"`       // 仓库名称
	Address      LogisticsWarehousesAddress `json:"address"`    // 仓库地址
	/*
		SubType：
		- DOMESTIC_WAREHOUSE ;既是目标市场又是卖方基地的国家/地区的仓库。
		- of a domestic seller ;国内卖家。
		- CB_OVERSEA_WAREHOUSE ;对于跨境卖家，目标市场的本地仓库。
		- CB_DIRECT_SHIPPING_WAREHOUSE ;对于跨境卖家，卖家在基地国家/地区（例如中国大陆或香港）的仓库
	*/
	SubType string `json:"sub_type"`
	/*
		Type：
		- SALES_WAREHOUSE：运输产品的仓库。
		- RETURN_WAREHOUSE：用于接收退货的仓库。
		您可以为发货和收货使用相同的仓库，但它们将是具有相同地址的不同仓库 ID
	*/
	Type string `json:"type"`
	/*
		EffectStatus:
			-ENABLED：(启用) 所有库存产品均可出售。
			-DISABLED：(已禁用) 所有库存产品均无法销售。
			-RESTRICTED：(受限) 仓库处于假期模式或订单限制状态。所有库存产品均无法销售。
			-Holiday mode ：(假期模式) 当卖家无法履行某个仓库的订单时，卖家可以在卖家中心为该仓库开启假期模式。
			-Order limit mode ：(订单限制模式) 当卖家违反抖音商城政策时，抖音商城将限制仓库可履行的订单量
	*/
	EffectStatus string `json:"effect_status"`
}
type LogisticsWarehousesAddress struct {
	City          string `json:"city"`           // 国家
	ContactPerson string `json:"contact_person"` // 联络人
	Distict       string `json:"distict"`        // 地区
	FullAddress   string `json:"full_address"`   // 完整地址
	PhoneNumber   string `json:"phone_number"`   // 手机号
	PostalCode    string `json:"postal_code"`    // 邮政编码
	Region        string `json:"region"`         // 地区
	RegionCode    string `json:"region_code"`    // 地区编码
	State         string `json:"state"`          // 洲
	Town          string `json:"town"`           // 镇
}

// 获取卖家指定仓库订阅的配送选项列表
type LogisticsWarehousesDeliveryOptionsResult struct {
	DeliveryOptions []LogisticsWarehousesDeliveryOptions `json:"delivery_options"`
}
type LogisticsWarehousesDeliveryOptions struct {
	DeliveryOptionsId string                        `json:"id"`
	Description       string                        `json:"description"` // 描述
	Name              string                        `json:"name"`        // 名称
	Type              string                        `json:"type"`
	WeightLimit       DeliveryOptionsWeightLimit    `json:"weight_limit"`
	DimensionLimit    DeliveryOptionsDimensionLimit `json:"dimension_limit"`
}
type DeliveryOptionsDimensionLimit struct {
	MaxHeight int    `json:"max_height"` // 最大高度
	MaxLength int    `json:"max_length"` // 最大长度
	MaxWidth  int    `json:"max_width"`  // 最大宽度
	Unit      string `json:"unit"`       // 单位 - CM - INCH
}
type DeliveryOptionsWeightLimit struct {
	MaxWeight int    `json:"max_weight"` // 最大宽度
	MinWeight int    `json:"min_weight"` // 最小宽度
	Unit      string `json:"unit"`       // 单位 - GRAM - POUND
}

// 获取指定配送选项对应的配送商
type LogisticsWarehousesDeliveryOptionsShipResult struct {
	ShippingProviders []LogisticsWarehousesDeliveryOptionsShip `json:"shipping_providers"`
}
type LogisticsWarehousesDeliveryOptionsShip struct {
	ShippingProviderId string `json:"id"`   // 供应商ID
	Name               string `json:"name"` // 供应商名称
}

// 获取包裹详情
type PackageDetail struct {
	CreateTime         int    `json:"create_time"`
	DeliveryOptionId   string `json:"delivery_option_id"`
	DeliveryOptionName string `json:"delivery_option_name"`
	Dimension          struct {
		Height string `json:"height"`
		Length string `json:"length"`
		Unit   string `json:"unit"`
		Width  string `json:"width"`
	} `json:"dimension"`
	HandoverMethod   string   `json:"handover_method"`
	HasMultiSkus     bool     `json:"has_multi_skus"`
	NoteTag          string   `json:"note_tag"`
	OrderLineItemIds []string `json:"order_line_item_ids"`
	Orders           []struct {
		Id   string `json:"id"`
		Skus []struct {
			Id       string `json:"id"`
			ImageUrl string `json:"image_url"`
			Name     string `json:"name"`
			Quantity int    `json:"quantity"`
		} `json:"skus"`
	} `json:"orders"`
	PackageId     string `json:"package_id"`
	PackageStatus string `json:"package_status"`
	PickupSlot    struct {
		EndTime   int `json:"end_time"`
		StartTime int `json:"start_time"`
	} `json:"pickup_slot"`
	RecipientAddress struct {
		AddressDetail string `json:"address_detail"`
		AddressLine1  string `json:"address_line1"`
		AddressLine2  string `json:"address_line2"`
		AddressLine3  string `json:"address_line3"`
		AddressLine4  string `json:"address_line4"`
		FullAddress   string `json:"full_address"`
		Name          string `json:"name"`
		PhoneNumber   string `json:"phone_number"`
		PostalCode    string `json:"postal_code"`
		RegionCode    string `json:"region_code"`
	} `json:"recipient_address"`
	SenderAddress struct {
		AddressDetail string `json:"address_detail"`
		AddressLine1  string `json:"address_line1"`
		AddressLine2  string `json:"address_line2"`
		AddressLine3  string `json:"address_line3"`
		AddressLine4  string `json:"address_line4"`
		FullAddress   string `json:"full_address"`
		Name          string `json:"name"`
		PhoneNumber   string `json:"phone_number"`
		PostalCode    string `json:"postal_code"`
		RegionCode    string `json:"region_code"`
	} `json:"sender_address"`
	ShippingProviderId   string `json:"shipping_provider_id"`
	ShippingProviderName string `json:"shipping_provider_name"`
	ShippingType         string `json:"shipping_type"`
	SplitAndCombineTag   string `json:"split_and_combine_tag"`
	TrackingNumber       string `json:"tracking_number"`
	UpdateTime           int    `json:"update_time"`
	Weight               struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"weight"`
}

type ShippingServices struct {
	Dimension struct {
		Height string `json:"height"`
		Length string `json:"length"`
		Unit   string `json:"unit"`
		Width  string `json:"width"`
	} `json:"dimension"`
	OrderId          string   `json:"order_id"`
	OrderLineId      []string `json:"order_line_id"`
	ShippingServices []struct {
		Currency             string `json:"currency"`
		EarliestDeliveryDays int    `json:"earliest_delivery_days"`
		Id                   string `json:"id"`
		IsDefault            bool   `json:"is_default"`
		LatestDeliveryDays   int    `json:"latest_delivery_days"`
		Name                 string `json:"name"`
		Price                string `json:"price"`
		ShippingProviderId   string `json:"shipping_provider_id"`
		ShippingProviderName string `json:"shipping_provider_name"`
	} `json:"shipping_services"`
	Weight struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"weight"`
}

type GlobalProductDetailResponse struct {
	Brand struct {
		Id string `json:"id"`
	} `json:"brand"`
	Category struct {
		Id string `json:"id"`
	} `json:"category"`
	Certifications []struct {
		Files []struct {
			Format string   `json:"format"`
			Id     string   `json:"id"`
			Name   string   `json:"name"`
			Urls   []string `json:"urls"`
		} `json:"files"`
		Id     string `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			Uri    string `json:"uri"`
			Width  int    `json:"width"`
		} `json:"images"`
		Title string `json:"title"`
	} `json:"certifications"`
	CreateTime     int    `json:"create_time"`
	Description    string `json:"description"`
	GlobalSellerId string `json:"global_seller_id"`
	Id             string `json:"id"`
	MainImages     []struct {
		Height int    `json:"height"`
		Uri    string `json:"uri"`
		Width  int    `json:"width"`
	} `json:"main_images"`
	Manufacturer struct {
		Address     string `json:"address"`
		Email       string `json:"email"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"manufacturer"`
	PackageDimensions struct {
		Height string `json:"height"`
		Length string `json:"length"`
		Unit   string `json:"unit"`
		Width  string `json:"width"`
	} `json:"package_dimensions"`
	PackageWeight struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"package_weight"`
	ProductAttributes []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Values []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"product_attributes"`
	Products []struct {
		Id     string `json:"id"`
		Region string `json:"region"`
	} `json:"products"`
	SizeChart struct {
		Image struct {
			Height int    `json:"height"`
			Uri    string `json:"uri"`
			Width  int    `json:"width"`
		} `json:"image"`
		Template struct {
			Id string `json:"id"`
		} `json:"template"`
	} `json:"size_chart"`
	Skus []struct {
		GlobalQuantity int    `json:"global_quantity"`
		Id             string `json:"id"`
		IdentifierCode struct {
			Code string `json:"code"`
			Type string `json:"type"`
		} `json:"identifier_code"`
		Inventory []struct {
			GlobalWarehouseId string `json:"global_warehouse_id"`
			Quantity          int    `json:"quantity"`
		} `json:"inventory"`
		Price struct {
			Amount    string `json:"amount"`
			Currency  string `json:"currency"`
			UnitPrice string `json:"unit_price"`
		} `json:"price"`
		SalesAttributes []struct {
			Id     string `json:"id"`
			Name   string `json:"name"`
			SkuImg struct {
				Height    int      `json:"height"`
				ThumbUrls []string `json:"thumb_urls"`
				Uri       string   `json:"uri"`
				Urls      []string `json:"urls"`
				Width     int      `json:"width"`
			} `json:"sku_img"`
			ValueId   string `json:"value_id"`
			ValueName string `json:"value_name"`
		} `json:"sales_attributes"`
		SellerSku    string `json:"seller_sku"`
		SkuUnitCount string `json:"sku_unit_count"`
	} `json:"skus"`
	Title      string `json:"title"`
	UpdateTime int    `json:"update_time"`
	Video      struct {
		Id string `json:"id"`
	} `json:"video"`
}

type CreateProduct struct {
	ProductId string `json:"product_id"`
	Skus      []struct {
		ExternalSkuId   string `json:"external_sku_id"`
		Id              string `json:"id"`
		SalesAttributes []struct {
			Id      string `json:"id"`
			ValueId string `json:"value_id"`
		} `json:"sales_attributes"`
		SellerSku string `json:"seller_sku"`
	} `json:"skus"`
	Warnings []struct {
		Message string `json:"message"`
	} `json:"warnings"`
}

type GlobalSku struct {
	Id              string `json:"id"`
	SalesAttributes []struct {
		Id      string `json:"id"`
		ValueId string `json:"value_id"`
	} `json:"sales_attributes"`
	SellerSku string `json:"seller_sku"`
}

type UploadProductImage struct {
	Height  int64  `json:"height"`
	Uri     string `json:"uri"`
	Url     string `json:"url"`
	UseCase string `json:"use_case"`
	Width   int64  `json:"width"`
}

type PublishGlobalProduct struct {
	Products []struct {
		Id     string `json:"id"`
		Region string `json:"region"`
		ShopId string `json:"shop_id"`
		Skus   []struct {
			Id                 string `json:"id"`
			RelatedGlobalSkuId string `json:"related_global_sku_id"`
			SaleAttributes     []struct {
				Id      string `json:"id"`
				ValueId string `json:"value_id"`
			} `json:"sale_attributes"`
			SellerSku string `json:"seller_sku"`
		} `json:"skus"`
	} `json:"products"`
	PublishResult []struct {
		FailReasons []struct {
			Message string `json:"message"`
		} `json:"fail_reasons"`
		Region string `json:"region"`
		Status string `json:"status"`
	} `json:"publish_result"`
}

type CategoriesAttributes struct {
	Attributes []struct {
		Id                  string `json:"id"`
		IsCustomizable      bool   `json:"is_customizable"`
		IsMultipleSelection bool   `json:"is_multiple_selection"`
		IsRequried          bool   `json:"is_requried"`
		Name                string `json:"name"`
		Type                string `json:"type"`
		Values              []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"attributes"`
}

type GetGlobalProductsCategories struct {
	Id                 string   `json:"id"`
	IsLeaf             bool     `json:"is_leaf"`
	LocalName          string   `json:"local_name"`
	ParentId           string   `json:"parent_id"`
	PermissionStatuses []string `json:"permission_statuses"`
}

type GlobalProduct struct {
	CreateTime int    `json:"create_time"`
	Id         string `json:"id"`
	Skus       []struct {
		Id        string `json:"id"`
		SellerSku string `json:"seller_sku"`
	} `json:"skus"`
	Status     string `json:"status"`
	Title      string `json:"title"`
	UpdateTime int    `json:"update_time"`
}
