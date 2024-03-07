package helipay

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

/*
 *商户下单（转账）请求接口
 */

type AccountPaySubRequest struct {
	P1BizType   string `json:"Pt1_bizType"`   //交易类型
	P2SignCode  string `json:"Pt2_signType"`  //签名类型
	P3Timestamp string `json:"Pt3_timestamp"` //时间戳

	P4OrderId        string `json:"Pt4_orderId"`        //商户订单号
	P5CustomerNumber string `json:"Pt5_customerNumber"` //付款商户商编
	P6Ext            *P6Ext `json:"P6_ext"`             //拓展参数
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
	P1BizType   string `json:"Pt1_bizType"`   //交易类型
	P2SignCode  string `json:"Pt2_signType"`  //签名类型
	P3Timestamp string `json:"Pt3_timestamp"` //时间戳

	P4OrderId        string `json:"Pt4_orderId"`        //商户订单号
	P5CustomerNumber string `json:"Pt5_customerNumber"` //商户商编
}
