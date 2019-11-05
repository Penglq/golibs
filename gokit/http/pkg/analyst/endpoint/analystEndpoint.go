package endpoint

import (
	"context"
	"github.com/penglq/QLog"
	"time"
)

func AnalystFirstEndpoint(ctx context.Context, req *interfaces.AnalystFirstRequest) (response interface{}, err error) {
	method := "AnalystFirstEndpoint"
	res := interfaces.AnalystFirstResponse{}
	_, ok, err := model.GetUserByYrdUid(req.YrdUid)
	if err != nil {
		QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询数据", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	} else if ok {
		res.First = "no"
		return new(utils.OutResponse).ResponseSuccService(res), nil
	}
	res.First = "yes"
	userModel := model.AnalystUser{
		YrdUid:     req.YrdUid,
		PassportId: req.PassportId,
		Mobile:     req.Mobile,
		Name:       req.Name,
	}
	_, err = model.InsertData(&userModel)
	if err != nil {
		err = nil
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "插入失败", "data", userModel)
	}
	return new(utils.OutResponse).ResponseSuccService(res), nil
}

// 强推/买入页
func AnalystAnalyzeEndpoint(ctx context.Context, req *interfaces.AnalystAnalyzeRequest) (response interface{}, err error) {
	method := "AnalystAnalyzeEndpoint"
	t := time.Now()
	date, err := checkLastDate(t, t, req.TabType)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询股票列表", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	}
	stock, err := model.GetStockByDateYrdUid(date, req.YrdUid, req.TabType, req.Page, req.PageSize)
	if err != nil {
		return
	}
	res := &interfaces.AnalystAnalyzeResponse{Date: date}
	res.Data = make([]interfaces.Analyze, len(stock))
	handleAnalyze(ctx, res, stock)
	return new(utils.OutResponse).ResponseSuccService(res), nil
}

func AnalystHistoryEndpoint(ctx context.Context, req *interfaces.AnalystHistoryRequest) (response interface{}, err error) {
	method := "AnalystHistoryEndpoint"
	// t, err := time.Parse("20060102", req.Date)
	// if err != nil {
	// 	QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询股票列表", "error", err)
	// 	return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	// }
	// 修改，如果请求日期没有推荐的基金则往前找到最近有的那天
	// date, err := checkLastDate(t, t, req.TabType)
	// if err != nil {
	// 	QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询股票列表", "error", err)
	// 	return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	// }
	date := req.Date
	stock, err := model.GetStockByDateYrdUid(date, req.YrdUid, req.TabType, req.Page, req.PageSize)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询股票列表", "error", err)
		return
	}
	res := &interfaces.AnalystHistoryResponse{Date: date}
	res.Data = make([]interfaces.Analyze, len(stock))
	handleAnalyzeHistory(ctx, res, stock)
	return new(utils.OutResponse).ResponseSuccService(res), nil
}

func AnalystCollectEndpoint(ctx context.Context, req *interfaces.AnalystCollectRequest) (response interface{}, err error) {
	method := "AnalystCollectEndpoint"
	// 是否有此股票数据id
	_, ok, err := model.GetUserById(req.StocksId)
	if err != nil {
		QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询数据", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	} else if !ok {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询数据", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamStock, value.TypeNotExist), nil
	}

	// 是否有过收藏记录
	collect, ok, err := model.GetCollectByYrdUidStockId(req.YrdUid, req.StocksId)
	if err != nil {
		QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询数据", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	} else if ok {
		if collect.Collection == req.Collection && collect.Collection == config.CollectYes {
			return new(utils.OutResponse).ResponseSuccService(struct{}{}), nil
		} else if collect.Collection == req.Collection && collect.Collection == config.CollectNo {
			return new(utils.OutResponse).ResponseSuccService(struct{}{}), nil
		} else {
			data := new(model.AnalystCollect)
			data.Collection = req.Collection
			affected, err := model.UpdateCollectByYrdUidStockId(req.YrdUid, req.StocksId, data)
			if err != nil {
				QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "更新数据", "error", err)
				return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
			}
			if affected < 1 {
				QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "更新数据影响行数小于1", "affected", affected)
				return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamUpdate, value.TypeFail), nil
			}
			return new(utils.OutResponse).ResponseSuccService(struct{}{}), nil
		}
	}
	data := model.AnalystCollect{
		YrdUid:     req.YrdUid,
		StockId:    req.StocksId,
		Collection: req.Collection,
	}
	_, err = model.InsertData(&data)
	if err != nil {
		QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "插入数据", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	}
	return new(utils.OutResponse).ResponseSuccService(struct{}{}), nil
}

