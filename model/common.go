package model

type StoreTokenResponse struct {
	AccessToken        string `json:"accessToken"`        // 刷新令牌
	AccessTokenExpire  int64  `json:"accessTokenExpire"`  // 访问令牌有效期(时间戳)
	RefreshToken       string `json:"refreshToken"`       // 刷新令牌
	RefreshTokenExpire int64  `json:"refreshTokenExpire"` // 刷新令牌有效期(时间戳)
}
