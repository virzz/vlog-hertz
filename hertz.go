package vloghertz

import (
	"log/slog"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/virzz/vlog"
)

const (
	LevelTrace  = slog.Level(-8)
	LevelNotice = slog.Level(2)
	LevelFatal  = slog.Level(12)
)

func hLevelToSLevel(level hlog.Level) (lvl slog.Level) {
	switch level {
	case hlog.LevelTrace:
		lvl = LevelTrace
	case hlog.LevelDebug:
		lvl = slog.LevelDebug
	case hlog.LevelInfo:
		lvl = slog.LevelInfo
	case hlog.LevelWarn:
		lvl = slog.LevelWarn
	case hlog.LevelNotice:
		lvl = LevelNotice
	case hlog.LevelError:
		lvl = slog.LevelError
	case hlog.LevelFatal:
		lvl = LevelFatal
	default:
		lvl = slog.LevelWarn
	}
	return
}

func NewConfig(opts ...vlog.Option) *vlog.Config {
	lvl := &slog.LevelVar{}
	lvl.Set(hLevelToSLevel(hlog.LevelInfo))
	config := &vlog.Config{
		Level:              lvl,
		HandlerOptions:     &slog.HandlerOptions{Level: lvl},
		WithLevel:          false,
		WithHandlerOptions: false,
		Output:             os.Stdout,
	}
	for _, opt := range opts {
		opt.Apply(config)
	}
	if !config.WithLevel && config.WithHandlerOptions && config.HandlerOptions.Level != nil {
		lvl := &slog.LevelVar{}
		lvl.Set(config.HandlerOptions.Level.Level())
		config.Level = lvl
	}
	config.HandlerOptions.Level = config.Level
	var replaceAttrDefined bool
	if config.HandlerOptions.ReplaceAttr == nil {
		replaceAttrDefined = false
	} else {
		replaceAttrDefined = true
	}
	replaceFun := config.HandlerOptions.ReplaceAttr
	replaceAttr := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.LevelKey {
			level := a.Value.Any().(slog.Level)
			switch level {
			case LevelTrace:
				a.Value = slog.StringValue("Trace")
			case slog.LevelDebug:
				a.Value = slog.StringValue("Debug")
			case slog.LevelInfo:
				a.Value = slog.StringValue("Info")
			case LevelNotice:
				a.Value = slog.StringValue("Notice")
			case slog.LevelWarn:
				a.Value = slog.StringValue("Warn")
			case slog.LevelError:
				a.Value = slog.StringValue("Error")
			case LevelFatal:
				a.Value = slog.StringValue("Fatal")
			default:
				a.Value = slog.StringValue("Warn")
			}
		}
		if replaceAttrDefined {
			return replaceFun(groups, a)
		} else {
			return a
		}
	}
	config.HandlerOptions.ReplaceAttr = replaceAttr
	return config
}

func NewHLog(log *slog.Logger, opts ...vlog.Option) *HLog {
	config := NewConfig(opts...)
	if log == nil {
		log = slog.Default()
	}
	return &HLog{l: log.WithGroup("hlog"), cfg: config}
}
