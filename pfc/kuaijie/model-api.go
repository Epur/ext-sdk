package kuaijie

// 账户余额查询请求报文
type QueryCustomAcctInfoRequest struct {
	AcctNo string `json:"acctNo"` // 平台账户号， 业务账户，用于资金处理
}

// 账户余额查询响应报文
type QueryCustomAcctInfoResponse struct {
	AcctNo        string `json:"acctNo"`        // 平台账户号 业务账户，用于资金处理
	AcctBalance   int64  `json:"acctBalance"`   // 总余额,单位：分 总余额=其他账户金额之和
	AcctValid     int64  `json:"acctValid"`     // 可用金额 单位：分
	AcctHold      int64  `json:"acctHold"`      // 冻结金额 单位：分
	AcctWaitSett  int64  `json:"acctWaitSett"`  // 待结算账户金额 单位：分
	AcctWaitSplit int64  `json:"acctWaitSplit"` // 待分账账户金额 单位：分 开通分账功能后才有该值
	AcctCharge    int64  `json:"acctCharge"`    // 手续费支出账户金额 单位：分 开通指定账户后才有该值
	AcctPromotion int64  `json:"acctPromotion"` // 营销费用账户金额 单位：分 开通指定账户后才有该值
}

// 银行账户绑定
type CustomBankAcctBindRequest struct {
	AcctNo string `json:"acctNo,omitempty"` // 平台账户号， 业务账户，用于资金处理
}

// 银行账户绑定
type CustomBankAcctBindProtected struct {
	BankCode           string `json:"bankCode,omitempty"`           // 开户行代码
	BankName           string `json:"bankName,omitempty"`           // 开户行名称
	BankAcctType       string `json:"bankAcctType,omitempty"`       // 账户类型 pCard：个人借记卡（户名为法人姓名） eGeneral：对公一般户（户名为企业证照名称）
	BankAcctNo         string `json:"bankAcctNo,omitempty"`         // 账号
	BankAcctName       string `json:"bankAcctName,omitempty"`       // 户名
	BankProvince       string `json:"bankProvince,omitempty"`       // 开户省份
	BankCity           string `json:"bankCity,omitempty"`           // 开户城市
	BankAcctBindMobile string `json:"bankAcctBindMobile,omitempty"` // 银行预留的 手机号
}

// 银行账户绑定
type CustomBankAcctBindResponse struct {
	TxnState  string `json:"txnState"`  // 处理状态 1：成功 4：失败
	RspRemark string `json:"rspRemark"` // 处理详情 失败原因
}

// 银行账户绑定查询
type QueryCustomBankAcctBindInfoRequest struct {
	AcctNo string `json:"acctNo,omitempty"` // 平台账户号， 业务账户，用于资金处理
}

// 银行账户绑定查询
type QueryCustomBankAcctBindInfoResponse struct {
	Accounts []QueryCustomBankAcctBindInfo `json:"accounts"` // 结算账户
}

// 银行账户绑定查询
type QueryCustomBankAcctBindInfo struct {
	BankCode           string `json:"bankCode"`           // 开户行代码
	BankName           string `json:"bankName"`           // 开户行名称
	BankAcctType       string `json:"bankAcctType"`       // 账户类型 pCard：个人借记卡（户名为法人姓名）eGeneral：对公一般户（户名为企业证照名称）
	BankAcctNo         string `json:"bankAcctNo"`         // 账号
	BankAcctName       string `json:"bankAcctName"`       // 户名
	BankProvince       string `json:"bankProvince"`       // 开户省份
	BankCity           string `json:"bankCity"`           // 开户城市
	BankAcctBindMobile string `json:"bankAcctBindMobile"` // 银行预留的手机号
	Default            string `json:"default"`            // 默认0 0:否 1:是
}

