package model

import (
	"github.com/deatil/go-cryptobin/gm/sm2"
	//"github.com/deatil/go-cryptobin/gm/sm2"
)

type SM4 struct {
	PrivateKey    string // 客户私钥
	PublicKey     string // 客户公钥
	BankPublicKey string // 银行公钥
	SymKey        string // 对称密钥
}

type SM2 struct {
	//PrivateKey string
	//PublicKey  string
	PrivateKey *sm2.PrivateKey // 商户私钥，商户加签
	PublicKey  *sm2.PublicKey  // 合利宝公钥，验签

	MerchantKey string // 签名密钥 （进件文档接口、余额查询文档、对账下载接口）
	EncryptKey  string // 加密密钥 （进件文档接口、余额查询文档、对账下载接口）
	UserId      string
}

type Setting struct {
	Id              *string //专用tk企业资源规划的id(逸采)
	Key             *string // Key
	Secret          *string // 密钥
	Shop            *string //店铺名称（适用于shopify)
	ShopId          *string // 店铺ID
	RetryCount      *int    // 重试次数
	AccessToken     *string // 刷新令牌
	RefreshToken    *string //刷新令牌
	ServerUrl       *string // 服务链接
	AuthCallbackUrl *string // 授权回调地址
	SiteNo          *string // 站点(目前Lazada切换访问域名用的)
	MerchantId      *string // 商户
	ShopCipher      *string // 商店密码(目前是Tk在用)
	IsMerchant      bool
	UserId          *string
	CustomTraceNo   string

	RsaPublicKey  *string
	RsaPrivateKey *string
	RsaEncryptKey *string // 对称加密密钥明文
	RsaEncryptIV  *string // 对称加密初始向量明文
	SM4           *SM4
	SM2           *SM2    //合利宝使用
	DevelopId     *string //连连国际使用
	MasterToken   *string //连连国际使用
}

func (c *Setting) SetSM4(data *SM4) *Setting {
	c.SM4 = data
	return c
}

func (c *Setting) SetSM2(data *SM2) *Setting {
	c.SM2 = data
	return c
}

/*
 *连连使用的开发者id
 */

func (c *Setting) SetDevelopId(developId string) *Setting {
	c.DevelopId = &developId
	return c
}

/*
 *连连使用的主token
 */

func (c *Setting) SetMasterToken(masterToken string) *Setting {
	c.MasterToken = &masterToken
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

func (c *Setting) SetRsaEncryptKey(data string) *Setting {
	if len(data) > 0 {
		c.RsaEncryptKey = &data
	}
	return c
}
func (c *Setting) SetRsaEncryptIV(data string) *Setting {
	if len(data) > 0 {
		c.RsaEncryptIV = &data
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

func (c *Setting) SetShop(data string) *Setting {
	if len(data) > 0 {
		c.Shop = &data
	}
	return c
}

func (c *Setting) SetRetryCount(data int) *Setting {
	if data > 0 {
		c.RetryCount = &data
	}
	return c
}
