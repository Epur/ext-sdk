package cmbchina

import "github.com/shopspring/decimal"

// 1.可经办业务模式查询DCLISMOD
type DCLISMODResponse struct {
	Ntqmdlstz []DCLISMODNtqmdlstz `json:"ntqmdlstz"`
}

// 1.可经办业务模式查询DCLISMOD
type DCLISMODNtqmdlstz struct {
	Busmod string `json:"busmod"`
	Modals string `json:"modals"`
}

// 2.查询可经办的账户列表 DCLISACC
type DCLISACCResponse struct {
	Ntqpyaccz []DCLISACCNtqpyaccz `json:"ntqpyaccz"`
}

// 2.查询可经办的账户列表 DCLISACC
type DCLISACCNtqpyaccz struct {
	Ccynbr  string `json:"ccynbr"`   // 币种
	Bbknbr  string `json:"bbknbr"`   // 分行号
	Accnbr  string `json:"accnbr"`   // 帐号
	Accnam  string `json:"accnam"`   // 户名
	Relnbr  string `json:"relnbr"`   // 企业编号
	CRelnbr string `json:"c_relnbr"` // 公司名称
	MAccnbr string `json:"m_accnbr"` // 账户信息
	Accflg  string `json:"accflg"`
}

// 3.账户详细信息查询NTQACINF
type NTQACINFResponse struct {
	Ntqacinfz []NTQACINFNtqacinfz `json:"ntqacinfz"`
}

// 3.账户详细信息查询NTQACINF
type NTQACINFNtqacinfz struct {
	Ccynbr string          `json:"ccynbr"` // 币种 A.3 货币代码表
	Accitm string          `json:"accitm"` // 科目  科目代码
	Bbknbr string          `json:"bbknbr"` // 分行号 A.1 招商分行
	Accnbr string          `json:"accnbr"` // 帐号
	Accnam string          `json:"accnam"` // 注解  一般为户名
	Accblv decimal.Decimal `json:"accblv"` // 上日余额 // 当intcod ='S'时，这个字段的值显示为"头寸额度（集团支付子公司余额）"是子公司的虚拟余额
	Onlblv decimal.Decimal `json:"onlblv"` // 联机余额
	Hldblv decimal.Decimal `json:"hldblv"` // 冻结余额
	Avlblv decimal.Decimal `json:"avlblv"` // 可用余额 可用余额=联机余额+预期余额-冻结金额+透支额度
	Lmtovr string          `json:"lmtovr"` // 透支额度
	Stscod string          `json:"stscod"` // 状态 A=活动，B=冻结，C=关户
	Intcod string          `json:"intcod"` // 利息码 S=子公司虚拟余额
	Intrat decimal.Decimal `json:"intrat"` // 年利率
	Opndat string          `json:"opndat"` // 开户日
	Mutdat string          `json:"mutdat"` // 到期日
	Inttyp string          `json:"inttyp"` // 利率类型 ZZZ=不计息 TD2=定期不滚积数 TD1=定期滚积数 D99=手工计息 B01=凭证式国债 D12=挂牌对公活期保证金关户结息 D11=挂牌对公活期保证金按季结息C01=协议计息D32=离岸挂牌个人活期存款D03=挂牌同业活期存款D02=在岸挂牌个人活期存款D31=离岸挂牌对公活期存款D01=在岸挂牌对公活期存款
	Dpstxt string          `json:"dpstxt"` // 存期 定期时，取值： 一天 七天 一个月 三个月 六个月 一年 二年 三年 四年 五年
	Dactyp string          `json:"dactyp"`
	Opnbbk string          `json:"opnbbk"`
	Opnbrn string          `json:"opnbrn"`
	Relnbr string          `json:"relnbr"` // 客户关系号
}

// 4.查询账户历史余额 NTQABINF
type NTQABINFResponse struct {
	Ntqabinfz []QueryAccountBalHistory1 `json:"ntqabinfz"` // 账户信息ntqabinfz（多记录）
	Ntqabinfy []QueryAccountBalHistory2 `json:"ntqabinfy"` // 续传接口ntqabinfy（单记录）
}

// 4.查询账户历史余额NTQABINF
type QueryAccountBalHistory1 struct {
	Bbknbr  string          `json:"bbknbr"`   // 分行号
	Accnbr  string          `json:"accnbr"`   // 帐号
	Trsdat  string          `json:"trsdat"`   // 交易日期
	Balamt  decimal.Decimal `json:"balamt"`   // 联机余额
	CCcynbr string          `json:"c_ccynbr"` // 币种名称
	Rsv30z  string          `json:"rsv30Z"`   // 保留字段
}

// 4.查询账户历史余额NTQABINF
type QueryAccountBalHistory2 struct {
	Bbknbr string `json:"bbknbr"` // 分行号
	Accnbr string `json:"accnbr"` // 帐号
	Bgndat string `json:"bgndat"` // 起始日期 格式：YYYYMMDD,开始日期与结束日期的间隔不能超过31天
	Enddat string `json:"enddat"` // 结束日期 格式：YYYYMMDD
	Ctnkey string `json:"ctnkey"` // 续传字段 首次传空。当响应报文中含有ntqabinfy接口且该接口中ctnkey字段不为空时表明需要续传。将响应报文中ntqabinfy接口中的内容放入下次请求报文的ntqabinfy接口中获取后续数据。

}

// 5. 查询分行号信息NTACCBBK 账户信息ntaccbbkz（多记录）
type QueryAcctOpenBankResponse struct {
	Ntaccbbky []QueryAcctOpenBankInfo `json:"ntaccbbky"` // 续传接口ntaccbbky，当响应报文中含有ntaccbbky接口且该接口中ctnkey字段不为空时标明需要续传，将响应报文中ntaccbbky接口中的内容放入下次请求的请求报文中获取后续数据。
	Ntaccbbkz []QueryAcctOpenBankInfo `json:"ntaccbbkz"` // 账户信息ntaccbbkz（多记录）
}

