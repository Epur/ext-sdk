package main

import (
	"errors"
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/pfc/lazada"
	"github.com/tangchen2018/eshop-sdk/utils"
)

type LazadaTest struct {
	api *lazada.Api
}

func main() {
	api := lazada.New(
		new(model.Setting).
			SetKey("").
			SetSecret("").
			SetAuthCallbackUrl("").
			SetAccessToken(``),
	)
	testApi := LazadaTest{api: api}
	testApi.GetOrderDetail()
}

func (p *LazadaTest) GetAuthUrl() {
	result := p.api.GetAuthUrl("123")
	fmt.Println(result)
}

func (p *LazadaTest) GetToken() {
	c := p.api.GetToken(model.BodyMap{"code": ""})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(lazada.GetTokenResponse)
	fmt.Println(utils.ToJson(result))
}

func (p *LazadaTest) RefreshToken() {
	c := p.api.RefreshToken(model.BodyMap{"refresh_token": ""})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(lazada.GetTokenResponse)
	fmt.Println(utils.ToJson(result))
}

func (p *LazadaTest) GetSeller() {

	p.api.Setting.SetSiteNo("th")

	c := p.api.GetSeller(nil)
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(lazada.GetSellerResponse)
	fmt.Println(result)
}

func (p *LazadaTest) GetOrderDetail() {

	p.api.Setting.SetSiteNo("th")

	c := p.api.GetOrder(model.BodyMap{"order_id": "690646034119032"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(lazada.GetOrderResponse)
	fmt.Println(result)
}