// 发起转账-无验证码
type AcctTransferAmtNoIdentCodeRequest struct {
	AcctNoFrom string `json:"acctNoFrom,omitempty"`              // 转出方平台账户号
	AcctNoTo   string `json:"acctNoTo,omitempty"`                // 转入方平台账户号
	TxnAmt     int64  `json:"txnAmt,omitempty" validate:"gte=0"` // 转账金额 单位：分 数值必须大于 0
	TxnRemark  string `json:"txnRemark,omitempty"`               // 订单附言
	NotifyUrl  string `json:"notifyUrl,omitempty"`               // 后台通知地址
	Memo       string `json:"memo,omitempty"`                    // 附加数据
	BsnExt     string `json:"bsnExt,omitempty"`                  // 业务扩展信息
}

// 发起转账-无验证码
type AcctTransferAmtNoIdentCodeResponse struct {
	TxnState  string `json:"txnState"`  // 订单状态0：已受理 1：成功 4：失败
	TxnAmtOut int64  `json:"txnAmtOut"` // (付款账户) 扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn  int64  `json:"txnAmtIn"`  // (收款账户) 入账金额 单位：分 txnState 为 1 时必填
	RspRemark string `json:"rspRemark"` // 处理详情 失败原因等
}

// 发起转账-转账结果
type AcctTransferAmtResultRequest struct {
	OrigTxnDate    string `json:"origTxnDate,omitempty"` // 转账订单日期,格式:yyyyMMdd
	OrigCusTraceNo string `json:"origCusTraceNo"`        // 转账订单的客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo"`        // 转账订单的系统跟踪号
}

// 发起转账-转账结果
type AcctTransferAmtResultResponse struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
	TxnState       string `json:"txnState"`                 // 订单状态 0：已受理 1：成功 4：失败
	OrigRspCode    string `json:"origRspCode"`              // 提现订单应答码 txnState 为 1 或 4 时必填
	OrigRspMsg     string `json:"origRspMsg"`               // 提现订单应答信息 txnState 为 1 或 4 时必填
	TxnAmtOut      int64  `json:"txnAmtOut"`                // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn       int64  `json:"txnAmtIn"`                 // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	RspRemark      string `json:"rspRemark"`                // 处理详情 失败原因等
	Memo           string `json:"memo"`                     // 附加数据 原订单上送的自定义数据
}

// 发起提现-无验证码
type AcctWithdrawNoIdentCodeRequest struct {
	AcctNo     string `json:"acctNo,omitempty"`                  // 平台账户号
	SettAcctNo string `json:"settAcctNo,omitempty"`              // 提现账号
	TxnAmt     int64  `json:"txnAmt,omitempty" validate:"gte=0"` // 提现金额 单位：分 数值必须大于 0
	TxnRemark  string `json:"txnRemark,omitempty"`               // 订单附言
	NotifyUrl  string `json:"notifyUrl,omitempty"`               // 后台通知地址
	Memo       string `json:"memo,omitempty"`                    // 附加数据
	BsnExt     string `json:"bsnExt,omitempty"`                  // 业务扩展信息
}

// 发起提现-无验证码
type AcctWithdrawNoIdentCodeResponse struct {
	TxnState     string `json:"txnState"`     // 订单状态0：已受理 1：成功 4：失败
	TxnAmtOut    int64  `json:"txnAmtOut"`    // (付款账户) 扣账金额 单位：分 txnState为1时必填
	TxnAmtIn     int64  `json:"txnAmtIn"`     // (收款账户) 入账金额 单位：分 txnState为1时必填
	BankAcctNo   string `json:"bankAcctNo"`   // 收款银行账号 txnState为1时必填
	BankAcctName string `json:"bankAcctName"` // 收款银行户名 txnState为1时必填
	RspRemark    string `json:"rspRemark"`    // 处理详情 失败原因等
}

// 提现结果查询
type AcctWithdrawResultRequest struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
}