// 5. 查询分行号信息NTACCBBK
type QueryAcctOpenBankInfo struct {
	Bbknbr string `json:"bbknbr"` // 分行号
	Fctval string `json:"fctval"` // 账号
}

// 6. 批量查询余额NTQADINF
type QueryAcctBalBatchRequest struct {
	Ntqadinfx []QueryAcctBalBatchNX `json:"ntqadinfx"`
}

// 6. 批量查询余额NTQADINF
type QueryAcctBalBatchNX struct {
	Accnbr string `json:"accnbr"` // 账号
	Bbknbr string `json:"bbknbr"` // 分行号
}

// 6. 批量查询余额NTQADINF
type QueryAcctBalBatchResponse struct {
	Ntqacinfz []QueryAcctBalBatchInfo `json:"ntqacinfz"` // 账户信息ntqacinfz（多记录）
}

// 6. 批量查询余额NTQADINF
type QueryAcctBalBatchInfo struct {
	Ccynbr string          `json:"ccynbr"` // 币种 A.3 货币代码表
	Accitm string          `json:"accitm"` // 科目  科目代码
	Bbknbr string          `json:"bbknbr"` // 分行号 A.1 招商分行
	Accnbr string          `json:"accnbr"` // 帐号
	Accnam string          `json:"accnam"` // 注解  一般为户名
	Relnbr string          `json:"relnbr"` // 客户关系号
	Accblv string          `json:"accblv"` // 上日余额 // 当intcod ='S'时，这个字段的值显示为"头寸额度（集团支付子公司余额）"是子公司的虚拟余额
	Onlblv decimal.Decimal `json:"onlblv"` // 联机余额
	Hldblv decimal.Decimal `json:"hldblv"` // 冻结余额
	Avlblv decimal.Decimal `json:"avlblv"` // 可用余额 可用余额=联机余额+预期余额-冻结金额+透支额度
	Lmtovr string          `json:"lmtovr"` // 透支额度
	Stscod string          `json:"stscod"` // 状态 A=活动，B=冻结，C=关户
	Intcod string          `json:"intcod"` // 利息码 S=子公司虚拟余额
	Intrat decimal.Decimal `json:"intrat"` // 年利率
	Opndat string          `json:"opndat"` // 开户日
	Mutdat string          `json:"mutdat"` // 到期日
	Inttyp string          `json:"inttyp"` // 利率类型 ZZZ=不计息 TD2=定期不滚积数 TD1=定期滚积数 D99=手工计息 B01=凭证式国债 D12=挂牌对公活期保证金关户结息 D11=挂牌对公活期保证金按季结息C01=协议计息D32=离岸挂牌个人活期存款D03=挂牌同业活期存款D02=在岸挂牌个人活期存款D31=离岸挂牌对公活期存款D01=在岸挂牌对公活期存款
	Dpstxt string          `json:"dpstxt"` // 存期 定期时，取值： 一天 七天 一个月 三个月 六个月 一年 二年 三年 四年 五年
	Errcod string          `json:"errcod"` // 错误码
	Errtxt string          `json:"errtxt"` // 错误说明
}

// 7.账户交易信息查询trsQryByBreakPoint
type TranPointY struct {
	AcctNbr            string `json:"acctNbr"`            // 子账号
	TransDate          string `json:"transDate"`          // 交易日期
	ExpectNextSequence string `json:"expectNextSequence"` // 期望下一记账序号
}

// 7.账户交易信息查询trsQryByBreakPoint
type QueryAccountTranInfoRequest struct {
	X1 []TranPointX `json:"TRANSQUERYBYBREAKPOINT_X1"`
	//Y1 []TranPointY `json:"TRANSQUERYBYBREAKPOINT_Y1"` // 首次传空，执行断点或续传请求时必传，将接口返回的TRANSQUERYBYBREAKPOINT_Y1值传入
}

// 7.账户交易信息查询trsQryByBreakPoint
type TranPointX struct {
	CardNbr             string `json:"cardNbr"`             // 账号 网银账务查询业务模式下的可用账号
	BeginDate           string `json:"beginDate"`           // 开始日期 查询交易时间段的开始日期
	EndDate             string `json:"endDate"`             // 结束日期
	TransactionSequence string `json:"transactionSequence"` // 起始记账序号 仅在不传入续传键值时有效，表示初始查询从第几笔开始查询，默认从第1笔开始查询。
	CurrencyCode        string `json:"currencyCode"`        // 币种
	QueryAcctNbr        string `json:"queryAcctNbr"`        // 继续查询账号
	Reserve             string `json:"reserve"`             // 保留字段
}

// 7.账户交易信息查询trsQryByBreakPoint
type QueryAccountTranInfoResponse struct {
	Y1 []TranPointY  `json:"TRANSQUERYBYBREAKPOINT_Y1"` // 首次传空，执行断点或续传请求时必传，将接口返回的TRANSQUERYBYBREAKPOINT_Y1值传入
	Z1 []TranPointZ1 `json:"TransQueryByBreakPoint_Z1"` //
	Z2 []TranPointZ2 `json:"TransQueryByBreakPoint_Z2"` //
}

// 7.账户交易信息查询trsQryByBreakPoint
type TranPointZ1 struct {
	CtnFlag      string          `json:"ctnFlag"`      // Y|N，Y表示还有记录需要查询，N表示当前已查询完毕
	QueryAcctNbr string          `json:"queryAcctNbr"` // 当ctnFlag为Y时，下一次请求时携带该值进行下一次查询
	DebitNums    string          `json:"debitNums"`    // 借方笔数
	DebitAmount  decimal.Decimal `json:"debitAmount"`  // 借方金额
	CreditNums   string          `json:"creditNums"`   // 贷方笔数
	CreditAmount decimal.Decimal `json:"creditAmount"` // 贷方金额
	Reserve      string          `json:"reserve"`      // 保留字
}

