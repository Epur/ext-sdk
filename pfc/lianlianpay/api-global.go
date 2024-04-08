package lianlianpay

import "github.com/Epur/ext-sdk/model"

/*
	用户授权申请
	Url : 沙盒环境https://global-api-sandbox.lianlianpay-inc.com/fund/userAuthApply
				正式环境 https://global-api.lianlianpay.com/fund/userAuthApply
	Response:
*/

func (p *Api) UserAuthApply(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/fund/userAuthApply`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("prodType", "extApplyId", "returnUrl"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := UserAuthApplyResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	用户授权查询
	Url : 沙盒环境/https	https://global-api-sandbox.lianlianpay-inc.com/fund/userAuthQuery
正式环境/https	https://global-api.lianlianpay.com/fund/userAuthQuery
	Response:
*/

func (p *Api) UserAuthQuery(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/fund/userAuthQuery`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("userId", "prodType", "extApplyId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := UserAuthQueryResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	放款申请
	Url : 沙盒环境/https	https://global-api-sandbox.lianlianpay-inc.com/fund/loanApply
			正式环境/https	https://global-api.lianlianpay.com/fund/loanApply
	Response:
*/

func (p *Api) LoanApply(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/fund/loanApply`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("userId",
		"prodType",
		"fundParty",
		"extLoanId",
		"loanAmount",
		"currency",
		"borrowerName",
		"borrowerType",
		"borrowerCertNo",
		"repayType"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := LoanApplyResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	放款确认
	Url : 沙盒环境/https	https://global-api-sandbox.lianlianpay-inc.com/fund/loanConfirm
		  正式环境/https	https://global-api.lianlianpay.com/fund/loanConfirm
	Response:
*/

func (p *Api) LoanConfirm(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/fund/loanConfirm`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("userId",
		"prodType",
		"loanId",
		"extLoanId",
		"loanAmount",
		"loanBeginDate",
		"loanEndDate",
		"currency",
		"status",
	); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := LoanConfirmResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	贷款结清
	Url : 沙盒环境/https	https://global-api-sandbox.lianlianpay-inc.com/fund/loanFinished
			正式环境/https	https://global-api.lianlianpay.com/fund/loanFinished
	Response:
*/

func (p *Api) LoanFinished(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/fund/loanFinished`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("userId",
		"prodType",
		"loanId",
		"finishDate",
	); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := LoanFinishedResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	强制还款
	Url : 沙盒环境/https	https://global-api-sandbox.lianlianpay-inc.com/fund/forceRepayment
			正式环境/https	https://global-api.lianlianpay.com/fund/forceRepayment
	Response:
*/

func (p *Api) ForceRepayment(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/fund/forceRepayment`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("userId",
		"prodType",
		"loanId",
		"fundParty",
		"status",
		"deductAmount",
		"notifyTime",
		"deductId",
	); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := ForceRepaymentResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	还款单通知
	Url : 沙盒环境/https	https://global-api-sandbox.lianlianpay-inc.com/gateway/fund/notifyRepayment
			正式环境/https	https://global-api.lianlianpay.com/gateway/fund/notifyRepayment
	Response:
*/

func (p *Api) NotifyRepayment(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/fund/notifyRepayment`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("userId",
		"loanId",
		"extRepaymentId",
		"extLoanId",
		"repaymentAmount",
		"currency",
		"repaymentTime",
		"notifyTime"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := NotifyRepaymentResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
