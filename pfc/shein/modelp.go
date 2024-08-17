package shein

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type CommonResponse struct {
	Code string `json:"code" validate:"required"`
	Msg  string `json:"msg"`
}

type Response struct {
	CommonResponse
	Info json.RawMessage `json:"info"`
}

type OrderListResponse struct {
	CommonResponse
	Info *Info `json:"info"`
}

type Info struct {
	Count     int64   `json:"count"`
	OrderList []Order `json:"orderList"`
}
type Order struct {
	OrderNo         string `json:"orderNo"`
	OrderStatus     string `json:"orderStatus"`
	OrderCreateTime string `json:"orderCreateTime"`
	OrderUpdateTime string `json:"orderUpdateTime"`
}

type OrderDetailResponse struct {
	CommonResponse
	Info []OrderDetail `json:"info"`
}

type OrderDetail struct {
	OrderNo                       string                  `json:"orderNo"`
	UnProcessReason               []int64                 `json:"unProcessReason"`
	IsOverLimitOrder              int64                   `json:"isOverLimitOrder"`
	OrderType                     int64                   `json:"orderType"`
	PerformanceType               int8                    `json:"performanceType"`
	OrderStatus                   int8                    `json:"orderStatus"`
	IsCod                         int8                    `json:"isCod"`
	OrderTag                      int8                    `json:"orderTag"`
	DeliveryType                  int8                    `json:"deliveryType"`
	PrintOrderStatus              int8                    `json:"printOrderStatus"`
	InvoiceStatus                 int8                    `json:"invoiceStatus"`
	SettleActuallyPrice           decimal.Decimal         `json:"settleActuallyPrice"`
	OrderGoodsInfoList            []OrderGoodsInfo        `json:"orderGoodsInfoList"`
	PackageWaybillList            []PackageWaybill        `json:"packageWaybillList"`
	PackageInvoiceProblems        []PackageInvoiceProblem `json:"packageInvoiceProblems"`
	OrderCurrency                 string                  `json:"orderCurrency"`
	ProductTotalPrice             decimal.Decimal         `json:"productTotalPrice"`
	StoreDiscountTotalPrice       decimal.Decimal         `json:"storeDiscountTotalPrice"`
	PromotionDiscountTotalPrice   decimal.Decimal         `json:"promotionDiscountTotalPrice"`
	TotalCommission               decimal.Decimal         `json:"totalCommission"`
	TotalServiceCharge            decimal.Decimal         `json:"totalServiceCharge"`
	TotalPerformanceServiceCharge decimal.Decimal         `json:"totalPerformanceServiceCharge"`
	EstimatedGrossIncome          decimal.Decimal         `json:"estimatedGrossIncome"`
	TotalSaleTax                  decimal.Decimal         `json:"totalSaleTax"`
	OrderAllocateTime             string                  `json:"orderAllocateTime"`
	RequestDeliveryTime           string                  `json:"requestDeliveryTime"`
	PrintingTime                  string                  `json:"printingTime"`
	ScheduleDeliveryTime          string                  `json:"scheduleDeliveryTime"`
	PickUpTime                    string                  `json:"pickUpTime"`
	OrderReturnTime               string                  `json:"orderReturnTime"`
	OrderMsgUpdateTime            string                  `json:"orderMsgUpdateTime"`
	OrderTime                     string                  `json:"orderTime"`
	PaymentTime                   string                  `json:"paymentTime"`
	SellerDeliveryTime            string                  `json:"sellerDeliveryTime"`
	WarehouseDeliveryTime         string                  `json:"warehouseDeliveryTime"`
	OrderRejectionTime            string                  `json:"orderRejectionTime"`
	OrderReportedLossTime         string                  `json:"orderReportedLossTime"`
	BillNo                        string                  `json:"billNo"`
	SalesArea                     int8                    `json:"salesArea"`
	StockMode                     int8                    `json:"stockMode"`
	SalesSite                     string                  `json:"salesSite"`
	StoreCode                     int64                   `json:"storeCode"`
	ExpectedCollectTime           string                  `json:"expectedCollectTime"`
}

