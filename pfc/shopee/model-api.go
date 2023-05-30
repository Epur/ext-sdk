package shopee

type GetTokenResponse struct {
	*Response
	RefreshToken string `json:"refresh_token"` //响应信息
	AccessToken  string `json:"access_token"`  //响应信息
	ExpireIn     int64  `json:"expire_in"`     //响应信息
	MerchantId   int64  `json:"merchant_id"`   //响应信息

	ShopId    int64 `json:"shop_id"`    //响应信息
	PartnerId int64 `json:"partner_id"` //响应信息

	ShopIdList     []int64 `json:"shop_id_list"`     //响应信息
	MerchantIdList []int64 `json:"merchant_id_list"` //响应信息
}

type GetSellerResponse struct {
	*Response
	ShopName     string `json:"shop_name"`
	Region       string `json:"region"`
	Status       string `json:"status"`
	SipAffiShops string `json:"sip_affi_shops"`
	IsCb         bool   `json:"is_cb"`
	IsCnsc       bool   `json:"is_cnsc"`
	ShopCbsc     string `json:"shop_cbsc"`
	AuthTime     int64  `json:"auth_time"`
	ExpireTime   int64  `json:"expire_time"`
	IsSip        bool   `json:"is_sip"`
}

type GetMerchantResponse struct {
	*Response
	Name       string `json:"merchant_name"`
	IsCnsc     bool   `json:"is_cnsc"`
	AuthTime   int64  `json:"auth_time"`
	ExpireTime int64  `json:"expire_time"`
	Currency   string `json:"merchant_currency"`
	Region     string `json:"merchant_region"`
}

type GetOrderDetailResponse struct {
	Details []*OrderItem `json:"order_list"` //订单列表
}
