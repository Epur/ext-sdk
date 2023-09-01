package kuaijie

import (
	"fmt"
	"github.com/Epur/ext-sdk/logger"
	"github.com/Epur/ext-sdk/model"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	// 初始化日志服务
	logger.New("logs", "info")
	return &Api{Setting: setting}
}

/*
	账户余额查询
	Repsonse : QueryCustomAcctInfoResponse
*/

func (p *Api) QueryCustomAcctInfo(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/balance/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("acctNo"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := QueryCustomAcctInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response
	fmt.Println("AcctBalance:", response.AcctBalance)
	fmt.Println("AcctValid:", response.AcctValid)

	return &c.Client
}

/*
	银行账户绑定查询
	Repsonse : QueryCustomBankAcctBindInfoResponse
*/

func (p *Api) QueryCustomBankAcctBindInfo(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/bind/account/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("acctNo"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := QueryCustomBankAcctBindInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response
	fmt.Printf("AcctBalance:%+v", response.Accounts)

	return &c.Client
}

/*
	银行账户绑定查询
	Repsonse : CashSweepQueryAcctBalInfoResponse
*/

func (p *Api) CashSweepQueryAcctBalInfo(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/bank/balance/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("issAcctNo"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := CashSweepQueryAcctBalInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response
	//fmt.Printf("response:%+v", response)

	return &c.Client
}
