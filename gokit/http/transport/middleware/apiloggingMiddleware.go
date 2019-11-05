package middleware

import (
	"context"
	"git/miniTools/data-service/utils"
	"github.com/penglq/QLog"
	"net/http"
	"time"
)

func ApiLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // allow any origin domain
		// if options.Origin != "" {
		// 	w.Header().Set("Access-Control-Allow-Origin", options.Origin)
		// }
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Auth-Token, X-Auth-UUID, X-Auth-Openid, referrer, Authorization, x-client-id, x-client-version, x-client-type")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("content-type", "application/json; charset=utf-8")

		traceId := r.Header.Get("minitools-trace-id")
		if traceId == "" {
			traceId = utils.CreateTraceId()
		}
		ctx := utils.SetTraceIdContext(context.Background(), traceId)
		r = r.WithContext(ctx)
		defer func(t time.Time) {
			QLog.GetLogger().Info(
				utils.TraceKey, utils.GetTraceIdFromCTX(ctx),
				//"header", r.Header,
				"remoteAddr", r.RemoteAddr,
				"requestURI", r.RequestURI,
				"took/s", time.Since(t).Seconds(),
			)
		}(time.Now())
		next.ServeHTTP(w, r)
	})
}
