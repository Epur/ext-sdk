package lianlianpay

import "encoding/json"

const (
	// APIGatewaySG endpoint
	APIGateway = "https://gtest-open-api.lianlianpay-inc.com/api"
	AuthURL    = "https://gtest.lianlianpay-inc.com/openapi/"
)

const (
	AccessTokenURL = "/token"
	//强制还款状态
	REPAY_FORCECANCEL        = "1" //取消强制扣款
	REPAY_MATYDEDUT_USER     = "2" //强制扣款（用户维度）
	REPAY_MATYDEDUT_PLATFORM = "3" //强制扣款（平台维度）

	//还款方式
	REPAY_TYPE_LSR       = "LUMP_SUM_REPAY"   //一次性还本付息
	REPAY_TYPE_UNLIMITED = "UNLIMITED"        //随借随还
	REPAY_TYPE_ET        = "EVEN_TOTAL"       //:等额本息;
	REPAY_TYPE_EP        = "EVEN_PRINCIPAL"   //等额本金
	REPAY_TYPE_MI        = "MONTHLY_INTEREST" //先息后本

	//币种
	CURRENCY_RMB = "CNY" //人民币
	CURRENCY_US  = "USD" //美元
	CURRENCY_GB  = "GBP" //英镑
	CURRENCY_EU  = "EUR" //欧元
	CURRENCY_JP  = "JPY" //日元
	CURRENCY_CAD = "CAD" //加元
	//RefreshURL     = "/token"
)

