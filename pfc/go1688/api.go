package go1688

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	api := &Api{Setting: setting}
	if api.Setting.ServerUrl == nil || len(*api.Setting.ServerUrl) <= 0 {
		api.Setting.SetServerUrl("https://gw.open.1688.com/openapi")
	}
	return api
}

func (p *Api) GetAuthUrl(callbackParams string) string {

	return fmt.Sprintf("%s?%s", AuthURL, model.BodyMap{}.
		Set("client_id", *p.Setting.Key).
		Set("redirect_uri", fmt.Sprintf("%s?callbackParams=%s", *p.Setting.AuthCallbackUrl, url.QueryEscape(callbackParams))).
		Set("site", "1688").
		Set("state", "state").EncodeURLParams(),
	)
}

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	Body.Set("grant_type", "authorization_code").
		Set("need_refresh_token", "true").
		Set("client_id", *p.Setting.Key).
		Set("client_secret", *p.Setting.Secret)

	c := NewClient(p.Setting)
	c.SetPath(AccessTokenURL).
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("code"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	if c.Response.Success {
		c.Client.Response.Response.Data = c.HttpReq.Result

		tmpResponse := map[string]interface{}{}
		if c.Err = json.Unmarshal(c.HttpReq.Result, &tmpResponse); c.Err != nil {
			return &c.Client
		}

		response := GetTokenResponse{
			AliId:         tmpResponse["aliId"].(string),
			AccessToken:   tmpResponse["access_token"].(string),
			RefreshToken:  tmpResponse["refresh_token"].(string),
			ResourceOwner: tmpResponse["resource_owner"].(string),
			MemberId:      tmpResponse["memberId"].(string),
		}
		response.ExpiresIn, _ = strconv.ParseInt(tmpResponse["expires_in"].(string), 10, 64)

		aTmp := strings.Split(tmpResponse["refresh_token_timeout"].(string), "+")
		location, _ := time.LoadLocation("Asia/Shanghai")
		s, _ := time.ParseInLocation(fmt.Sprintf("20060102150405000+%s", aTmp[1]),
			tmpResponse["refresh_token_timeout"].(string), location)
		response.RefreshExpiresIn = s.UnixMilli() / 1000

		c.Response.Response.DataTo = response
	}

	return &c.Client
}

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {

	Body.Set("grant_type", "refresh_token").
		Set("client_id", *p.Setting.Key).
		Set("client_secret", *p.Setting.Secret)

	c := NewClient(p.Setting)
	c.SetPath(AccessTokenURL).
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("refresh_token"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	fmt.Println(c.Response.Success)

	if c.Response.Success {

		c.Client.Response.Response.Data = c.HttpReq.Result

		tmpResponse := map[string]interface{}{}
		if c.Err = json.Unmarshal(c.HttpReq.Result, &tmpResponse); c.Err != nil {
			return &c.Client
		}

		response := GetTokenResponse{
			AliId:         tmpResponse["aliId"].(string),
			AccessToken:   tmpResponse["access_token"].(string),
			ResourceOwner: tmpResponse["resource_owner"].(string),
			MemberId:      tmpResponse["memberId"].(string),
		}
		response.ExpiresIn, _ = strconv.ParseInt(tmpResponse["expires_in"].(string), 10, 64)
		c.Response.Response.DataTo = response
	}

	return &c.Client
}

func (p *Api) StoreRefreshToken(Body model.BodyMap) *model.Client {

	currTime := utils.TimestampSecond()

	c := p.RefreshToken(Body)

	if c.Response.Response.DataTo != nil {
		response := c.Response.Response.DataTo.(GetTokenResponse)
		c.Response.Response.DataTo = model.StoreTokenResponse{
			AccessToken:       response.AccessToken,
			AccessTokenExpire: response.ExpiresIn + currTime,
			//RefreshToken:      response.RefreshToken,
			//RefreshTokenExpire: response.RefreshExpiresIn + currTime,
		}
	}

	return c
}

// 跨境场景下将商品加入铺货列表
func (p *Api) PushedProductList(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.product.push:alibaba.cross.syncProductListPushed-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("productIdList"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// alibaba.cross.productInfo
func (p *Api) GetProductInfo(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.product:alibaba.cross.productInfo-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("productId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 跨境场景获取商品列表
func (p *Api) GetProductList(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.product:alibaba.cross.productList-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("productIdList"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 根据地址解析地区码
func (p *Api) ParseAddressCode(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.trade:alibaba.trade.addresscode.parse-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("addressInfo"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 跨境订单创建
func (p *Api) CreateCrossOrder(Body model.BodyMap) *model.Client {
	//fmt.Printf("===== body:%+v\n", Body)
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.trade:alibaba.trade.createCrossOrder-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("flow"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 获取买家视角订单信息
func (p *Api) GetBuyerView(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.trade:alibaba.trade.get.buyerView-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("webSite", "orderId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 获取支付链接
func (p *Api) GetPayUrl(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.trade:alibaba.alipay.url.get-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("orderIdList"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 调用免密支付
func (p *Api) PreparePay(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay-1").
		SetMethod(http.POST).
		SetParams(Body)

	if c.Err = Body.CheckEmptyError("tradeWithholdPreparePayParam"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 查看是否开启免密支付
func (p *Api) DisPasswdIsOpen(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen-1").
		SetMethod(http.POST).
		SetParams(Body)

	//if c.Err = Body.CheckEmptyError("tradeWithholdPreparePayParam"); c.Err != nil {
	//	return &c.Client
	//}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}

// 图搜跨境商品
func (p *Api) SimilarOfferSearch(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath("com.alibaba.linkplus:alibaba.cross.similar.offer.search-1").
		SetMethod(http.POST).
		SetParams(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	return &c.Client
}
