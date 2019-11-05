package endpoint

import (

)

var answerSep = `|`

func ConsultantGetAllListEndpoint(ctx context.Context, req *interfaces.ConsultantGetAllListRequest) (response interface{}, err error) {
	method := "ConsultantGetListEndpoint"
	total, err := model.GetConsultantCount()
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询保单记录总数", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	}
	consultant, err := model.GetConsultantAllList(req.Page, req.PageSize)
	if err != nil {
		QLog.GetLogger().Alert(utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "查询保单列表", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	}
	resData := interfaces.ConsultantGetAllListResponse{
		Page:           req.Page,
		PageSize:       req.PageSize,
		Total:          total,
		PageCount:      utils.Ceil(total, req.PageSize),
		ConsultantInfo: make([]map[string]interface{}, 0, len(consultant)),
	}
	// fmt.Println(">>>>>", consultant)
	for i := 0; i < len(consultant); i++ {
		info := make(map[string]interface{})
		info["id"] = consultant[i].Id
		info["yrdUid"] = consultant[i].YrdUid
		info["name"] = consultant[i].Name
		info["mobile"] = consultant[i].Mobile
		info["reservationTime"] = consultant[i].ReservationTime
		info["sex"] = consultant[i].Sex
		info["age"] = consultant[i].Age
		info["money"] = consultant[i].Money
		HandleDetailAnswers2Str(handleDetailQuestions(&consultant[i]), &info)
		resData.ConsultantInfo = append(resData.ConsultantInfo, info)
	}
	return new(utils.OutResponse).ResponseSuccService(resData), nil
}

func HandleDetailAnswers2Str(consultantDetail map[string][]string, info *map[string]interface{}) {
	conf := config.GetGlobalConfig().Questions
	for i := 0; i < len(conf); i++ {
		// fmt.Println(">>>>", i, consultantDetail[conf[i].QuestionId], conf[i].QuestionId)
		if v, ok := consultantDetail[conf[i].QuestionId]; ok {
			ans := []string{}
			for j := 0; j < len(v); j++ {
				if _, ok := conf[i].Answers[v[j]]; ok {
					ans = append(ans, conf[i].Answers[v[j]])
				}
			}
			(*info)[conf[i].QuestionId] = strings.Join(ans, " "+answerSep+" ")
		}
	}
	// fmt.Println(">>>>2", *info)
}

func handleDetailQuestions(consultantDetail *model.ToolsConsultant) (resData map[string][]string) {
	resData = make(map[string][]string)
	resData[consultantDetail.Question1] = strings.Split(consultantDetail.Answer1, answerSep)
	resData[consultantDetail.Question2] = strings.Split(consultantDetail.Answer2, answerSep)
	resData[consultantDetail.Question3] = strings.Split(consultantDetail.Answer3, answerSep)
	resData[consultantDetail.Question4] = strings.Split(consultantDetail.Answer4, answerSep)
	resData[consultantDetail.Question5] = strings.Split(consultantDetail.Answer5, answerSep)
	resData[consultantDetail.Question6] = strings.Split(consultantDetail.Answer6, answerSep)
	delete(resData, "")
	return
}
