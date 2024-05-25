package amazon

import "encoding/json"

type Response struct {
	Errors  []Error         `json:"errors"`
	Payload json.RawMessage `json:"payload"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

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

type GetOrdersResponse struct {
	*Response
	Payload OrdersPayload `json:"payload"`
}

type GetOrderResponse struct {
	*Response
	Payload Order `json:"payload"`
}

type GetOrderDetailResponse struct {
	*Response
	OrderDetail OrderDetail `json:"order_detail"`
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

type OrdersPayload struct {
	NextToken         string  `json:"NextToken"`
	CreatedBefore     string  `json:"CreatedBefore"`
	LastUpdatedBefore string  `json:"LastUpdatedBefore"`
	OrderList         []Order `json:"Orders"`
}

type Order struct {
	AmazonOrderId      string `json:"AmazonOrderId"`
	SellerOrderId      string `json:"SellerOrderId"`
	PurchaseDate       string `json:"PurchaseDate"`
	LastUpdateDate     string `json:"LastUpdateDate"`
	OrderStatus        string `json:"OrderStatus"`
	FulfillmentChannel string `json:"FulfillmentChannel"`
	SalesChannel       string `json:"SalesChannel"`
	OrderChannel       string `json:"OrderChannel"`
	ShipServiceLevel   string `json:"ShipServiceLevel"`
	OrderTotal         struct {
		CurrencyCode string `json:"CurrencyCode"`
		Amount       string `json:"Amount"`
	} `json:"OrderTotal"`
	NumberOfItemsShipped           int64    `json:"NumberOfItemsShipped"`
	NumberOfItemsUnshipped         int64    `json:"NumberOfItemsUnshipped"`
	PaymentExecutionDetail         string   `json:"PaymentExecutionDetail"`
	PaymentMethod                  string   `json:"PaymentMethod"`
	PaymentMethodDetails           []string `json:"PaymentMethodDetails"`
	MarketplaceId                  string   `json:"MarketplaceId"`
	ShipmentServiceLevelCategory   string   `json:"ShipmentServiceLevelCategory"`
	EasyShipShipmentStatus         string   `json:"EasyShipShipmentStatus"`
	CbaDisplayableShippingLabel    string   `json:"CbaDisplayableShippingLabel"`
	OrderType                      string   `json:"OrderType"`
	EarliestShipDate               string   `json:"EarliestShipDate"`
	LatestShipDate                 string   `json:"LatestShipDate"`
	EarliestDeliveryDate           string   `json:"EarliestDeliveryDate"`
	LatestDeliveryDate             string   `json:"LatestDeliveryDate"`
	IsBusinessOrder                bool     `json:"IsBusinessOrder"`
	IsPrime                        bool     `json:"IsPrime"`
	IsPremiumOrder                 bool     `json:"IsPremiumOrder"`
	IsGlobalExpressEnabled         bool     `json:"IsGlobalExpressEnabled"`
	ReplacedOrderId                string   `json:"ReplacedOrderId"`
	IsReplacementOrder             bool     `json:"IsReplacementOrder"`
	PromiseResponseDueDate         string   `json:"PromiseResponseDueDate"`
	IsEstimatedShipDateSet         bool     `json:"IsEstimatedShipDateSet"`
	IsSoldByAB                     bool     `json:"IsSoldByAB"`
	IsIBA                          bool     `json:"IsIBA"`
	DefaultShipFromLocationAddress struct {
		Name          string `json:"Name"`
		AddressLine1  string `json:"AddressLine1"`
		City          string `json:"City"`
		StateOrRegion string `json:"StateOrRegion"`
		PostalCode    string `json:"PostalCode"`
		CountryCode   string `json:"CountryCode"`
		Phone         string `json:"Phone"`
		AddressType   string `json:"AddressType"`
	} `json:"DefaultShipFromLocationAddress"`
	BuyerInvoicePreference string `json:"BuyerInvoicePreference"`
	BuyerTaxInformation    string `json:"BuyerTaxInformation"`
	FulfillmentInstruction struct {
		FulfillmentSupplySourceId string `json:"FulfillmentSupplySourceId"`
	} `json:"FulfillmentInstruction"`
	IsISPU                    bool   `json:"IsISPU"`
	IsAccessPointOrder        bool   `json:"IsAccessPointOrder"`
	MarketplaceTaxInfo        string `json:"MarketplaceTaxInfo"`
	SellerDisplayName         string `json:"SellerDisplayName"`
	ShippingAddress           string `json:"ShippingAddress"`
	BuyerInfo                 string `json:"BuyerInfo"`
	AutomatedShippingSettings struct {
		HasAutomatedShippingSettings bool `json:"HasAutomatedShippingSettings"`
	} `json:"AutomatedShippingSettings"`
	HasRegulatedItems       string `json:"HasRegulatedItems"`
	ElectronicInvoiceStatus string `json:"ElectronicInvoiceStatus"`
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
