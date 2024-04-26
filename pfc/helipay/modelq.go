package helipay

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

// 交易类型
const (
	BIZ_TYPE_MS  = "MerchantSettlement"      //结算接口
	BIZ_TYPE_MSQ = "MerchantSettlementQuery" //结算查询接口
	BIZ_TYPE_APS = "AccountPaySub"           //子商户支付接口
	BIZ_TYPE_APQ = "AccountPayQuery"         //单笔查询接口
	BIZ_TYPE_MAQ = "MerchantAccountQuery"    //商户余额查询
	BIZ_TYPE_AP  = "AccountPay"              //支付接口通知
	BIZ_TYPE_QR  = "AppPay"                  //扫码/下单接口

)

var QRPAY_REQ_FIELDS = []string{
	"P1_bizType",
	"P2_orderId",
	"P3_customerNumber",
	"P4_payType",
	"P5_orderAmount",
	"P6_currency",
	"P7_authcode",
	"P8_appType",
	"P9_notifyUrl",
	"P10_successToUrl",
	"P11_orderIp",
	"P12_goodsName",
	"P13_goodsDetail",
	"P14_desc",
}

/*
 *商户下单（转账）请求接口
 */

type AccountPaySubRequest struct {
	P1BizType   string `json:"P1_bizType"`   //交易类型
	P2SignType  string `json:"P2_signType"`  //签名类型
	P3Timestamp string `json:"P3_timestamp"` //时间戳

	P4OrderId        string `json:"P4_orderId"`        //商户订单号
	P5CustomerNumber string `json:"P5_customerNumber"` //付款商户商编
	P6Ext            *P6Ext `json:"P6_ext"`            //拓展参数
}

type P6Ext struct {
	InMerchantNo      string            `json:"inMerchantNo"`                //收款商户商编
	OrderType         string            `json:"orderType"`                   //订单类型
	Amount            decimal.Decimal   `json:"amount"`                      //订单金额
	ServerCallbackUrl string            `json:"serverCallbackUrl,omitempty"` //服务器通知地址
	GoodsName         string            `json:"goodsName"`                   //商品名称(活动补贴)
	OrderDesc         string            `json:"orderDesc,omitempty"`         //订单备注
	ProductType       string            `json:"productType,omitempty"`       //原订单产品类型
	AssociatedOrderNo string            `json:"associatedOrderNo,omitempty"` //原订单商户订单号
	InEscrow          string            `json:"inEscrow,omitempty"`          //担保交易标识
	SplitBillRules    []json.RawMessage `json:"splitBillRules,omitempty"`    //收款分账规则串
	BelongsType       string            `json:"belongsType,omitempty"`       //手续费承担方向
}

/*
 *商户订单查询接口
 */

type AccountPayQueryRequest struct {
	P1BizType   string `json:"P1_bizType"`   //交易类型
	P2SignType  string `json:"P2_signType"`  //签名类型
	P3Timestamp string `json:"P3_timestamp"` //时间戳

	P4OrderId        string `json:"P4_orderId"`        //商户订单号
	P5CustomerNumber string `json:"P5_customerNumber"` //商户商编
}

/*
 *商户提现接口
 */
type MerchantSettlementRequest struct {
	P1BizType string `json:"P1_bizType"` //交易类型
	SignType  string `json:"signType"`   //签名类型
	//P3Timestamp string `json:"Pt3_timestamp"` //时间戳

	P2OrderId        string          `json:"P2_orderId"`             //商户订单号
	P3CustomerNumber string          `json:"P3_customerNumber"`      //商户商编
	P4Amount         decimal.Decimal `json:"P4_amount"`              //金额
	P5Summary        string          `json:"P5_summary,omitempty"`   //备注
	P6NotifyUrl      string          `json:"P6_notifyUrl,omitempty"` //结果通知地址
}

/*
 *商户提现查询接口
 */

