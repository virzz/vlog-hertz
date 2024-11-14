package vloghertz

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/virzz/vlog"
)

// Logger slog impl
type HLog struct {
	l   *slog.Logger
	cfg *vlog.Config
}

func (l *HLog) Logger() *slog.Logger { return l.l }

func (l *HLog) SetLevel(level hlog.Level) {
	lvl := hLevelToSLevel(level)
	l.cfg.Level.Set(lvl)
}
func (l *HLog) SetOutput(writer io.Writer) {
	l.cfg.Output = writer
	l.l = slog.New(slog.NewJSONHandler(writer, l.cfg.HandlerOptions))
}
func (l *HLog) log(level hlog.Level, v ...any) {
	l.l.Log(context.TODO(), hLevelToSLevel(level), fmt.Sprint(v...))
}
func (l *HLog) logf(level hlog.Level, format string, kvs ...any) {
	l.l.Log(context.TODO(), hLevelToSLevel(level), fmt.Sprintf(format, kvs...))
}
func (l *HLog) ctxLogf(level hlog.Level, ctx context.Context, format string, v ...any) {
	l.l.Log(ctx, hLevelToSLevel(level), fmt.Sprintf(format, v...))
}
func (l *HLog) Trace(v ...any)             { l.log(hlog.LevelTrace, v...) }
func (l *HLog) Debug(v ...any)             { l.log(hlog.LevelDebug, v...) }
func (l *HLog) Info(v ...any)              { l.log(hlog.LevelInfo, v...) }
func (l *HLog) Notice(v ...any)            { l.log(hlog.LevelNotice, v...) }
func (l *HLog) Warn(v ...any)              { l.log(hlog.LevelWarn, v...) }
func (l *HLog) Error(v ...any)             { l.log(hlog.LevelError, v...) }
func (l *HLog) Fatal(v ...any)             { l.log(hlog.LevelFatal, v...) }
func (l *HLog) Tracef(f string, v ...any)  { l.logf(hlog.LevelTrace, f, v...) }
func (l *HLog) Debugf(f string, v ...any)  { l.logf(hlog.LevelDebug, f, v...) }
func (l *HLog) Infof(f string, v ...any)   { l.logf(hlog.LevelInfo, f, v...) }
func (l *HLog) Noticef(f string, v ...any) { l.logf(hlog.LevelNotice, f, v...) }
func (l *HLog) Warnf(f string, v ...any)   { l.logf(hlog.LevelWarn, f, v...) }
func (l *HLog) Errorf(f string, v ...any)  { l.logf(hlog.LevelError, f, v...) }
func (l *HLog) Fatalf(f string, v ...any)  { l.logf(hlog.LevelFatal, f, v...) }

func (l *HLog) CtxTracef(ctx context.Context, f string, v ...any) {
	l.ctxLogf(hlog.LevelDebug, ctx, f, v...)
}
func (l *HLog) CtxDebugf(ctx context.Context, f string, v ...any) {
	l.ctxLogf(hlog.LevelDebug, ctx, f, v...)
}
func (l *HLog) CtxInfof(ctx context.Context, f string, v ...any) {
	l.ctxLogf(hlog.LevelInfo, ctx, f, v...)
}
func (l *HLog) CtxNoticef(ctx context.Context, f string, v ...any) {
	l.ctxLogf(hlog.LevelNotice, ctx, f, v...)
}
func (l *HLog) CtxWarnf(ctx context.Context, f string, v ...any) {
	l.ctxLogf(hlog.LevelWarn, ctx, f, v...)
}
func (l *HLog) CtxErrorf(ctx context.Context, f string, v ...any) {
	l.ctxLogf(hlog.LevelError, ctx, f, v...)
}
func (l *HLog) CtxFatalf(ctx context.Context, f string, v ...any) {
	l.ctxLogf(hlog.LevelFatal, ctx, f, v...)
}

var _ hlog.FullLogger = (*HLog)(nil)
