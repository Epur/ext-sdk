package kuaijie

// 账户余额查询请求报文
type QueryCustomAcctInfoRequest struct {
	AcctNo string `json:"acctNo"` // 平台账户号， 业务账户，用于资金处理
}

// 账户余额查询响应报文
type QueryCustomAcctInfoResponse struct {
	AcctNo        string `json:"acctNo"`        // 平台账户号 业务账户，用于资金处理
	AcctBalance   int64  `json:"acctBalance"`   // 总余额,单位：分 总余额=其他账户金额之和
	AcctValid     int64  `json:"acctValid"`     // 可用金额 单位：分
	AcctHold      int64  `json:"acctHold"`      // 冻结金额 单位：分
	AcctWaitSett  int64  `json:"acctWaitSett"`  // 待结算账户金额 单位：分
	AcctWaitSplit int64  `json:"acctWaitSplit"` // 待分账账户金额 单位：分 开通分账功能后才有该值
	AcctCharge    int64  `json:"acctCharge"`    // 手续费支出账户金额 单位：分 开通指定账户后才有该值
	AcctPromotion int64  `json:"acctPromotion"` // 营销费用账户金额 单位：分 开通指定账户后才有该值
}

// 银行账户绑定查询
type QueryCustomBankAcctBindInfoRequest struct {
	AcctNo string `json:"acctNo"` // 平台账户号， 业务账户，用于资金处理
}

// 银行账户绑定查询
type QueryCustomBankAcctBindInfoResponse struct {
	Accounts []QueryCustomBankAcctBindInfo `json:"accounts"` // 结算账户
}

// 银行账户绑定查询
type QueryCustomBankAcctBindInfo struct {
	BankCode           string `json:"bankCode"`           // 开户行代码
	BankName           string `json:"bankName"`           // 开户行名称
	BankAcctType       string `json:"bankAcctType"`       // 账户类型 pCard：个人借记卡（户名为法人姓名）eGeneral：对公一般户（户名为企业证照名称）
	BankAcctNo         string `json:"bankAcctNo"`         // 账号
	BankAcctName       string `json:"bankAcctName"`       // 户名
	BankProvince       string `json:"bankProvince"`       // 开户省份
	BankCity           string `json:"bankCity"`           // 开户城市
	BankAcctBindMobile string `json:"bankAcctBindMobile"` // 银行预留的手机号
	Default            string `json:"default"`            // 默认0 0:否 1:是
}

// 资金归集--账户余额查询
type CashSweepQueryAcctBalInfoRequest struct {
	IssAcctNo string `json:"issAcctNo"` // 平台账户号 资金归集账户，用于接收外部平台资金
}

// 资金归集--账户余额查询
type CashSweepQueryAcctBalInfoResponse struct {
	IssAcctNo   string `json:"issAcctNo"`   // 平台账户号
	AcctBalance int64  `json:"acctBalance"` // 总余额
	AcctValid   int64  `json:"acctValid"`   // 可用金额
	AcctHold    int64  `json:"acctHold"`    // 冻结金额
}
