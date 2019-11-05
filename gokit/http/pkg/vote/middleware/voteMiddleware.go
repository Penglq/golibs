package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"git/miniTools/data-service/common/interfaces"
	"git/miniTools/data-service/common/value"
	"git/miniTools/data-service/utils"
	"github.com/penglq/QLog"
)

func VoteCountMiddleware(next interfaces.VoteCountEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "VoteCountMiddleware"
		req, ok := request.(interfaces.VoteCountRequest)
		if !ok {
			QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}
		return next(ctx, &req)
	}
}

func VoteIncreaseMiddleware(next interfaces.VoteIncreaseEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "VoteIncreaseMiddleware"
		req, ok := request.(interfaces.VoteIncreaseRequest)
		if !ok {
			QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}
		return next(ctx, &req)
	}
}
