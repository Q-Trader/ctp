package ctp

import "gopkg.in/logger.v1"

//RspInfo CTP服务器响应错误码
type RspInfo struct {
	ErrID  int
	ErrMsg string
}

//NewRspInfo 响应信息
func NewRspInfo(errID int, errMsg string) *RspInfo {
	return &RspInfo{errID, errMsg}
}

//UserInfo 用户信息
type UserInfo struct {
	BrokerID   string //经纪公司代码
	InvestorID string //投资者代码
}

//SettlementInfo 投资者结算结果信息
type SettlementInfo struct {
	TradingDay   string //交易日
	SettlementID int    //结算编号
	BrokerID     string //经纪公司代码
	InvestorID   string //投资者代码
	SequenceNo   int    //序号
	Content      string //消息正文[501]
}

//SettlementConfirmInfo 投资者结算结果确认信息
type SettlementConfirmInfo struct {
	BrokerAndInvestor
	ConfirmDate string ///确认日期
	ConfirmTime string ///确认时间
}

//TradingAccount 交易账号
type TradingAccount struct {
	BrokerID string //经纪公司代码
	//投资者帐号
	AccountID string
	//上次质押金额
	PreMortgage float64
	// 上次信用额度
	PreCredit float64
	// 上次存款额
	PreDeposit float64
	// 上次结算准备金
	PreBalance float64
	// 上次占用的保证金
	PreMargin float64
	// 利息基数
	InterestBase float64
	// 利息收入
	Interest float64
	// 入金金额
	Deposit float64
	// 出金金额
	Withdraw float64
	///冻结的保证金
	FrozenMargin float64
	// 冻结的资金
	FrozenCash float64

	///冻结的手续费
	FrozenCommission float64
	///当前保证金总额
	CurrMargin float64
	///资金差额
	CashIn float64
	///手续费
	Commission float64
	///平仓盈亏
	CloseProfit float64
	///持仓盈亏
	PositionProfit float64
	///期货结算准备金
	Balance float64
	///可用资金
	Available float64
	///可取资金
	WithdrawQuota float64
	///基本准备金
	Reserve float64
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///信用额度
	Credit float64
	///质押金额
	Mortgage float64
	///交易所保证金
	ExchangeMargin float64
	///投资者交割保证金
	DeliveryMargin float64
	///交易所交割保证金
	ExchangeDeliveryMargin float64
	///保底期货结算准备金
	ReserveBalance float64
	///币种代码
	CurrencyID string
	///上次货币质入金额
	PreFundMortgageIn float64
	///上次货币质出金额
	PreFundMortgageOut float64
	///货币质入金额
	FundMortgageIn float64
	///货币质出金额
	FundMortgageOut float64
	///货币质押余额
	FundMortgageAvailable float64
	///可质押货币金额
	MortgageableFund float64
	///特殊产品占用保证金
	SpecProductMargin float64
	///特殊产品冻结保证金
	SpecProductFrozenMargin float64
	///特殊产品手续费
	SpecProductCommission float64
	///特殊产品冻结手续费
	SpecProductFrozenCommission float64
	///特殊产品持仓盈亏
	SpecProductPositionProfit float64
	///特殊产品平仓盈亏
	SpecProductCloseProfit float64
	///根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg float64
	///特殊产品交易所保证金
	SpecProductExchangeMargin float64
}

//RspUserLoginInfo 用户登录响应信息
type RspUserLoginInfo struct {
	///交易日
	TradingDay string
	///登录成功时间
	LoginTime string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///交易系统名称
	SystemName string
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///最大报单引用
	MaxOrderRef string
	///上期所时间
	SHFETime string
	///大商所时间
	DCETime string
	///郑商所时间
	CZCETime string
	///中金所时间
	FFEXTime string
}

//ReqUserLogoutInfo 用户登出请求
type ReqUserLogoutInfo struct {
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
}

//SpecificInstrument 指定的合约
type SpecificInstrument struct {
	///合约代码
	InstrumentID string
}

