package tasks

import (
	"context"
	"git/miniTools/data-service/config"
	"git/miniTools/data-service/pkg/model"
	"git/miniTools/data-service/pkg/net"
	"git/miniTools/data-service/utils"
	"github.com/penglq/QLog"
	"strings"
	"time"
)

func AnalystStock(ctx context.Context, request interface{}) (response interface{}, err error) {
	ctx = utils.SetTraceIdContext(ctx, utils.CreateTraceId())
	defer func() {
		if r := recover(); r != nil {
			QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "error", err)
		}
	}()
	date := time.Now().Format("20060102")
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "日期", date)
	err = Stockdata(ctx, date)
	return
}

func Stockdata(ctx context.Context, date string) (err error) {
	method := "Stockdata"
	url := config.GetGlobalConfig().AnalystUrl.StockCompany + "?date=" + date + "&count=20"
	resStockCompany, err := net.GetStock(ctx, url)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "处理股票数据", "url", url, "error", err)
		return
	}
	err = InsertStockCompany(ctx, resStockCompany)

	url = config.GetGlobalConfig().AnalystUrl.StockIndustry + "?date=" + date + "&count=20"
	resStockIndustry, err := net.GetStock(ctx, url)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "处理股票数据", "url", url, "error", err)
		return
	}
	err = InsertStockIndustry(ctx, resStockIndustry)

	return
}
func InsertStockCompany(ctx context.Context, resStock *net.ResStock) (err error) {
	method := "InsertStockCompany"
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "处理股票数据", "data", *resStock)
	for i := 0; i < len(resStock.Results); i++ {
		stockDataModel := new(model.AnalystStocksdata)
		stockDataModel.StockType = config.StockCompany
		switch resStock.Results[i].SratingName {
		case "买入":
			stockDataModel.SratingName = config.SratingNamePurchase
		case "强推", "强烈推荐":
			stockDataModel.SratingName = config.SratingNameRecommend
		}
		stockDataModel.Date = resStock.Results[i].Date
		stockDataModel.Code = resStock.Results[i].Code
		stockDataModel.CodeName = resStock.Results[i].Codename
		stockDataModel.Actualdprice = resStock.Results[i].Actualdprice
		// 后端同一code可能有多条数据(多个不同对报告),只取最新一条。
		// 根据stock_type,date,code,SratingName查询是否数据已经存在
		_, ok, err := model.GetStockByDateCodeSratingNameStockType(stockDataModel.Date, stockDataModel.Code, stockDataModel.SratingName)
		if err != nil {
			QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询数据", "error", err)
			continue
		} else if ok {
			continue
		}

		// 取分析报告
		if len(resStock.Results[i].Attach) > 0 {
			stockDataModel.FileName = resStock.Results[i].Attach[0].Name
			stockDataModel.FileUrl = resStock.Results[i].Attach[0].URL
		}
		// 取一位作者
		for j := 0; j < len(resStock.Results[i].AuthorList); j++ {
			if _, ok := config.AuthorList[resStock.Results[i].AuthorList[j].Auth]; ok {
				stockDataModel.AuthorName = resStock.Results[i].AuthorList[j].Auth
				stockDataModel.AuthorInfo = resStock.Results[i].AuthorList[j].Authprizeinfo
				break
			}
		}
		// 计算潜在涨幅，没有目标价无潜在涨幅;行业的，交通运输和金融就没有日k线 300036.SZ
		if resStock.Results[i].Actualdprice != 0 {
			url := config.GetGlobalConfig().AnalystUrl.StockKLine + "?code=" + resStock.Results[i].Code + "&count=5&period=86400"
			resKLine, err := net.GetKLine(ctx, url)
			if err != nil {
				return err
			}
			if len(resKLine.Results.Datas) > 1 {
				stockDataModel.PotentialGain = int(float64(resStock.Results[i].Actualdprice)-resKLine.Results.Datas[1].ClosePrice) * 10000 / int(resKLine.Results.Datas[1].ClosePrice*100)
				QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "计算涨幅", "目标价格", resStock.Results[i].Actualdprice, "前一交易日收盘价格", resKLine.Results.Datas[1].ClosePrice)
			}
		}
		// 查询基金简称和code
		// 注：查询如果出错，继续往下走，可以后续补数据
		fund, ok, err := model.GetFundByCodeOne(resStock.Results[i].Code)
		if err != nil {
			QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "获取股票windcode", "error", err)
			// return err
		} else if ok {
			funds, ok, err := model.GetFundByWindCode(fund.SInfoWindcode)
			if err != nil {
				QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "获取股票windcode", "error", err)
				// return err
			} else if ok {
				stockDataModel.FundName = funds.FInfoName
				fundCodes := strings.Split(fund.SInfoWindcode, `.`)
				if len(fundCodes) > 0 {
					stockDataModel.FundCode = fundCodes[0]
				}
			} else {
				QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查基金数据", "error", "未查到基金数据", "data", funds.FInfoName)
			}
		} else {
			QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查基金数据", "error", "未查到基金数据", "data", resStock.Results[i].Code)
		}
		_, err = model.InsertData(stockDataModel)
		if err != nil {
			QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "插入数据错误", "error", err)
			// return err
		}
	}
	return
}

