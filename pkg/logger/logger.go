package logger

import (
	"context"
	"github.com/mdobak/go-xerrors"
	"log/slog"
	"os"
	"sync"
)

var once sync.Once
var l *slog.Logger

func New() *slog.Logger {
	once.Do(func() {
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		slog.SetDefault(l)
	})
	return l
}

func Info(ctx context.Context, msg string, info any) {
	l.InfoContext(ctx, msg, "info", info)
}

func Debug(ctx context.Context, msg string, info any) {
	l.DebugContext(ctx, msg, "debug", info)
}

func Error(ctx context.Context, msg string, info any) {
	infoStr, ok := info.(string)
	if ok {
		l.ErrorContext(ctx, msg, "error", infoStr)
	} else {
		trace := xerrors.StackTrace(info.(error))
		l.ErrorContext(ctx, msg, "error", trace)
	}
}
