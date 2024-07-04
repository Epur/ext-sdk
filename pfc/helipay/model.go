package helipay

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

/*
 * 下单接口返回结构化结构
 */
type Response struct {
	Rt1BizType   string `json:"rt1_bizType"`   //交易类型
	Rt2SignCode  string `json:"rt2_signType"`  //签名类型
	Rt3Timestamp string `json:"rt3_timestamp"` //时间戳
	Rt4Success   string `json:"rt4_success"`   //是否需要验签
	Rt5RetCode   string `json:"rt5_retCode"`   //响应码

	Rt6RetMsg         string `json:"rt6_retMsg"`         //返回信息
	Rt7OrderId        string `json:"rt7_orderId"`        //商户订单号
	Rt8CustomerNumber string `json:"rt8_customerNumber"` //商户编号

	Rt9SerialNumber string          `json:"rt9_serialNumber"` //平台流水号
	Rt10OrderStatus string          `json:"rt10_orderStatus"` //订单状态
	Rt11Ext         json.RawMessage `json:"rt11_ext"`         //拓展参数(AccountPaySub)
	Rt12Reason      string          `json:"rt12_reason"`      //拓展参数
	Rt13Ext         json.RawMessage `json:"rt13_ext"`         //拓展参数(AccountPayQuery)
	Sign            string          `json:"sign"`             //签名
}

type PayResponse = Response

/*
 * 下单结果查询接口返回结构化结构
 */

type PayQueryResponse = Response

/*
 * 商户结算接口响应结构
 */
type MerchantSettlementResponse struct {
	Rt1BizType        string `json:"rt1_bizType"`        //交易类型
	Rt2RetCode        string `json:"rt2_retCode"`        //响应码
	Rt3RetMsg         string `json:"rt3_retMsg"`         //返回信息
	Rt4CustomerNumber string `json:"rt4_customerNumber"` //商户编号
	Rt5OrderId        string `json:"rt5_orderId"`        //商户订单号
	Sign              string `json:"sign"`               //签名
}

/*
 * 商户结算查询接口响应结构
 */
type MerchantSettlementQueryResponse struct {
	Rt1BizType       string          `json:"rt1_bizType"`       //交易类型
	Rt2RetCode       string          `json:"rt2_retCode"`       //响应码
	Rt3RetMsg        string          `json:"rt3_retMsg"`        //返回信息
	Rt4SettleRecords json.RawMessage `json:"rt4_settleRecords"` //结算记录
	Sign             string          `json:"sign"`              //签名
}

/*
 * 结算结果查询返回的rt4_settleRecords字段对应的结构
 */

type SettleRecord struct {
	CustomerNumber   string          `json:"customerNumber"`   //商户编号
	OrderId          string          `json:"orderId"`          //订单号
	SettleDate       string          `json:"settleDate"`       //结算日期
	SettlementAmount decimal.Decimal `json:"settlementAmount"` //结算金额
	SettleFee        decimal.Decimal `json:"settleFee"`        //结算手续费

	D0Amount         decimal.Decimal `json:"d0Amount"`         //加急资金
	D0Fee            decimal.Decimal `json:"d0Fee"`            //加急手续费
	TodayAmount      decimal.Decimal `json:"todayAmount"`      //当日加急资金
	TodayFee         decimal.Decimal `json:"todayFee"`         //当日加急手续费
	NonTodayD0Amount decimal.Decimal `json:"nonTodayD0Amount"` //非当日加急资金

	NonTodayD0Fee decimal.Decimal `json:"nonTodayD0Fee"` //非当日加急手续费
	SettleType    string          `json:"settleType"`    //结算类型（T1 / D1）
	OrderStatus   string          `json:"orderStatus"`   //结算状态
	Reason        string          `json:"reason"`        //失败原因
	CompleteDate  string          `json:"completeDate"`  //完成时间
}

/*
 * 商户余额查询接口响应结构
 */
