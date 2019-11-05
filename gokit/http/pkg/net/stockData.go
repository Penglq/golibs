package net

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

func GetStock(ctx context.Context, url string) (resStock *ResStock, err error) {
	resStock = new(ResStock)
	originCompanyStockByte, err := utils.NewHttpClient(&ctx,
		utils.UrlOptions(url),
		utils.TryOptions(1, time.Second)).HttpGet()
	if err != nil {
		return
	}
	err = json.Unmarshal(originCompanyStockByte, &resStock)
	if err != nil {
		return
	}
	if resStock.ErrorCode != "0" {
		err = errors.New("数据源响应不为成功态:" + resStock.ErrorCode)
	}
	return
}

func GetKLine(ctx context.Context, url string) (resKLine *ResKLine, err error) {
	resKLine = new(ResKLine)
	originCompanyStockByte, err := utils.NewHttpClient(&ctx,
		utils.UrlOptions(url),
		utils.TryOptions(1, time.Second)).HttpGet()
	if err != nil {
		return
	}
	err = json.Unmarshal(originCompanyStockByte, &resKLine)
	if err != nil {
		return
	}
	if resKLine.ErrorCode != "0" {
		err = errors.New("数据源响应不为成功态:" + resKLine.ErrorCode)
	}
	return
}

func GetIndustryKLine(ctx context.Context, url string) (resKLine *ResIndustryKLine, err error) {
	resKLine = new(ResIndustryKLine)
	originIndustryStockByte, err := utils.NewHttpClient(&ctx,
		utils.UrlOptions(url),
		utils.TryOptions(1, time.Second)).HttpGet()
	if err != nil {
		return
	}
	err = json.Unmarshal(originIndustryStockByte, &resKLine)
	if err != nil {
		return
	}
	if resKLine.ErrorCode != "0" {
		err = errors.New("数据源响应不为成功态:" + resKLine.ErrorCode)
	}
	return
}

type ResStock struct {
	Count     int         `json:"count"`
	Results   []Stockdata `json:"results"`
	ErrorCode string      `json:"errorCode"` // "0"成功
}

type ResKLine struct {
	Count   int `json:"count"`
	Results struct {
		Code  string  `json:"code"`
		Datas []KLine `json:"datas"`
	} `json:"results"`
	ErrorCode string `json:"errorCode"` // "0"成功
}

type ResIndustryKLine struct {
	Count   int `json:"count"`
	Results struct {
		Code  string          `json:"code"`
		Name  string          `json:"name"`
		Datas []IndustryKLine `json:"datas"`
	} `json:"results"`
	ErrorCode string `json:"errorCode"` // "0"成功
}

// 股票
type Stockdata struct {
	Attach []struct {
		// FileSize string `json:"fileSize"`
		// Filetype string `json:"filetype"`
		Name string `json:"name"`
		// Pagenum  string `json:"pagenum"`
		// Seq      int    `json:"seq"`
		URL string `json:"url"`
	} `json:"attach"`
	AuthorList []struct {
		Auth          string `json:"auth"`
		Authprizeinfo string `json:"authprizeinfo"`
		// Authcode      string `json:"authcode"`
		// IsWealth      string `json:"isWealth"`
	} `json:"authorList"`
	Actualdprice    float64 `json:"actualdprice"`
	Code            string  `json:"code"`
	Codename        string  `json:"codename"`
	Date            string  `json:"date"`
	Industry        string  `json:"industry"`
	IndustryCode    string  `json:"industrycode"`
	ReportTimestamp int     `json:"reportTimestamp"`
	SratingName     string  `json:"sratingName"`
	// Cprice          string  `json:"cprice"`
	// Change          string `json:"change"`
	// Dprice          string `json:"dprice"`
	// ID              string `json:"id"`
	// Kcode           string `json:"kcode"`
	// Kname           string `json:"kname"`
	// Ktype           string `json:"ktype"`
	// Org             string `json:"org"`
	// Orgcode         string `json:"orgcode"`
	// Orgprizeinfo    string `json:"orgprizeinfo"`
	// Priority        int    `json:"priority"`
	// Rate            string `json:"rate"`
	// ReportDate      string `json:"reportDate"`
	// Rtype           string `json:"rtype"`
	// Rtypecode       string `json:"rtypecode"`
	// StockType       string `json:"stockType"`
	// Title           string `json:"title"`
}

// 日k线
type KLine struct {
	ClosePrice float64 `json:"close_price"`
	HighPrice  float64 `json:"high_price"`
	LowPrice   float64 `json:"low_price"`
	OpenPrice  float64 `json:"open_price"`
	// AveragePrice    int     `json:"average_price"`
	// PriceChange     float64 `json:"price_change"`
	// PriceChangeRate float64 `json:"price_change_rate"`
	// TickTime        string  `json:"tick_time"`
	TickTimestamp int64 `json:"tick_timestamp"`
	// TurnoverValue   float64 `json:"turnover_value"`
	// TurnoverVolume  int     `json:"turnover_volume"`
}

// 行业k线
type IndustryKLine struct {
	Dt    string  `json:"dt"`
	Close float64 `json:"close"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Open  float64 `json:"open"`
}
