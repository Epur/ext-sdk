package amazon

import (
	"encoding/json"
	"time"
)

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

type GetOrderItemsResponse struct {
	*Response
	Payload OrderItemsPayload `json:"payload"`
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

type OrderItemsPayload struct {
	NextToken     string      `json:"NextToken"`
	AmazonOrderId string      `json:"AmazonOrderId"`
	OrderItems    []OrderItem `json:"OrderItems"`
}

type OrderItem struct {
	ASIN            string        `json:"asin"`
	SellerSKU       string        `json:"sellerSKU,omitempty"`
	OrderItemId     string        `json:"orderItemId"`
	AssociatedItems []interface{} `json:"associatedItems,omitempty"`
	Title           string        `json:"title,omitempty"`
	QuantityOrdered int           `json:"quantityOrdered"`
	QuantityShipped int           `json:"quantityShipped,omitempty"`
	ProductInfo     struct {
		NumberOfItems string `json:"numberOfItems"`
	} `json:"productInfo,omitempty"`
	PointsGranted struct {
		PointsNumber        int   `json:"pointsNumber"`
		PointsMonetaryValue Money `json:"pointsMonetaryValue"`
	} `json:"pointsGranted,omitempty"`
	ItemPrice                  Money       `json:"itemPrice,omitempty"`
	ShippingPrice              Money       `json:"shippingPrice,omitempty"`
	ItemTax                    Money       `json:"itemTax,omitempty"`
	ShippingTax                Money       `json:"shippingTax,omitempty"`
	ShippingDiscount           Money       `json:"shippingDiscount,omitempty"`
	ShippingDiscountTax        Money       `json:"shippingDiscountTax,omitempty"`
	PromotionDiscount          Money       `json:"promotionDiscount,omitempty"`
	PromotionDiscountTax       Money       `json:"promotionDiscountTax,omitempty"`
	PromotionIds               []string    `json:"promotionIds,omitempty"`
	CODFee                     Money       `json:"codFee,omitempty"`
	CODFeeDiscount             Money       `json:"codFeeDiscount,omitempty"`
	IsGift                     string      `json:"isGift,omitempty"`
	ConditionNote              string      `json:"conditionNote,omitempty"`
	ConditionId                string      `json:"conditionId,omitempty"`
	ConditionSubtypeId         string      `json:"conditionSubtypeId,omitempty"`
	ScheduledDeliveryStartDate *time.Time  `json:"scheduledDeliveryStartDate,omitempty"`
	ScheduledDeliveryEndDate   *time.Time  `json:"scheduledDeliveryEndDate,omitempty"`
	PriceDesignation           string      `json:"priceDesignation,omitempty"`
	TaxCollection              interface{} `json:"taxCollection,omitempty"`
	SerialNumberRequired       bool        `json:"serialNumberRequired,omitempty"`
	IsTransparency             bool        `json:"isTransparency,omitempty"`
	IossNumber                 string      `json:"iossNumber,omitempty"`
	StoreChainStoreId          string      `json:"storeChainStoreId,omitempty"`
	DeemedResellerCategory     string      `json:"deemedResellerCategory,omitempty"`
	BuyerInfo                  struct {
		BuyerEmail          string      `json:"buyerEmail"`
		BuyerName           string      `json:"buyerName"`
		BuyerCounty         string      `json:"buyerCounty"`
		BuyerTaxInfo        interface{} `json:"buyerTaxInfo"`
		PurchaseOrderNumber string      `json:"purchaseOrderNumber"`
	} `json:"buyerInfo,omitempty"`
	BuyerRequestedCancel struct {
		IsBuyerRequestedCancel string `json:"isBuyerRequestedCancel"`
		BuyerCancelReason      string `json:"buyerCancelReason"`
	} `json:"buyerRequestedCancel,omitempty"`
	SerialNumbers           []string    `json:"serialNumbers,omitempty"`
	SubstitutionPreferences interface{} `json:"substitutionPreferences,omitempty"`
	Measurement             string      `json:"measurement,omitempty"`
	ShippingConstraints     interface{} `json:"shippingConstraints,omitempty"`
}

type Order struct {
	AmazonOrderId                  string      `json:"AmazonOrderId"`
	SellerOrderId                  string      `json:"SellerOrderId"`
	PurchaseDate                   string      `json:"PurchaseDate"`
	LastUpdateDate                 string      `json:"LastUpdateDate"`
	OrderStatus                    string      `json:"OrderStatus"`
	FulfillmentChannel             string      `json:"FulfillmentChannel"`
	SalesChannel                   string      `json:"SalesChannel"`
	OrderChannel                   string      `json:"OrderChannel"`
	ShipServiceLevel               string      `json:"ShipServiceLevel"`
	OrderTotal                     Money       `json:"OrderTotal"`
	NumberOfItemsShipped           int64       `json:"NumberOfItemsShipped"`
	NumberOfItemsUnshipped         int64       `json:"NumberOfItemsUnshipped"`
	PaymentExecutionDetail         string      `json:"PaymentExecutionDetail"`
	PaymentMethod                  string      `json:"PaymentMethod"`
	PaymentMethodDetails           []string    `json:"PaymentMethodDetails"`
	MarketplaceId                  string      `json:"MarketplaceId"`
	ShipmentServiceLevelCategory   string      `json:"ShipmentServiceLevelCategory"`
	EasyShipShipmentStatus         string      `json:"EasyShipShipmentStatus"`
	CbaDisplayableShippingLabel    string      `json:"CbaDisplayableShippingLabel"`
	OrderType                      string      `json:"OrderType"`
	EarliestShipDate               string      `json:"EarliestShipDate"`
	LatestShipDate                 string      `json:"LatestShipDate"`
	EarliestDeliveryDate           string      `json:"EarliestDeliveryDate"`
	LatestDeliveryDate             string      `json:"LatestDeliveryDate"`
	IsBusinessOrder                bool        `json:"IsBusinessOrder"`
	IsPrime                        bool        `json:"IsPrime"`
	IsPremiumOrder                 bool        `json:"IsPremiumOrder"`
	IsGlobalExpressEnabled         bool        `json:"IsGlobalExpressEnabled"`
	ReplacedOrderId                string      `json:"ReplacedOrderId"`
	IsReplacementOrder             interface{} `json:"IsReplacementOrder"`
	PromiseResponseDueDate         string      `json:"PromiseResponseDueDate"`
	IsEstimatedShipDateSet         bool        `json:"IsEstimatedShipDateSet"`
	IsSoldByAB                     bool        `json:"IsSoldByAB"`
	IsIBA                          bool        `json:"IsIBA"`
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
	IsISPU             bool   `json:"IsISPU"`
	IsAccessPointOrder bool   `json:"IsAccessPointOrder"`
	MarketplaceTaxInfo string `json:"MarketplaceTaxInfo"`
	SellerDisplayName  string `json:"SellerDisplayName"`
	ShippingAddress    struct {
		StateOrRegion string `json:"StateOrRegion"`
		PostalCode    string `json:"PostalCode"`
		City          string `json:"City"`
		CountryCode   string `json:"CountryCode"`
	} `json:"ShippingAddress"`
	BuyerInfo struct {
		BuyerEmail string `json:"BuyerEmail"`
	} `json:"BuyerInfo"`
	AutomatedShippingSettings struct {
		HasAutomatedShippingSettings bool `json:"HasAutomatedShippingSettings"`
	} `json:"AutomatedShippingSettings"`
	HasRegulatedItems       bool   `json:"HasRegulatedItems"`
	ElectronicInvoiceStatus string `json:"ElectronicInvoiceStatus"`
}

type Money struct {
	CurrencyCode string `json:"CurrencyCode"`
	Amount       string `json:"Amount"`
}

// 通用的通知结构体
type Notification struct {
	NotificationVersion  string      `json:"notificationVersion"`
	NotificationType     string      `json:"notificationType"` // ORDER_CHANGE 等
	PayloadVersion       string      `json:"payloadVersion"`
	EventTime            time.Time   `json:"eventTime"`
	Payload              interface{} `json:"payload"` // Payload可以是ReportProcessingFinishedPayload或OrderChangePayload
	NotificationMetadata `json:"notificationMetadata"`
}

// 通用的通知元数据结构体
type NotificationMetadata struct {
	ApplicationId  string    `json:"applicationId"`
	SubscriptionId string    `json:"subscriptionId"`
	PublishTime    time.Time `json:"publishTime"`
	NotificationId string    `json:"notificationId"`
}

// 报告完成的通知的Payload结构体
type ReportProcessingFinishedPayload struct {
	ReportProcessingFinishedNotification struct {
		SellerId         string `json:"sellerId"`
		AccountId        string `json:"accountId"`
		ReportId         string `json:"reportId"`
		ReportType       string `json:"reportType"`
		ProcessingStatus string `json:"processingStatus"`
		ReportDocumentId string `json:"reportDocumentId"`
	} `json:"reportProcessingFinishedNotification"`
}

// 轮换LWA Secret的Payload结构体
type ApplicationOauthClientSecretExpiryPayload struct {
	ApplicationOAuthClientSecretExpiry struct {
		ClientId                 string `json:"clientId"`
		ClientSecretExpiryTime   string `json:"clientSecretExpiryTime"`
		ClientSecretExpiryReason string `json:"clientSecretExpiryReason"`
	} `json:"applicationOAuthClientSecretExpiry"`
}

// 获取最新LWA Secret 的Payload结构体
type ApplicationOAuthClientNewSecretPayload struct {
	ApplicationOAuthClientNewSecret struct {
		ClientId                  string `json:"clientId"`
		NewClientSecret           string `json:"newClientSecret"`
		NewClientSecretExpiryTime string `json:"newClientSecretExpiryTime"`
		OldClientSecretExpiryTime string `json:"oldClientSecretExpiryTime"`
	} `json:"applicationOAuthClientNewSecret"`
}

// 订单变更的通知的Payload结构体
type OrderChangePayload struct {
	OrderChangeNotification struct {
		NotificationLevel  string `json:"notificationLevel"`
		SellerId           string `json:"sellerId"`
		AmazonOrderId      string `json:"amazonOrderId"`
		OrderChangeType    string `json:"orderChangeType"`
		OrderChangeTrigger struct {
			TimeOfOrderChange time.Time `json:"timeOfOrderChange"`
			ChangeReason      string    `json:"changeReason"`
		} `json:"orderChangeTrigger"`
		Summary struct {
			MarketplaceId         string        `json:"marketplaceId"`
			OrderStatus           string        `json:"orderStatus"`
			PurchaseDate          time.Time     `json:"purchaseDate"`
			DestinationPostalCode *string       `json:"destinationPostalCode"`
			FulfillmentType       string        `json:"fulfillmentType"`
			OrderType             string        `json:"orderType"`
			OrderPrograms         []interface{} `json:"orderPrograms"`
			ShippingPrograms      []interface{} `json:"shippingPrograms"`
			OrderItems            []struct {
				OrderItemId    string  `json:"orderItemId"`
				SellerSKU      string  `json:"sellerSKU"`
				SupplySourceId *string `json:"supplySourceId"`
				Quantity       int     `json:"quantity"`
			} `json:"orderItems"`
		} `json:"summary"`
	} `json:"orderChangeNotification"`
}

// 根据NotificationType确定Payload的具体类型
func (n *Notification) UnmarshalJSON(data []byte) error {
	// 定义一个辅助结构体，用于先解析除Payload外的其他字段
	type Alias Notification
	aux := &struct {
		Payload json.RawMessage `json:"payload"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// 根据NotificationType字段来决定Payload的类型
	switch n.NotificationType {
	case REPORT_PROCESSING_FINISHED:
		n.Payload = &ReportProcessingFinishedPayload{}
	case APPLICATION_OAUTH_CLIENT_SECRET_EXPIRY:
		n.Payload = &ApplicationOauthClientSecretExpiryPayload{}
	case APPLICATION_OAUTH_CLIENT_NEW_SECRET:
		n.Payload = &ApplicationOAuthClientNewSecretPayload{}
	case ORDER_CHANGE:
		n.Payload = &OrderChangePayload{}
	default:
		n.Payload = nil
	}

	// 如果Payload不为空，则继续解析Payload字段
	if n.Payload != nil {
		return json.Unmarshal(aux.Payload, n.Payload)
	}

	return nil
}
