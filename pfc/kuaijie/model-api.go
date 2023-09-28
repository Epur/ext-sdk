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

// 银行账户绑定
type CustomBankAcctBindRequest struct {
	AcctNo string `json:"acctNo,omitempty"` // 平台账户号， 业务账户，用于资金处理
}

// 银行账户绑定
type CustomBankAcctBindProtected struct {
	BankCode           string `json:"bankCode,omitempty"`           // 开户行代码
	BankName           string `json:"bankName,omitempty"`           // 开户行名称
	BankAcctType       string `json:"bankAcctType,omitempty"`       // 账户类型 pCard：个人借记卡（户名为法人姓名） eGeneral：对公一般户（户名为企业证照名称）
	BankAcctNo         string `json:"bankAcctNo,omitempty"`         // 账号
	BankAcctName       string `json:"bankAcctName,omitempty"`       // 户名
	BankProvince       string `json:"bankProvince,omitempty"`       // 开户省份
	BankCity           string `json:"bankCity,omitempty"`           // 开户城市
	BankAcctBindMobile string `json:"bankAcctBindMobile,omitempty"` // 银行预留的 手机号
}

// 银行账户绑定
type CustomBankAcctBindResponse struct {
	TxnState  string `json:"txnState"`  // 处理状态 1：成功 4：失败
	RspRemark string `json:"rspRemark"` // 处理详情 失败原因
}

// 银行账户绑定查询
type QueryCustomBankAcctBindInfoRequest struct {
	AcctNo string `json:"acctNo,omitempty"` // 平台账户号， 业务账户，用于资金处理
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

// 发起转账-无验证码
type AcctTransferAmtNoIdentCodeRequest struct {
	AcctNoFrom string `json:"acctNoFrom,omitempty"`              // 转出方平台账户号
	AcctNoTo   string `json:"acctNoTo,omitempty"`                // 转入方平台账户号
	TxnAmt     int64  `json:"txnAmt,omitempty" validate:"gte=0"` // 转账金额 单位：分 数值必须大于 0
	TxnRemark  string `json:"txnRemark,omitempty"`               // 订单附言
	NotifyUrl  string `json:"notifyUrl,omitempty"`               // 后台通知地址
	Memo       string `json:"memo,omitempty"`                    // 附加数据
	BsnExt     string `json:"bsnExt,omitempty"`                  // 业务扩展信息
}

// 发起转账-无验证码
type AcctTransferAmtNoIdentCodeResponse struct {
	TxnState  string `json:"txnState"`  // 订单状态0：已受理 1：成功 4：失败
	TxnAmtOut int64  `json:"txnAmtOut"` // (付款账户) 扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn  int64  `json:"txnAmtIn"`  // (收款账户) 入账金额 单位：分 txnState 为 1 时必填
	RspRemark string `json:"rspRemark"` // 处理详情 失败原因等
}

// 发起转账-转账结果
type AcctTransferAmtResultRequest struct {
	OrigTxnDate    string `json:"origTxnDate,omitempty"` // 转账订单日期,格式:yyyyMMdd
	OrigCusTraceNo string `json:"origCusTraceNo"`        // 转账订单的客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo"`        // 转账订单的系统跟踪号
}

// 发起转账-转账结果
type AcctTransferAmtResultResponse struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
	TxnState       string `json:"txnState"`                 // 订单状态 0：已受理 1：成功 4：失败
	OrigRspCode    string `json:"origRspCode"`              // 提现订单应答码 txnState 为 1 或 4 时必填
	OrigRspMsg     string `json:"origRspMsg"`               // 提现订单应答信息 txnState 为 1 或 4 时必填
	TxnAmtOut      int64  `json:"txnAmtOut"`                // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn       int64  `json:"txnAmtIn"`                 // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	RspRemark      string `json:"rspRemark"`                // 处理详情 失败原因等
	Memo           string `json:"memo"`                     // 附加数据 原订单上送的自定义数据
}

// 发起提现-无验证码
type AcctWithdrawNoIdentCodeRequest struct {
	AcctNo     string `json:"acctNo,omitempty"`                  // 平台账户号
	SettAcctNo string `json:"settAcctNo,omitempty"`              // 提现账号
	TxnAmt     int64  `json:"txnAmt,omitempty" validate:"gte=0"` // 提现金额 单位：分 数值必须大于 0
	TxnRemark  string `json:"txnRemark,omitempty"`               // 订单附言
	NotifyUrl  string `json:"notifyUrl,omitempty"`               // 后台通知地址
	Memo       string `json:"memo,omitempty"`                    // 附加数据
	BsnExt     string `json:"bsnExt,omitempty"`                  // 业务扩展信息
}

