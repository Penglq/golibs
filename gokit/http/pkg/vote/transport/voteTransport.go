package transport

import (
	"context"
	"encoding/json"
	"git/miniTools/data-service/common/interfaces"
	"git/miniTools/data-service/utils"
	"github.com/penglq/QLog"
	"net/http"
)

func VoteCountDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfaces.VoteCountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析VoteCountRequest失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}

func VoteIncreaseDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfaces.VoteIncreaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析VoteIncreaseRequest失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}
