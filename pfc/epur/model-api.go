package epur

type GetTokenResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	ClientId         string `json:"client_id"`
	Scope            string `json:"scope"`
	Openid           string `json:"openid"`
}

type GetSellerResponse struct {
	Bid           int    `json:"bid"`
	BName         string `json:"bName"`
	Channel       string `json:"channel"`
	Spid          int    `json:"spid"`
	AdminName     string `json:"adminName"`
	AdminMobile   string `json:"adminMobile"`
	Contacts      string `json:"contacts"`
	Phone         string `json:"phone"`
	Status        string `json:"status"`
	Status1       string `json:"status1"`
	Ctime         int    `json:"ctime"`
	UserType      string `json:"userType"`
	Operator      string `json:"operator"`
	Addr          string `json:"addr"`
	MergePrice    string `json:"mergePrice"`
	TranPrice     string `json:"tranPrice"`
	FreeMergeRate string `json:"freeMergeRate"`
	StmtFlag      string `json:"stmtFlag"`
	StmtCron      string `json:"stmtCron"`
}