// 提现结果查询
type AcctWithdrawResultResponse struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
	TxnState       string `json:"txnState"`                 // 订单状态 0：已受理 1：成功 4：失败
	OrigRspCode    string `json:"origRspCode"`              // 提现订单应答码 txnState 为 1 或 4 时必填
	OrigRspMsg     string `json:"origRspMsg"`               // 提现订单应答信息 txnState 为 1 或 4 时必填
	TxnAmtOut      int64  `json:"txnAmtOut"`                // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn       int64  `json:"txnAmtIn"`                 // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	BankAcctNo     string `json:"bankAcctNo"`               // 收款银行账号 txnState 为 1 时必填
	BankAcctName   string `json:"bankAcctName"`             // 收款银行户名 txnState 为 1 时必填
	RspRemark      string `json:"rspRemark"`                // 处理详情 失败原因等
	Memo           string `json:"memo"`                     // 附加数据 原订单上送的自定义数据
}

// 发起付款-无验证码
type AcctPayAmtNoIdentCodeRequest struct {
	AcctNo     string `json:"acctNo,omitempty"`                  // 平台账户号
	BankAcctNo string `json:"bankAcctNo,omitempty"`              // 银行账号 该银行账号必须事先通过“银行账户绑定配置”接口完成绑定
	TxnAmt     int64  `json:"txnAmt,omitempty" validate:"gte=0"` // 提现金额 单位：分 数值必须大于 0
	TxnRemark  string `json:"txnRemark,omitempty"`               // 订单附言
	NotifyUrl  string `json:"notifyUrl,omitempty"`               // 后台通知地址
	Memo       string `json:"memo,omitempty"`                    // 附加数据
	BsnExt     string `json:"bsnExt,omitempty"`                  // 业务扩展信息
}

// 发起付款-无验证码
type AcctPayAmtNoIdentCodeResponse struct {
	TxnState     string `json:"txnState"`     // 订单状态0：已受理 1：成功 4：失败
	TxnAmtOut    int64  `json:"txnAmtOut"`    // (付款账户) 扣账金额 单位：分 txnState为1时必填
	TxnAmtIn     int64  `json:"txnAmtIn"`     // (收款账户) 入账金额 单位：分 txnState为1时必填
	BankAcctNo   string `json:"bankAcctNo"`   // 收款银行账号 txnState为1时必填
	BankAcctName string `json:"bankAcctName"` // 收款银行户名 txnState为1时必填
	RspRemark    string `json:"rspRemark"`    // 处理详情 失败原因等
}

// 付款结果查询
type AcctPayAmtResultRequest struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
}

// 付款结果查询
type AcctPayAmtResultResponse struct {
	OrigTxnDate    string `json:"origTxnDate"`              // 转账订单日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 转账订单的 客户订单号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 转账订单的 系统跟踪号
	TxnState       string `json:"txnState"`                 // 订单状态 0：已受理 1：成功 4：失败
	OrigRspCode    string `json:"origRspCode"`              // 提现订单应答码 txnState 为 1 或 4 时必填
	OrigRspMsg     string `json:"origRspMsg"`               // 提现订单应答信息 txnState 为 1 或 4 时必填
	TxnAmtOut      string `json:"txnAmtOut"`                // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	TxnAmtIn       string `json:"txnAmtIn"`                 // (付款账户)扣账金额 单位：分 txnState 为 1 时必填
	BankAcctNo     string `json:"bankAcctNo"`               // 收款银行账号 txnState 为 1 时必填
	BankAcctName   string `json:"bankAcctName"`             // 收款银行户名 txnState 为 1 时必填
	RspRemark      string `json:"rspRemark"`                // 处理详情 失败原因等
	Memo           string `json:"memo"`                     // 附加数据 原订单上送的自定义数据
}

// 资金归集--账户余额查询
type CashSweepQueryAcctBalInfoRequest struct {
	IssAcctNo string `json:"issAcctNo"` // 平台账户号 资金归集账户，用于接收外部平台资金
}

// 资金归集--账户余额查询
type CashSweepQueryAcctBalInfoResponse struct {
	IssAcctNo   string `json:"issAcctNo"`   // 平台账户号
	AcctBalance int64  `json:"acctBalance"` // 总余额
	AcctValid   int64  `json:"acctValid"`   // 可用金额
	AcctHold    int64  `json:"acctHold"`    // 冻结金额
}

// 客户入网

