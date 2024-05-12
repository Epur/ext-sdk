package helipay

import (
	"github.com/Epur/ext-sdk/logger"
	"github.com/Epur/ext-sdk/model"
)

type Api struct {
	Setting *model.Setting
}

/*
 * 设置并获取合利宝API(接口层)
 */

func New(setting *model.Setting) *Api {
	// 初始化日志服务
	logger.New("logs", "info")
	return &Api{Setting: setting}
}

/*
 *结算下单（转账）
 */
func (p *Api) AccountPay(Body model.BodyMap) *model.Client {
	logger.CmbcLogger.Info("合利宝下单接口...")

	c := NewClient(p.Setting)
	c.SetPath(`/trx/accountPay/interface.action`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("P1_bizType", "P2_signType", "P3_timestamp",
		"P4_orderId", "P5_customerNumber", "P6_ext"); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := PayResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
 *扫码/刷卡下单
 */
func (p *Api) QrcodePay(Body model.BodyMap) *model.Client {
	logger.CmbcLogger.Info("扫码/刷卡下单接口...")

	c := NewClient(p.Setting)
	c.SetPath(`/trx/app/interface.action`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("P1_bizType", "P2_orderId", "P3_customerNumber",
		"P4_payType", "P5_orderAmount", "P6_currency",
		"P7_authcode",
		"P8_appType",
		"P12_goodsName",
		"signatureType"); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AppPayResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
 *公众号/JS/服务窗预下单接口
 */

func (p *Api) PrePay(Body model.BodyMap) *model.Client {
	logger.CmbcLogger.Info("公众号/JS/服务窗预下单接口...")

	c := NewClient(p.Setting)
	c.SetPath(`/trx/app/interface.action`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("P1_bizType", "P2_orderId", "P3_customerNumber",
		"P4_payType", "P5_appid", "P8_openid",
		"P9_orderAmount",
		"P10_currency",
		"P11_appType",
		"P14_orderIp",
		"P15_goodsName",
		"signatureType"); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AppPayPublicResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
 * 订单结果查询（转账结果查询）
 */
func (p *Api) AccountPayQuery(Body model.BodyMap) *model.Client {
	logger.CmbcLogger.Info("合利宝订单结果查询接口...")

	c := NewClient(p.Setting)
	c.SetPath(`/trx/accountPay/interface.action`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("P1_bizType", "P2_signType", "P3_timestamp",
		"P4_orderId", "P5_customerNumber"); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	response := PayQueryResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
 * 结算查询(提现结果查询)
 */

func (p *Api) MerchantSettlementQuery(Body model.BodyMap) *model.Client {
	logger.HeliLogger.Info("合利宝结算结果查询接口...")

	c := NewClient(p.Setting)
	c.SetPath(`/trx/transfer/interface.action`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("P1_bizType", "P2_orderId", "P3_customerNumber",
		"P4_settleDate", "signType"); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := MerchantSettlementQueryResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}

/*
 * 结算接口(提现)
 */

func (p *Api) MerchantSettlement(Body model.BodyMap) *model.Client {
	logger.CmbcLogger.Info("合利宝结算接口...")

	c := NewClient(p.Setting)
	c.SetPath(`/trx/transfer/interface.action`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("P1_bizType", "P2_orderId", "P3_customerNumber",
		"P4_amount", "signType"); c.Err != nil {
		logger.CmbcLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.CmbcLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := MerchantSettlementResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.CmbcLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
 * 商户余额查询接口
 */

func (p *Api) MerchantAccountQuery(Body model.BodyMap) *model.Client {
	logger.CmbcLogger.Info("合利宝商户余额查询接口...")

	c := NewClient(p.Setting)
	c.SetPath(`/trx/merchant/interface.action`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("P1_bizType", "P2_customerNumber", "P3_timestamp"); c.Err != nil {
		logger.CmbcLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.CmbcLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := MerchantAccountQueryResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.CmbcLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
**产品手续费收取方式修改
 */
func (p *Api) MdfyPdConf(Body model.BodyMap, otherParam model.BodyMap) *model.Client {
	logger.CmbcLogger.Info("产品手续费收取方式修改接口...")

	c := NewClient(p.Setting)
	c.SetPath(BIZ_TXN_PMENRY).
		SetMethod("POST").
		SetBody(Body)

	//if c.Err = Body.CheckEmptyError("merchantNo", "type", "productType",
	//	"value"); c.Err != nil {
	//	logger.HeliLogger.Error("ERROR:", c.Err.Error())
	//	return &c.Client
	//}

	if c.Err = Body.CheckEmptyError("interfaceName", "merchantNo", "body"); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}

	response := AppPayPublicResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		logger.HeliLogger.Error("ERROR:", c.Err.Error())
		return &c.Client
	}
	c.Response.Response.DataTo = response

	return &c.Client
}
