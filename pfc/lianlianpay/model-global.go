package lianlianpay

// 用户授权响应

type UserAuthApplyResponse struct {
	ExtApplyId string `json:"extApplyId"`
	ApplyId    string `json:"applyId"`
	GatewayUrl string `json:"gatewayUrl"`
}

//用户授权查询响应

type UserAuthQueryResponse struct {
	LoginId      string `json:"loginId"`
	UserId       string `json:"userId"`
	ExtApplyId   string `json:"extApplyId"`
	ApplyId      string `json:"applyId"`
	UserName     string `json:"userName"`
	IdType       string `json:"idType"`
	IdNo         string `json:"idNo"`
	LegalIdType  string `json:"legalIdType"`
	LegalIdNo    string `json:"legalIdNo"`
	LegalName    string `json:"legalName"`
	ApplyResult  string `json:"applyResult"`
	ApplyDate    string `json:"applyDate"`
	Mobile       string `json:"mobile"`
	UserType     string `json:"userType"`
	DistrictType string `json:"districtType"`
	Memo         string `json:"memo"`
}

//放款申请响应

type LoanApplyResponse struct {
	LoanId     string `json:"loanId"`     //支用单号
	UserId     string `json:"userId"`     //用户id
	ExtApplyId string `json:"extApplyId"` //资方放款单号
}

//放款确认响应

type LoanConfirmResponse struct {
	UserId string `json:"userId"` //用户类型
	//ProdType      string `json:"prodType"`      //产品类型
	LoanId    string `json:"loanId"`    //支用单号
	ExtLoanId string `json:"extLoanId"` //资方放款号
	//LoanAmount    string `json:"loanAmount"`    //放款金额
	//LoanBeginDate string `json:"loanBeginDate"` //贷款起始日期
	//LoanEndDate   string `json:"loanEndDate"`   //贷款结束日期
	//Currency      string `json:"currency"`      //币种
	//Status        string `json:"status"`        //状态
	//Memo          string `json:"memo"`          //备注
}

//贷款结清响应

type LoanFinishedResponse struct {
	//UserId string `json:"userId"` //用户类型
	////ProdType      string `json:"prodType"`      //产品类型
	//LoanId    string `json:"loanId"`    //支用单号
	//ExtLoanId string `json:"extLoanId"` //资方放款号
	////LoanAmount    string `json:"loanAmount"`    //放款金额
	////LoanBeginDate string `json:"loanBeginDate"` //贷款起始日期
	////LoanEndDate   string `json:"loanEndDate"`   //贷款结束日期
	////Currency      string `json:"currency"`      //币种
	////Status        string `json:"status"`        //状态
	////Memo          string `json:"memo"`          //备注
}

//强制还款响应

type ForceRepaymentResponse struct {
	UserId string `json:"userId"` //用户类型
	//ProdType      string `json:"prodType"`      //产品类型
	LoanId string `json:"loanId"` //支用单号
}

//还款单通知

type NotifyRepaymentResponse struct {
	UserId      string `json:"userId"`      //用户类型
	RepaymentId string `json:"repaymentId"` //产品类型
	LoanId      string `json:"loanId"`      //支用单号
}
