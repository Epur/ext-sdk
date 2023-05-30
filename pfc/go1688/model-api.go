package go1688

type GetTokenResponse struct {
	*Response

	AccessToken      string `json:"access_token"`
	AliId            string `json:"aliId"`
	RefreshToken     string `json:"refresh_token"`
	ResourceOwner    string `json:"resource_owner"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	MemberId         string `json:"member_id"`
}
