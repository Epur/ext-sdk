package lazada

type GetTokenResponse struct {
	*Response
	ExpiresIn        int64                      `json:"expires_in"`
	AccountId        string                     `json:"account_id"`
	Country          string                     `json:"country"`
	CountryUserInfo  []*ResponseConutryUserInfo `json:"country_user_info"`
	AccountPlatform  string                     `json:"account_platform"`
	AccessToken      string                     `json:"access_token"`
	Account          string                     `json:"account"`
	RefreshExpiresIn int64                      `json:"refresh_expires_in"`
	RefreshToken     string                     `json:"refresh_token"`
}

type GetSellerResponse struct {
	NameCompany string `json:"name_company"`
	Name        string `json:"name"`
	Verified    bool   `json:"verified"`
	SellerId    int64  `json:"seller_id"`
	Email       string `json:"email"`
	ShortCode   string `json:"short_code"`
	Cb          bool   `json:"cb"`
	Status      string `json:"status"`
}

type GetOrderDetailResponse []OrderDetailResponse

type GetOrderResponse OrderResponse
