package md

// #include "mdbridge.h"
import "C"
import (
	"fmt"

	"gopkg.in/iconv.v1"

	. "qtrx.io/ctp/types"
)

func toGoRspUserLoginField(pCLoginFiled *C.CThostFtdcRspUserLoginField) (pLoginField *RspUserLoginInfo) {
	if pCLoginFiled == nil {
		return
	}
	pLoginField = &RspUserLoginInfo{}
	pLoginField.TradingDay = C.GoString(&pCLoginFiled.TradingDay[0])
	pLoginField.LoginTime = C.GoString(&pCLoginFiled.LoginTime[0])
	pLoginField.BrokerID = C.GoString(&pCLoginFiled.BrokerID[0])
	pLoginField.UserID = C.GoString(&pCLoginFiled.UserID[0])
	pLoginField.SystemName = C.GoString(&pCLoginFiled.SystemName[0])
	pLoginField.FrontID = int(pCLoginFiled.FrontID)
	pLoginField.SessionID = int(pCLoginFiled.SessionID)
	pLoginField.MaxOrderRef = C.GoString(&pCLoginFiled.MaxOrderRef[0])
	pLoginField.SHFETime = C.GoString(&pCLoginFiled.SHFETime[0])
	pLoginField.DCETime = C.GoString(&pCLoginFiled.DCETime[0])
	pLoginField.CZCETime = C.GoString(&pCLoginFiled.CZCETime[0])
	pLoginField.FFEXTime = C.GoString(&pCLoginFiled.FFEXTime[0])
	return
}

func toGoUserLogoutField(pUserLogout *C.CThostFtdcUserLogoutField) (pGoUserLogout *ReqUserLogoutInfo) {
	if pUserLogout == nil {
		return
	}
	pGoUserLogout = &ReqUserLogoutInfo{
		BrokerID: C.GoString(&pUserLogout.BrokerID[0]),
		UserID:   C.GoString(&pUserLogout.UserID[0])}
	return
}

func toGoSpecificInstrumentField(pSpecificInstrument *C.CThostFtdcSpecificInstrumentField) (pGoSpecificInstrument *SpecificInstrument) {
	if pSpecificInstrument == nil {
		return
	}
	pGoSpecificInstrument = &SpecificInstrument{
		C.GoString(&pSpecificInstrument.InstrumentID[0])}
	return
}

func toGoDepthMarketDataField(pDepthMarketData *C.CThostFtdcDepthMarketDataField) (pGoDepthMarketData *DepthMarketData) {
	if pDepthMarketData == nil {
		return nil
	}
	pGoDepthMarketData = &DepthMarketData{
		TradingDay:         C.GoString(&pDepthMarketData.TradingDay[0]),
		InstrumentID:       C.GoString(&pDepthMarketData.InstrumentID[0]),
		ExchangeID:         C.GoString(&pDepthMarketData.ExchangeID[0]),
		ExchangeInstID:     C.GoString(&pDepthMarketData.ExchangeInstID[0]),
		LastPrice:          float64(pDepthMarketData.LastPrice),
		PreSettlementPrice: float64(pDepthMarketData.PreSettlementPrice),
		PreClosePrice:      float64(pDepthMarketData.PreClosePrice),
		PreOpenInterest:    float64(pDepthMarketData.PreOpenInterest),
		OpenPrice:          float64(pDepthMarketData.OpenPrice),
		HighestPrice:       float64(pDepthMarketData.HighestPrice),
		LowestPrice:        float64(pDepthMarketData.LowestPrice),
		Volume:             int(pDepthMarketData.Volume),
		Turnover:           float64(pDepthMarketData.Turnover),
		OpenInterest:       float64(pDepthMarketData.OpenInterest),
		SettlementPrice:    float64(pDepthMarketData.SettlementPrice),
		UpperLimitPrice:    float64(pDepthMarketData.UpperLimitPrice),
		LowerLimitPrice:    float64(pDepthMarketData.LowerLimitPrice),
		PreDelta:           float64(pDepthMarketData.PreDelta),
		CurrDelta:          float64(pDepthMarketData.CurrDelta),
		UpdateTime:         C.GoString(&pDepthMarketData.UpdateTime[0]),
		UpdateMillisec:     int(pDepthMarketData.UpdateMillisec),
		BidPrice1:          float64(pDepthMarketData.BidPrice1),
		BidVolume1:         float64(pDepthMarketData.BidVolume1),
		AskPrice1:          float64(pDepthMarketData.AskPrice1),
		AskVolume1:         float64(pDepthMarketData.AskVolume1),
		BidPrice2:          float64(pDepthMarketData.BidPrice2),
		BidVolume2:         float64(pDepthMarketData.BidVolume2),
		AskPrice2:          float64(pDepthMarketData.AskPrice2),
		AskVolume2:         float64(pDepthMarketData.AskVolume2),
		BidVolume3:         float64(pDepthMarketData.BidVolume3),
		AskPrice3:          float64(pDepthMarketData.AskPrice3),
		AskVolume3:         float64(pDepthMarketData.AskVolume3),
		BidPrice4:          float64(pDepthMarketData.BidPrice4),
		BidVolume4:         float64(pDepthMarketData.BidVolume4),
		AskPrice4:          float64(pDepthMarketData.AskPrice4),
		AskVolume4:         float64(pDepthMarketData.AskVolume4),
		BidPrice5:          float64(pDepthMarketData.BidPrice5),
		BidVolume5:         float64(pDepthMarketData.BidVolume5),
		AskPrice5:          float64(pDepthMarketData.AskPrice5),
		AskVolume5:         float64(pDepthMarketData.AskVolume5),
		AveragePrice:       float64(pDepthMarketData.AveragePrice)}

	return
}

//convert c RspInfo to go struct
func convertRspInfo(pRspInfo *C.CThostFtdcRspInfoField) (rspInfo *RspInfo, err error) {
	cd, err := iconv.Open("utf-8", "gbk") // convert gbk to utf-8
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()
	errMsg := cd.ConvString(C.GoString(&pRspInfo.ErrorMsg[0]))
	rspInfo = NewRspInfo(int(pRspInfo.ErrorID), errMsg)
	return
}