//Instrument 合约
type Instrument struct {
	///合约代码
	InstrumentID string
	///交易所代码
	ExchangeID string
	///合约名称
	InstrumentName string
	///合约在交易所的代码
	TThostFtdcExchangeInstIDType string
	///产品代码
	ProductID string
	///产品类型
	ProductClass string
	///交割年份
	DeliveryYear int
	///交割月
	DeliveryMonth int
	///市价单最大下单量
	MaxMarketOrderVolume int
	///市价单最小下单量
	MinMarketOrderVolume int
	///限价单最大下单量
	MaxLimitOrderVolume int
	///限价单最小下单量
	MinLimitOrderVolume int
	///合约数量乘数
	VolumeMultiple int
	///最小变动价位
	PriceTick float64
	///创建日
	CreateDate string
	///上市日
	OpenDate string
	///到期日
	ExpireDate string
	///开始交割日
	StartDelivDate string
	///结束交割日
	EndDelivDate string

	/**
	/////////////////////////////////////////////////////////////////////////
	///InstLifePhase是一个合约生命周期状态类型
	/////////////////////////////////////////////////////////////////////////
	///未上市
	#define THOST_FTDC_IP_NotStart '0'
	///上市
	#define THOST_FTDC_IP_Started '1'
	///停牌
	#define THOST_FTDC_IP_Pause '2'
	///到期
	#define THOST_FTDC_IP_Expired '3'
	*/
	///合约生命周期状态
	InstLifePhase string
	///当前是否交易
	IsTrading bool

	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcPositionTypeType是一个持仓类型类型
	/////////////////////////////////////////////////////////////////////////
	///净持仓
	#define THOST_FTDC_PT_Net '1'
	///综合持仓
	#define THOST_FTDC_PT_Gross '2'
	*/
	///持仓类型
	PositionType string

	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcPositionDateTypeType是一个持仓日期类型类型
	/////////////////////////////////////////////////////////////////////////
	///使用历史持仓
	#define THOST_FTDC_PDT_UseHistory '1'
	///不使用历史持仓
	#define THOST_FTDC_PDT_NoUseHistory '2'
	*/
	///持仓日期类型
	PositionDateType string
	///多头保证金率
	LongMarginRatio float64
	///空头保证金率
	ShortMarginRatio float64
}

//InvestorPosition 投资者持仓
type InvestorPosition struct {
	BrokerAndInvestor
	///合约代码
	InstrumentID string
	///持仓多空方向
	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcPosiDirectionType是一个持仓多空方向类型
	/////////////////////////////////////////////////////////////////////////
	///净
	#define THOST_FTDC_PD_Net '1'
	///多头
	#define THOST_FTDC_PD_Long '2'
	///空头
	#define THOST_FTDC_PD_Short '3'
	*/
	PosiDirection string
	///投机套保标志
	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcHedgeFlagType是一个投机套保标志类型
	/////////////////////////////////////////////////////////////////////////
	///投机
	#define THOST_FTDC_HF_Speculation '1'
	///套保
	#define THOST_FTDC_HF_Hedge '3'
	*/
	HedgeFlag string
	///持仓日期类型
	/////////////////////////////////////////////////////////////////////////
	///TFtdcPositionDateType是一个持仓日期类型
	/////////////////////////////////////////////////////////////////////////
	///今日持仓
	//#define THOST_FTDC_PSD_Today '1'
	///历史持仓
	//#define THOST_FTDC_PSD_History '2'

	PositionDate string
	///上日持仓
	YdPosition int
	///今日持仓
	Position int
	///多头冻结
	LongFrozen int
	///空头冻结
	ShortFrozen int
	///开仓冻结金额
	LongFrozenAmount float64
	///开仓冻结金额
	ShortFrozenAmount float64
	///开仓量
	OpenVolume int
	///平仓量
	CloseVolume int
	///开仓金额
	OpenAmount float64
	///平仓金额
	CloseAmount float64
	///持仓成本
	PositionCost float64
	///上次占用的保证金
	PreMargin float64
	///占用的保证金
	UseMargin float64
	///冻结的保证金
	FrozenMargin float64
	///冻结的资金
	FrozenCash float64
	///冻结的手续费
	FrozenCommission float64
	///资金差额
	CashIn float64
	///手续费
	Commission float64
	///平仓盈亏
	CloseProfit float64
	///持仓盈亏
	PositionProfit float64
	///上次结算价
	PreSettlementPrice float64
	///本次结算价
	SettlementPrice float64
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///开仓成本
	OpenCost float64
	///交易所保证金
	ExchangeMargin float64
	///组合成交形成的持仓
	CombPosition int
	///组合多头冻结
	CombLongFrozen int
	///组合空头冻结
	CombShortFrozen int
	///逐日盯市平仓盈亏
	CloseProfitByDate float64
	///逐笔对冲平仓盈亏
	CloseProfitByTrade float64
	///今日持仓
	TodayPosition int
	///保证金率
	MarginRateByMoney float64
	///保证金率(按手数)
	MarginRateByVolume float64
}

