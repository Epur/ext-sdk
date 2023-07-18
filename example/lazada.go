package main

import (
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/pfc/lazada"
	"github.com/Epur/ext-sdk/utils"
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
	testApi.GetOrderList()
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

func (p *LazadaTest) GetOrderList() {

	p.api.Setting.SetSiteNo("th")

	c := p.api.GetOrderList(model.BodyMap{"limit": "10", "created_after": "2020-01-02T15:04:05.000+07:00"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(lazada.GetOrderListResponse)
	for _, item := range result.List {
		fmt.Println(item)
	}
}
