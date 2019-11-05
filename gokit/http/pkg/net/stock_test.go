package net

import (
	"context"
	"git/miniTools/data-service/config"
	"git/miniTools/data-service/utils"
	"testing"
)

func TestGetStock(t *testing.T) {
	initConfig()
	url := config.GetGlobalConfig().AnalystUrl.StockCompany + "?date=20190524"
	ctx := context.Background()
	utils.SetTraceIdContext(ctx, utils.CreateTraceId())
	data, err := GetStock(ctx, url)
	t.Logf("%+v\n%+v", data, err)
}

func TestGetKLine(t *testing.T) {
	initConfig()
	url := config.GetGlobalConfig().AnalystUrl.StockKLine + "?code=300036.SZ&count=20&period=86400"
	ctx := context.Background()
	utils.SetTraceIdContext(ctx, utils.CreateTraceId())
	data, err := GetKLine(ctx, url)
	t.Logf("%+v\n%+v", data, err)
}

func TestGetIndustryKLine(t *testing.T) {
	initConfig()
	url := config.GetGlobalConfig().AnalystUrl.StockIndustryKLine + "?code=017026"
	ctx := context.Background()
	utils.SetTraceIdContext(ctx, utils.CreateTraceId())
	data, err := GetIndustryKLine(ctx, url)
	t.Logf("%+v\n%+v", data, err)
}