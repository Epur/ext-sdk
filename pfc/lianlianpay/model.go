package lianlianpay

import "encoding/json"

const (
	// APIGatewaySG endpoint
	APIGateway = "https://gtest-open-api.lianlianpay-inc.com/api"
	AuthURL    = "https://gtest.lianlianpay-inc.com/openapi/"
)

const (
	AccessTokenURL = "/token"
	//RefreshURL     = "/token"
)

type Response struct {
	Code      string          `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
	Result    json.RawMessage `json:"result"`
}

type getExchangeRateResponse struct {
	Rate           string `json:"rate"`
	SourceCurrency string `json:"sourceCurrency"`
	TargetCurrency string `json:"targetCurrency"`
	RateTime       string `json:"rateTime"`
	BaseCurrency   string `json:"baseCurrency"`
}

type getHolderResponse struct {
	HolderId     string `json:"holderId"`
	HoldType     int64  `json:"holdType"`
	HolderName   string `json:"holderName"`
	HolderNameEn string `json:"holderNameEn"`
}

type getAccountBalanceResponse struct {
	Currency   string `json:"currency"`
	Balance    string `json:"balance"`
	AccountNum int64  `json:"accountNum"`
}

type getAccountResponse struct {
	Currency          string  `json:"currency"`
	Platform          string  `json:"platform"`
	AccountId         string  `json:"accountId"`
	HolderName        string  `json:"holderName"`
	CreateTime        string  `json:"createTime"`
	Balance           string  `json:"balance"`
	BalanceUpdateTime string  `json:"balanceUpdateTime"`
	ShopName          string  `json:"shopName"`
	ShopStatus        int     `json:"shopStatus"`
	VirtualCardNo     string  `json:"virtualCardNo"`
	AppendInfo        string  `json:"appendInfo"`
	FeeRate           float64 `json:"feeRate"`
	SellerId          string  `json:"sellerId"`
	Region            string  `json:"region"`
}

type getTransactionResponse struct {
	TransId         string      `json:"transId"`
	TransType       string      `json:"transType"`
	Platform        string      `json:"platform"`
	Region          string      `json:"region"`
	Currency        string      `json:"currency"`
	AccountId       string      `json:"accountId"`
	VirtualCardNo   string      `json:"virtualCardNo"`
	Amount          string      `json:"amount"`
	Balance         string      `json:"balance"`
	BusiType        string      `json:"busiType"`
	TransTime       string      `json:"transTime"`
	Remark          interface{} `json:"remark"`
	MarketplaceName interface{} `json:"marketplaceName"`
}

type getTransactionEntryListResponse struct {
	TransId        string      `json:"transId"`
	CreditStatus   string      `json:"creditStatus"`
	Platform       string      `json:"platform"`
	ShopName       string      `json:"shopName"`
	VirtualCardNo  string      `json:"virtualCardNo"`
	AccountId      string      `json:"accountId"`
	Currency       string      `json:"currency"`
	Amount         string      `json:"amount"`
	CreateTime     string      `json:"createTime"`
	AcctDate       string      `json:"acctDate"`
	FundsId        string      `json:"fundsId"`
	Region         string      `json:"region"`
	OriginAmount   interface{} `json:"originAmount"`
	OriginCurrency interface{} `json:"originCurrency"`
}

type getWithdrawRecordListResponse struct {
	TransId            string `json:"transId"`
	ExternalTransId    string `json:"externalTransId"`
	WithdrawStatus     int    `json:"withdrawStatus"`
	Platform           string `json:"platform"`
	Region             string `json:"region"`
	WithdrawAmount     string `json:"withdrawAmount"`
	WithdrawCurrency   string `json:"withdrawCurrency"`
	ArrivalAmount      string `json:"arrivalAmount"`
	ArrivalCurrency    string `json:"arrivalCurrency"`
	FeeAmount          string `json:"feeAmount"`
	ActualFeeAmount    string `json:"actualFeeAmount"`
	FeeCurrency        string `json:"feeCurrency"`
	CreateTime         string `json:"createTime"`
	Rate               string `json:"rate"`
	RateTime           string `json:"rateTime"`
	BankAccountName    string `json:"bankAccountName"`
	BankName           string `json:"bankName"`
	BankCardNo         string `json:"bankCardNo"`
	WithdrawDetailList []struct {
		DetailTransId   string `json:"detailTransId"`
		ShopName        string `json:"shopName"`
		AccountId       string `json:"accountId"`
		Amount          string `json:"amount"`
		Currency        string `json:"currency"`
		FeeRate         string `json:"feeRate"`
		FeeAmount       string `json:"feeAmount"`
		ActualFeeAmount string `json:"actualFeeAmount"`
	} `json:"withdrawDetailList"`
	FeeRate string `json:"feeRate"`
}

type getBankcardListResponse struct {
	CardId          string `json:"cardId"`
	BankCardNo      string `json:"bankCardNo"`
	BankAccountName string `json:"bankAccountName"`
	BankName        string `json:"bankName"`
	Currency        string `json:"currency"`
	District        string `json:"district"`
	CardType        string `json:"cardType"`
}
