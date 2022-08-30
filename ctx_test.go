package golibs

import (
	"context"
	"testing"
)

func TestCTX(t *testing.T) {
	ctx := context.Background()
	ctx = SetTraceIdContext(ctx, CreateTraceId())
	ctx = SetTraceIdContext(ctx, CreateTraceId())
	t.Log(GetTraceIdFromCTX(ctx))
}
