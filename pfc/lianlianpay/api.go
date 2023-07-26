package lianlianpay

import (
	"github.com/Epur/ext-sdk/model"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

//
//func (p *Api) GetAuthUrl(callbackParams string) string {
//
//	ss := model.BodyMap{}.
//		Set("response_type", "code").
//		Set("client_id", *p.Setting.Key).
//		Set("state", "state").
//		Set("redirect_uri", fmt.Sprintf("%s?callbackParams=%s", *p.Setting.AuthCallbackUrl, url.QueryEscape(callbackParams)))
//
//	if p.Setting.UserId != nil {
//		ss.Set("user_id", *p.Setting.UserId)
//	}
//
//	return fmt.Sprintf("%s?%s", AuthURL, ss.EncodeURLParams())
//}

/*
	持有人查询
	Url : https://developer.lianlianglobal.com/docs/llp-api/09455dd3980c0-
	Response: GetHolderResponse
*/

func (p *Api) GetHolder(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/holder/query/list`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("holderType"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetHolderResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	用户收款账户余额查询
	Url : https://developer.lianlianglobal.com/docs/llp-api/6e3c95aa7b658-
	Response: GetAccountBalanceResponse
*/

func (p *Api) GetAccountBalance(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/account/balance/query`).
		SetMethod("POST").
		SetBody(Body)

	//if c.Err = Body.CheckEmptyError("currency"); c.Err != nil {
	//	return &c.Client
	//}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetAccountBalanceResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	用户详情查询
	Url : https://developer.lianlianglobal.com/docs/llp-api/5f2e875a7a8d2-
	Response: GetAccountResponse
*/

func (p *Api) GetAccount(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/account/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("pageNum", "pageSize"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetAccountResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	资金流水查询
	Url : https://developer.lianlianglobal.com/docs/llp-api/9250a1f8d5c54-
	Response: GetTransactionResponse
*/

func (p *Api) GetTransaction(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/myshop/transaction`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("pageNum", "pageSize", "startTime", "endTime", "accountId"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetTransactionResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	获取实时汇率
	Url : https://developer.lianlianglobal.com/docs/llp-api/06fa3fdadbcae-
	Response: GetExchangeRateResponse
*/

func (p *Api) GetExchangeRate(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/exchangeRate/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("sourceCurrency", "targetCurrency"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetExchangeRateResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	入账交易查询
	Url : https://developer.lianlianglobal.com/docs/llp-api/e361747518aa8-
	Response: GetTransactionEntryListResponse
*/

func (p *Api) GetTransactionEntryList(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/transaction/entry/list`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("pageNum", "pageSize"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetTransactionEntryListResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	提现交易查询
	Url : https://developer.lianlianglobal.com/docs/llp-api/3f7ebd879c7c2-
	Response: GetWithdrawRecordListResponse
*/

func (p *Api) GetWithdrawRecordList(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/withdraw/record/list`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("pageNum", "pageSize"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetWithdrawRecordListResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	提现银行卡查询
	Url : https://developer.lianlianglobal.com/docs/llp-api/152b55312f128-
	Response: GetBankcardListResponse
*/

func (p *Api) GetBankcardList(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/bankcard/query/list`).
		SetMethod("POST").
		SetBody(Body)

	//if c.Err = Body.CheckEmptyError("pageNum", "pageSize"); c.Err != nil {
	//	return &c.Client
	//}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetBankcardListResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	用户资金操作-发送验证码
	Url : https://developer.lianlianglobal.com/docs/llp-api/152b55312f128-
*/

func (p *Api) SendCaptchaMoBileWithdrawCombine(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/captcha/mobile/withdraw/combine`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("captchaType"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	return &c.Client
}

/*
	用户资金提现
	Url : https://developer.lianlianglobal.com/docs/llp-api/e7dba8775d756-
	Response : WithdrawCombineSubmitResponse
*/

func (p *Api) WithdrawCombineSubmit(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/withdraw/combine/submit`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("externalTransId", "withdrawCurrency", "arrivalCurrency", "cardId", "captcha", "withdrawDetailList"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := WithdrawCombineSubmitResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	查询用户KYC状态
	Url : https://developer.lianlianglobal.com/docs/llp-api/989d595e47877-kyc
	Response : GetCollectionKycResponse
*/

func (p *Api) GetCollectionKyc(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/collection/v1/kyc/status`).
		SetMethod("GET").
		SetBody(Body)

	//if c.Err = Body.CheckEmptyError("externalTransId", "withdrawCurrency", "arrivalCurrency", "cardId", "captcha", "withdrawDetailList"); c.Err != nil {
	//	return &c.Client
	//}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetCollectionKycResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
