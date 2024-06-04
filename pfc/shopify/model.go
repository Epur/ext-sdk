package shopify

import (
	"encoding/json"
)

// 交互令牌请求
type ExchangeTokenRequest struct {
	ClientId           string `json:"client_id"`
	ClientSecret       string `json:"client_secret"`
	GrantType          string `json:"grant_type"`
	SubjectToken       string `json:"subject_token"`
	SubjectTokenType   string `json:"subject_token_type"`
	RequestedTokenType string `json:"requested_token_type"`
}

// 交互令牌响应
type ExchangeTokenResponse struct {
	AccessToken         string          `json:"access_token"`
	Scope               string          `json:"scope"`
	ExpiresIn           string          `json:"expires_in"`
	AssociatedUserScope string          `json:"associated_user_scope"`
	associatedUser      json.RawMessage `json:"associated_user"`
}

// 访问令牌请求（在线）
type AccessTokenResponse struct {
	AccessToken         string          `json:"access_token"`
	Scope               string          `json:"scope"`
	ExpiresIn           string          `json:"expires_in"`
	AssociatedUserScope string          `json:"associated_user_scope"`
	associatedUser      json.RawMessage `json:"associated_user"`
}

// 检索订单请求
type OrderListRequest struct {
	ApiVersion       string `json:"api_version" validate:"required"`
	AttributionAppId string `json:"attribution_app_id,omitempty"` //Show orders attributed to a certain app, specified by the app ID. Set as current to show orders for the app currently consuming the API.
	CreatedAtMax     string `json:"created_at_max,omitempty"`     //Show orders created at or before date.
	CreatedAtMin     string `json:"created_at_min,omitempty"`     //Show orders created at or after date.
	Fields           string `json:"fields,omitempty"`             //Retrieve only certain fields, specified by a comma-separated list of fields names.
	FinancialStatus  string `json:"financial_status,omitempty"`   //authorized: Show only authorized orders
	//pending: Show only pending orders
	//paid: Show only paid orders
	//partially_paid: Show only partially paid orders
	//refunded: Show only refunded orders
	//voided: Show only voided orders
	//partially_refunded: Show only partially refunded orders
	//any: Show orders of any financial status.
	//unpaid: Show authorized and partially paid orders.
	FulfillmentStatus string `json:"fulfillment_status,omitempty"` //shipped: Show orders that have been shipped. Returns orders with fulfillment_status of fulfilled.
	//partial: Show partially shipped orders.
	//unshipped: Show orders that have not yet been shipped. Returns orders with fulfillment_status of null.
	//any: Show orders of any fulfillment status.
	//unfulfilled: Returns orders with fulfillment_status of null or partial.
	Ids            string `json:"ids,omitempty"`
	Limit          int64  `json:"limit,omitempty"`            //INT32 ≤ 250 default 50
	ProcessedAtMax string `json:"processed_at_max,omitempty"` //Show orders imported at or before date.
	ProcessedAtMin string `json:"processed_at_min,omitempty"` //Show orders imported at or after date.
	SinceId        int64  `json:"since_id,omitempty"`         //Show orders after the specified ID.
}

// 检索特定订单详情
type OrderDetailRequest struct {
	ApiVersion string `json:"api_version" validate:"required"`
	OrderId    string `json:"order_id" validate:"required"`
	Fields     string `json:"fields,omitempty"`
}
