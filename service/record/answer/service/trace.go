package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

type Trace struct {
}

func (t *Trace) StartSpanFromContext(ctx context.Context, funcName string) (func(), context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, funcName)
	return func() {
		span.Finish()
	}, ctx
}
