package cmbchina

import "github.com/Epur/ext-sdk/model"

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

/*
	1.可经办业务模式查询 DCLISMOD
	Repsonse : DCLISMODResponse
*/

func (p *Api) DCLISMOD(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`DCLISMOD`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("buscod"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := DCLISMODResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 2.查询可经办的账户列表 DCLISACC
func (p *Api) DCLISACC(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`DCLISACC`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("buscod", "busmod"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := DCLISACCResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	3 账户详细信息查询NTQACINF
	Repsonse : NTQACINFResponse
*/

func (p *Api) NTQACINF(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`NTQACINF`).
		SetMethod("POST").
		SetBody(Body)

	//if c.Err = Body.CheckEmptyError("buscod"); c.Err != nil {
	//	return &c.Client
	//}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := NTQACINFResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 4.查询账户历史余额 NTQABINF
func (p *Api) NTQABINF(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`NTQABINF`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("ntqabinfy"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := NTQABINFResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 5. 查询分行号信息NTACCBBK
func (p *Api) NTACCBBK(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`NTACCBBK`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("fctval"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	//response := QueryAcctOpenBankResponse{}
	//if c.Err = c.Client.Response.To(&response); c.Err != nil {
	//	return &c.Client
	//}
	//c.Response.Response.DataTo = response
	return &c.Client
}

// 6. 批量查询余额NTQADINF
func (p *Api) NTQADINF(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`NTQADINF`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("ntqadinfx"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryAcctBalBatchResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 7.账户交易信息查询trsQryByBreakPoint
func (p *Api) TrsQryByBreakPoint(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`trsQryByBreakPoint`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("TRANSQUERYBYBREAKPOINT_X1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryAccountTranInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 8.PDF文件对账单获取DCTRSPDF
func (p *Api) DCTRSPDF(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`DCTRSPDF`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("sdktsinfx"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := CheckOrderPdfResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 9.OFD文件对账单获取issueBillOfd
func (p *Api) IssueBillOfd(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`issueBillOfd`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("CprDirectIssueBillOfdX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := IssueBillOfdResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 10.OFD文件对账单获取结果查询queryBillOfd
func (p *Api) QueryBillOfd(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`queryBillOfd`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("CprDirectQueryBillOfdX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryBillOfdResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 11.单笔回单查询DCSIGREC
func (p *Api) DCSIGREC(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`DCSIGREC`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("eacnbr", "quedat", "trsseq"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := DCSIGRECResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 12. 电子回单异步查询ASYCALHD
func (p *Api) ASYCALHD(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`ASYCALHD`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("primod", "eacnbr", "begdat", "enddat", "rrcflg"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := ASYCALHDResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 13. 异步打印结果查询DCTASKID
func (p *Api) DCTASKID(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`DCTASKID`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("taskid"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := DCTASKIDResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 2. 企银支付单笔经办BB1PAYOP
func (p *Api) BB1PAYOP(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`BB1PAYOP`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("bb1paybmx1", "bb1payopx1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := BusiPayResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 3.企银支付业务查询BB1PAYQR
func (p *Api) BB1PAYQR(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`BB1PAYQR`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("bb1payqrx1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := BusiPayQueryResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 4. 企银支付批量经办BB1PAYBH
func (p *Api) BB1PAYBH(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`BB1PAYBH`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("bb1bmdbhx1", "bb1paybhx1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := BusiPayBatchResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 5. 企银批量支付批次查询BB1QRYBT
func (p *Api) BB1QRYBT(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`BB1QRYBT`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("bb1qrybtx1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := BusiQueryBatchResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 6.企银批量支付明细查询BB1QRYBD
func (p *Api) BB1QRYBD(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`BB1QRYBD`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("bb1qrybdy1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := BusiQueryBatchPayListResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 7. 支付退票明细查询BB1PAYQB
func (p *Api) BB1PAYQB(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`BB1PAYQB`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("bb1payqby1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := BusiQueryReturnListResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 1.待授权个人户口信息录入ADDPREAUTHINFO
func (p *Api) ADDPREAUTHINFO(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`ADDPREAUTHINFO`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("PreAuthInfoX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	//response := BusiQueryReturnListResponse{}
	//if c.Err = c.Client.Response.To(&response); c.Err != nil {
	//	return &c.Client
	//}
	//c.Response.Response.DataTo = response
	return &c.Client
}

// 2.待授权信息查询QUERYPREAUTHINFO
func (p *Api) QUERYPREAUTHINFO(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`QUERYPREAUTHINFO`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("QPreAuthInfoX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryPrivateAuthInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 3.已授权信息查询QUERYAUTHINFO
func (p *Api) QUERYAUTHINFO(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`QUERYAUTHINFO`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("QAuthInfoX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryAuthInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 4.查询个人账户人民币余额QUERYAUTHACCBALAMT
func (p *Api) QUERYAUTHACCBALAMT(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`QUERYAUTHACCBALAMT`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("QAuthAccBalAmtX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryPrivateAcctInfoResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 5.查询个人账户交易明细QUERYAUTHACCTRA
func (p *Api) QUERYAUTHACCTRA(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`QUERYAUTHACCTRA`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("QAuthAccTraX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := QueryPrivateAcctTranListResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

// 6.公私一网通支付经办提交PAYOPR
func (p *Api) PAYOPR(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`PAYOPR`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("FBPayOprX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	//response := QueryPrivateAcctTranListResponse{}
	//if c.Err = c.Client.Response.To(&response); c.Err != nil {
	//	return &c.Client
	//}
	//c.Response.Response.DataTo = response
	return &c.Client
}

// 7.公私一网通支付结果查询PAYQUERY
func (p *Api) PAYQUERY(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`PAYQUERY`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("FBPayQueryX1"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	//response := QueryPrivateAcctTranListResponse{}
	//if c.Err = c.Client.Response.To(&response); c.Err != nil {
	//	return &c.Client
	//}
	//c.Response.Response.DataTo = response
	return &c.Client
}
