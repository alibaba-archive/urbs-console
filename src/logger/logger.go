package logger

import (
	"context"
	"os"

	"github.com/mushroomsir/logger/pkg"
	"github.com/teambition/gear"
)

// Default ...
var Default = pkg.New(os.Stderr, pkg.Options{
	EnableFileLine: true,
	EnableJSON:     true,
	Skip:           3,
})

// local ...
var local = pkg.New(os.Stderr, pkg.Options{
	EnableFileLine: true,
	EnableJSON:     true,
	Skip:           4,
})

// SetLevel ...
func SetLevel(level string) *pkg.Logger {
	Default.SetLoggerLevel(level)
	return local.SetLoggerLevel(level)
}

// SetJSONLog set the logger writing JSON string log.
func SetJSONLog() *pkg.Logger {
	Default.SetJSONLog()
	return local.SetJSONLog()
}

// Level ...
func Level() uint32 {
	return local.Level()
}

// GetRequestId ...
func GetRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	gearCtx, ok := ctx.(*gear.Context)
	if !ok {
		return ""
	}
	if s := gearCtx.GetHeader(gear.HeaderXRequestID); s != "" {
		return s
	}
	if gearCtx.Res != nil {
		s := gearCtx.Res.Get(gear.HeaderXRequestID)
		return s
	}
	return ""
}

// Debug ...
func Debug(ctx context.Context, message string, kv ...interface{}) {
	kv = append(kv, "message", message)
	if s := GetRequestId(ctx); s != "" {
		kv = append(kv, "xRequestId", s)
	}
	local.Debug(kv...)
}

// Info ...
func Info(ctx context.Context, message string, kv ...interface{}) {
	kv = append(kv, "message", message)
	if s := GetRequestId(ctx); s != "" {
		kv = append(kv, "xRequestId", s)
	}
	local.Info(kv...)
}

// Warning ...
func Warning(ctx context.Context, message string, kv ...interface{}) {
	kv = append(kv, "message", message)
	if s := GetRequestId(ctx); s != "" {
		kv = append(kv, "xRequestId", s)
	}
	local.Warning(kv...)
}

// Err ...
func Err(ctx context.Context, message string, kv ...interface{}) {
	kv = append(kv, "message", message)
	if s := GetRequestId(ctx); s != "" {
		kv = append(kv, "xRequestId", s)
	}
	local.Err(kv...)
}
