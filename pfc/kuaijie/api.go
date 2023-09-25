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

func (p *Api) SetUserId(userId string) error {
	p.Setting.SetUserId(userId)
	return nil
}
func (p *Api) SetCustomTraceNo(traceNo string) {
	p.Setting.CustomTraceNo = traceNo
}

/*
	账户余额查询
	Response : QueryCustomAcctInfoResponse
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
	银行账户绑定
	Response : CustomBankAcctBindResponse
*/

func (p *Api) CustomBankAcctBind(Body, Protected model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/bind/account`).
		SetMethod("POST").
		SetBody(Body).
		SetProtected(Protected)

	if c.Err = Body.CheckEmptyError("acctNo"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	if c.Err = Protected.CheckEmptyError("bankCode", "bankAcctType", "bankAcctNo", "bankAcctName"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := CustomBankAcctBindResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response
	fmt.Printf("TxnState:%+v", response)

	return &c.Client
}

/*
	银行账户绑定查询
	Response : QueryCustomBankAcctBindInfoResponse
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
	发起转账-无验证码
	Response : AcctTransferAmtNoIdentCodeResponse
*/

func (p *Api) AcctTransferAmtNoIdentCode(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/transfer/silent`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("acctNoFrom", "acctNoTo", "txnAmt", "txnRemark"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AcctTransferAmtNoIdentCodeResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
	发起结果查询
	Response: AcctTransferAmtResult
*/

func (p *Api) AcctTransferAmtResult(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/transfer/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("origTxnDate"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AcctTransferAmtResultResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
	发起提现-无验证码
	Response: AcctCashOutNoIdentCodeResponse
*/

func (p *Api) AcctWithdraw(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/withdraw/silent`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("acctNo", "settAcctNo", "txnAmt", "txnRemark", "notifyUrl", "memo"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AcctWithdrawNoIdentCodeResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
	提现结果查询
	Response: AcctWithdrawResultResponse
*/

func (p *Api) AcctWithdrawResult(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/withdraw/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("origTxnDate"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AcctWithdrawResultResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
	付款要主账户发起交易，发起付款可以向非同名账户提现
	发起付款-无验证码
	Response: AcctPayAmtNoIdentCodeResponse
*/

func (p *Api) AcctPayAmt(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/pay/silent`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("acctNo", "bankAcctNo", "txnAmt", "txnRemark", "notifyUrl", "memo"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AcctPayAmtNoIdentCodeResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
	提现结果查询
	Response: AcctPayAmtResultResponse
*/

func (p *Api) AcctPayAmtResult(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/forward/da/txn/v2/da/third/pay/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("origTxnDate"); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AcctPayAmtResultResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.KuaijieLoger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
	银行账户绑定查询
	Response : CashSweepQueryAcctBalInfoResponse
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