// 7.账户交易信息查询trsQryByBreakPoint
type TranPointZ2 struct {
	TransDate           string          `json:"transDate"`           // 交易日
	TransSequenceIdn    string          `json:"transSequenceIdn"`    // 流水号
	TransTime           string          `json:"transTime"`           // 交易时间
	ValueDate           string          `json:"valueDate"`           // 起息日
	LoanCode            string          `json:"loanCode"`            // 借贷码
	TransAmount         decimal.Decimal `json:"transAmount"`         // 交易金额
	CurrencyNbr         string          `json:"currencyNbr"`         // 币种
	TextCode            string          `json:"textCode"`            // 交易类型
	BillNumber          string          `json:"billNumber"`          // 票据号
	RemarkTextClt       string          `json:"remarkTextClt"`       // 你方摘要
	ReversalFlag        string          `json:"reversalFlag"`        // 冲帐标志 *为冲帐，X为补帐 （冲账交易与原交易借贷相反）
	AcctOnlineBal       decimal.Decimal `json:"acctOnlineBal"`       // 余额
	ExtendedRemark      string          `json:"extendedRemark"`      // 扩展摘要
	CtpAcctNbr          string          `json:"ctpAcctNbr"`          // 收付方帐号
	CtpAcctName         string          `json:"ctpAcctName"`         // 收付方名称
	CtpBankName         string          `json:"ctpBankName"`         // 收付方开户行行名
	CtpBankAddress      string          `json:"ctpBankAddress"`      // 收付方开户行地址
	FatOrSonAccount     string          `json:"fatOrSonAccount"`     // 母子公司帐号
	FatOrSonCompanyName string          `json:"fatOrSonCompanyName"` // 母子公司名称
	FatOrSonBankName    string          `json:"fatOrSonBankName"`    // 母子公司开户行行名
	FatOrSonBankAddress string          `json:"fatOrSonBankAddress"` // 母子公司开户行地址
	InfoFlag            string          `json:"infoFlag"`            // 信息标志 用于标识收/付方帐号和母/子公司的信息。 为空表示付方帐号和子公司； 为“1”表示收方帐号和子公司；为“2”表示收方帐号和母公司；为“3”表示原收方帐号和子公司
	BusinessName        string          `json:"businessName"`        // 业务名称
	BusinessText        string          `json:"businessText"`        // 网银业务摘要
	RequestNbr          string          `json:"requestNbr"`          // 网银流程实例号
	YurRef              string          `json:"yurRef"`              // 网银业务参考号
	VirtualNbr          string          `json:"virtualNbr"`          // 虚拟户编号
	MchOrderNbr         string          `json:"mchOrderNbr"`         // 商务支付订单号
	Reserve             string          `json:"reserve"`             // 保留字
}

// 8.PDF文件对账单获取DCTRSPDF
type CheckOrderPdfRequest struct {
	SX []CheckOrderPdfNFX `json:"sdktsinfx"`
	//NY []CheckOrderPdfTNY `json:"ntqacctny"`
}

// 8.PDF文件对账单获取DCTRSPDF
type CheckOrderPdfNFX struct {
	Bbknbr string `json:"bbknbr"` // 分行号
	Accnbr string `json:"accnbr"` // 账号 网银账务查询业务模式下的可用账号
	Bgndat string `json:"bgndat"` // 开始日期 查询交易时间段的开始日期
	Enddat string `json:"enddat"` // 结束日期
	Lowamt string `json:"lowamt"` // 最小金额
	Hghamt string `json:"hghamt"` // 最大金额
	Amtcdr string `json:"amtcdr"` // 借贷码 C:贷；D:借
	Genble string `json:"genble"` // 通知地址URL类型
}

// 8.PDF文件对账单获取DCTRSPDF
type CheckOrderPdfTNY struct {
	Ctndta string `json:"ctndta"` // 续传字段 首次传空，当返回接口有ntqacctny并且ntqacctny的续传字段(ctndta)有值时，将返回的ntqacctny放入请求报文body中继续查询，直到返回的不返回ntqacctny或者ntqacctny的续传字段为空，续传结束。
	Pagnbr string `json:"pagnbr"` // 页数
}

// 8.PDF文件对账单获取DCTRSPDF
type CheckOrderPdfResponse struct {
	DZ []CheckOrderPdfDFZ `json:"dctrspdfz"`
	NY []CheckOrderPdfTNY `json:"ntqacctny"`
}

// 8.PDF文件对账单获取DCTRSPDF
type CheckOrderPdfDFZ struct {
	Printid string `json:"printid"` // 返回的任务ID，请根据该ID查询PDF文件。
	Taskid  string `json:"taskid"`  // 如果使用主动通知，可以和通知结果关联
}

// 9.OFD文件对账单获取issueBillOfd
type IssueBillOfdRequest struct {
	CprDirectIssueBillOfdX1 []IssueBillOfdOfdX1 `json:"CprDirectIssueBillOfdX1"`
}
type IssueBillOfdOfdX1 struct {
	Accnbr    string `json:"accnbr"`
	Ccynbr    string `json:"ccynbr"`
	BillYear  string `json:"billYear"`
	BillMonth string `json:"billMonth"`
}

// 9.OFD文件对账单获取issueBillOfd
type IssueBillOfdResponse struct {
	CprDirectIssueBillOfdZ1 []struct {
		FileName string `json:"fileName"`
		IssueKey string `json:"issueKey"`
		SpecCode string `json:"specCode"`
	} `json:"CprDirectIssueBillOfdZ1"`
}

