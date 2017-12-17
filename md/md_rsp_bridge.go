package md

//#include "mdbridge.h"
import "C"
import (
	"gopkg.in/logger.v1"
)

//export __go__on_connected__
func __go__on_connected__() {
	log.Debug("Run go c++ bridge function:__go__on_connected__")
	mdUserSpi.OnFrontConnected()
	//spi.OnFrontConnected()
}

//export __go_on_rsp_user_login__
func __go_on_rsp_user_login__(pRspUserLogin *C.CThostFtdcRspUserLoginField, pRspInfo *C.CThostFtdcRspInfoField,
	nRequestID int, bIsLast bool) {
	log.Debug("Run go c++ bridge function:__go_on_rsp_user_login__")

	loginInfo := toGoRspUserLoginField(pRspUserLogin)
	//spi := getMdUserSpi(pGoMdAPI)
	pGoRspInfo, err := convertRspInfo(pRspInfo)
	if (err != nil) {
		//Todo
	}
	mdUserSpi.OnRspUserLogin(loginInfo, pGoRspInfo, nRequestID, bIsLast)

}

//export __go_on_rsp_sub_market_data__
func __go_on_rsp_sub_market_data__(pDepthMarketData *C.CThostFtdcDepthMarketDataField) {
	log.Debug("Run go c++ bridge function:__go_on_rsp_sub_market_data__")

	pGoDepthMarketData := toGoDepthMarketDataField(pDepthMarketData)
	//spi := getMdUserSpi(pGoMdAPI)
	mdUserSpi.OnRtnDepthMarketData(pGoDepthMarketData)
}

//func getMdUserSpi(pGoMdAPI unsafe.Pointer) GoMdSpi {
//	pMdApi := (*GoMdApi)(pGoMdAPI)
//	spi := (* GoMdSpi)(unsafe.Pointer(pMdApi.Spi))
//	return *spi
//}

//export __go_on_rsp_un_sub_market_data__
func __go_on_rsp_un_sub_market_data__(pSpecificInstrument *C.CThostFtdcSpecificInstrumentField,
	pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	log.Debug("Run go c++ bridge function:__go_on_rsp_un_sub_market_data__")
	//spi := getMdUserSpi(pGoMdAPI)
	pGoSpecificInstrument := toGoSpecificInstrumentField(pSpecificInstrument)
	pGoRspInfo, err := convertRspInfo(pRspInfo)
	if (err != nil) {
		//Todo
	}
	mdUserSpi.OnRspUnSubMarketData(pGoSpecificInstrument, pGoRspInfo, nRequestID, bIsLast)
}