// 发起提现-无验证码
type AcctWithdrawNoIdentCodeResponse struct {
	TxnState     string `json:"txnState"`     // 订单状态0：已受理 1：成功 4：失败
	TxnAmtOut    int64  `json:"txnAmtOut"`    // (付款账户) 扣账金额 单位：分 txnState为1时必填
	TxnAmtIn     int64  `json:"txnAmtIn"`     // (收款账户) 入账金额 单位：分 txnState为1时必填
	BankAcctNo   string `json:"bankAcctNo"`   // 收款银行账号 txnState为1时必填
	BankAcctName string `json:"bankAcctName"` // 收款银行户名 txnState为1时必填
	RspRemark    string `json:"rspRemark"`    // 处理详情 失败原因等
}

// 提现结果查询
type AcctWithdrawResultRequest struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
}

// 提现结果查询
type AcctWithdrawResultResponse struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
	TxnState       string `json:"txnState"`                 // 订单状态 0：已受理 1：成功 4：失败
	OrigRspCode    string `json:"origRspCode"`              // 提现订单应答码 txnState 为 1 或 4 时必填
	OrigRspMsg     string `json:"origRspMsg"`               // 提现订单应答信息 txnState 为 1 或 4 时必填
	TxnAmtOut      int64  `json:"txnAmtOut"`                // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn       int64  `json:"txnAmtIn"`                 // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	BankAcctNo     string `json:"bankAcctNo"`               // 收款银行账号 txnState 为 1 时必填
	BankAcctName   string `json:"bankAcctName"`             // 收款银行户名 txnState 为 1 时必填
	RspRemark      string `json:"rspRemark"`                // 处理详情 失败原因等
	Memo           string `json:"memo"`                     // 附加数据 原订单上送的自定义数据
}

// 发起付款-无验证码
type AcctPayAmtNoIdentCodeRequest struct {
	AcctNo     string `json:"acctNo,omitempty"`                  // 平台账户号
	BankAcctNo string `json:"bankAcctNo,omitempty"`              // 银行账号 该银行账号必须事先通过“银行账户绑定配置”接口完成绑定
	TxnAmt     int64  `json:"txnAmt,omitempty" validate:"gte=0"` // 提现金额 单位：分 数值必须大于 0
	TxnRemark  string `json:"txnRemark,omitempty"`               // 订单附言
	NotifyUrl  string `json:"notifyUrl,omitempty"`               // 后台通知地址
	Memo       string `json:"memo,omitempty"`                    // 附加数据
	BsnExt     string `json:"bsnExt,omitempty"`                  // 业务扩展信息
}

// 发起付款-无验证码
type AcctPayAmtNoIdentCodeResponse struct {
	TxnState     string `json:"txnState"`     // 订单状态0：已受理 1：成功 4：失败
	TxnAmtOut    int64  `json:"txnAmtOut"`    // (付款账户) 扣账金额 单位：分 txnState为1时必填
	TxnAmtIn     int64  `json:"txnAmtIn"`     // (收款账户) 入账金额 单位：分 txnState为1时必填
	BankAcctNo   string `json:"bankAcctNo"`   // 收款银行账号 txnState为1时必填
	BankAcctName string `json:"bankAcctName"` // 收款银行户名 txnState为1时必填
	RspRemark    string `json:"rspRemark"`    // 处理详情 失败原因等
}

// 付款结果查询
type AcctPayAmtResultRequest struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
}

// 付款结果查询
type AcctPayAmtResultResponse struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
	TxnState       string `json:"txnState"`                 // 订单状态 0：已受理 1：成功 4：失败
	OrigRspCode    string `json:"origRspCode"`              // 提现订单应答码 txnState 为 1 或 4 时必填
	OrigRspMsg     string `json:"origRspMsg"`               // 提现订单应答信息 txnState 为 1 或 4 时必填
	TxnAmtOut      string `json:"txnAmtOut"`                // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn       string `json:"txnAmtIn"`                 // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	BankAcctNo     string `json:"bankAcctNo"`               // 收款银行账号 txnState 为 1 时必填
	BankAcctName   string `json:"bankAcctName"`             // 收款银行户名 txnState 为 1 时必填
	RspRemark      string `json:"rspRemark"`                // 处理详情 失败原因等
	Memo           string `json:"memo"`                     // 附加数据 原订单上送的自定义数据
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
