package msgfactory

import (
	"context"
)

type Trace struct {
}

func NewTrace() *Trace {
	return &Trace{}
}

func (t *Trace) StartSpanFromContext(ctx context.Context, funcName string) (func(), context.Context) {
	//TODO trace impl
	return func() {
	}, ctx
}
