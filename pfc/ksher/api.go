package ksher

import "github.com/Epur/ext-sdk/model"

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

/*
	客户信息登记结果查询
	Response: GetMerchantResponse
*/

func (p *Api) GetMerchant(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/global/v1/merchant/query`).
		SetMethod("POST").
		SetBody(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetMerchantResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
