package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strings"
)

const TraceKey = "trace-id"
const UserId = "userId"

var (
	prefix = "traceid-"
)

func SetContext(ctx context.Context, k, v string) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.New(map[string]string{k: v}))
}
func GetContext(ctx context.Context, k string) interface{} {
	return getCTXValue(ctx, k)
}
func CreateTraceId(val ...string) string {
	return prefix + strings.Join(append(val, GetNowDateStr("20060102150405"), GetRandomString(6)), `-`)
}

func GetTraceIdFromCTX(ctx context.Context) string {
	return getMetadataCTXValue(ctx, TraceKey)
}

func SetTraceIdContext(ctx context.Context, traceId string) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.New(map[string]string{TraceKey: traceId}))
}

func getMetadataCTXValue(ctx context.Context, k string) (str string) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if v, ok := md[k]; ok {
			if len(v) > 0 {
				str = v[0]
			}
		}
	}
	return
}
func GetUserIdFromCTX(ctx context.Context) int {
	return getCTXValue(ctx, UserId).(int)
}
func SetUserIdContext(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, UserId, userId)
}
func getCTXValue(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}
