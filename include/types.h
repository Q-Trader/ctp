/////////////////////////////////////////////////////////////////////////
///@system XTrader
///@company rxforce.io
///@file types.h
///@brief 定义了导出数据类型
///@history
///20170903	李伟		创建该文件
/////////////////////////////////////////////////////////////////////////
#ifdef __cplusplus
extern "C" {
#endif

#if !defined(XTRADER_TYPES_H)
#define XTRADER_TYPES_H
#endif

#if _MSC_VER > 1000
#pragma once
#endif // _MSC_VER > 1000



//请求查询资金账户报文类型
typedef struct CThostFtdcQryTradingAccountField CThostFtdcQryTradingAccountField;
typedef struct CThostFtdcTradingAccountField CThostFtdcTradingAccountField;

typedef struct CThostFtdcRspInfoField CThostFtdcRspInfoField;
typedef struct CThostFtdcRspUserLoginField CThostFtdcRspUserLoginField;
typedef struct CThostFtdcSettlementInfoConfirmField CThostFtdcSettlementInfoConfirmField;
typedef struct CThostFtdcUserLogoutField CThostFtdcUserLogoutField;
typedef struct CThostFtdcSpecificInstrumentField CThostFtdcSpecificInstrumentField;
typedef struct CThostFtdcDepthMarketDataField CThostFtdcDepthMarketDataField;
typedef struct CThostFtdcQrySettlementInfoField CThostFtdcQrySettlementInfoField;
typedef struct CThostFtdcSettlementInfoField CThostFtdcSettlementInfoField;


#ifdef __cplusplus
}
#endif