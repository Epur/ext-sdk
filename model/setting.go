package model

type Setting struct {
	Key             *string // Key
	Secret          *string // 密钥
	ShopId          *string // 店铺ID
	RetryCount      *int    // 重试次数
	AccessToken     *string // 刷新令牌
	ServerUrl       *string // 服务链接
	AuthCallbackUrl *string // 授权回调地址
	SiteNo          *string // 站点(目前Lazada切换访问域名用的)
	MerchantId      *string // 商户
	IsMerchant      bool
}

func (c *Setting) SetIsMerchant(data bool) *Setting {
	c.IsMerchant = data
	return c
}

func (c *Setting) SetMerchantId(data string) *Setting {
	if len(data) > 0 {
		c.MerchantId = &data
	}
	return c
}

func (c *Setting) SetSiteNo(data string) *Setting {
	if len(data) > 0 {
		c.SiteNo = &data
	}
	return c
}

func (c *Setting) SetAuthCallbackUrl(data string) *Setting {
	if len(data) > 0 {
		c.AuthCallbackUrl = &data
	}
	return c
}

func (c *Setting) SetAccessToken(data string) *Setting {
	if len(data) > 0 {
		c.AccessToken = &data
	}
	return c
}

func (c *Setting) SetServerUrl(data string) *Setting {
	if len(data) > 0 {
		c.ServerUrl = &data
	}
	return c
}

func (c *Setting) SetShopId(data string) *Setting {
	if len(data) > 0 {
		c.ShopId = &data
	}
	return c
}

func (c *Setting) SetKey(data string) *Setting {
	if len(data) > 0 {
		c.Key = &data
	}
	return c
}

func (c *Setting) SetSecret(data string) *Setting {
	if len(data) > 0 {
		c.Secret = &data
	}
	return c
}

func (c *Setting) SetRetryCount(data int) *Setting {
	if data > 0 {
		c.RetryCount = &data
	}
	return c
}