func AnalystCollectHistoryEndpoint(ctx context.Context, req *interfaces.AnalystCollectHistoryRequest) (response interface{}, err error) {
	method := "AnalystCollectHistoryEndpoint"
	total, err := model.GetCollectCount(req.TabType, req.YrdUid)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询保单记录总数", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	}
	resp := interfaces.AnalystCollectHistoryResponse{}
	resp.TotalPage = total/req.PageSize + 1
	resp.Data = []interfaces.CollectHistory{}
	if total == 0 {
		return new(utils.OutResponse).ResponseSuccService(resp), nil
	}
	collect, err := model.GetCollectStockByDateYrdUid(req.TabType, req.YrdUid, req.Page, req.PageSize)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询股票列表", "error", err)
		return
	}
	res := make([]interfaces.CollectHistory, len(collect))
	for i := 0; i < len(collect); i++ {
		res[i].Date = collect[i].AnalystStocksdata.Date
		res[i].AuthorName = collect[i].AnalystStocksdata.AuthorName
		res[i].Collection = collect[i].AnalystCollect.Collection
		res[i].FundName = collect[i].AnalystStocksdata.FundName
		res[i].FundURL = config.GetGlobalConfig().AnalystUrl.FundAppBaseURL + collect[i].AnalystStocksdata.FundCode
		res[i].Name = collect[i].AnalystStocksdata.CodeName
		res[i].PotentialGain = collect[i].AnalystStocksdata.PotentialGain
		res[i].RecommendFileName = collect[i].AnalystStocksdata.FileName
		res[i].RecommendFileURL = collect[i].AnalystStocksdata.FileUrl
		res[i].StocksID = collect[i].AnalystStocksdata.Id
		res[i].StocksType = collect[i].AnalystStocksdata.StockType
	}
	resp.Data = res
	return new(utils.OutResponse).ResponseSuccService(resp), nil
}

// 找到最近一天有数据的日期
func checkLastDate(originT, t time.Time, tabType string) (d string, err error) {
	tStr := t.Format("20060102")
	if originT.Sub(t).Seconds()/(3600*24) > 10 {
		return tStr, nil
	}
	stock, err := model.GetStockByDate(tStr, tabType)
	if err != nil {
		return
	}
	if len(stock) > 0 {
		return tStr, err
	}
	d, err = checkLastDate(originT, t.AddDate(0, 0, -1), tabType)
	return
}

// 处理股票数据
func handleAnalyze(ctx context.Context, res *interfaces.AnalystAnalyzeResponse, stocks []model.Collect) {
	for i := 0; i < len(stocks); i++ {
		res.Data[i].StocksID = stocks[i].AnalystStocksdata.Id
		res.Data[i].StocksType = stocks[i].AnalystStocksdata.StockType
		if stocks[i].AnalystStocksdata.FundName != "" {
			res.Data[i].FundName = stocks[i].AnalystStocksdata.FundName
			res.Data[i].FundURL = config.GetGlobalConfig().AnalystUrl.FundAppBaseURL + stocks[i].FundCode
		}
		if stocks[i].AnalystStocksdata.StockType == config.StockCompany {
			res.Data[i].Name = stocks[i].AnalystStocksdata.CodeName
		} else {
			res.Data[i].Name = stocks[i].AnalystStocksdata.IndustryName
		}
		res.Data[i].AuthorName = stocks[i].AnalystStocksdata.AuthorName
		res.Data[i].PotentialGain = stocks[i].AnalystStocksdata.PotentialGain
		res.Data[i].RecommendFileName = stocks[i].AnalystStocksdata.FileName
		res.Data[i].RecommendFileURL = stocks[i].AnalystStocksdata.FileUrl
		if stocks[i].AnalystCollect.Collection == "" {
			res.Data[i].Collection = config.CollectNo
		} else {
			res.Data[i].Collection = stocks[i].AnalystCollect.Collection
		}
		if stocks[i].AnalystStocksdata.PotentialGain != 0 {
			// 强烈推荐有k线
			if stocks[i].AnalystStocksdata.SratingName == config.SratingNameRecommend {
				// todo 给kline加redis
				switch stocks[i].AnalystStocksdata.StockType {
				case config.StockCompany:
					res.Data[i].KData = handleKLine(ctx, stocks[i].AnalystStocksdata.Code)
				case config.StockIndustry:
					res.Data[i].KData = handleIndustryKLine(ctx, stocks[i].AnalystStocksdata.Code)
				}
			} else {
				res.Data[i].KData = []interfaces.KData{}
			}
		} else {
			res.Data[i].KData = []interfaces.KData{}
		}
	}
}