// 10.OFD文件对账单获取结果查询queryBillOfd
type QueryBillOfdResponse struct {
	CprDirectQueryBillOfdZ1 []struct {
		FinishFlag string `json:"finishFlag"`
		SpecCode   string `json:"specCode"`
	} `json:"CprDirectQueryBillOfdZ1"`
	CprDirectQueryBillOfdZ2 []struct {
		FileKey  string `json:"fileKey"`
		FileSeq  string `json:"fileSeq"`
		FileUrl  string `json:"fileUrl"`
		SpecCode string `json:"specCode"`
	} `json:"CprDirectQueryBillOfdZ2"`
}

// 11.单笔回单查询DCSIGREC
type DCSIGRECResponse struct {
	Checod string `json:"checod"`
	Fildat string `json:"fildat"`
	Istnbr string `json:"istnbr"`
}

// 12. 电子回单异步查询ASYCALHD
type ASYCALHDResponse struct {
	Asycalhdz1 struct {
		Rtncod string `json:"rtncod"`
		Rtnmsg string `json:"rtnmsg"`
		Rtndat string `json:"rtndat"`
	} `json:"asycalhdz1"`
	Ctnkeyz2 struct {
		Begamt string `json:"begamt"`
		Begdat string `json:"begdat"`
		Daltag string `json:"daltag"`
		Eacnbr string `json:"eacnbr"`
		Endamt string `json:"endamt"`
		Enddat string `json:"enddat"`
		Nxtdat string `json:"nxtdat"`
		Nxtnbr string `json:"nxtnbr"`
		Nxttim string `json:"nxttim"`
		Oprtyp string `json:"oprtyp"`
		Pagcnt string `json:"pagcnt"`
		Pattyp string `json:"pattyp"`
		Predat string `json:"predat"`
		Prenbr string `json:"prenbr"`
		Pretim string `json:"pretim"`
		Rrccod string `json:"rrccod"`
		Rrcflg string `json:"rrcflg"`
		Spc100 string `json:"spc100"`
	} `json:"ctnkeyz2"`
}

// 13. 异步打印结果查询DCTASKID
type DCTASKIDResponse struct {
	Fileurl string `json:"fileurl"`
	Fintim  string `json:"fintim"`
}

// 2. 企银支付单笔经办BB1PAYOP
type BusiPayRequest struct {
	M1 []BusiPayBmx1 `json:"bb1paybmx1"`
	X1 []BusiPayOpx1 `json:"bb1payopx1"`
	//X5 []BusiPayOpx5 `json:"bb1payopx5"`
}

// 2. 企银支付单笔经办BB1PAYOP
type BusiPayBmx1 struct {
	BusMod string `json:"busMod"` // 业务模式（模式编号） 可通过“可经办业务模式查询(DCLISMOD)”接口获得，也可通过前置机获得。
	BusCod string `json:"busCod"` // 业务类型（业务代码） N02030:支付
}

// 2. 企银支付单笔经办BB1PAYOP
// 收方行联行号 1、非标准银联卡，必须要填行名称或者联行号； 其他情况可空，但优先使用客户传入
// 2、若设置了收方限制，则行名联行号都必传；
// 3、若同时输入行名和联行号，则以联行号为准进行汇出；
// 4、若收方非招行户，请尽量补充完整账户信息以提高支付成功率; 若客户未填写收方开户行联行号，计费将根据系统识别的行号进行开户地判断，可能存在同城异地的误判，请知悉
type BusiPayOpx1 struct {
	DbtAcc string          `json:"dbtAcc"` // 转出帐号
	DmaNbr string          `json:"dmaNbr"` // 记账子单元编号
	CrtAcc string          `json:"crtAcc"` // 收方帐号
	CrtNam string          `json:"crtNam"` // 收方户名
	CrtBnk string          `json:"crtBnk"` // 收方开户行名称 1、非标准银联卡，必须要填行名称或者联行号；其他情况可空，但优先使用客户传入； 2、若设置了收方限制，则行名联行号都必传；3、若同时输入行名和联行号，则以联行号为准进行汇出；4、若收方非招行户，请尽量补充完整账户信息以提高支付成功率；
	CrtAdr string          `json:"crtAdr"` // 收方开户行地址 以下任意情况，收方开户行地址可不传： a.收方为招行账户；b.已输入收方开户行名称或收方开户行联行号（若客户未填写开户地，计费将根据行名/行号判断开户地进行同城异地判断，若行名中不含明确的地址信息，可能存在同城异地的误判，请知悉）；
	BrdNbr string          `json:"brdNbr"` // 收方行联行号
	CcyNbr string          `json:"ccyNbr"` // 币种 只支持10人民币
	TrsAmt decimal.Decimal `json:"trsAmt"` // 交易金额
	BnkFlg string          `json:"bnkFlg"` // 系统内标志 收方为招行户：传空或Y； 收方为他行户：传N；
	EptDat string          `json:"eptDat"` // 期望日
	EptTim string          `json:"eptTim"` // 期望时间
	StlChn string          `json:"stlChn"` // 结算通道
	NusAge string          `json:"nusAge"` // 用途
	CrtSqn string          `json:"crtSqn"` // 收方编号
	YurRef string          `json:"yurRef"` // 业务参考号 业务参考号，必须唯一。 业务参考号（YURREF）是每笔交易的唯一编号，是防止重复提交的重要手段，如需重复发送请求，请务必保证同一笔交易的业务参考号不变，否则会存在重复提交风险。
	BusNar string          `json:"busNar"` // 业务摘要
	NtfCh1 string          `json:"ntfCh1"` // 通知方式一（邮箱）
	NtfCh2 string          `json:"ntfCh2"` // 通知方式二（手机号）
	TrsTyp string          `json:"trsTyp"` // 业务种类 100001    普通汇兑 （默认值） 101001    慈善捐款 101002    其他 注：只有结算通道为“G 普通”或者“Q 快速”时，才支持101001、101002方式优惠手续费。
	RcvChk string          `json:"rcvChk"` // 行内收方账号户名校验 1：校验 空或者其他值：不校验 如果为1，行内收方账号与户名不相符则支付经办失败。
	DrpFlg string          `json:"drpFlg"` // 直汇普通标志 空时为A； A-普通 B-直汇（失败后不落人工处理）。 注：只有结算通道为“G 普通”或者“Q 快速”时，才支持。

}

