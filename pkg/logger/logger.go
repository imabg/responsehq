package logger

import (
	"context"
	"fmt"
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

func Error(ctx context.Context, msg string, info error) {
	trace := xerrors.StackTrace(xerrors.New(info))
	fmt.Println(trace)
	l.ErrorContext(ctx, msg, "error", info.Error())
}

func DBError(ctx context.Context, msg string, message string) {
	wrpErr := xerrors.New(message)
	l.ErrorContext(ctx, msg, "error", wrpErr)
}
