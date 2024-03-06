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