// 2. 企银支付单笔经办BB1PAYOP
type BusiPayOpx5 struct {
	CopNbr string `json:"copNbr"` // 优惠券编号
}

// 2. 企银支付单笔经办BB1PAYOP
type BusiPayResponse struct {
	Z1 []BusiPayOpz1 `json:"bb1payopz1"`
}

// 2. 企银支付单笔经办BB1PAYOP
type BusiPayOpz1 struct {
	ReqNbr string `json:"reqNbr"` // 流程实例号
	EvtIst string `json:"evtIst"` // 事件实例号
	ReqSts string `json:"reqSts"` // 请求状态 AUT	等待审批 NTE	终审完毕 BNK，WRF	银行处理中 FIN	完成 OPR	数据接收中
	RtnFlg string `json:"rtnFlg"` // 业务处理结果 reqSts =’FIN’时，rtnFlg才有意义； reqSts =’FIN’并且RTNFLG=’F’，表示支付失败；否则表示支付已被银行受理。返回结果：S	成功	银行支付成功 F	失败	银行支付失败 B	退票	银行支付被退票 R	否决	企业审批否决 D	过期	企业过期不审批 C	撤消	企业撤销 U	银行挂账
	OprSqn string `json:"oprSqn"` // 待处理操作序列
	OprAls string `json:"oprAls"` // 操作别名
	ErrCod string `json:"errCod"` // 错误码
	ErrTxt string `json:"errTxt"` // 错误文本
	MsgTxt string `json:"msgTxt"` // 提示文本
}

// 3.企银支付业务查询BB1PAYQR
type BusiPayQueryRequest struct {
	Bb1Payqrx1 []BusiPayQueryBb1Payqrx1 `json:"bb1payqrx1"`
}

// 3.企银支付业务查询BB1PAYQR
type BusiPayQueryBb1Payqrx1 struct {
	BusCod string `json:"busCod"` // 业务类型 N02030
	YurRef string `json:"yurRef"` // 业务参考号
}

// 3.企银支付业务查询BB1PAYQR
type BusiPayQueryResponse struct {
	Z1 []BusiPayQrz1 `json:"bb1payqrz1"` // 业务类型 N02030
}

// 3.企银支付业务查询BB1PAYQR
type BusiPayQrz1 struct {
	ReqNbr string `json:"reqNbr"`  // 流程实例号
	BusCod string `json:"busCod"`  // 业务编码 A4.业务代码
	BusMod string `json:"busMod"`  // 业务模式
	DbtBbk string `json:"dbtBbk"`  // 转出分行号 A.1 招商分行
	DbtAcc string `json:"dbtAcc"`  // 付方帐号
	DmaNbr string `json:"dmaNbr"`  // 付方记账子单元编号
	DbtNam string `json:"dbtNam"`  // 付方帐户名
	CrtBbk string `json:"crtBbk"`  // 收方分行号 A.1 招商分行
	CrtAcc string `json:"crtAcc"`  // 收方帐号
	CrtNam string `json:"crtNam"`  // 收方名称
	CrtBnk string `json:"crtBnk"`  // 收方行名称
	CrtAdr string `json:"crtAdr"`  // 收方行地址
	CcyNbr string `json:"ccyNbr"`  // 币种 A.3 货币代码表
	TrsAmt string `json:"trsAmt"`  // 交易金额
	EptDat string `json:"eptDat"`  // 期望日
	EptTim string `json:"eptTim"`  // 期望时间
	BnkFlg string `json:"bnkFlg"`  // 系统内外标志 Y 系统内 N系统外
	StlChn string `json:"stl_chn"` // 结算通路 G 普通 Q 快速 R 实时-超网
	NusAge string `json:"nusAge"`  // 用途
	NtfCh1 string `json:"ntfCh1"`  // 通知方式一
	NtfCh2 string `json:"ntfCh2"`  // 通知方式二
	OprDat string `json:"oprDat"`  // 经办日期
	YurRef string `json:"yurRef"`  // 参考号
	BusNar string `json:"busNar"`  // 业务摘要
	ReqSts string `json:"reqSts"`  // 请求状态 AUT等待审批 NTE终审完毕 BNK银行处理中 FIN完成 OPR数据接收中 APW银行人工审批 WRF可疑 ，表示状态未知，需要人工介入处理
	RtnFlg string `json:"rtnFlg"`  // 业务处理结果
	OprSqn string `json:"oprSqn"`  // 待处理操作序列
	OprAls string `json:"oprAls"`  // 操作别名 001：一级操作 002：二级操作 以此类推
	LgnNam string `json:"lgnNam"`  // 用户名
	UsrNam string `json:"usrNam"`  // 用户姓名
	RtnNar string `json:"rtnNar"`  // 失败原因
	AthFlg string `json:"athFlg"`  // 是否有附件信息
	RcvBrd string `json:"rcvBrd"`  // 收方大额行号
	TrsTyp string `json:"trsTyp"`  // 业务种类 100001    普通汇兑 （默认值） 101001    慈善捐款 101002    其他
	TrxSet string `json:"trxSet"`  // 账务套号
	TrxSeq string `json:"trxSeq"`  // 账务流水

}

