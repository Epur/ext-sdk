package lianlianpay_accp

import (
	"github.com/shopspring/decimal"
)

// 提现返回信息
type WithdrawResponse struct {
	RetCode    string `json:"ret_code"`    //请求结果代码
	RetMsg     string `json:"ret_msg"`     //请求结果描述
	OidPartner string `json:"oid_partner"` //ACCP系统分配给平台商户的唯一编号

	UserId      string          `json:"user_id"`      //用户编号
	TxnSeqNo    string          `json:"txn_seqno"`    //商户系统唯一交易流水号
	TotalAmount decimal.Decimal `json:"total_amount"` //订单总金额

	FeeAmount decimal.Decimal `json:"fee_amount"` //手续费金额，单位为元，精确到小数点后两位
	AccpTxnno string          `json:"accp_txno"`  //ACCP系统交易单号
	Token     string          `json:"token"`      //支付授权令牌，有效期30分钟。当交易需要二次验证的时候，需要通过token调用调用交易二次短信验证接口
}

// 快接支付返回信息
type PaymentResponse struct {
	RetCode    string `json:"ret_code"`    //请求结果代码
	RetMsg     string `json:"ret_msg"`     //请求结果描述
	OidPartner string `json:"oid_partner"` //ACCP系统分配给平台商户的唯一编号

	UserId      string          `json:"user_id"`      //用户编号
	TxnSeqNo    string          `json:"txn_seqno"`    //商户系统唯一交易流水号
	TotalAmount decimal.Decimal `json:"total_amount"` //订单总金额

	AccpTxnno string `json:"accp_txno"` //ACCP系统交易单号
	Token     string `json:"token"`     //支付授权令牌，有效期30分钟。当交易需要二次验证的时候，需要通过token调用调用交易二次短信验证接口

	LinkedAgrtno   string `json:"linked_agrtno"`   //绑卡协议号
	AccountingDate string `json:"accounting_date"` //账务日期。ACCP系统交易账务日期，交易成功时返回，格式：yyyyMMdd
	FinishTime     string `json:"finish_time"`     //支付完成时间，格式为yyyyMMddHHmmss
}

// 提现结果查询返回信息

type QueryWithdrawResponse struct {
	RetCode    string `json:"ret_code"`    //请求结果代码
	RetMsg     string `json:"ret_msg"`     //请求结果描述
	OidPartner string `json:"oid_partner"` //ACCP系统分配给平台商户的唯一编号

	LinkedAcctno   string `json:"linked_acctno"`   //个人用户绑定的银行卡号，企业用户绑定的银行帐号。
	AccountingDate string `json:"accounting_date"` //账务日期。ACCP系统交易账务日期，交易成功时返回，格式：yyyyMMdd
	FinishTime     string `json:"finish_time"`     //支付完成时间，格式为yyyyMMddHHmmss
	AccpTxnno      string `json:"accp_txno"`       //ACCP系统交易单号

	ChnlTxno  string `json:"chnl_txno"`  //渠道交易单号
	TxnStatus string `json:"txn_status"` //提现交易状态TRADE_SUCCESS：交易成功
	// TRADE_FAILURE：交易失败
	//TRADE_CANCEL：退汇
	//TRADE_PREPAID：预付完成。
	//交易最终状态以此为准，商户按该字段值进行后续业务逻辑处理。
	FilureReason string `json:"failure_reason"` //提现失败原因。当txn_status为FAILURE或CANCEL时返回具体失败原因信息。

	ChnlReason string `json:"chnl_reason"` //渠道原始原因。提现失败渠道原始原因。
	Bankcode   string `json:"bankcode"`    //渠道原始原因。提现失败渠道原始原因。

	OrderInfo *OrderInfoM `json:"orderInfo"` //订单信息
	PayerInfo *PayInfoM   `json:"payerInfo"` //付款方信息
}

type PayInfoM struct {
	PayerType string `json:"payer_type"` //订单信息
	PayerId   string `json:"payer_id"`   //付款方信息
}

type OrderInfoM struct {
	TxnSeqNo string `json:"txn_seqno"` //商户系统唯一交易流水号
	TxnTime  string `json:"txn_time"`  //商户系统交易时间。
	//格式：yyyyMMddHHmmss
	TotalAmount decimal.Decimal `json:"total_amount"` //订单总金额，单位为元，精确到小数点后两位。
	FeeAmount   decimal.Decimal `json:"fee_amount"`   //手续费金额，单位为元，精确到小数点后两位。回显提现接口传的手续费金额。
	OrderInfo   string          `json:"order_info"`   //用于订单说明，透传返回。
}

//支付结果查询返回信息

