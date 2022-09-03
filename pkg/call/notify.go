package call

import "context"

type StopType int32

const (
	Busy    StopType = 0
	Timeout StopType = 1
	Reject  StopType = 2
	Hangup  StopType = 3
	Cancel  StopType = 4
)

type SignalNotify interface {
	SendStartCallSignal(ctx context.Context, target string, traceId int64) error
	SendAcceptCallSignal(ctx context.Context, target string, traceId int64, ticket Ticket) error
	SendStopCallSignal(ctx context.Context, target string, traceId int64, stopType StopType) error
}
