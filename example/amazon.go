package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/pfc/amazon"
	"github.com/Epur/ext-sdk/utils"
)

type AmazonTest struct {
	api *amazon.Api
}

func main() {
	set := new(model.Setting).
		SetKey("amzn1.application-oa2-client.71425753b10e4462b6e29621f97301b2").
		SetSecret("amzn1.oa2-cs.v1.0f4a3c881354103d5571714d61ddd833f817579f7628aeecc856a86fee8a3b62").
		SetAuthCallbackUrl("https://dev2.web.epur.cn/openapi/admin/v1/erp/platform/cross/auth/callback/Amazon").
		SetAccessToken("Atza|IwEBIPG_JLxCAZmJHWLwsCWkNo_GLdGaqEmLcQ_g7Cm4AMVwEFFmrhrzH5q9S8_R2VqlGuKFO10GPeQnptDXzfldslHJCp3-Sfadfnrs5EkxMope2MP6m2lst1sYh5wgcQKE6NWsPpZW5F5IB5LpRv8NEbekr90o8MVQAHN1jZ3m3XSVed5yv55DYlT5IUpHoKJHxjlsl6ZEY5yLnubs9E15gvKQ-3Xoa_9ogFgbhAN7ZuufN43ZVacy7Sw6Qfc27XeONUbCknwDqOoXwIxsIV9ZXXepEfXOy-zf6qNW1btnXgmJTFj9A72DUQZ5gzsrEA6v9uw").
		SetServerUrl("https://dev2.web.epur.cn/#/to/enter-in").
		SetDevelopId("amzn1.sp.solution.02150dd3-4cd1-4ce4-bfa6-29655e709975")
	api := amazon.New(set)
	testApi := AmazonTest{api: api}
	//testApi.GetAuthUrl()
	//testApi.GetToken("")
	//testApi.RefreshToken("")
	testApi.GetSeller()
	//testApi.GetOrderList()
}

func (p *AmazonTest) GetAuthUrl() {
	callbackParam := map[string]interface{}{"link": 1, "sites": "US,MX,CA"}
	callbackParamStr, _ := json.Marshal(callbackParam)
	result := p.api.GetAuthUrl(string(callbackParamStr))
	fmt.Println(result)
}

func (p *AmazonTest) GetToken(code string) {
	c := p.api.GetToken(model.BodyMap{"code": code})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(amazon.GetTokenResponse)
	fmt.Println(utils.ToJson(result))
}

func (p *AmazonTest) RefreshToken(refreshToken string) {
	c := p.api.RefreshToken(model.BodyMap{"refresh_token": refreshToken})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(amazon.GetTokenResponse)
	fmt.Println(utils.ToJson(result))
}

func (p *AmazonTest) GetSeller() {

	p.api.Setting.SetSiteNo("US")

	c := p.api.GetSeller()
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(amazon.GetSellerResponse)
	fmt.Println(result)
}

func (p *AmazonTest) GetOrderDetail() {

	p.api.Setting.SetSiteNo("TH")

	c := p.api.GetOrder(model.BodyMap{"orderId": "690646034119032"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(amazon.GetOrderResponse)
	fmt.Println(result)
}

func (p *AmazonTest) GetOrderList() {

	p.api.Setting.SetSiteNo("US")

	c := p.api.GetOrderList(model.BodyMap{
		"MarketplaceIds": []string{"A1AM78C64UM0Y8", "ATVPDKIKX0DER"},
		"CreatedAfter":   "2024-01-02"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(amazon.GetOrderListResponse)
	for _, item := range result.OrderList {
		fmt.Println(item)
	}
}
