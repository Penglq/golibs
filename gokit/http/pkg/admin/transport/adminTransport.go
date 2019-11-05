package transport

import (
	"context"
	"encoding/json"
	"github.com/Penglq/QLog"
	"net/http"
)

func ConsultantGetAllListDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interfacess.ConsultantGetAllListRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "message:", "解析账户请求失败", "error:", err)
		return nil, nil
	}
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "request", request)
	return request, nil
}