type QueryPaymentResponse struct {
	RetCode    string `json:"ret_code"`    //请求结果代码
	RetMsg     string `json:"ret_msg"`     //请求结果描述
	OidPartner string `json:"oid_partner"` //ACCP系统分配给平台商户的唯一编号

	TxnType string `json:"txn_type"` //交易类型。
	//用户充值：USER_TOPUP
	//商户充值：MCH_TOPUP
	//普通消费：GENERAL_CONSUME
	//担保消费：SECURED_CONSUME
	//担保确认：SECURED_CONFIRM
	//内部代发：INNER_FUND_EXCHANGE
	//外部代发：OUTER_FUND_EXCHANGE
	AccountingDate string `json:"accounting_date"` //账务日期。ACCP系统交易账务日期，交易成功时返回，格式：yyyyMMdd
	FinishTime     string `json:"finish_time"`     //支付完成时间，格式为yyyyMMddHHmmss
	AccpTxnno      string `json:"accp_txno"`       //ACCP系统交易单号

	ChnlTxno  string `json:"chnl_txno"`  //渠道交易单号
	TxnStatus string `json:"txn_status"` //支付交易状态。
	//TRADE_WAIT_PAY:交易处理中
	//TRADE_SUCCESS:交易成功
	//TRADE_CLOSE:交易失败
	//支付交易状态以此为准，商户必须依据此字段值进行后续业务逻辑处理。
	FilureReason string `json:"pay_chnl_txno"` //渠道流水号。如微信支付单号。

	ChnlReason string `json:"sub_chnl_no"` //渠道商家订单号。如微信商家订单号。
	Bankcode   string `json:"bankcode"`    //渠道原始原因。提现失败渠道原始原因。

	OrderInfo *OrderInfoM `json:"orderInfo"` //订单信息
	PayerInfo *PayerInfoM `json:"payerInfo"` //付款方信息
	PayeeInfo *PayeeInfoM `json:"payeeInfo"` //收款方信息
}

// 付款方信息
type PayerInfoM struct {
	PayerType string `json:"payer_type"` //付款方类型。
	//用户：USER
	//平台商户：MERCHANT
	PayerId string `json:"payer_id"` //付款方标识。
	//付款方为用户时设置user_id
	//付款方为商户时设置平台商户号。

	Method      string          `json:"method"`        //付款方式。参见付款方式列表。。
	Amount      decimal.Decimal `json:"amount"`        //付款金额。付款方式对应的金额，单位为元，精确到小数点后两位。
	Acctno      string          `json:"acctno"`        //付款银行账号，付款方式为银行卡支付和企业网银B2B支付方式时返回，脱敏返回：卡号前六后四
	Acctname    string          `json:"acctname"`      //付款银行户名，付款方式为银行卡支付和企业网银B2B支付方式时返回
	ChnlPayerId string          `json:"chnl_payer_id"` //渠道付款方标识
}

// 收款方信息
type PayeeInfoM struct {
	PayeeType string `json:"payee_type"` //收款方类型。
	//用户：USER
	//平台商户：MERCHANT
	PayeeId string `json:"payee_id"` //收款方标识。用户的user_id或者平台商户商户号。

	Amount          decimal.Decimal `json:"amount"`            //付款金额。付款方式对应的金额，单位为元，精确到小数点后两位。
	AvailableAmtBal string          `json:"available_amt_bal"` //交易后可用账户余额。该交易收款方可用账户有资金变动时返回。
	UnsettledAmtBal string          `json:"unsettled_amt_bal"` //交易后待结算账户余额。该交易收款方待结算账户有资金变动时返回。
}

// 用户开户返回信息

type OpenAcctApplyesponse struct {
	RetCode    string `json:"ret_code"`    //请求结果代码
	RetMsg     string `json:"ret_msg"`     //请求结果描述
	OidPartner string `json:"oid_partner"` //ACCP系统分配给平台商户的唯一编号

	UserId     string `json:"user_id"`     //商户用户唯一编号。用户在商户系统中的唯一编号，要求该编号在商户系统能唯一标识用户。
	TxnSeqNo   string `json:"txn_seqno"`   //商户系统唯一交易流水号
	AccpTxnno  string `json:"accp_txno"`   //ACCP系统交易单号
	GatewayUrl string `json:"gateway_url"` //用户开户网关地址，用户跳转至该地址完成开户过程。跳转方式：商户前端Get请求该地址。
	//如果想要显示返回上一页按钮，可以在gateway_url后拼接header=Y，即可带上H5头部
	//示例：https://openweb.lianlianpay.com/accp-open-account/?token=96deb14a22eb4aa2dff2bcf81828fc7b&header=Y
}
