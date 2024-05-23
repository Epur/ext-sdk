package tiktokv2

type GetTokenResponse struct {
	AccessToken          string `json:"access_token"`            //访问令牌
	RefreshToken         string `json:"refresh_token"`           //刷新令牌
	AccessTokenExpireIn  int64  `json:"access_token_expire_in"`  //访问令牌有效期
	RefreshTokenExpireIn int64  `json:"refresh_token_expire_in"` //刷新令牌有效期
	OpenId               string `json:"open_id"`                 //用户唯一标识
	SellerName           string `json:"seller_name"`             //卖家名称
	SellerBaseRegion     string `json:"seller_base_region"`      //卖家所在区域，如ID:印尼
	UserType             int    `json:"user_type"`               //卖家所在区域，如ID:印尼
}

// 获取店铺信息
type ShopListResponse []Shop
type Shop struct {
	ShopCipher string `json:"cipher"`      // 商店密码(Tk用)
	Code       string `json:"code"`        // 店铺code
	ShopId     string `json:"id"`          // 店铺ID
	ShopName   string `json:"name"`        // 店铺名称
	Region     string `json:"region"`      // 店铺所在区域
	SellerType string `json:"seller_type"` // 跨境类型: LOCAL-本土店 CROSS_BORDER-跨境店
}

type GetOrderDetailResponse struct {
	List []OrderDetailResponse `json:"order_list"`
}

type GetOrderListResponse struct {
	Total int            `json:"total"`
	List  []OrderListRow `json:"list"`
}

type GetProductListResponse struct {
	List  []ProductListResponse `json:"products"`
	Total int                   `json:"total"`
}

type GetProductDetailResponse struct {
	ProductDetailResponse
}

type GetOrderPrintAwbResponse struct {
	DocUrl string `json:"doc_url"`
}
