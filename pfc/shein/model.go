package shein

// 签名因子
var SignFactors = []string{
	"x-lt-appid",
	"x-lt-openKeyId",
	"x-lt-timestamp",
}

type OrderListRequest struct {
	QueryType   int64  `json:"queryType" validate:"required"`
	StartTime   string `json:"startTime" validate:"required"`
	EndTime     string `json:"endTime" validate:"required"`
	Page        int64  `json:"page" validate:"required"`
	PageSize    int64  `json:"pageSize" validate:"required"`
	OrderStatus int64  `json:"orderStatus,omitempty"`
}
