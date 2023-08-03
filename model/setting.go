package model

type SM4 struct {
	PrivateKey    string // 客户私钥
	PublicKey     string // 客户公钥
	BankPublicKey string // 银行公钥
	SymKey        string // 对称密钥
}
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
	ShopCipher      *string // 商店密码(目前是Tk在用)
	IsMerchant      bool
	UserId          *string

	RsaPublicKey  *string
	RsaPrivateKey *string
	SM4           *SM4
}

func (c *Setting) SetSM4(data *SM4) *Setting {
	c.SM4 = data
	return c
}

func (c *Setting) SetRsaPrivateKey(data string) *Setting {
	if len(data) > 0 {
		c.RsaPrivateKey = &data
	}
	return c
}

func (c *Setting) SetRsaPublicKey(data string) *Setting {
	if len(data) > 0 {
		c.RsaPublicKey = &data
	}
	return c
}

func (c *Setting) SetUserId(data string) *Setting {
	if len(data) > 0 {
		c.UserId = &data
	}
	return c
}

func (c *Setting) SetShopCipher(data string) *Setting {
	if len(data) > 0 {
		c.ShopCipher = &data
	}
	return c
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
