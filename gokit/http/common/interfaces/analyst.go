package interfaces

import "context"

type AnalystFirstEndpoint func(context.Context, *AnalystFirstRequest) (interface{}, error)
type AnalystAnalyzeEndpoint func(context.Context, *AnalystAnalyzeRequest) (interface{}, error)
type AnalystHistoryEndpoint func(context.Context, *AnalystHistoryRequest) (interface{}, error)
type AnalystCollectEndpoint func(context.Context, *AnalystCollectRequest) (interface{}, error)
type AnalystCollectHistoryEndpoint func(context.Context, *AnalystCollectHistoryRequest) (interface{}, error)

// request
type AnalystFirstRequest struct {
	YrdUid     string `json:"yrdUid"`
	PassportId string `json:"passportId"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
}
type AnalystAnalyzeRequest struct {
	YrdUid   string `json:"yrdUid"`
	TabType  string `json:"tabType"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}
type AnalystHistoryRequest struct {
	YrdUid   string `json:"yrdUid"`
	Date     string `json:"date"`
	TabType  string `json:"tabType"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}
type AnalystCollectRequest struct {
	YrdUid     string `json:"yrdUid"`
	StocksId   int    `json:"stocksId"`
	Collection string `json:"collection"`
}
type AnalystCollectHistoryRequest struct {
	YrdUid   string `json:"yrdUid"`
	TabType  string `json:"tabType"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}

// response
type AnalystFirstResponse struct {
	First string `json:"first"`
}
type AnalystAnalyzeResponse struct {
	Date string    `json:"date"`
	Data []Analyze `json:"data"`
}
type AnalystHistoryResponse struct {
	Date string    `json:"date"`
	Data []Analyze `json:"data"`
}
type AnalystCollectResponse struct{}
type AnalystCollectHistoryResponse struct {
	TotalPage int64            `json:"totalPage"`
	Data      []CollectHistory `json:"data"`
}

type KData struct {
	ClosePrice float64 `json:"closePrice"`
	HighPrice  float64 `json:"highPrice"`
	LowPrice   float64 `json:"lowPrice"`
	OpenPrice  float64 `json:"openPrice"`
	Timestamp  int64   `json:"timestamp"`
}

type Analyze struct {
	Collection        string  `json:"collection"`
	FundName          string  `json:"fundName"`
	FundURL           string  `json:"fundUrl"`
	Name              string  `json:"name"`
	AuthorName        string  `json:"authorName"`
	PotentialGain     int     `json:"potentialGain"`
	RecommendFileName string  `json:"recommendFileName"`
	RecommendFileURL  string  `json:"recommendFileUrl"`
	StocksID          int     `json:"stocksId"`
	StocksType        string  `json:"stocksType"`
	KData             []KData `json:"kData"`
}

type CollectHistory struct {
	Date              string `json:"date"`
	AuthorName        string `json:"authorName"`
	Collection        string `json:"collection"`
	FundName          string `json:"fundName"`
	FundURL           string `json:"fundUrl"`
	Name              string `json:"name"`
	PotentialGain     int    `json:"potentialGain"`
	RecommendFileName string `json:"recommendFileName"`
	RecommendFileURL  string `json:"recommendFileUrl"`
	StocksID          int    `json:"stocksId"`
	StocksType        string `json:"stocksType"`
}
