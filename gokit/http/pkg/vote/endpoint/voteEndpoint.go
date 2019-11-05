package endpoint

import (
	"context"
	"git/miniTools/data-service/common/interfaces"
	"git/miniTools/data-service/common/value"
	"git/miniTools/data-service/pkg/model"
	"git/miniTools/data-service/utils"
	"git/miniTools/data-service/utils/db"
	"github.com/penglq/QLog"
	"log"
)

// 查询票数
func VoteCountEndpoint(ctx context.Context, req *interfaces.VoteCountRequest) (response interface{}, err error) {
	method := "VoteCountEndpoint"
	voteCount, err := model.GetVoteCountByVoteId(req.VoteId)
	if err != nil {
		QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "GetVoteCountById错误", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	}
	responseData := make([]interfaces.CurrentVoteStatus, len(voteCount))

	for i := 0; i < len(voteCount); i++ {
		responseData[i].VoteId = voteCount[i].VoteId
		responseData[i].ItemId = voteCount[i].ItemId
		responseData[i].Count = voteCount[i].Count
		log.Print(req.UserId)
		if req.UserId != "" {
			if has, _ := db.GetEngine().Exist(&model.VoteRecord{UserId: req.UserId, VoteId: req.VoteId, ItemId: voteCount[i].ItemId}); has {
				log.Print(has)
				responseData[i].CurrentStatus = true
			}
		}

	}
	log.Print(responseData)
	return new(utils.OutResponse).ResponseSuccService(responseData), nil
}

// 增加票数
func VoteIncreaseEndpoint(ctx context.Context, req *interfaces.VoteIncreaseRequest) (response interface{}, err error) {
	method := "VoteIncreaseEndpoint"

	if has, _ := db.GetEngine().Exist(&model.VoteRecord{UserId: req.UserId, VoteId: req.VoteId}); has {
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeIllegal), nil
	}

	err = model.IncreaseVoteCount(req.VoteId, req.ItemId, req.Count)
	if err != nil {
		QLog.GetLogger().AlertWithLevel(QLog.ALERTCRITICAL, utils.TraceKey, utils.GetTraceIdFromCTX(ctx), "method", method, "action", "增加票数错误", "error", err)
		return new(utils.OutResponse).ResponseErrorService(value.SourceDataservice, value.ParamNone, value.TypeService), nil
	}

	//投票记录
	db.GetEngine().Insert(&model.VoteRecord{UserId: req.UserId, VoteId: req.VoteId, ItemId: req.ItemId})

	return new(utils.OutResponse).ResponseSuccService(struct{}{}), nil
}