// 文件上传
type FileUploadRequest struct {
	FileName string `json:"fileName,omitempty"` // 文件名
	FileData string `json:"fileData,omitempty"` // 文件大小不超过 1M 将文件字节流，转为 base64 编码字符串(无需添加 “data:”标识)
}

// 文件上传
type FileUploadResponse struct {
	FileId    string `json:"fileId,omitempty"`    // 文件ID
	FileUrl   string `json:"fileUrl,omitempty"`   // 文件链接
	RspRemark string `json:"rspRemark,omitempty"` // 文件备注
}

// 企业入网
type CustomInNetEntRequest struct {
	InnetOwner     CustomInNetEntOwner      `json:"innetOwner,required"`     // 入网主体信息
	InnetJuridical CustomInNetEntJuridical  `json:"innetJuridical,required"` // 法人信息
	InnetContact   CustomInNetEntContact    `json:"innetContact,required"`   // 联系人信息
	InnetAddress   CustomInNetEntAddress    `json:"innetAddress,required"`   // 经营地址
	InnetIndustry  CustomInNetEntIndustry   `json:"innetIndustry,required"`  // 行业信息
	InnetAccount   CustomInNetEntAccount    `json:"innetAccount,required"`   // 结算账户
	Product        []CustomInNetEntProduct  `json:"product,omitempty"`       // 产品信息
	Function       []CustomInNetEntFunction `json:"function,omitempty"`      // 产品功能信息
	NotifyUrl      string                   `json:"notifyUrl,omitempty"`     // 后台通知地址
}

// 企业入网 -- 入网主体信息
type CustomInNetEntOwner struct {
	CifType               string `json:"cifType,required"`                // 主体类型 21：企业 22：个体工商户
	CifName               string `json:"cifName,required"`                // 主体名称
	CifBriefName          string `json:"cifBriefName,required"`           // 主体简称
	EntCreditCode         string `json:"entCreditCode,required"`          // 证件号码
	EntCertPhoto          string `json:"entCertPhoto,required"`           // 证件照片,调用“文件上传接口”后返回的文件 id
	TpBankPermit          string `json:"tpBankPermit,required"`           // 开户许可证编号
	TpBankPermitPhoto     string `json:"tpBankPermitPhoto,required"`      // 开户许可证照片,调用“文件上传接口”后返回的文件 id
	EntServicePhone       string `json:"entServicePhone,required"`        // 客服电话
	EntCertInHandPhoto    string `json:"entCertInHandPhoto,omitempty"`    // 手持证件照片
	EntOfficeOutsidePhoto string `json:"entOfficeOutsidePhoto,omitempty"` // 经营场所外景照
	EntOfficeInsidePhoto  string `json:"entOfficeInsidePhoto,omitempty"`  // 经营场所内景照
}

// 企业入网 -- 法人信息
type CustomInNetEntJuridical struct {
	JuridicalName      string `json:"juridicalName,required"`      // 法人姓名
	JuridicalCertType  string `json:"juridicalCertType,required"`  // 法人证件类型 默认 P01 P01：身份证
	JuridicalCertNo    string `json:"juridicalCertNo,required"`    // 法人证件号码
	JuridicalCertFront string `json:"juridicalCertFront,required"` // 法人证件人像面,调用“文件上传接口”后返回的文件 id
	JuridicalCertBack  string `json:"juridicalCertBack,required"`  // 法人证件非人像面,调用“文件上传接口”后返回的文件 id
	JuridicalPhone     string `json:"juridicalPhone,omitempty"`    // 法人手机号
}

// 企业入网 -- 联系人信息
type CustomInNetEntContact struct {
	Name     string `json:"name,required"`     // 联系人姓名
	CertType string `json:"certType,required"` // 联系人证件类型 默认 P01 P01：身份证
	CertNo   string `json:"certNo,required"`   // 联系人证件号码
	Mobile   string `json:"mobile,required"`   // 联系人手机号
	Email    string `json:"email,omitempty"`   // 联系人邮箱
}

