package routers

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"git/miniTools/data-service/common/interfaces"
	analystEndpoint "git/miniTools/data-service/pkg/analyst/endpoint"
	analystMiddleware "git/miniTools/data-service/pkg/analyst/middleware"
	analystTransport "git/miniTools/data-service/pkg/analyst/transport"
	"git/miniTools/data-service/transport/common"
)

func NewAnalystRouter(r *mux.Router) {
	interfaces.SubRouterOptions(r.PathPrefix(V1).Subrouter(),
		AnalystRouter,
	)
}
func AnalystRouter(r *mux.Router) {
	r.Handle(AnalystFirst, httptransport.NewServer(
		analystMiddleware.AnalystFirstMiddleware(analystEndpoint.AnalystFirstEndpoint),
		analystTransport.AnalystFirstDecodeRequest,
		common.EncodeResponse))
	r.Handle(AnalystAnalyze, httptransport.NewServer(
		analystMiddleware.AnalystAnalyzeMiddleware(analystEndpoint.AnalystAnalyzeEndpoint),
		analystTransport.AnalystAnalyzeDecodeRequest,
		common.EncodeResponse))
	r.Handle(AnalystHistory, httptransport.NewServer(
		analystMiddleware.AnalystHistoryMiddleware(analystEndpoint.AnalystHistoryEndpoint),
		analystTransport.AnalystHistoryDecodeRequest,
		common.EncodeResponse))
	r.Handle(AnalystCollect, httptransport.NewServer(
		analystMiddleware.AnalystCollectMiddleware(analystEndpoint.AnalystCollectEndpoint),
		analystTransport.AnalystCollectDecodeRequest,
		common.EncodeResponse))
	r.Handle(AnalystCollectHistory, httptransport.NewServer(
		analystMiddleware.AnalystCollectHistoryMiddleware(analystEndpoint.AnalystCollectHistoryEndpoint),
		analystTransport.AnalystCollectHistoryDecodeRequest,
		common.EncodeResponse))
}