// 4. 企银支付批量经办BB1PAYBH
type BusiPayBatchRequest struct {
	BX1 []BusiPayBatchBX1 `json:"bb1bmdbhx1"`
	PX1 []BusiPayBatchPX1 `json:"bb1paybhx1"`
}

// 4. 企银支付批量经办BB1PAYBH
type BusiPayBatchBX1 struct {
	BusMod string `json:"busMod"` // 业务模式
	BusCod string `json:"busCod"` // 业务编码 N02030:支付
	BthNbr string `json:"bthNbr"` // 批次编号
	DtlNbr string `json:"dtlNbr"` // 总明细笔数
	CtnFlg string `json:"ctnFlg"` // 总明细笔数 大于1000笔需要做续传 Y N
	CtnSts string `json:"ctnSts"` // 续传状态 当续传标志 =Y时必输； 1 批次开始 2 续传中3 批次结束
}

// 4. 企银支付批量经办BB1PAYBH
type BusiPayBatchPX1 struct {
	DbtAcc string          `json:"dbtAcc"` // 转出帐号
	DmaNbr string          `json:"dmaNbr"` // 记账子单元编号 目前生产上记账子单元编号为10位
	CrtAcc string          `json:"crtAcc"` // 收方帐号
	CrtNam string          `json:"crtNam"` // 收方户名
	CrtBnk string          `json:"crtBnk"` // 收方开户行名称 1、非标准银联卡，必须要填行名称或者联行号；其他情况可空，但优先使用客户传入； 2、若设置了收方限制，则行名联行号都必传； 3、若同时输入行名和联行号，则以联行号为准进行汇出；4、若收方非招行户，请尽量补充完整账户信息以提高支付成功率；
	CrtAdr string          `json:"crtAdr"` // 收方开户行地址 以下任意情况，收方开户行地址可不传： a.收方为招行账户；b.已输入收方开户行名称或收方开户行联行号（若客户未填写开户地，计费将根据行名/行号判断开户地进行同城异地判断，若行名中不含明确的地址信息，可能存在同城异地的误判，请知悉）。
	BrdNbr string          `json:"brdNbr"` // 收方行联行号
	CcyNbr string          `json:"ccyNbr"` // 币种
	TrsAmt decimal.Decimal `json:"trsAmt"` // 交易金额
	EptDat string          `json:"eptDat"` // 期望日
	EptTim string          `json:"eptTim"` // 期望时间
	BnkFlg string          `json:"bnkFlg"` // 系统内标志 收方为招行户：传空或Y； 收方为他行户：传N
	StlChn string          `json:"stlChn"` // 结算通道
	NusAge string          `json:"nusAge"` // 用途
	CrtSqn string          `json:"crtSqn"` // 收方编号
	YurRef string          `json:"yurRef"` // 业务参考号，必须唯一。 业务参考号（YURREF）是每笔交易的唯一编号，是防止重复提交的重要手段，如需重复发送请求，请务必保证同一笔交易的业务参考号不变，否则会存在重复提交风险。
	BusNar string          `json:"busNar"` // 业务摘要
	NtfCh1 string          `json:"ntfCh1"` // 通知方式一（邮箱）
	NtfCh2 string          `json:"ntfCh2"` // 通知方式二（手机号）
	TrsTyp string          `json:"trsTyp"` // 业务种类
	RcvChk string          `json:"rcvChk"` // 行内收方账号户名校验
	DrpFlg string          `json:"drpFlg"` // 直汇普通标志
	TrxAmt string          `json:"trxAmt"` // 相应金额
	CtrNbr string          `json:"ctrNbr"` // 合同号
	InvNbr string          `json:"invNbr"` // 发票号
	RsvAmt string          `json:"rsvAmt"` // 预留金额
	RsvNa1 string          `json:"rsvNa1"` // 预留摘要１
	RsvNa2 string          `json:"rsvNa2"` // 预留摘要２
	RsvNb1 string          `json:"rsvNb1"` // 预留编号１
	RsvNb2 string          `json:"rsvNb2"` // 预留编号２
	RemNbr string          `json:"remNbr"` // 非居民附言编号
	SplC80 string          `json:"splC80"` // 特殊码
}

// 4. 企银支付批量经办BB1PAYBH
type BusiPayBatchResponse struct {
	PZ1 []BusiPayBatchPZ1 `json:"bb1paybhz1"`
}

// 4. 企银支付批量经办BB1PAYBH
type BusiPayBatchPZ1 struct {
	BthNbr string `json:"bthNbr"` // 批次编号
	ReqSts string `json:"reqSts"` // 请求状态 AUT	等待审批 NTE终审完毕 BNK，WRF银行处理中 FIN完成 OPR数据接收中
	RtnFlg string `json:"rtnFlg"` // 业务处理结果 reqSts =’FIN’时，rtnFlg才有意义； reqSts =’FIN’并且RTNFLG=’F’，表示支付失败；否则表示支付已被银行受理。返回结果：S成功银行支付成功 F失败银行支付失败 B退票银行支付被退票 R否决企业审批否决 D过期企业过期不审批 C撤消企业撤销 U银行挂账
	ErrCod string `json:"errCod"` // 错误码
	ErrTxt string `json:"errTxt"` // 错误文本
	MsgTxt string `json:"msgTxt"` // 提示文本
}

// 5. 企银批量支付批次查询BB1QRYBT
type BusiQueryBatchRequest struct {
	QX1 []BusiQueryBatchQX1 `json:"bb1qrybtx1"`
}

