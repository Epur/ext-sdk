package kuaijie

import "encoding/json"

// 公共请求报文
type PublicRequest struct {
	Version    string `json:"version"`    // 版本号 默认填“2.0”
	CusReqTime string `json:"cusReqTime"` // 请求报文发送时间 格式：yyyyMMddHHmmss 请求报文发送时间与平台接收时间应在 300 秒内，否则将被平台视为无效请求
	CusTraceNo string `json:"cusTraceNo"` // 客户请求流水号 同一接入方一个自然日内唯一
	CusCode    string `json:"cusCode"`    // 客户代码 平台分配给客户的唯一识别号
}

type Response struct {
	//Code      string          `json:"code"`
	//Message   string          `json:"message"`
	//RequestID string          `json:"request_id"`
	//Data      json.RawMessage `json:"data"`
	//Result    json.RawMessage `json:"result"`
	MsgPublic  PublicResponse  `json:"msgPublic"`
	MsgPrivate json.RawMessage `json:"msgPrivate"`
}

// 公共应答报文
type PublicResponse struct {
	TxnDate    string `json:"txnDate"`    // 订单日期 格式：yyyyMMdd
	CusCode    string `json:"cusCode"`    // 客户代码 平台分配给客户的唯一识别号
	CusTraceNo string `json:"cusTraceNo"` // 客户请求流水号 同请求报文
	SysTraceNo string `json:"sysTraceNo"` // 系统跟踪号 平台唯一
	SysRspTime string `json:"sysRspTime"` // 应答报文发送时间 格式：yyyyMMddHHmmss
	RspCode    string `json:"rspCode"`    // 应答码  0000 表示查询成功
	RspMsg     string `json:"rspMsg"`     // 应答信息
}