//Order 报单
type Order struct {
	BrokerAndInvestor
	///合约代码
	InstrumentID string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///报单价格条件类型
	/////////////////////////////////////////////////////////////////////////
	///TFtdcOrderPriceTypeType是一个报单价格条件类型
	/////////////////////////////////////////////////////////////////////////
	///任意价
	/*#define THOST_FTDC_OPT_AnyPrice '1'
	///限价
	#define THOST_FTDC_OPT_LimitPrice '2'
	///最优价
	#define THOST_FTDC_OPT_BestPrice '3'
	///最新价
	#define THOST_FTDC_OPT_LastPrice '4'
	///最新价浮动上浮1个ticks
	#define THOST_FTDC_OPT_LastPricePlusOneTicks '5'
	///最新价浮动上浮2个ticks
	#define THOST_FTDC_OPT_LastPricePlusTwoTicks '6'
	///最新价浮动上浮3个ticks
	#define THOST_FTDC_OPT_LastPricePlusThreeTicks '7'
	///卖一价
	#define THOST_FTDC_OPT_AskPrice1 '8'
	///卖一价浮动上浮1个ticks
	#define THOST_FTDC_OPT_AskPrice1PlusOneTicks '9'
	///卖一价浮动上浮2个ticks
	#define THOST_FTDC_OPT_AskPrice1PlusTwoTicks 'A'
	///卖一价浮动上浮3个ticks
	#define THOST_FTDC_OPT_AskPrice1PlusThreeTicks 'B'
	///买一价
	#define THOST_FTDC_OPT_BidPrice1 'C'
	///买一价浮动上浮1个ticks
	#define THOST_FTDC_OPT_BidPrice1PlusOneTicks 'D'
	///买一价浮动上浮2个ticks
	#define THOST_FTDC_OPT_BidPrice1PlusTwoTicks 'E'
	///买一价浮动上浮3个ticks
	#define THOST_FTDC_OPT_BidPrice1PlusThreeTicks 'F'*/
	OrderPriceType string
	///买卖方向
	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcDirectionType是一个买卖方向类型
	/////////////////////////////////////////////////////////////////////////
	///买
	#define THOST_FTDC_D_Buy '0'
	///卖
	#define THOST_FTDC_D_Sell '1'

	*/
	Direction string
	///组合开平标志
	CombOffsetFlag string
	///组合投机套保标志
	CombHedgeFlag string
	///价格
	LimitPrice float64
	///数量
	VolumeTotalOriginal int
	///有效期类型

	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcTimeConditionType是一个有效期类型类型
	/////////////////////////////////////////////////////////////////////////
	///立即完成，否则撤销
	#define THOST_FTDC_TC_IOC '1'
	///本节有效
	#define THOST_FTDC_TC_GFS '2'
	///当日有效
	#define THOST_FTDC_TC_GFD '3'
	///指定日期前有效
	#define THOST_FTDC_TC_GTD '4'
	///撤销前有效
	#define THOST_FTDC_TC_GTC '5'
	///集合竞价有效
	#define THOST_FTDC_TC_GFA '6'
	*/
	TimeCondition string
	///GTD日期
	GTDDate string
	///成交量类型

	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcVolumeConditionType是一个成交量类型类型
	/////////////////////////////////////////////////////////////////////////
	///任何数量
	#define THOST_FTDC_VC_AV '1'
	///最小数量
	#define THOST_FTDC_VC_MV '2'
	///全部数量
	#define THOST_FTDC_VC_CV '3'

	*/
	VolumeCondition string
	///最小成交量
	MinVolume int
	///触发条件
	/*
	/////////////////////////////////////////////////////////////////////////
	///TFtdcContingentConditionType是一个触发条件类型
	/////////////////////////////////////////////////////////////////////////
	///立即
	#define THOST_FTDC_CC_Immediately '1'
	///止损
	#define THOST_FTDC_CC_Touch '2'
	///止赢
	#define THOST_FTDC_CC_TouchProfit '3'
	///预埋单
	#define THOST_FTDC_CC_ParkedOrder '4'
	///条件价大于最新价
	#define THOST_FTDC_CC_StopPriceGreaterThanLastPrice '5'
	///条件价大于等于最新价
	#define THOST_FTDC_CC_StopPriceGreaterEqualLastPrice '6'
	///条件价小于最新价
	#define THOST_FTDC_CC_StopPriceLesserThanLastPrice '7'
	///条件价小于等于最新价
	#define THOST_FTDC_CC_StopPriceLesserEqualLastPrice '8'
	///条件价大于卖一价
	#define THOST_FTDC_CC_StopPriceGreaterThanAskPrice '9'
	///条件价大于等于卖一价
	#define THOST_FTDC_CC_StopPriceGreaterEqualAskPrice 'A'
	///条件价小于卖一价
	#define THOST_FTDC_CC_StopPriceLesserThanAskPrice 'B'
	///条件价小于等于卖一价
	#define THOST_FTDC_CC_StopPriceLesserEqualAskPrice 'C'
	///条件价大于买一价
	#define THOST_FTDC_CC_StopPriceGreaterThanBidPrice 'D'
	///条件价大于等于买一价
	#define THOST_FTDC_CC_StopPriceGreaterEqualBidPrice 'E'
	///条件价小于买一价
	#define THOST_FTDC_CC_StopPriceLesserThanBidPrice 'F'
	///条件价小于等于买一价
	#define THOST_FTDC_CC_StopPriceLesserEqualBidPrice 'a'
	*/
	ContingentCondition string
	///止损价
	StopPrice float64
	///强平原因
	ForceCloseReason string
	///自动挂起标志
	IsAutoSuspend int
	///业务单元
	BusinessUnit string
	///请求编号
	RequestID int
	///本地报单编号
	OrderLocalID string
	///交易所代码
	ExchangeID string
	///会员代码
	ParticipantID string
	///客户代码
	ClientID string
	///合约在交易所的代码
	ExchangeInstID string
	///交易所交易员代码
	TraderID string
	///安装编号
	InstallID int
	///报单提交状态
	OrderSubmitStatus string
	///报单提示序号
	NotifySequence int
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///报单编号
	OrderSysID string
	///报单来源
	OrderSource string
	///报单状态
	OrderStatus string
	///报单类型
	OrderType string
	///今成交数量
	VolumeTraded int
	///剩余数量
	VolumeTotal int
	///报单日期
	InsertDate string
	///委托时间
	InsertTime string
	///激活时间
	ActiveTime string
	///挂起时间
	SuspendTime string
	///最后修改时间
	UpdateTime string
	///撤销时间
	CancelTime string
	///最后修改交易所交易员代码
	ActiveTraderID string
	///结算会员编号
	ClearingPartID string
	///序号
	SequenceNo int
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///用户端产品信息
	UserProductInfo string
	///状态信息
	StatusMsg string
	///用户强评标志
	UserForceClose bool
	///操作用户代码
	ActiveUserID string
	///经纪公司报单编号
	BrokerOrderSeq int
}