type MerchantAccountQueryResponse struct {
	Rt1BizType          string          `json:"rt1_bizType"`            //交易类型
	Rt2RetCode          string          `json:"rt2_retCode"`            //响应码
	Rt3RetMsg           string          `json:"rt3_retMsg"`             //返回信息
	Rt4CustomerNumber   string          `json:"rt4_customerNumber"`     //商户编号
	Rt5AccountStatus    string          `json:"rt5_accountStatus"`      //商户订单号
	Rt6Balance          string          `json:"rt6_balance"`            //账户余额
	Rt7FrozenBalance    string          `json:"rt7_frozenBalance"`      //账户余额
	Rt8D0Balance        decimal.Decimal `json:"rt8_d0Balance"`          //账户余额
	Rt9T1Balance        string          `json:"rt9_T1Balance"`          //账户余额
	Rt10Currency        string          `json:"rt10_currency"`          //账户余额
	Rt11CreateDate      string          `json:"rt11_createDate"`        //账户余额
	Rt12Desc            string          `json:"rt12_desc"`              //账户余额
	Rt13D1Balance       decimal.Decimal `json:"rt13_d1Balance"`         //账户余额
	Rt14RechargeBalance decimal.Decimal `json:"rt14_rechargeBalance"`   //账户余额
	Rt15AmountToSettled decimal.Decimal `json:"rt15_amountToBeSettled"` //账户余额
	Sign                string          `json:"sign"`                   //签名

}

/*
 * 公众号/JS/服务窗预下单接口
 */

type AppPayPublicResponse struct {
	Rt1BizType         string `json:"rt1_bizType"`         //交易类型
	Rt2RetCode         string `json:"rt2_retCode"`         //返回码
	Rt3RetMsg          string `json:"rt3_retMsg"`          //时间戳
	Rt4CustomerNumber  string `json:"rt4_customerNumber"`  //商户编号
	Rt5OrderId         string `json:"rt5_orderId"`         //商户订单号
	Rt6SerialNumber    string `json:"rt6_serialNumber"`    //平台流水号
	Rt7PayType         string `json:"rt7_payType"`         //支付类型
	Rt8Appid           string `json:"rt8_appid"`           //公众账号 ID
	Rt9TokenId         string `json:"rt9_tokenId"`         //动态口令
	Rt10PayInfo        string `json:"rt10_payInfo"`        //原生态js 支付信息
	Rt11OrderAmount    string `json:"rt11_orderAmount"`    //交易金额
	Rt12Currency       string `json:"rt12_currency"`       //币种类型
	Rt13ChannelRetCode string `json:"rt13_channelRetCode"` //上游返回码
	Rt14AppPayType     string `json:"rt14_appPayType"`     //客户端类型
	SubMerchantNo      string `json:"subMerchantNo"`       //渠道子商户号(U/A/T)
	SignatureType      string `json:"signatureType"`       //签名方式
	Sign               string `json:"sign"`                //签名
}

/*
 * 扫码/刷卡下单接口返回结构化结构
 */
