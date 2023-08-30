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
