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
		SetSecret("amzn1.oa2-cs.v1.3b1e4dc89b49e879d59de2161c848b851b961347844d145f23c5b540e99d1ce0").
		SetAuthCallbackUrl("https://dev2.web.epur.cn/openapi/admin/v1/erp/platform/cross/auth/callback/Amazon").
		SetAccessToken("Atza|IwEBIKq59lYMWbDNhf-vFTeKnd96BZ7EsTuZs9-9ZzQYKmQ1cCWBPMv_ychaedwdIgGPYqVUIthlUeyuFbrIkqFuVuEhgLVSctXM0fvrSIgn8nozKeosPmHX_EqmMhcxUtRpyIMwxgQNarVbXpEoGTHSo0KCnm5mIFc5zCzRcNrc26lrYBTl5P9A9nJ_b8gb3Uht_CxJuzPbvFOH5SYlttrZXCzcrT2yJlKe8Ib4FHWt7QycUM9VqvB-kCxEzwIXJkpQOxnTodEBuQ-QqGm-jOSGYovGhe6ZacjtfQ5MdoyZ7UVxs9pok-DwM7xXpZKTCg5sHKvHy0jPlpkizxwXqfVjywUv").
		SetServerUrl("https://dev2.web.epur.cn/#/to/enter-in").
		SetDevelopId("amzn1.sp.solution.3dbb2269-d7a9-415d-802d-bf0c9ed40fe4")
	api := amazon.New(set)
	testApi := AmazonTest{api: api}
	//testApi.GetAuthUrl()
	//testApi.GetToken("")
	//testApi.RefreshToken("Atzr|IwEBIJ7ZrNhqoJUmWD_s1RQC2hf8kS_j37fFAcNx0XYLWbUEA0r6GjXts5j45LGCS5mKpA1Hospv2ojIbHcp5Kn9ans61YF0p5WVMmKrDBXxvshzT0NB3EaY7g2YRwZiZ7iZnVdYKeyU273dzNEQKIALOld3kaNre_K8vbfO09tfPb3P_a4ZX240yUdMjDG3A2Jr_-z9q-j5tmkdK9-oISDnB-DCSdghtB_cRNvTXWOt5CM-M_MUS489AueTNuLwbsic2h6zt2FVpM4EYyjXAncfy0VgUtVVA_QYMdutJ_NKqD5zTeHrbDbzg04hsnlq6yaBeFw")
	//testApi.GetSeller()
	//testApi.GetOrderList()
	testApi.GetOrderDetail("TEST_CASE_200")
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

func (p *AmazonTest) GetOrderDetail(orderId string) {

	p.api.Setting.SetSiteNo("US")

	c := p.api.GetOrder(model.BodyMap{"orderId": orderId})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(amazon.GetOrderResponse)
	fmt.Println(result)
	fmt.Println(result.Payload.AmazonOrderId)
}

func (p *AmazonTest) GetOrderList() {

	p.api.Setting.SetSiteNo("US")

	c := p.api.GetOrderList(model.BodyMap{
		"MarketplaceIds": "ATVPDKIKX0DER",
		"CreatedAfter":   "TEST_CASE_200"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(amazon.GetOrdersResponse)
	for _, item := range result.Payload.OrderList {
		fmt.Println(item)
	}
}
