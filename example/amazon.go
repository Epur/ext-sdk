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
	spId := "amzn1.sp.solution.02150dd3-4cd1-4ce4-bfa6-29655e709975"
	set := new(model.Setting).
		SetKey("amzn1.application-oa2-client.71425753b10e4462b6e29621f97301b2").
		SetSecret("amzn1.oa2-cs.v1.f68fbe1d18251cb617b68e054b78fbaf874ae315a455061fb80b6b7278587f6c").
		SetAuthCallbackUrl("https://dev2.web.epur.cn/openapi/admin/v1/erp/platform/cross/auth/callback/Amazon").
		SetAccessToken("Atza|IwEBIOxZ65qC9690cQIC9KGC0sDvWHXZkL33o-1Kpa6xW69SFUJNxAGMlVQZOedVhg2OtWUtWErYcCJaCKkx0HalENEFA8c6c01T7SuWyHvj2AemhyyXey-AcjsHL7STokja13Vny4gunslgccvDuP5KYSyQglkBgsx4F8fg0eWGYA-LqFupr1aXl_VPS9Ip-M_XmXDaoXC939HkgPHw-I2-FiiHdsxwLBPk9FXxOAJBIlagWJEwIZ2bWaxRm0AX-uxGPivC68mQ6geMUTd5PB3cNX0-wm9GJatK92g68Fqdjjv-jYBy4DZyxfe6tjzI7iW_VM").
		SetServerUrl("https://dev2.web.epur.cn/#/to/enter-in").
		SetDevelopId(spId)
	api := amazon.New(set)
	testApi := AmazonTest{api: api}
	//testApi.GetAuthUrl()
	//testApi.GetToken("ANNLKqIpxXoJKeSrmxuc")
	testApi.RefreshToken("Atzr|IwEBIDQcDmEG1Mhe3kZkupBytw-2BQjdAVPZrXjZ2dsvd5UoAmpTTzGm2jJjcoPaK72WLHRR3Z9_D5nn2rqZA-msT0CJB7fIQaln0GSfSKovhUeHs9fOt3Sq5IrhTate_ffJ_bhc7fp7dwupivkbdlTfxlNXRvs76rlq9ciyNgl6WmvFupRlfOSucaCwUmdSRz1uYvs1y2XbO9mc8wEG8mzkDTjEnp5FbA1kXDGRLXUfhx9rd0iRSstvuj1OWnwjeqbEQ-OXC5h2ekZAhMBP9joTk76ekv_6gJuJYMYWKPrgxOnL-JCVNQM_uCDqjPt1JtJCGo8")
	//testApi.GetOrderList()
}

func (p *AmazonTest) GetAuthUrl() {
	callbackParam := map[string]interface{}{"link": 1, "sites": []string{"MX"}}
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

	p.api.Setting.SetSiteNo("us")

	c := p.api.GetSeller()
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(amazon.GetSellerResponse)
	fmt.Println(result)
}

func (p *AmazonTest) GetOrderDetail() {

	p.api.Setting.SetSiteNo("th")

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

	p.api.Setting.SetSiteNo("us")

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
