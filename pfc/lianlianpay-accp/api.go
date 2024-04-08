package lianlianpay_accp

import (
	"github.com/Epur/ext-sdk/logger"
	"github.com/Epur/ext-sdk/model"
)

type Api struct {
	Setting *model.Setting
}

/*
 * 设置并获取连连API(接口层)
 */

func New(setting *model.Setting) *Api {
	// 初始化日志服务
	logger.New("logs", "info")
	return &Api{Setting: setting}
}

/*
 * 提现 reference:https://open.lianlianpay.com/docs/accp/accpstandard/withdrawal.html
 */

func (p *Api) withdraw(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(`/v1/txn/withdrawal`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("timestamp", "oid_partner", "risk_item", "orderInfo", "payerInfo"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := WithdrawResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
 * 快捷支付接口 reference:https://open.lianlianpay.com/docs/accp/accpstandard/payment-bankcard.html
 */

func (p *Api) AccpPay(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(`/v1/txn/payment-bankcard`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("timestamp", "oid_partner", "txn_seqno", "total_amount", "risk_item", "payerInfo", "payMethods"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := PaymentResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
 * 快捷支付结果查询 reference:https://open.lianlianpay.com/docs/accp/accpstandard/query-payment.html
 */

func (p *Api) QueryPayment(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/v1/txn/query-payment`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("timestamp", "oid_partner"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryPaymentResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
 * 提现结果查询接口 reference:https://open.lianlianpay.com/docs/accp/accpstandard/query-withdrawal.html
 */

func (p *Api) QueryWithdraw(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(`/v1/txn/query-withdrawal`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("timestamp", "oid_partner"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryWithdrawResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
 * 用户开户 reference:https://open.lianlianpay.com/docs/accp/accpstandard/openacct-apply.html
 */

func (p *Api) OpenAcctApply(Body model.BodyMap) *model.Client {
	c := NewClient(p.Setting)
	c.SetPath(`/v1/acctmgr/openacct-apply`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("timestamp", "oid_partner",
		"user_id",
		"txn_seqno",
		"txn_time",
		"flag_chnl",
		"notify_url",
		"user_type"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := OpenAcctApplyesponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
