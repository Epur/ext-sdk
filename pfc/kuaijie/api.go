package kuaijie

import "github.com/Epur/ext-sdk/model"

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

/*
	客户结算账户信息查询
	Repsonse : QueryCustomAcctInfoResponse
*/

func (p *Api) QueryCustomAcctInfo(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/balance/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("acctNo"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryCustomAcctInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
