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
		SetKey("amzn1.application-oa2-client.60110a1e40a041779f11bb6c49c1fe26").
		SetSecret("amzn1.oa2-cs.v1.fa9a3df6962f8443287b84628aff9dc91dc45624eb6eee5331f12e4c50aaa428").
		SetAuthCallbackUrl("https://dev2.web.epur.cn/openapi/admin/v1/erp/platform/cross/auth/callback/Amazon").
		SetAccessToken("Atza|IwEBIEOl7Kpnz-A1Jcucsxqj2KrC2e3fSudTEmE2ijRDSEff3W0oojDeUWTT2DIzzAWzn8PRPYIFIMjWbhhxbipN8wJMk8quyf5GITlzBO-vSMAP7f3WkUPNsDAADbSWYD8555O9C7i6NTr6RvTI4hcpeaog4RUB67QnKo5Olvi3nWhqKTzI8tUXPnnWm67h4NZ82IeOInllvEF2rDwVESE2I7Kn8iMiHwxf4PVErfH9m1EOPktgeVzW4TWqfnVyKG8WBDoz4dP30czW1vo90hU0usPW5M75PlIfH4bqEKop2pJ87IKDoeDn-11CpFmQcN5wLnXcFyPKKCjwiIf_V3TDlnF7").
		SetServerUrl("https://dev2.web.epur.cn/#/to/enter-in").
		SetDevelopId("amzn1.sp.solution.3dbb2269-d7a9-415d-802d-bf0c9ed40fe4")
	api := amazon.New(set)
	testApi := AmazonTest{api: api}
	//testApi.GetAuthUrl()
	//testApi.GetToken("")
	testApi.RefreshToken("Atzr|IwEBIJ7ZrNhqoJUmWD_s1RQC2hf8kS_j37fFAcNx0XYLWbUEA0r6GjXts5j45LGCS5mKpA1Hospv2ojIbHcp5Kn9ans61YF0p5WVMmKrDBXxvshzT0NB3EaY7g2YRwZiZ7iZnVdYKeyU273dzNEQKIALOld3kaNre_K8vbfO09tfPb3P_a4ZX240yUdMjDG3A2Jr_-z9q-j5tmkdK9-oISDnB-DCSdghtB_cRNvTXWOt5CM-M_MUS489AueTNuLwbsic2h6zt2FVpM4EYyjXAncfy0VgUtVVA_QYMdutJ_NKqD5zTeHrbDbzg04hsnlq6yaBeFw")
	//testApi.GetSeller()
	//testApi.GetOrderList()
	//testApi.GetOrder("TEST_CASE_200")
	//testApi.GetOrder("112-2041974-3583467")
	//testApi.GetOrderItems("112-2041974-3583467")
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

	p.api.Setting.SetSiteNo("UK")

	c := p.api.GetSeller()
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(amazon.GetSellerResponse)
	fmt.Println(result)
}

func (p *AmazonTest) GetOrders() {

	p.api.Setting.SetSiteNo("US")

	c := p.api.GetOrders(model.BodyMap{
		"MarketplaceIds": "ATVPDKIKX0DER",
		"AmazonOrderIds": "114-0263527-6375432,112-2041974-3583467",
		"CreatedAfter":   "2024-05-01",
		//"CreatedAfter":   "TEST_CASE_200",
	})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(amazon.GetOrdersResponse)
	for _, item := range result.Payload.OrderList {
		by, _ := json.Marshal(item)
		fmt.Println(string(by))
	}
	fmt.Printf("共%d条", len(result.Payload.OrderList))
}

func (p *AmazonTest) GetOrder(orderId string) {

	p.api.Setting.SetSiteNo("US")

	c := p.api.GetOrder(model.BodyMap{"orderId": orderId})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(amazon.GetOrderResponse)
	by, _ := json.Marshal(result)
	fmt.Println(string(by))
	fmt.Println(result.Payload.AmazonOrderId)
}

func (p *AmazonTest) GetOrderItems(orderId string) {

	p.api.Setting.SetSiteNo("US")

	c := p.api.GetOrderItems(model.BodyMap{"orderId": orderId})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(amazon.GetOrderItemsResponse)
	by, _ := json.Marshal(result)
	fmt.Println(string(by))
	fmt.Println(result.Payload.AmazonOrderId)
}
