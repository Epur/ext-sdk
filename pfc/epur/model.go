package epur

import "encoding/json"

const (
	AUTHSITE      = `/openapi/uaa/auth/oauth2/authorize`
	ACCESSTOKEN   = `/openapi/uaa/auth/oauth2/accessToken`
	REFRESHTOKEN  = `/openapi/uaa/auth/oauth2/refreshToken`
	PAYAPPLY      = `/openapi/pur/v1/order/pay/apply`
	MERCHANT_INFO = `/openapi/usercenter/v1/merchant/user`
)

type Response struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}
