package routers

const (
	Admin = "/admin"
	V1    = "/v1"
)
const (
	// admin
	AdminConsultantGetAllList = "/consultant/getAllList"
	// 金牌分析师
	AnalystFirst          = "/analyst/first"
	AnalystAnalyze        = "/analyst/analyze"
	AnalystHistory        = "/analyst/history"
	AnalystCollect        = "/analyst/collect"
	AnalystCollectHistory = "/analyst/collectHistory"

	// 保险管理
	InsuranceCompanyV1   = "/insurance/company"
	InsuranceInsertV1    = "/insurance/insert"
	InsuranceUpdateV1    = "/insurance/update"
	InsuranceDeleteV1    = "/insurance/delete"
	InsuranceGetListV1   = "/insurance/getList"
	InsuranceGetDetailV1 = "/insurance/getDetail"
	// 专属理财顾问
	ConsultantQuestionV1   = "/consultant/question"
	ConsultantInsertV1     = "/consultant/insert"
	ConsultantGetDetailV1  = "/consultant/getLastDetail"
	ConsultantGetListV1    = "/consultant/getList"
	ConsultantGetAllListV1 = "/consultant/getAllList"

	//投票
	VoteIncreaseV1 = "/vote/increase"
	VoteGetCountV1 = "/vote/getCount"

	//指数低估值
	IndexValuationGetListV1 = "/indexValuation/getList"

	//猜涨跌
	GuessRiseFallGetTradingDay  = "/guessRiseFall/getTradingDay"
	GuessRiseFallGetCurrentInfo = "/guessRiseFall/getCurrentInfo"
	GuessRiseFallGuess          = "/guessRiseFall/guess"
	GuessRiseFallGetGuessRecord = "/guessRiseFall/getGuessRecord"
	GuessRiseFallRead           = "/guessRiseFall/read"
	GuessRiseFallGetGuessRank   = "/guessRiseFall/getGuessRank"

	//端外投放保险1v1链接
	Insurance1v1AddV1 = "/insurance1v1/add"
)