type MerchantSettlementQueryRequest struct {
	P1BizType string `json:"P1_bizType"` //交易类型
	SignType  string `json:"signType"`   //签名类型
	//P3Timestamp string `json:"Pt3_timestamp"` //时间戳

	P2OrderId        string `json:"P2_orderId,omitempty"` //商户订单号
	P3CustomerNumber string `json:"P3_customerNumber"`    //商户商编
	P4SettleDate     string `json:"P4_settleDate"`        //结算日期
}

/*
 *商户余额查询接口
 */

type MerchantAccountQueryRequest struct {
	P1BizType string `json:"P1_bizType"` //交易类型

	P2CustomerNumber string `json:"P2_customerNumber"` //商户商编号
	P3Timestamp      string `json:"P3_timestamp"`      //时间戳
}

/*
 *扫码/刷单接口
 */

type AppPayRequest struct {
	P1BizType string `json:"P1_bizType"` //交易类型

	P2OrderId        string  `json:"P2_orderId"`        //订单号
	P3CustomerNumber string  `json:"P3_customerNumber"` //商户编号
	P4PayType        string  `json:"P4_payType"`        //支付类型
	P5OrderAmount    float64 `json:"P5_orderAmount"`    //订单额
	P6Currency       string  `json:"P6_currency"`       //币种
	P7Authcode       string  `json:"P7_authcode"`       //授权码
	P8AppType        string  `json:"P8_appType"`        //类型
	P9NotifyUrl      string  `json:"P9_notifyUrl"`      //通知地址
	P10SuccessToUrl  string  `json:"P10_successToUrl"`  //返回地址
	P11OrderIp       string  `json:"P11_orderIp"`       //订单ID
	P12GoodsName     string  `json:"P12_goodsName"`     //商品名称
	P13GoodsDetail   string  `json:"P13_goodsDetail"`   //商品明细
	P14Desc          string  `json:"P14_desc"`          //描述
	P15SubMerchantId string  `json:"P15_subMerchantId"` //描述
	//P16AppId           string  `json:"P16_appId"`          //公众号id
	//P17LimitCreditPay  string  `json:"P17_limitCreditPay"` //是否限制借贷记
	//P18GoodsTag        string  `json:"P18_goodsTag"`       //商品标记
	//P19Guid            string  `json:"P19_guid"`           //微信上送的唯一号
	//P20MarketingRule   string  `json:"P20_marketingRule"`  //营销参数规则
	//P21Identity        string  `json:"P21_identity"`       //实名参数
	//SplitBillType      string  `json:"splitBillType"`      //分账类型
	//RuleJson           string  `json:"ruleJson"`           //分账规则
	//HbfqNum            string  `json:"hbfqNum"`            //分期数
	//DeviceInfo         string  `json:"deviceInfo"`         //终端号
	//StoreId            string  `json:"storeId"`            //商户门店编号
	//AlipayStoreId      string  `json:"alipayStoreId"`      //支付宝店铺编号
	//TimeExpire         string  `json:"timeExpire"`         //超时时间
	//IndustryRefluxInfo string  `json:"industryRefluxInfo"` //支付宝行业数据回流信息
	//TermInfo           string  `json:"termInfo"`           //银联终端信息
	//OpenId             string  `json:"openId"`             //用户id
	//AuthConfirmMode    string  `json:"authConfirmMode"`    //预授权确认模式
	//TerminalSysBindNo  string  `json:"terminalSysBindNo"`  //终端绑定号
	//EncryptRandNum     string  `json:"encryptRandNum"`     //加密随机因子
	//SecretText         string  `json:"secretText"`         //密文数据
	//SceneInfo          string  `json:"sceneInfo"`          //场景信息
	//EduSubject         string  `json:"eduSubject"`         //学校名称、场景名称
	//BusinessParams     string  `json:"businessParams"`     //商户传入业务信息，具体值要和支付宝约定
	//ExtendParams       string  `json:"extendParams"`       //业务扩展参数
	//Pid                string  `json:"pid"`                //服务商pid
	//EncryptionKey      string  `json:"encryptionKey"`      //加密密钥
	SignatureType string `json:"signatureType"` //签名方式
}