type Response struct {
	Code      string          `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
	Result    json.RawMessage `json:"result"`
}

//用户授权请求

type UserAuthApplyRequest struct {
	ProdType   string `json:"prodType" validate:"required"`
	ExtApplyId string `json:"extApplyId" validate:"required"`
	ReturnUrl  string `json:"returnUrl" validate:"required"`
}

//用户授权查询
//注：userId与extApplyId不能全部为空。

type UserAuthQueryRequest struct {
	UserId     string `json:"userId"`
	ProdType   string `json:"ProdType" validate:"required"`
	ExtApplyId string `json:"extApplyId"`
}

//放款申请

type LoanApplyRequest struct {
	UserId         string `json:"userId" validate:"required"`
	ProdType       string `json:"prodType" validate:"required"`
	FundParty      string `json:"fundParty" validate:"required"`
	ExtLoanId      string `json:"extLoanId" validate:"required"`
	LoanAmount     string `json:"loanAmount" validate:"required"`
	Currency       string `json:"currency" validate:"required"`
	BorrowerName   string `json:"borrowerName" validate:"required"`
	BorrowerType   string `json:"borrowerType" validate:"required"`
	BorrowerCertNo string `json:"borrowerCertNo" validate:"required"`
	RepayType      string `json:"repayType" validate:"required"`
}

//放款确认

type LoanConfirmRequest struct {
	UserId        string `json:"userId" validate:"required"`
	ProdType      string `json:"prodType" validate:"required"`
	LoanId        string `json:"loanId" validate:"required"`
	ExtLoanId     string `json:"extLoanId" validate:"required"`
	LoanAmount    string `json:"loanAmount" validate:"required"`
	LoanBeginDate string `json:"loanBeginDate" validate:"required"`
	LoanEndDate   string `json:"loanEndDate" validate:"required"`
	Currency      string `json:"currency" validate:"required"`
	Status        string `json:"status" validate:"required"`
	Memo          string `json:"memo"`
}

//贷款结清

type LoanFinishedRequest struct {
	UserId     string `json:"userId" validate:"required"`
	ProdType   string `json:"prodType" validate:"required"`
	LoanId     string `json:"loanId" validate:"required"`
	ExtLoanId  string `json:"extLoanId"`
	FinishDate string `json:"finishDate" validate:"required"`
	Memo       string `json:"memo"`
}

//强制还款

type ForceRepaymentRequest struct {
	UserId       string `json:"userId" validate:"required"`
	ProdType     string `json:"prodType" validate:"required"`
	LoanId       string `json:"loanId" validate:"required"`
	FundParty    string `json:"fundParty" validate:"required"`
	Status       string `json:"status"`
	DeductAmount string `json:"deductAmount" validate:"required"`
	NotifyTime   string `json:"notifyTime" validate:"required"`
	DeductId     string `json:"deductId" validate:"required"`
}

//强制还款

type NotifyRepaymentRequest struct {
	UserId          string `json:"userId" validate:"required"`
	LoanId          string `json:"loanId" validate:"required"`
	ExtRepaymentId  string `json:"extRepaymentId" validate:"required"`
	CapitalAmount   string `json:"capitalAmount"`
	InterestAmount  string `json:"interestAmount"`
	FineAmount      string `json:"fineAmount"`
	LoanFeeAmount   string `json:"loanFeeAmount"`
	RepaymentAmount string `json:"repaymentAmount" validate:"required"`
	Currency        string `json:"currency" validate:"required"`
	RepaymentTime   string `json:"repaymentTime" validate:"required"`
	NotifyTime      string `json:"notifyTime" validate:"required"`
}

//type getExchangeRateResponse struct {
//	Rate           string `json:"rate"`
//	SourceCurrency string `json:"sourceCurrency"`
//	TargetCurrency string `json:"targetCurrency"`
//	RateTime       string `json:"rateTime"`
//	BaseCurrency   string `json:"baseCurrency"`
//}
//
//type getHolderResponse struct {
//	HolderId     string `json:"holderId"`
//	HoldType     int64  `json:"holdType"`
//	HolderName   string `json:"holderName"`
//	HolderNameEn string `json:"holderNameEn"`
//}
//
//type getAccountBalanceResponse struct {
//	Currency   string `json:"currency"`
//	Balance    string `json:"balance"`
//	AccountNum int64  `json:"accountNum"`
//}
//
//type getAccountResponse struct {
//	Currency          string  `json:"currency"`
//	Platform          string  `json:"platform"`
//	AccountId         string  `json:"accountId"`
//	HolderName        string  `json:"holderName"`
//	CreateTime        string  `json:"createTime"`
//	Balance           string  `json:"balance"`
//	BalanceUpdateTime string  `json:"balanceUpdateTime"`
//	ShopName          string  `json:"shopName"`
//	ShopStatus        int     `json:"shopStatus"`
//	VirtualCardNo     string  `json:"virtualCardNo"`
//	AppendInfo        string  `json:"appendInfo"`
//	FeeRate           float64 `json:"feeRate"`
//	SellerId          string  `json:"sellerId"`
//	Region            string  `json:"region"`
//}
//
//type getTransactionResponse struct {
//	TransId         string      `json:"transId"`
//	TransType       string      `json:"transType"`
//	Platform        string      `json:"platform"`
//	Region          string      `json:"region"`
//	Currency        string      `json:"currency"`
//	AccountId       string      `json:"accountId"`
//	VirtualCardNo   string      `json:"virtualCardNo"`
//	Amount          string      `json:"amount"`
//	Balance         string      `json:"balance"`
//	BusiType        string      `json:"busiType"`
//	TransTime       string      `json:"transTime"`
//	Remark          interface{} `json:"remark"`
//	MarketplaceName interface{} `json:"marketplaceName"`
//}
//
//type getTransactionEntryListResponse struct {
//	TransId        string      `json:"transId"`
//	CreditStatus   string      `json:"creditStatus"`
//	Platform       string      `json:"platform"`
//	ShopName       string      `json:"shopName"`
//	VirtualCardNo  string      `json:"virtualCardNo"`
//	AccountId      string      `json:"accountId"`
//	Currency       string      `json:"currency"`
//	Amount         string      `json:"amount"`
//	CreateTime     string      `json:"createTime"`
//	AcctDate       string      `json:"acctDate"`
//	FundsId        string      `json:"fundsId"`
//	Region         string      `json:"region"`
//	OriginAmount   interface{} `json:"originAmount"`
//	OriginCurrency interface{} `json:"originCurrency"`
//}
//
//type getWithdrawRecordListResponse struct {
//	TransId            string `json:"transId"`
//	ExternalTransId    string `json:"externalTransId"`
//	WithdrawStatus     int    `json:"withdrawStatus"`
//	Platform           string `json:"platform"`
//	Region             string `json:"region"`
//	WithdrawAmount     string `json:"withdrawAmount"`
//	WithdrawCurrency   string `json:"withdrawCurrency"`
//	ArrivalAmount      string `json:"arrivalAmount"`
//	ArrivalCurrency    string `json:"arrivalCurrency"`
//	FeeAmount          string `json:"feeAmount"`
//	ActualFeeAmount    string `json:"actualFeeAmount"`
//	FeeCurrency        string `json:"feeCurrency"`
//	CreateTime         string `json:"createTime"`
//	Rate               string `json:"rate"`
//	RateTime           string `json:"rateTime"`
//	BankAccountName    string `json:"bankAccountName"`
//	BankName           string `json:"bankName"`
//	BankCardNo         string `json:"bankCardNo"`
//	WithdrawDetailList []struct {
//		DetailTransId   string `json:"detailTransId"`
//		ShopName        string `json:"shopName"`
//		AccountId       string `json:"accountId"`
//		Amount          string `json:"amount"`
//		Currency        string `json:"currency"`
//		FeeRate         string `json:"feeRate"`
//		FeeAmount       string `json:"feeAmount"`
//		ActualFeeAmount string `json:"actualFeeAmount"`
//	} `json:"withdrawDetailList"`
//	FeeRate string `json:"feeRate"`
//}
//
//type getBankcardListResponse struct {
//	CardId          string `json:"cardId"`
//	BankCardNo      string `json:"bankCardNo"`
//	BankAccountName string `json:"bankAccountName"`
//	BankName        string `json:"bankName"`
//	Currency        string `json:"currency"`
//	District        string `json:"district"`
//	CardType        string `json:"cardType"`
//}
