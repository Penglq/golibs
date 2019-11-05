package common

import (
	"context"
	"encoding/json"
	"git/miniTools/data-service/utils"
	"github.com/penglq/QLog"
	"net/http"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	QLog.GetLogger().Info(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "response", response)
	return json.NewEncoder(w).Encode(response)
}