type OrderGoodsInfo struct {
	GoodsId        int64  `json:"goodsId"`
	SkuCode        string `json:"skuCode"`
	Skc            string `json:"skc"`
	GoodsSn        string `json:"goodsSn"`
	SellerSku      string `json:"sellerSku"`
	GoodsStatus    int8   `json:"goodsStatus"`
	NewGoodsStatus int64  `json:"newGoodsStatus"`

	SkuAttribute       []SkuAttribute `json:"skuAttribute"`
	GoodsTitle         string         `json:"goodsTitle"`
	SpuPicURL          string         `json:"spuPicURL"`
	StorageTag         int64          `json:"storageTag"`
	PerformanceTag     int64          `json:"performanceTag"`
	GoodsExchangeTag   int64          `json:"goodsExchangeTag"`
	BeExchangeEntityId int64          `json:"beExchangeEntityId"`

	OrderCurrency                 string          `json:"orderCurrency"`
	SellerCurrencyPrice           decimal.Decimal `json:"sellerCurrencyPrice"`
	OrderCurrencyStoreCouponPrice decimal.Decimal `json:"orderCurrencyStoreCouponPrice"`
	OrderCurrencyPromotionPrice   decimal.Decimal `json:"orderCurrencyPromotionPrice"`
	Commission                    decimal.Decimal `json:"commission"`
	CommissionRate                decimal.Decimal `json:"commissionRate"`
	ServiceCharge                 decimal.Decimal `json:"serviceCharge"`
	PerformanceServiceCharge      decimal.Decimal `json:"performanceServiceCharge"`
	EstimatedIncome               decimal.Decimal `json:"estimatedIncome"`
	SpuName                       string          `json:"spuName"`
	SaleTax                       decimal.Decimal `json:"saleTax"`
	WarehouseCode                 string          `json:"warehouseCode"`
	WarehouseName                 string          `json:"warehouseName"`
	SellerCurrencyDiscountPrice   decimal.Decimal `json:"sellerCurrencyDiscountPrice"`
	UnpackingGroupNo              string          `json:"unpackingGroupNo"`
	UnpackingGroupInvoiceStatus   int8            `json:"unpackingGroupInvoiceStatus"`
}
type PackageWaybill struct {
	PackageNo   string `json:"packageNo"`
	WaybillNo   string `json:"waybillNo"`
	Carrier     string `json:"carrier"`
	CarrierCode string `json:"carrierCode"`

	ProductInventoryList []ProductInventory `json:"productInventoryList"`
	PackageLabel         string             `json:"packageLabel"`
	SortingCode          string             `json:"sortingCode"`
	ExpressSortingCode   string             `json:"expressSortingCode"`
	IsCutOffSeller       int64              `json:"isCutOffSeller"`
}
type ProductInventory struct {
	ProductId string `json:"productId"`
}
type PackageInvoiceProblem struct {
	ProblemCode        string `json:"problemCode"`
	ProblemDescEnglish string `json:"problemDescEnglish"`
	ProblemField       string `json:"problemField"`
	ProposalEnglish    string `json:"proposalEnglish"`
	PackageNo          string `json:"packageNo"`
}

type SkuAttribute struct {
	AttrValueId string `json:"attrValueId"`
	AttrName    string `json:"attrName"`
	Language    string `json:"language"`
}

type GetByTokenResponse struct {
	CommonResponse
	Info *SecretKey `json:"info"`
}
type SecretKey struct {
	SecretKey string `json:"secretKey" validate:"required"`
	OpenKeyId string `json:"openKeyId" validate:"required"`
	Appid     string `json:"appid" validate:"required"`
	State     string `json:"state,omitempty"`
}