//TradeInfo 成交
type TradeInfo struct {
	BrokerAndInvestor
	///合约代码
	InstrumentID string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///交易所代码
	ExchangeID string
	///成交编号
	TradeID string
	///买卖方向
	Direction string
	///报单编号
	OrderSysID string
	///会员代码
	ParticipantID string
	///客户代码
	ClientID string
	///交易角色
	TradingRole string
	///合约在交易所的代码
	ExchangeInstID string
	///开平标志
	OffsetFlag string
	///投机套保标志
	HedgeFlag string
	///价格
	Price float64
	///数量
	Volume int
	///成交时期
	TradeDate string
	///成交时间
	TradeTime string
	///成交类型
	TradeType string
	///成交价来源
	PriceSource string
	///交易所交易员代码
	TraderID string
	///本地报单编号
	OrderLocalID string
	///结算会员编号
	ClearingPartID string
	///业务单元
	BusinessUnit string
	///序号
	SequenceNo int
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///经纪公司报单编号
	BrokerOrderSeq int
}

//DepthMarketData 深度行情
type DepthMarketData struct {
	///交易日
	TradingDay string
	///合约代码
	InstrumentID string
	///交易所代码
	ExchangeID string
	///合约在交易所的代码
	ExchangeInstID string
	///最新价
	LastPrice float64
	///上次结算价
	PreSettlementPrice float64
	///昨收盘
	PreClosePrice float64
	///昨持仓量
	PreOpenInterest float64
	///今开盘
	OpenPrice float64
	///最高价
	HighestPrice float64
	///最低价
	LowestPrice float64
	///数量
	Volume int
	///成交金额
	Turnover float64
	///持仓量
	OpenInterest float64
	///今收盘
	ClosePrice float64
	///本次结算价
	SettlementPrice float64
	///涨停板价
	UpperLimitPrice float64
	///跌停板价
	LowerLimitPrice float64
	///昨虚实度
	PreDelta float64
	///今虚实度
	CurrDelta float64
	///最后修改时间
	UpdateTime string
	///最后修改毫秒
	UpdateMillisec int
	///申买价一
	BidPrice1 float64
	///申买量一
	BidVolume1 float64
	///申卖价一
	AskPrice1 float64
	///申卖量一
	AskVolume1 float64
	///申买价二
	BidPrice2 float64
	///申买量二
	BidVolume2 float64
	///申卖价二
	AskPrice2 float64
	///申卖量二
	AskVolume2 float64
	///申买价三
	BidPrice3 float64
	///申买量三
	BidVolume3 float64
	///申卖价三
	AskPrice3 float64
	///申卖量三
	AskVolume3 float64
	///申买价四
	BidPrice4 float64
	///申买量四
	BidVolume4 float64
	///申卖价四
	AskPrice4 float64
	///申卖量四
	AskVolume4 float64
	///申买价五
	BidPrice5 float64
	///申买量五
	BidVolume5 float64
	///申卖价五
	AskPrice5 float64
	///申卖量五
	AskVolume5 float64
	///当日均价
	AveragePrice float64
}

