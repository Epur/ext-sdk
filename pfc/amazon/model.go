package amazon

import (
	"fmt"
	"net/url"
)

const (
	APIGatewayUS = "https://sellingpartnerapi-na.amazon.com"
	APIGatewayEU = "https://sellingpartnerapi-eu.amazon.com"
	APIGatewayFE = "https://sellingpartnerapi-fe.amazon.com"
	//APIGatewayUS = "https://sandbox.sellingpartnerapi-na.amazon.com"
	//APIGatewayEU = "https://sandbox.sellingpartnerapi-eu.amazon.com"
	//APIGatewayFE = "https://sandbox.sellingpartnerapi-fe.amazon.com"
	TokenURL = "https://api.amazon.com/auth/o2/token"
	// AuthURL 默认的授权链接，其他分区域的，通过实现 getAuthBase() 调用得到
	AuthURL = "https://sellercentral.amazon.com/apps/authorize/consent"
)

type Setting struct {
	ClientID        *string
	ClientSecret    *string
	AuthCallbackUrl *string
	SiteNo          *string
	AccessToken     *string
	RefreshToken    *string
}

type BodyMap map[string]interface{}

func (bm BodyMap) Set(key string, value interface{}) BodyMap {
	bm[key] = value
	return bm
}

func (bm BodyMap) CheckEmptyError(keys ...string) error {
	for _, key := range keys {
		if bm[key] == nil {
			return fmt.Errorf("%s is empty", key)
		}
	}
	return nil
}

func (bm BodyMap) EncodeURLParams() string {
	values := url.Values{}
	for key, value := range bm {
		values.Set(key, fmt.Sprintf("%v", value))
	}
	return values.Encode()
}