// 企业入网 -- 经营地址
type CustomInNetEntAddress struct {
	Province string `json:"province,required"` // 实际经营地 所在省 国家行政区划代码
	City     string `json:"city,required"`     // 实际经营地 所在市 国家行政区划代码
	District string `json:"district,required"` // 实际经营地 所在区 国家行政区划代码
	Address  string `json:"address,required"`  // 实际经营地 所在详细地址 不含省市区
}

// 企业入网 -- 行业信息
type CustomInNetEntIndustry struct {
	IndustryCategory string `json:"industryCategory,required"` // 一级行业分类 按照实际经营场景选择对应的行业分类编码 具体编码详见附录
	IndustryType     string `json:"industryType,required"`     // 一级行业分类 按照实际经营场景选择对应的行业分类编码 具体编码详见附录
	Mcc              string `json:"mcc,required"`              // 银联商户 类型代码
}

// 企业入网 -- 结算账户
type CustomInNetEntAccount struct {
	BankCode           string `json:"bankCode,required"`            // 开户行代码
	BankName           string `json:"bankName,omitempty"`           // 开户行名称
	BankAcctType       string `json:"bankAcctType,required"`        // 账户类型 pCard：个人借记卡（户名为法人姓名） eGeneral：对公一般户（户名为企业证照名称）
	BankAcctNo         string `json:"bankAcctNo,required"`          // 账号
	BankAcctName       string `json:"bankAcctName,omitempty"`       // 户名
	BankProvince       string `json:"bankProvince,omitempty"`       // 开户省份
	BankCity           string `json:"bankCity,omitempty"`           // 开户城市
	BankAcctBindMobile string `json:"bankAcctBindMobile,omitempty"` // 银行预留的 手机号
}

// 企业入网 -- 产品信息
type CustomInNetEntProduct struct {
	ProductCode string `json:"productCode,required"` // 产品代码 详见附录
	SettPeriod  string `json:"settPeriod,omitempty"` // 结算周期 默认 T1 D1：D+1 T1：T+1 M：手动结算
	SceneCode   string `json:"sceneCode,required"`   // 场景代码 付款产品必填 详见附录
	SceneFile   string `json:"sceneFile,required"`   // 场景资质文件 付款产品必填 调用“文件上传接口”后返回的文件 id
}

// 企业入网 -- 产品功能信息
type CustomInNetEntFunction struct {
	FunctionCode string `json:"functionCode,required"` // 功能代码 目前仅支持分账功能 split：分账
	SceneCode    string `json:"sceneCode,omitempty"`   // 场景代码 开通分账能力必填 详见附录
	SceneFile    string `json:"sceneFile,omitempty"`   // 场景资质文件 付款产品必填 调用“文件上传接口”后返回的文件 id
}

// 企业入网 -- 返回信息
type CustomInNetEntResponse struct {
	SubCusCode string `json:"subCusCode"` // 入网客户代码 受理成功时必填
	RspRemark  string `json:"rspRemark"`  // 入网详情 失败原因
}

// 小微/个人入网
type CustomInNetMicroRequest struct {
	InnetOwner     CustomInNetMicroOwner      `json:"innetOwner,required"`     // 入网主体信息
	InnetJuridical CustomInNetMicroJuridical  `json:"innetJuridical,required"` // 法人信息
	InnetAddress   CustomInNetMicroAddress    `json:"innetAddress,required"`   // 经营地址
	InnetAccount   CustomInNetMicroAccount    `json:"innetAccount,required"`   // 结算账户
	Product        []CustomInNetMicroProduct  `json:"product,omitempty"`       // 产品信息
	Function       []CustomInNetMicroFunction `json:"function,omitempty"`      // 产品功能信息
	NotifyUrl      string                     `json:"notifyUrl,omitempty"`     // 后台通知地址
}

// 小微入网 -- 入网主体信息
type CustomInNetMicroOwner struct {
	CifName      string `json:"cifName,required"`      // 主体名称
	CifBriefName string `json:"cifBriefName,required"` // 主体简称
}

