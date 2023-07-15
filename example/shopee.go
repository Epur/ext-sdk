package main

import (
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/pfc/shopee"
	"github.com/Epur/ext-sdk/utils"
)

type ShopeeTest struct {
	api *shopee.Api
}

func main() {
	api := shopee.New(
		new(model.Setting).
			SetKey("2005467").
			SetSecret("4a7a70474173707972707366424d5654517248664a4b41776568734b706f7152").
			SetAuthCallbackUrl("").
			SetServerUrl("").
			SetShopId("988362252").
			SetMerchantId("3902713").
			SetAccessToken(`646e4a5462684f466a5263714c585265`),
	)
	testApi := ShopeeTest{api: api}
	testApi.GetOrderList()

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

func (p *ShopeeTest) GetOrderList() {

	c := p.api.GetOrderListAsc(model.BodyMap{"time_from": "1563159542", "page_size": "10", "time_to": "1689389942", "time_range_field": "create_time"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(shopee.GetOrderListResponse)
	for _, item := range result.List {
		fmt.Println(item.OrderSn)
	}
}
