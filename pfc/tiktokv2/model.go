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

	/*
		status
			1-draft、
			2-pending、
			3-failed(initial creation)、
			4-live、
			5-seller_deactivated、
			6-platform_deactivated、
			7-freeze
			8-deleted
	*/
	ProductId     string `json:"product_id"`
	ProductStatus int64  `json:"product_status"`
	ProductName   string `json:"product_name"`
	CategoryList  []struct {
		Id       string `json:"id"`
		ParentId string `json:"parent_id"`
		Leaf     bool   `json:"is_leaf"`
		Name     string `json:"local_display_name"`
	} `json:"category_list"`
	Brand struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Status int64  `json:"status"`
	} `json:"brand"`
	Images []ImageEx `json:"images"`
	Video  struct {
		Id         string      `json:"id"`
		Duration   interface{} `json:"duration"`
		PostUrl    string      `json:"post_url"`
		MediaType  string      `json:"media_type"`
		VideoInfos []struct {
			MainUrl   string `json:"main_url"`
			BackupUrl string `json:"backup_url"`
			UrlExpire int64  `json:"url_expire"`
			Width     int64  `json:"width"`
			Height    int64  `json:"height"`
			FileHash  string `json:"file_hash"`
			Format    string `json:"format"`
			Size      int64  `json:"size"`
			Bitrate   int64  `json:"bitrate"`
		} `json:"video_infos"`
	} `json:"video"`
	Description    string `json:"description"`
	WarrantyPeriod struct {
		WarrantyId          int64  `json:"warranty_id"`
		WarrantyDescription string `json:"warranty_description"`
	} `json:"warranty_period"`
	WarrantyPolicy string `json:"warranty_policy"`
	PackageLength  int64  `json:"package_length"`
	PackageWidth   int64  `json:"package_width"`
	PackageHeight  int64  `json:"package_height"`
	PackageWeight  string `json:"package_weight"`
	Skus           []struct {
		Id        string `json:"id"`
		SellerSku string `json:"seller_sku"`
		Price     struct {
			OriginalPrice   decimal.Decimal `json:"original_price"`
			PriceIncludeVat string          `json:"price_include_vat"`
			Currency        string          `json:"currency"`
		} `json:"price"`
		StockInfos []struct {
			AvailableStock int64  `json:"available_stock"`
			WarehouseId    string `json:"warehouse_id"`
		} `json:"stock_infos"`
		SalesAttributes []struct {
			Id        string  `json:"id"`
			Name      string  `json:"name"`
			ValueName string  `json:"value_name"`
			ValueId   string  `json:"value_id"`
			SkuImg    ImageEx `json:"sku_img"`
		} `json:"sales_attributes"`
	} `json:"skus"`
	ProductCertifications []struct {
		Files []struct {
			Id   string   `json:"id"`
			Name string   `json:"name"`
			Type string   `json:"type"`
			List []string `json:"list"`
		}
		Id     string    `json:"id"`
		Title  string    `json:"title"`
		Images []ImageEx `json:"images"`
	} `json:"product_certifications"`
	SizeChart         ImageEx `json:"size_chart"`
	IsCodOpen         bool    `json:"is_cod_open"`
	ProductAttributes []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Values []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"product_attributes"`
	UpdateTime int64 `json:"update_time"`
	CreateTime int64 `json:"create_time"`
	QcReasons  []struct {
		Reason     string   `json:"reason"`
		SubReasons []string `json:"sub_reasons"`
	} `json:"qc_reasons"`

	DeliveryServices []struct {
		DeliveryServiceId     interface{} `json:"delivery_service_id"`
		DeliveryOptionName    string      `json:"delivery_option_name"`
		DeliveryServiceStatus bool        `json:"delivery_service_status"`
	} `json:"delivery_services"`
	ExemptionOfIdentifierCode struct {
		ExemptionReason []int64 `json:"exemption_reason"`
	} `json:"exemption_of_identifier_code"`
	PackageDimensionUnit string `json:"package_dimension_unit"`
}

type ImageEx struct {
	Height       int64    `json:"height"`
	Width        int64    `json:"width"`
	ThumbUrlList []string `json:"thumb_url_list"`
	Id           string   `json:"id"`
	UrlList      []string `json:"url_List"`
}
