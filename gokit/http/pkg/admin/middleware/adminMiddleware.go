package middleware

import (
	"context"
	"github.com/Penglq/QLog"
	"github.com/go-kit/kit/endpoint"
	"git/miniTools/data-service/common/interfaces"
)

func ConsultantGetAllListMiddleware(next interfaces.ConsultantGetAllListEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method := "ConsultantGetAllListMiddleware"
		req, ok := request.(interfaces.ConsultantGetAllListRequest)
		if !ok {
			QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "接收参数出错", "error", ok)
			return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
		}

		return next(ctx, &req)
	}
}