type AppPayResponse struct {
	Rt1BizType        string `json:"rt1_bizType"`        //交易类型
	Rt2RetCode        string `json:"rt2_retCode"`        //签名类型
	Rt3RetMsg         string `json:"rt3_retMsg"`         //时间戳
	Rt4CustomerNumber string `json:"rt4_customerNumber"` //商户编号
	Rt5OrderId        string `json:"rt5_orderId"`        //商户订单号
	Rt6SerialNumber   string `json:"rt6_serialNumber"`   //平台流水号
	Rt7PayType        string `json:"rt7_payType"`        //是否需要验签
	Rt8Qrcode         string `json:"rt8_qrcode"`         //响应码

	Rt9Wapurl string `json:"rt9_wapurl"` //返回信息

	Rt10OrderAmount           string `json:"rt10_orderAmount"`                  //订单状态
	Rt11Currency              string `json:"rt11_currency"`                     //拓展参数(AccountPaySub)
	Rt12OpenId                string `json:"rt12_openId"`                       //拓展参数
	Rt13OrderStatus           string `json:"rt13_orderStatus"`                  //拓展参数(AccountPayQuery)
	Rt14FundBillList          string `json:"rt14_fundBillList"`                 //拓展参数(AccountPayQuery)
	Rt15ChannelRetCode        string `json:"rt15_channelRetCode"`               //拓展参数(AccountPayQuery)
	Rt16OutTransactionOrderId string `json:"rt16_outTransactionOrderId"`        //拓展参数(AccountPayQuery)
	Rt17BankType              string `json:"rt17_bankType"`                     //拓展参数(AccountPayQuery)
	Rt18SubOpenId             string `json:"rt18_subOpenId"`                    //拓展参数(AccountPayQuery)
	Rt19OrderAttribute        string `json:"rt19_orderAttribute"`               //拓展参数(AccountPayQuery)
	Rt20MarketingRule         string `json:"rt20_marketingRule"`                //拓展参数(AccountPayQuery)
	Rt21PromotionDetail       string `json:"rt21_promotionDetail"`              //拓展参数(AccountPayQuery)
	Rt22CreditAmount          string `json:"rt22_creditAmount,omitempty"`       //拓展参数(AccountPayQuery)
	Rt23PaymentAmount         string `json:"rt23_paymentAmount,omitempty"`      //拓展参数(AccountPayQuery)
	Rt24OrderCompleteDate     string `json:"rt24_orderCompleteDate"`            //拓展参数(AccountPayQuery)
	Rt25AppPayType            string `json:"rt25_appPayType"`                   //拓展参数(AccountPayQuery)
	Rt26AppId                 string `json:"rt26_appId"`                        //拓展参数(AccountPayQuery)
	RuleJson                  string `json:"ruleJson"`                          //拓展参数(AccountPayQuery)
	ProductFee                string `json:"productFee,omitempty"`              //拓展参数(AccountPayQuery)
	ChannelSettlementAmount   string `json:"channelSettlementAmount,omitempty"` //拓展参数(AccountPayQuery)
	RealCreditAmount          string `json:"realCreditAmount,omitempty"`        //拓展参数(AccountPayQuery)
	TradeType                 string `json:"tradeType"`                         //拓展参数(AccountPayQuery)
	ChargeFlag                string `json:"chargeFlag"`                        //拓展参数(AccountPayQuery)
	UpAddData                 string `json:"upAddData"`                         //拓展参数(AccountPayQuery)
	ResvData                  string `json:"resvData"`                          //拓展参数(AccountPayQuery)
	OnlineCardType            string `json:"onlineCardType"`                    //拓展参数(AccountPayQuery)
	SubMerchantNo             string `json:"subMerchantNo"`                     //拓展参数(AccountPayQuery)
	FeeRate                   string `json:"feeRate,omitempty"`                 //拓展参数(AccountPayQuery)
	FeeAccountAmt             string `json:"feeAccountAmt,omitempty"`           //拓展参数(AccountPayQuery)
	VoucherDetailList         string `json:"voucherDetailList"`                 //拓展参数(AccountPayQuery)
	ReceiverFee               string `json:"receiverFee,omitempty"`             //拓展参数(AccountPayQuery)
	OfflineFee                string `json:"offlineFee,omitempty"`              //拓展参数(AccountPayQuery)
	SignatureType             string `json:"signatureType"`                     //拓展参数(AccountPayQuery)
	Sign                      string `json:"sign"`                              //拓展参数(AccountPayQuery)

}

/*
 * 进件接口
 */

type MerchantEntryResponse struct {
	Success       bool            `json:"success"`       //交易类型
	Code          string          `json:"code"`          //签名类型
	Message       string          `json:"message"`       //时间戳
	Data          json.RawMessage `json:"data"`          //是否需要验签
	Sign          string          `json:"sign"`          //响应码
	Hostname      string          `json:"hostname"`      //响应码
	SignType      string          `json:"signType"`      //响应码
	EncryptionKey string          `json:"encryptionKey"` //响应码
}

type MdPdConfResponse struct {
	MerchantNo string `json:"merchantNo"` //子商户编号
	Status     string `json:"status"`     //状态
}