// 5. 企银批量支付批次查询BB1QRYBT
type BusiQueryBatchQX1 struct {
	BegDat string `json:"begDat"` // 起始日期 当前日期-1个月
	EndDat string `json:"endDat"` // 结束日期 默认当前日期
	AutStr string `json:"autStr"` // 请求状态 为空时不控制，查全部； OPR接收中 NTE待处理 FIN经办受理完成
	RtnStr string `json:"rtnStr"` // 处理结果 为空时不控制，查全部； F失败 S成功
}

// 5. 企银批量支付批次查询BB1QRYBT
type BusiQueryBatchResponse struct {
	QTZ1 []BusiQueryBatchQTZ1 `json:"bb1qrybtz1"`
}

// 5. 企银批量支付批次查询BB1QRYBT
type BusiQueryBatchQTZ1 struct {
	BthNbr string `json:"bthNbr"` // 批次编号
	TrsDat string `json:"trsDat"` // 经办日期
	TrsTim string `json:"trsTim"` // 经办时间
	BusCod string `json:"busCod"` // 业务类型
	BusMod string `json:"busMod"` // 业务模式
	ReqSts string `json:"reqSts"` // 请求状态 OPR接收中 NTE待处理 FIN完成
	RtnFlg string `json:"rtnFlg"` // 业务处理结果 F失败 S成功
	ErrTxt string `json:"errTxt"` // 失败描述
	DtlAmt string `json:"dtlAmt"` // 批次总金额 总金额超过9999999999999.99则默认只取9999999999999.99
	DtlNum string `json:"dtlNum"` // 批次总笔数
	SucAmt string `json:"sucAmt"` // 提交成功总金额 总金额超过9999999999999.99则默认只取9999999999999.99
	SucNum string `json:"sucNum"` // 提交成功总笔数
}

// 6.企银批量支付明细查询BB1QRYBD
type BusiQueryBatchPayListRequest struct {
	QY1 []BusiQueryBatchQY1 `json:"bb1qrybdy1"`
}

// 6.企银批量支付明细查询BB1QRYBD
type BusiQueryBatchQY1 struct {
	BthNbr string `json:"bthNbr"` // 批次编号
	AutStr string `json:"autStr"` // 请求状态 为空时不控制，查全部； OPR接收中 NTE待处理 FIN经办受理完成
	RtnStr string `json:"rtnStr"` // 处理结果 为空时不控制，查全部； F失败 S成功
	CtnKey string `json:"ctnKey"` // 续传键值 当返回报文有bb1qrybdy1接口返回且ctnKey字段有值时需要续传。请将返回报文的bb1qrybdy1内容放到下次请求报文中继续查询
}

// 6.企银批量支付明细查询BB1QRYBD
type BusiQueryBatchPayListResponse struct {
	QDZ1 []BusiQueryBatchQDZ1 `json:"bb1qrybdz1"` // bb1qrybtz1
}

// 6.企银批量支付明细查询BB1QRYBD
type BusiQueryBatchQDZ1 struct {
	BthNbr         string          `json:"bthNbr"`         // 批次编号
	DbtBbk         string          `json:"dbtBbk"`         // 转出分行号
	DbtAcc         string          `json:"dbtAcc"`         // 转出帐号
	DmaNbr         string          `json:"dmaNbr"`         // 虚拟户编号
	DmaNam         string          `json:"dmaNam"`         // 虚拟户名称
	CrtBbk         string          `json:"crtBbk"`         // 收方分行号
	CrtAcc         string          `json:"crtAcc"`         // 收方帐号
	CrtNam         string          `json:"crtNam"`         // 收方户名
	CrtBnk         string          `json:"crtBnk"`         // 收方行名称
	CrtAdr         string          `json:"crtAdr"`         // 收方行地址
	CcyNbr         string          `json:"ccyNbr"`         // 币种
	TrsAmt         decimal.Decimal `json:"trsAmt"`         // 交易金额
	EptDat         string          `json:"eptDat"`         // 期望日
	EptTim         string          `json:"eptTim"`         // 期望时间
	BnkFlg         string          `json:"bnkFlg"`         // 系统内标志 Y系统内 N系统外
	StlChn         string          `json:"stlChn"`         // 结算通道 G普通 Q快速 R实时-超网
	NusAge         string          `json:"nusAge"`         // 用途
	CrtSqn         string          `json:"crtSqn"`         // 收方编号
	YurRef         string          `json:"yurRef"`         // 参考号
	BusNar         string          `json:"busNar"`         // 业务摘要
	NtfCh1         string          `json:"ntfCh1"`         // 通知方式一
	NtfCh2         string          `json:"ntfCh2"`         // 通知方式二
	CtyCod         string          `json:"ctyCod"`         // 城市码
	TrsTyp         string          `json:"trsTyp"`         // 业务种类 100001普通汇兑 （默认值） 101001慈善捐款 101002其他
	BrdNbr         string          `json:"brdNbr"`         // 收方行号
	PasNbr         string          `json:"pasNbr"`         // 通道号
	RcvChk         string          `json:"rcvChk"`         // 行内收方账号户名校验
	ReqSts         string          `json:"reqSts"`         // 处理状态 AUT等待审批 NTE终审完毕 BNK，WRF银行处理中 FIN完成 OPR数据接收中 APW银行人工审批 WRF 可疑，表示状态未知，需要人工介入处理
	RtnFlg         string          `json:"rtnFlg"`         // 业务处理结果 reqSts =’FIN’时，rtnFlg才有意义； reqSts =’FIN’并且RTNFLG=’F’，表示支付失败；否则表示支付已被银行受理。返回结果：S成功银行支付成功 F失败银行支付失败 B退票银行支付被退票 R否决企业审批否决 D过期企业过期不审批 C撤消企业撤销 U银行挂账
	ErrTxt         string          `json:"errTxt"`         // 失败原因
	MsgTxt         string          `json:"msgTxt"`         // 提示信息
	DrpFlg         string          `json:"drpFlg"`         // 直汇普通标志  A-普通 B-直汇
	CnvNbr         string          `json:"cnvNbr"`         // 网银互联协议号
	NpsTyp         string          `json:"npsTyp"`         // 网银互联业务类型编码
	TrsCat         string          `json:"trsCat"`         // 业务种类编码
	RemNbr         string          `json:"remNbr"`         // 非居民附言编号
	CopNbr         string          `json:"copNbr"`         // 优惠劵编号
	TrxSet         string          `json:"trxSet"`         // 账务套号
	TrxSeq         string          `json:"trxSeq"`         // 账务流水
	ReqNbr         string          `json:"reqNbr"`         // 流程实例号
	TrxSeqBackward string          `json:"trxSeqBackward"` // 账务流水号
}

