package amazon

import "encoding/json"

type GetTokenResponse struct {
	ExpiresIn        int64  `json:"expires_in"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type GetSellerResponse struct {
	*Response
	Payload []MarketplaceParticipation `json:"payload"`
}

type GetOrderResponse struct {
	*Response
	AmazonOrderId                  string `json:"AmazonOrderId"`
	SellerOrderId                  string `json:"SellerOrderId"`
	PurchaseDate                   string `json:"PurchaseDate"`
	LastUpdateDate                 string `json:"LastUpdateDate"`
	OrderStatus                    string `json:"OrderStatus"`
	FulfillmentChannel             string `json:"FulfillmentChannel"`
	SalesChannel                   string `json:"SalesChannel"`
	OrderChannel                   string `json:"OrderChannel"`
	ShipServiceLevel               string `json:"ShipServiceLevel"`
	OrderTotal                     string `json:"OrderTotal"`
	NumberOfItemsShipped           string `json:"NumberOfItemsShipped"`
	NumberOfItemsUnshipped         string `json:"NumberOfItemsUnshipped"`
	PaymentExecutionDetail         string `json:"PaymentExecutionDetail"`
	PaymentMethod                  string `json:"PaymentMethod"`
	PaymentMethodDetails           string `json:"PaymentMethodDetails"`
	MarketplaceId                  string `json:"MarketplaceId"`
	ShipmentServiceLevelCategory   string `json:"ShipmentServiceLevelCategory"`
	EasyShipShipmentStatus         string `json:"EasyShipShipmentStatus"`
	CbaDisplayableShippingLabel    string `json:"CbaDisplayableShippingLabel"`
	OrderType                      string `json:"OrderType"`
	EarliestShipDate               string `json:"EarliestShipDate"`
	LatestShipDate                 string `json:"LatestShipDate"`
	EarliestDeliveryDate           string `json:"EarliestDeliveryDate"`
	LatestDeliveryDate             string `json:"LatestDeliveryDate"`
	IsBusinessOrder                string `json:"IsBusinessOrder"`
	IsPrime                        string `json:"IsPrime"`
	IsPremiumOrder                 string `json:"IsPremiumOrder"`
	IsGlobalExpressEnabled         string `json:"IsGlobalExpressEnabled"`
	ReplacedOrderId                string `json:"ReplacedOrderId"`
	IsReplacementOrder             string `json:"IsReplacementOrder"`
	PromiseResponseDueDate         string `json:"PromiseResponseDueDate"`
	IsEstimatedShipDateSet         string `json:"IsEstimatedShipDateSet"`
	IsSoldByAB                     string `json:"IsSoldByAB"`
	IsIBA                          string `json:"IsIBA"`
	DefaultShipFromLocationAddress string `json:"DefaultShipFromLocationAddress"`
	BuyerInvoicePreference         string `json:"BuyerInvoicePreference"`
	BuyerTaxInformation            string `json:"BuyerTaxInformation"`
	FulfillmentInstruction         string `json:"FulfillmentInstruction"`
	IsISPU                         string `json:"IsISPU"`
	IsAccessPointOrder             string `json:"IsAccessPointOrder"`
	MarketplaceTaxInfo             string `json:"MarketplaceTaxInfo"`
	SellerDisplayName              string `json:"SellerDisplayName"`
	ShippingAddress                string `json:"ShippingAddress"`
	BuyerInfo                      string `json:"BuyerInfo"`
	AutomatedShippingSettings      string `json:"AutomatedShippingSettings"`
	HasRegulatedItems              string `json:"HasRegulatedItems"`
	ElectronicInvoiceStatus        string `json:"ElectronicInvoiceStatus"`
}

type GetOrderListResponse struct {
	*Response
	OrderList  []Order `json:"order_list"`
	Count      int     `json:"count"`
	CountTotal int     `json:"count_total"`
}

type Order struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

type GetOrderDetailResponse struct {
	*Response
	OrderDetail OrderDetail `json:"order_detail"`
}

type OrderDetail struct {
	OrderId     string    `json:"order_id"`
	ProductList []Product `json:"product_list"`
}

type Product struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

type Response struct {
	Errors  []Error         `json:"errors"`
	Payload json.RawMessage `json:"payload"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type MarketplaceParticipation struct {
	Marketplace   Marketplace   `json:"marketplace"`
	Participation Participation `json:"participation"`
}

type Marketplace struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	CountryCode         string `json:"countryCode"`
	DefaultCurrencyCode string `json:"defaultCurrencyCode"`
	DefaultLanguageCode string `json:"defaultLanguageCode"`
	DomainName          string `json:"domainName"`
}

type Participation struct {
	IsParticipating      bool `json:"isParticipating"`
	HasSuspendedListings bool `json:"hasSuspendedListings"`
}
