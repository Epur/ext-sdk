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
	set := new(model.Setting)
	api := amazon.New(set)
	testApi := AmazonTest{api: api}
	//testApi.GetAuthUrl()
	//testApi.GetToken("")
	//testApi.RefreshToken("Atzr|IwEBIJ7ZrNhqoJUmWD_s1RQC2hf8kS_j37fFAcNx0XYLWbUEA0r6GjXts5j45LGCS5mKpA1Hospv2ojIbHcp5Kn9ans61YF0p5WVMmKrDBXxvshzT0NB3EaY7g2YRwZiZ7iZnVdYKeyU273dzNEQKIALOld3kaNre_K8vbfO09tfPb3P_a4ZX240yUdMjDG3A2Jr_-z9q-j5tmkdK9-oISDnB-DCSdghtB_cRNvTXWOt5CM-M_MUS489AueTNuLwbsic2h6zt2FVpM4EYyjXAncfy0VgUtVVA_QYMdutJ_NKqD5zTeHrbDbzg04hsnlq6yaBeFw")
	//testApi.GetSeller()
	//testApi.GetOrders()
	//testApi.GetOrder("TEST_CASE_200")
	//testApi.GetOrder("112-2041974-3583467")
	//testApi.GetOrderItems("112-2041974-3583467")
	testApi.GetDestSub("Test")
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

func (p *AmazonTest) GetDestSub(env string) {
	p.api.Setting.SetSiteNo("US")
	c := p.api.NoToken()
	accessToken := *p.api.Setting.AccessToken
	noAccessToken := c.GetResponseTo().(amazon.GetTokenResponse).AccessToken
	// 创建dest 无授权
	p.api.Setting.SetAccessToken(noAccessToken)
	c = p.api.GetDestinations()
	destArr := c.GetResponseTo().(amazon.GetDestinationsResponse)
	var destId string
	if destArr.Payload == nil || len(destArr.Payload) < 1 {
		cc := p.api.CreateDestination(model.BodyMap{
			"name": "SQS-" + env,
			"resourceSpecification": map[string]interface{}{
				"arn": "arn:aws:sqs:us-east-1:381492270878:SQS-" + env},
		})
		destId = cc.GetResponseTo().(amazon.CreateDestinationResponse).Payload.DestinationId
	} else {
		destId = destArr.Payload[len(destArr.Payload)-1].DestinationId
	}
	//创建sub 授权
	p.api.Setting.SetAccessToken(accessToken)
	ccc := p.api.CreateSubscription(model.BodyMap{
		"notificationType": amazon.ORDER_CHANGE,
		"destinationId":    destId,
		"payloadVersion":   "1.0",
		"processingDirective": map[string]interface{}{
			"eventFilter": map[string]interface{}{
				"eventFilterType":  amazon.ORDER_CHANGE,
				"orderChangeTypes": []string{"BuyerRequestedChange", "OrderStatusChange"}}},
	})
	csr := ccc.GetResponseTo().(amazon.CreateSubscriptionResponse)
	if !ccc.Response.Success {
		fmt.Println(ccc.Response)
		if ccc.Response.HttpStatus == 409 {

		} else if ccc.Response.HttpStatus == 403 {

		}
	} else {
		fmt.Println(csr)
	}
}