// 小微入网 -- 法人信息
type CustomInNetMicroJuridical struct {
	JuridicalName      string `json:"juridicalName,required"`      // 法人姓名
	JuridicalCertType  string `json:"juridicalCertType,required"`  // 法人证件类型 默认 P01 P01：身份证
	JuridicalCertNo    string `json:"juridicalCertNo,required"`    // 法人证件号码
	JuridicalCertFront string `json:"juridicalCertFront,required"` // 法人证件人像面,调用“文件上传接口”后返回的文件 id
	JuridicalCertBack  string `json:"juridicalCertBack,required"`  // 法人证件非人像面,调用“文件上传接口”后返回的文件 id
	JuridicalPhone     string `json:"juridicalPhone,omitempty"`    // 法人手机号
}

// 小微入网 -- 经营地址
type CustomInNetMicroAddress struct {
	Province string `json:"province,required"` // 实际经营地 所在省 国家行政区划代码
	City     string `json:"city,required"`     // 实际经营地 所在市 国家行政区划代码
	District string `json:"district,required"` // 实际经营地 所在区 国家行政区划代码
	Address  string `json:"address,required"`  // 实际经营地 所在详细地址 不含省市区
}

// 小微入网 -- 结算账户
type CustomInNetMicroAccount struct {
	BankCode           string `json:"bankCode,required"`            // 开户行代码
	BankName           string `json:"bankName,omitempty"`           // 开户行名称
	BankAcctType       string `json:"bankAcctType,required"`        // 账户类型 pCard：个人借记卡（户名为法人姓名） eGeneral：对公一般户（户名为企业证照名称）
	BankAcctNo         string `json:"bankAcctNo,required"`          // 账号
	BankAcctName       string `json:"bankAcctName,omitempty"`       // 户名
	BankProvince       string `json:"bankProvince,omitempty"`       // 开户省份
	BankCity           string `json:"bankCity,omitempty"`           // 开户城市
	BankAcctBindMobile string `json:"bankAcctBindMobile,omitempty"` // 银行预留的 手机号
}

// 小微入网 -- 产品信息
type CustomInNetMicroProduct struct {
	ProductCode string `json:"productCode,required"` // 产品代码 详见附录
	SettPeriod  string `json:"settPeriod,omitempty"` // 结算周期 默认 T1 D1：D+1 T1：T+1 M：手动结算
	//SceneCode   string `json:"sceneCode,required"`   // 场景代码 付款产品必填 详见附录
	//SceneFile   string `json:"sceneFile,required"`   // 场景资质文件 付款产品必填 调用“文件上传接口”后返回的文件 id
}

// 小微入网 -- 产品功能信息
type CustomInNetMicroFunction struct {
	FunctionCode string `json:"functionCode,required"` // 功能代码 目前仅支持分账功能 split：分账
	SceneCode    string `json:"sceneCode,omitempty"`   // 场景代码 开通分账能力必填 详见附录
	SceneFile    string `json:"sceneFile,omitempty"`   // 场景资质文件 付款产品必填 调用“文件上传接口”后返回的文件 id
}

// 小微入网 -- 返回信息
type CustomInNetMicroResponse struct {
	SubCusCode string `json:"subCusCode"` // 入网客户代码 受理成功时必填
	RspRemark  string `json:"rspRemark"`  // 入网详情 失败原因
}

// 入网结果查询请问报文
type CustomInNetResultRequest struct {
	OrigTxnDate    string `json:"origTxnDate,omitempty"`    // 入网请求日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 入网请求的客户流水号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 入网请求的系统跟踪号
}

// 入网结果查询响应报文
type CustomInNetResultResponse struct {
	OrigTxnDate    string `json:"origTxnDate,omitempty"`    // 入网请求日期
	OrigCusTraceNo string `json:"origCusTraceNo,omitempty"` // 入网请求的客户流水号
	OrigSysTraceNo string `json:"origSysTraceNo,omitempty"` // 入网请求的系统跟踪号
	TxnState       string `json:"txnState"`                 // 入网状态 1：成功 2：审核中 4：失败
	SubCusCode     string `json:"subCusCode"`               // 入网客户代码 受理成功时必填
	RspRemark      string `json:"rspRemark"`                // 入网详情 失败原因
}
