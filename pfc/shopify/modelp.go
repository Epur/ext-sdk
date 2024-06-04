package shopify

import "github.com/shopspring/decimal"

// 订单列表响应
type OrderListResponse struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	TotalPrice string `json:"total_price"`
}

// 订单详情响应

type OrderDetailResponse struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	TotalPrice string     `json:"total_price"`
	LineItems  []LineItem `json:"line_items"`
}

type LineItem struct {
	Id                         int64                `json:"id"`
	AdminGraphqlApiId          string               `json:"admin_graphql_api_id"`
	CurrentQuantity            int64                `json:"current_quantity"`
	FulfillableQuantity        int64                `json:"fulfillable_quantity"`
	FulfillmentService         string               `json:"fulfillment_service"`
	FulfillmentStatus          string               `json:"fulfillment_status"`
	GiftCard                   bool                 `json:"gift_card"`
	Grams                      int64                `json:"grams"`
	Name                       string               `json:"name"`
	Price                      string               `json:"price"`
	PriceSet                   *PriceSet            `json:"price_set"`
	ProductExists              bool                 `json:"product_exists"`
	ProductId                  int64                `json:"product_id"`
	Properties                 []Property           `json:"properties"`
	Quantity                   int64                `json:"quantity"`
	RequiresShipping           bool                 `json:"requires_shipping"`
	Sku                        string               `json:"sku"`
	Taxable                    bool                 `json:"taxable"`
	Title                      string               `json:"title"`
	TotalDiscount              string               `json:"total_discount"`
	TotalDiscountSet           *PriceSet            `json:"total_discount_set"`
	VariantId                  int64                `json:"variant_id"`
	VariantInventoryManagement string               `json:"variant_inventory_management"`
	VariantTitle               string               `json:"variant_title"`
	Vendor                     string               `json:"vendor"`
	TaxLines                   []TaxLine            `json:"tax_lines"`
	Duties                     []string             `json:"duties"`
	DiscountAllocations        []DiscountAllocation `json:"discount_allocations"`
}

type PriceSet struct {
	ShopMoney        *Money `json:"shop_money"`
	PresentmentMoney *Money `json:"presentment_money"`
}

type Money struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TaxLine struct {
	Name     string          `json:"channel_liable"`
	Price    string          `json:"price"`
	PriceSet *PriceSet       `json:"price_set"`
	Rate     decimal.Decimal `json:"rate"`
	Title    string          `json:"title"`
}

type DiscountAllocation struct {
	Amount                   string    `json:"amount"`
	AmountSet                *PriceSet `json:"amount_set"`
	DiscountApplicationIndex int       `json:"discount_application_index"`
}

// 店铺配置信息响应

type ShopConfigResponse struct {
	Id                                   int64           `json:"id"`
	Name                                 string          `json:"name"`
	Email                                string          `json:"email,omitempty"`
	Domain                               string          `json:"domain"`
	Province                             string          `json:"province"`
	Country                              string          `json:"country,omitempty"`
	Address                              string          `json:"address1"`
	Zip                                  string          `json:"zip"`
	City                                 string          `json:"city,omitempty"`
	Source                               string          `json:"source"`
	Phone                                string          `json:"phone"`
	Latitude                             decimal.Decimal `json:"latitude,omitempty"`
	Longitude                            decimal.Decimal `json:"longitude"`
	PrimaryLocale                        string          `json:"primary_locale"`
	Address2                             string          `json:"address2,omitempty"`
	CreateAt                             string          `json:"create_at"`
	UpdateAt                             string          `json:"update_at"`
	CountryCode                          string          `json:"country_code,omitempty"`
	CountryName                          string          `json:"country_name"`
	Currency                             string          `json:"currency"`
	CustomerEmail                        string          `json:"customer_email,omitempty"`
	Timezone                             string          `json:"timezone"`
	IanaTimezone                         string          `json:"iana_timezone"`
	ShopOwner                            string          `json:"shop_owner,omitempty"`
	MoneyFormat                          string          `json:"money_format"`
	MoneyWithCurrencyFormat              string          `json:"money_with_currency_format"`
	WeightUnit                           string          `json:"weight_unit,omitempty"`
	ProvinceCode                         string          `json:"province_code"`
	TaxesInclude                         string          `json:"taxes_include"`
	AutoConfigureTaxInclusivity          string          `json:"auto_configure_tax_inclusivity"`
	TaxShipping                          string          `json:"tax_shipping"`
	CountryTaxes                         string          `json:"county_taxes"`
	PlanDisplayName                      string          `json:"plan_display_name"`
	PlanName                             string          `json:"plan_name"`
	HasDiscounts                         bool            `json:"has_discounts"`
	HasGiftCards                         bool            `json:"has_gift_cards"`
	MyshopifyDomain                      string          `json:"myshopify_domain,omitempty"`
	GoogleAppsDomain                     string          `json:"google_apps_domain"`
	GoogleAppsLoginEnabled               string          `json:"goole_apps_login_enabled"`
	MoneyInEmailsFormat                  string          `json:"money_in_emails_format,omitempty"`
	MoneyWithCurrencyInEmailsFormat      string          `json:"money_with_currency_in_emails_format,omitempty"`
	EligibleForPayments                  bool            `json:"eligible_for_payments"`
	RequiresExtraPaymentsAgreement       bool            `json:"requires_extra_payments_agreement"`
	PasswordEnabled                      bool            `json:"password_enabled"`
	HasStorefront                        bool            `json:"has_storefront,omitempty"`
	Finances                             bool            `json:"finances"`
	PrimaryLocationId                    int64           `json:"primary_location_id"`
	CheckoutApiSupported                 bool            `json:"checkout_api_supported,omitempty"`
	MultiLocationEnabled                 bool            `json:"multi_location_enabled"`
	SetupRequired                        bool            `json:"setup_required"`
	PreLaunchEnabled                     bool            `json:"pre_launch_enabled,omitempty"`
	EnabledPresentmentCurrencies         []string        `json:"enabled_presentment_currencies,omitempty"`
	TransactionalSmsDisabled             bool            `json:"transactional_sms_disabled,omitempty"`
	MarketingSmsConsentEnabledAtCheckout bool            `json:"marketing_sms_consent_enabled_at_checkout,omitempty"`
}