// 7. 支付退票明细查询BB1PAYQB
type BusiQueryReturnListRequest struct {
	PBY1 []BusiQueryReturnListPBY1 `json:"bb1payqby1"`
}

// 7. 支付退票明细查询BB1PAYQB
type BusiQueryReturnListPBY1 struct {
	BbkNbr string `json:"bbkNbr"` // 分行号
	AccNbr string `json:"accNbr"` // 账号
	BgnDat string `json:"bgnDat"` // 开始日期
	EndDat string `json:"endDat"` // 结束日期
	ReqNbr string `json:"reqNbr"` // 流程实例号
	CtnKey string `json:"ctnKey"` // 续传键值 当返回报文有bb1qrybdy1接口返回且ctnKey字段有值时需要续传。请将返回报文的bb1qrybdy1内容放到下次请求报文中继续查询
	Rsv50z string `json:"rsv50Z"` // 保留字段
}

// 7. 支付退票明细查询BB1PAYQB
type BusiQueryReturnListResponse struct {
	PBZ1 []BusiQueryReturnListPBZ1 `json:"bb1payqbz1"`
	PBY1 []BusiQueryReturnListPBY1 `json:"bb1payqby1"`
}

// 7. 支付退票明细查询BB1PAYQB
type BusiQueryReturnListPBZ1 struct {
	ReqNbr string          `json:"reqNbr"` // 流程实例号
	YurRef string          `json:"yurRef"` // 业务参考号
	BusNbr string          `json:"busNbr"` // 汇款编号
	OutTyp string          `json:"outTyp"` // 汇款方式
	BusTyp string          `json:"busTyp"` // 转账汇款种类
	BusLvl string          `json:"busLvl"` // 汇款优先级
	BusSts string          `json:"busSts"` // 汇款业务状态
	ClrSts string          `json:"clrSts"` // 清算状态
	IsuCnl string          `json:"isuCnl"` // 汇款发起通道
	IsuDat string          `json:"isuDat"` // 发起日期
	TrsBbk string          `json:"trsBbk"` // 处理分行
	TrsBrn string          `json:"trsBrn"` // 处理机构
	CcyNbr string          `json:"ccyNbr"` // 交易货币
	TrsAmt decimal.Decimal `json:"trsAmt"` // 金额
	CcyTyp string          `json:"ccyTyp"` // 钞汇标志
	SndEac string          `json:"sndEac"` // 付方户口号
	SndEan string          `json:"sndEan"` // 付方户名
	SndClt string          `json:"sndClt"` // 付方客户号
	SndBrn string          `json:"sndBrn"` // 付方开户机构
	SndEab string          `json:"sndEab"` // 付方开户行
	SndEaa string          `json:"sndEaa"` // 付方开户地
	RcvEac string          `json:"rcvEac"` // 收方户口号
	RcvEan string          `json:"rcvEan"` // 收方户名
	RcvBbk string          `json:"rcvBbk"` // 收方分行号
	RcvBrd string          `json:"rcvBrd"` // 收方支付行号
	RcvEab string          `json:"rcvEab"` // 收方开户行
	RcvEaa string          `json:"rcvEaa"` // 收方开户地
	RcvFlg string          `json:"rcvFlg"` // 收方同城查询标志
	RcvRef string          `json:"rcvRef"` // 收方查询键值
	NarTxt string          `json:"narTxt"` // 摘要
	TrnBrn string          `json:"trnBrn"` // 转汇机构
	FeeAmt decimal.Decimal `json:"feeAmt"` // 费用总额
	FeeCcy string          `json:"feeCcy"` // 币种
	FeeTyp string          `json:"feeTyp"` // 收费方式
	PsbDat string          `json:"psbDat"` // 凭证日期
	PsbTyp string          `json:"psbTyp"` // 提出凭证种类
	PsbNbr string          `json:"psbNbr"` // 凭证号码
	TrsTyp string          `json:"trsTyp"` // 行内业务种类
	CtyFlg string          `json:"ctyFlg"` // 同城异地标志
	SysFlg string          `json:"sysFlg"` // 系统内外标志
	RcvTyp string          `json:"rcvTyp"` // 收方公私标志
	PrcTrs string          `json:"prcTrs"` // 当前交易流水
	RegTrs string          `json:"regTrs"` // 登记交易流水
	TrsPch string          `json:"trsPch"` // 提出通道
	KpsNbr string          `json:"kpsNbr"` // 当前支付系统键值
	RefNbr string          `json:"refNbr"` // 流水号
	WatRcn string          `json:"watRcn"` // 资金停留原因
	WatTrs string          `json:"watTrs"` // 资金停留流水
	RcdSts string          `json:"rcdSts"` // 记录状态
	UpdDat string          `json:"updDat"` // 更新日期
	RtnCod string          `json:"rtnCod"` // 退票代码
	RtnTxt string          `json:"rtnTxt"` // 退票原因
	SplC80 string          `json:"splC80"` // 特殊码

}