//BrokerAndInvestor 经济公司代码、投资者代码
type BrokerAndInvestor struct {
	BrokerID   string ///经纪公司代码
	InvestorID string //投资者代码
}

//SessionInfo 会话信息
type SessionInfo struct {
	BrokerAndInvestor
	RequestID  int    ///请求编号
	FrontID    int    ///前置编号
	SessionID  int    ///会话编号
	ExchangeID string ///交易所代码
	UserID     string ///用户代码
}

//InputOrderAction 报单引用
type InputOrderAction struct {
	BrokerAndInvestor
	OrderActionRef int     //报单操作引用
	OrderRef       string  //报单引用
	OrderSysID     string  //报单编号
	ActionFlag     string  ///操作标志
	LimitPrice     float64 //价格
	VolumeChange   int     //数量变化
	InstrumentID   string  //合约代码
}

//InputOrder 报单
type InputOrder struct {
	BrokerAndInvestor
	///合约代码
	InstrumentID string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///报单价格条件
	OrderPriceType string
	///买卖方向
	Direction string
	///组合开平标志
	CombOffsetFlag string
	///组合投机套保标志
	CombHedgeFlag string
	///价格
	LimitPrice float64
	///数量
	VolumeTotalOriginal int
	///有效期类型
	TimeCondition string
	///GTD日期
	GTDDate string

	/*	///任何数量
		#define THOST_FTDC_VC_AV '1'
		///最小数量
		#define THOST_FTDC_VC_MV '2'
		///全部数量
		#define THOST_FTDC_VC_CV '3'*/
	VolumeCondition string ///成交量类型
	///最小成交量
	MinVolume int
	/*
	/////////////////////////////////////////////////////////////////////////
	///触发条件类型
	/////////////////////////////////////////////////////////////////////////
	///立即
	#define THOST_FTDC_CC_Immediately '1'
	///止损
	#define THOST_FTDC_CC_Touch '2'
	///止赢
	#define THOST_FTDC_CC_TouchProfit '3'
	///预埋单
	#define THOST_FTDC_CC_ParkedOrder '4'
	///条件价大于最新价
	#define THOST_FTDC_CC_StopPriceGreaterThanLastPrice '5'
	///条件价大于等于最新价
	#define THOST_FTDC_CC_StopPriceGreaterEqualLastPrice '6'
	///条件价小于最新价
	#define THOST_FTDC_CC_StopPriceLesserThanLastPrice '7'
	///条件价小于等于最新价
	#define THOST_FTDC_CC_StopPriceLesserEqualLastPrice '8'
	///条件价大于卖一价
	#define THOST_FTDC_CC_StopPriceGreaterThanAskPrice '9'
	///条件价大于等于卖一价
	#define THOST_FTDC_CC_StopPriceGreaterEqualAskPrice 'A'
	///条件价小于卖一价
	#define THOST_FTDC_CC_StopPriceLesserThanAskPrice 'B'
	///条件价小于等于卖一价
	#define THOST_FTDC_CC_StopPriceLesserEqualAskPrice 'C'
	///条件价大于买一价
	#define THOST_FTDC_CC_StopPriceGreaterThanBidPrice 'D'
	///条件价大于等于买一价
	#define THOST_FTDC_CC_StopPriceGreaterEqualBidPrice 'E'
	///条件价小于买一价
	#define THOST_FTDC_CC_StopPriceLesserThanBidPrice 'F'
	///条件价小于等于买一价
	#define THOST_FTDC_CC_StopPriceLesserEqualBidPrice 'a'
	*/
	///触发条件
	ContingentCondition string
	///止损价
	StopPrice float64

	/*

	/////////////////////////////////////////////////////////////////////////
	///TFtdcForceCloseReasonType是一个强平原因类型
	/////////////////////////////////////////////////////////////////////////
	///非强平
	#define THOST_FTDC_FCC_NotForceClose '0'
	///资金不足
	#define THOST_FTDC_FCC_LackDeposit '1'
	///客户超仓
	#define THOST_FTDC_FCC_ClientOverPositionLimit '2'
	///会员超仓
	#define THOST_FTDC_FCC_MemberOverPositionLimit '3'
	///持仓非整数倍
	#define THOST_FTDC_FCC_NotMultiple '4'
	///违规
	#define THOST_FTDC_FCC_Violation '5'
	///其它
	#define THOST_FTDC_FCC_Other '6'
	*/

	///强平原因
	ForceCloseReason string
	///自动挂起标志
	IsAutoSuspend string
	///业务单元
	BusinessUnit int
	///请求编号
	RequestID int
	///用户强评标志
	UserForceClose int
}

//IsErrorRspInfo 判断是否为错误响应
func IsErrorRspInfo(pRspInfo *RspInfo) (bResult bool) {
	bResult = false
	log.Debugf("pRspInfo is Nil ?: %t", pRspInfo == nil)
	if pRspInfo != nil {
		//if rspInfo, err := ConvertRspInfo(pRspInfo) err == nil {
		// 如果ErrorID != 0, 说明收到了错误的响应
		bResult = (pRspInfo != nil) && (pRspInfo.ErrID != 0)
		if bResult {
			log.Infof("收到错误的响应：ErrorID= %d, ErrorMsg = %s", pRspInfo.ErrID, pRspInfo.ErrMsg)
			return
		}
		log.Debugf("ErrorID = %d, ErrorMsg = %s", pRspInfo.ErrID, pRspInfo.ErrMsg)
	}

	return

}
