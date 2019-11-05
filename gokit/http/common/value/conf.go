package value

// 错误码相关
const (
	SourceSuccess = "0"
	SourceFailed  = "-1"
	// 错误源key
	SourceDataservice = "data-service"

	// 参数key
	ParamNone          = "None"
	ParamDt            = "Dt"
	ParamCompany       = "Company"
	ParamInsuranceType = "InsuranceType"
	ParamRequest       = "Request"
	ParamResponse      = "Response"
	ParamUpdate        = "Update"
	ParamDelete        = "Delete"
	ParamCombine       = "Combine"
	ParamReservation   = "Reservation"
	ParamStock         = "Stock"

	// 错误类型
	TypeRequired      = "required"
	TypeIllegal       = "illegal"
	TypeNotExist      = "notExist"
	TypeFail          = "serviceFail"
	TypeService       = "serviceExceptions"
	TypeNotSuccStatus = "notSuccStatus"
	TypeConflict      = "conflict"
)
