package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/penglq/QLog"
)

func AnalystFirstMiddleware(next interfacess.AnalystFirstEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "AnalystFirstMiddleware"
		req, ok := request.(interfacess.AnalystFirstRequest)
		if !ok {
			QLog.GetLogger().Error(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}

		return next(ctx, &req)
	}
}

func AnalystAnalyzeMiddleware(next interfacess.AnalystAnalyzeEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "AnalystAnalyzeMiddleware"
		req, ok := request.(interfacess.AnalystAnalyzeRequest)
		if !ok {
			QLog.GetLogger().Error(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}

		return next(ctx, &req)
	}
}
func AnalystHistoryMiddleware(next interfacess.AnalystHistoryEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "AnalystHistoryMiddleware"
		req, ok := request.(interfacess.AnalystHistoryRequest)
		if !ok {
			QLog.GetLogger().Error(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}
		if req.Date == "" {
			QLog.GetLogger().Error(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数", "error", "date不可为空")
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamDt, value.TypeRequired), nil
		}
		return next(ctx, &req)
	}
}
func AnalystCollectMiddleware(next interfacess.AnalystCollectEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "AnalystCollectMiddleware"
		req, ok := request.(interfacess.AnalystCollectRequest)
		if !ok {
			QLog.GetLogger().Error(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}
		if req.YrdUid == "" {
			QLog.GetLogger().Error(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}
		return next(ctx, &req)
	}
}
func AnalystCollectHistoryMiddleware(next interfacess.AnalystCollectHistoryEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "AnalystCollectHistoryMiddleware"
		req, ok := request.(interfacess.AnalystCollectHistoryRequest)
		if !ok {
			QLog.GetLogger().Error(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}

		return next(ctx, &req)
	}
}