func InsertStockIndustry(ctx context.Context, resStock *net.ResStock) (err error) {
	method := "InsertStockIndustry"
	for i := 0; i < len(resStock.Results); i++ {
		stockDataModel := new(model.AnalystStocksdata)
		stockDataModel.StockType = config.StockIndustry
		stockDataModel.Date = resStock.Results[i].Date
		stockDataModel.IndustryCode = resStock.Results[i].IndustryCode
		stockDataModel.IndustryName = resStock.Results[i].Industry
		stockDataModel.Actualdprice = resStock.Results[i].Actualdprice
		switch resStock.Results[i].SratingName {
		case "买入":
			stockDataModel.SratingName = config.SratingNamePurchase
		case "强推", "强烈推荐":
			stockDataModel.SratingName = config.SratingNameRecommend
		}
		// 后端同一code可能有多条数据(多个不同对报告),只取最新一条。
		// 根据stock_type,date,code,SratingName查询是否数据已经存在
		_, ok, err := model.GetStockByDateIndustryCodeSratingNameStockType(stockDataModel.Date, stockDataModel.IndustryCode, stockDataModel.SratingName)
		if err != nil {
			QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询数据", "error", err)
			continue
		} else if ok {
			continue
		}

		// 取分析报告
		if len(resStock.Results[i].Attach) > 0 {
			stockDataModel.FileName = resStock.Results[i].Attach[0].Name
			stockDataModel.FileUrl = resStock.Results[i].Attach[0].URL
		}
		// 取一位作者
		for j := 0; j < len(resStock.Results[i].AuthorList); j++ {
			if _, ok := config.AuthorList[resStock.Results[i].AuthorList[j].Auth]; ok {
				stockDataModel.AuthorName = resStock.Results[i].AuthorList[j].Auth
				stockDataModel.AuthorInfo = resStock.Results[i].AuthorList[j].Authprizeinfo
				break
			}
		}
		// 计算潜在涨幅，没有目标价无潜在涨幅;行业的，交通运输和金融就没有日k线 300036.SZ
		if resStock.Results[i].Actualdprice != 0 {
			url := config.GetGlobalConfig().AnalystUrl.StockKLine + "?code=" + resStock.Results[i].Code + "&limit=5"
			resKLine, err := net.GetKLine(ctx, url)
			if err != nil {
				return err
			}
			if len(resKLine.Results.Datas) > 1 {
				stockDataModel.PotentialGain = int(float64(resStock.Results[i].Actualdprice)-resKLine.Results.Datas[1].ClosePrice) * 10000 / int(resKLine.Results.Datas[1].ClosePrice*100)
				QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "计算涨幅", "目标价格", resStock.Results[i].Actualdprice, "前一交易日收盘价格", resKLine.Results.Datas[1].ClosePrice)
			}
		}
		// 查询基金简称和code
		// 注：查询如果出错，继续往下走，可以后续补数据
		name := resStock.Results[i].Industry
		switch name {
		case "公用事业":
			name = `电力、热力、燃气及水生产和供应业`
		}
		fund, ok, err := model.GetFundByIndustryNameOne(name)
		if err != nil {
			QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "获取行业windcode", "error", err)
			// return err
		} else if ok {
			funds, ok, err := model.GetFundByWindCode(fund.SInfoWindcode)
			if err != nil {
				QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "获取windcode", "error", err)
				// return err
			} else if ok {
				stockDataModel.FundName = funds.FInfoName
				fundCodes := strings.Split(fund.SInfoWindcode, `.`)
				if len(fundCodes) > 0 {
					stockDataModel.FundCode = fundCodes[0]
				}
			} else {
				QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查基金数据", "error", "未查到基金数据", "data", funds.FInfoName)
			}
		} else {
			QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查基金数据", "error", "未查到基金数据", "data", name)
		}

		_, err = model.InsertData(stockDataModel)
		if err != nil {
			QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "插入数据错误", "error", err)
			// return err
		}
	}
	return
}
