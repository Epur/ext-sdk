package main

import (
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/pfc/shopee"
	"github.com/tangchen2018/eshop-sdk/utils"
)

type ShopeeTest struct {
	api *shopee.Api
}

func main() {
	api := shopee.New(
		new(model.Setting).
			SetKey("").
			SetSecret("").
			SetAuthCallbackUrl("").
			SetServerUrl("").
			SetShopId("").
			SetMerchantId("").
			SetAccessToken(``),
	)
	testApi := ShopeeTest{api: api}
	testApi.GetOrderDetail()

	//4e6344676e4b4b6c4774474756544465
}

func (p *ShopeeTest) GetAuthUrl() {
	result := p.api.GetAuthUrl("123")
	fmt.Println(result)
}

func (p *ShopeeTest) GetToken() {
	c := p.api.GetToken(
		model.BodyMap{"code": "",
			"main_account_id": 0})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(shopee.Response)
	fmt.Println(utils.ToJson(result))
}

func (p *ShopeeTest) GetSeller() {

	c := p.api.GetSeller(nil)
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(shopee.GetSellerResponse)
	fmt.Println(utils.ToJson(result))
}

func (p *ShopeeTest) GetMerchant() {

	c := p.api.GetMerchant(nil)
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(shopee.GetMerchantResponse)
	fmt.Println(utils.ToJson(result))
}

func (p *ShopeeTest) RefreshToken() {

	c := p.api.RefreshToken(model.BodyMap{"refresh_token": ""})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(shopee.GetTokenResponse)
	fmt.Println(utils.ToJson(result))
}

func (p *ShopeeTest) GetOrderDetail() {

	c := p.api.GetOrderDetail(model.BodyMap{"order_sn_list": "2305061NBXB2RY"})

	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(shopee.GetOrderDetailResponse)
	fmt.Println(utils.ToJson(result))
}
