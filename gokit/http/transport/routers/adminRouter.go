package routers

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"git/miniTools/data-service/common/interfaces"
	adminEndpoint "git/miniTools/data-service/pkg/admin/endpoint"
	adminMiddleware "git/miniTools/data-service/pkg/admin/middleware"
	adminTransport "git/miniTools/data-service/pkg/admin/transport"
	"git/miniTools/data-service/transport/common"
)

func NewAdminRouter(r *mux.Router) {
	interfaces.SubRouterOptions(r.PathPrefix(Admin).Subrouter(),
		AdminRouter,
	)
}

func AdminRouter(r *mux.Router) {
	r.Handle(AdminConsultantGetAllList, httptransport.NewServer(
		adminMiddleware.ConsultantGetAllListMiddleware(adminEndpoint.ConsultantGetAllListEndpoint),
		adminTransport.ConsultantGetAllListDecodeRequest,
		common.EncodeResponse))
}
