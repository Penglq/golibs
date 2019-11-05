package transport

import (
	"context"
	"encoding/json"
	"github.com/penglq/QLog"
	"net/http"
)

func AnalystFirstDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfacess.AnalystFirstRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析账户请求失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}
func AnalystAnalyzeDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfacess.AnalystAnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析账户请求失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}
func AnalystHistoryDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfacess.AnalystHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析账户请求失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}
func AnalystCollectDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfacess.AnalystCollectRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析账户请求失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}
func AnalystCollectHistoryDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfacess.AnalystCollectHistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析账户请求失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}