// 处理股票历史数据
func handleAnalyzeHistory(ctx context.Context, res *interfaces.AnalystHistoryResponse, stocks []model.Collect) {
	for i := 0; i < len(stocks); i++ {
		res.Data[i].StocksID = stocks[i].AnalystStocksdata.Id
		res.Data[i].StocksType = stocks[i].AnalystStocksdata.StockType
		if stocks[i].AnalystStocksdata.FundName != "" {
			res.Data[i].FundName = stocks[i].AnalystStocksdata.FundName
			res.Data[i].FundURL = config.GetGlobalConfig().AnalystUrl.FundAppBaseURL + stocks[i].FundCode
		}
		res.Data[i].AuthorName = stocks[i].AnalystStocksdata.AuthorName
		res.Data[i].PotentialGain = stocks[i].AnalystStocksdata.PotentialGain
		res.Data[i].RecommendFileName = stocks[i].AnalystStocksdata.FileName
		res.Data[i].RecommendFileURL = stocks[i].AnalystStocksdata.FileUrl
		if stocks[i].AnalystCollect.Collection == "" {
			res.Data[i].Collection = config.CollectNo
		} else {
			res.Data[i].Collection = config.CollectYes
		}
		if stocks[i].StockType == config.StockIndustry {
			res.Data[i].Name = stocks[i].AnalystStocksdata.IndustryName
		} else {
			res.Data[i].Name = stocks[i].AnalystStocksdata.CodeName
		}
		res.Data[i].KData = []interfaces.KData{}
	}
}

func handleKLine(ctx context.Context, code string) []interfaces.KData {
	method := "handleKLine"
	datas := []interfaces.KData{}
	url := config.GetGlobalConfig().AnalystUrl.StockKLine + "?count=30&period=86400&code=" + code
	data, err := net.GetKLine(ctx, url)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "请求k线数据", "error", err)
		return datas
	}
	for i := 0; i < len(data.Results.Datas); i++ {
		datas = append(datas, interfaces.KData{
			ClosePrice: data.Results.Datas[i].ClosePrice,
			OpenPrice:  data.Results.Datas[i].OpenPrice,
			HighPrice:  data.Results.Datas[i].HighPrice,
			LowPrice:   data.Results.Datas[i].LowPrice,
			Timestamp:  data.Results.Datas[i].TickTimestamp,
		})
	}
	return datas
}

func handleIndustryKLine(ctx context.Context, code string) []interfaces.KData {
	method := "handleIndustryKLine"
	datas := []interfaces.KData{}
	url := config.GetGlobalConfig().AnalystUrl.StockIndustryKLine + "?limit=30&code=" + code
	data, err := net.GetIndustryKLine(ctx, url)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "请求k线数据", "error", err)
		return datas
	}
	for i := 0; i < len(data.Results.Datas); i++ {
		t, err := time.Parse("2006-01-02", data.Results.Datas[i].Dt)
		if err != nil {
			continue
		}
		datas = append(datas, interfaces.KData{
			ClosePrice: data.Results.Datas[i].Close,
			OpenPrice:  data.Results.Datas[i].Open,
			HighPrice:  data.Results.Datas[i].High,
			LowPrice:   data.Results.Datas[i].Low,
			Timestamp:  t.Unix(),
		})
	}
	return datas
}
