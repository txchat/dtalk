package logger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/txchat/dtalk/pkg/api"
	"github.com/txchat/dtalk/pkg/api/trace"

	"github.com/rs/zerolog"
)

func New(env, srvName string) zerolog.Logger {
	var logger zerolog.Logger
	var out io.Writer
	out = os.Stdout
	//log init
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if env == DEBUG {
		out = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if env == BENCHMARK {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
	logger = zerolog.New(out).With().Timestamp().Str("srvName", srvName).Logger()

	return logger
}

func NewLogWithCtx(ctx context.Context, log zerolog.Logger) zerolog.Logger {
	addr := api.NewAddrWithContext(ctx)

	traceId := trace.NewTraceIdWithContext(ctx)

	return log.With().Str("opeId", addr).Str("trace", traceId).Logger()
}
