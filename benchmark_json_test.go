package logger_test

import (
	"context"
	"io"
	"testing"

	"github.com/thinkgos/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var testNativeZapEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "ts",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	FunctionKey:    zapcore.OmitKey,
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func newDiscardLogger(format string) *logger.Log {
	return logger.NewLogger(
		logger.WithAdapter("custom", io.Discard),
		logger.WithFormat(format),
	)
}
func dfltCtx(ctx context.Context) logger.Field {
	return zap.String("dflt_key", "dflt_value")
}

func newDfltHook() logger.Hook {
	return &logger.ImmutableString{
		Key:   "dflt_key",
		Value: "dflt_value",
	}
}

func Benchmark_Json_NativeLogger(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(testNativeZapEncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	l := zap.New(core)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.Info("success",
			zap.String("name", "jack"),
			zap.Int("age", 18),
			dfltCtx(ctx),
		)
	}
}

func Benchmark_Json_Logger(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			String("name", "jack").
			Int("age", 18).
			With(dfltCtx(ctx)).
			Msg("success")
	}
}

func Benchmark_Json_Logger_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson)
	l.ExtendDefaultHook(newDfltHook())
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			String("name", "jack").
			Int("age", 18).
			Msg("success")
	}
}

func Benchmark_Json_NativeSugar(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(testNativeZapEncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	l := zap.New(core).Sugar()
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.Infow("success",
			"name", "jack",
			"age", 18,
			dfltCtx(ctx),
		)
	}
}

func Benchmark_Json_Use_WithFields(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			With(
				logger.String("name", "jack"),
				logger.Int("age", 18),
				dfltCtx(ctx),
			).
			Msg("success")
	}
}

func Benchmark_Json_Use_WithFields_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson)
	l.ExtendDefaultHook(newDfltHook())
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			With(
				logger.String("name", "jack"),
				logger.Int("age", 18),
			).
			Msg("success")
	}
}

func Benchmark_Json_Use_ExtendHook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			ExtendHook(
				&logger.ImmutableString{"name", "jack"},
				&logger.ImmutableInt{"age", 18},
				newDfltHook(),
			).
			Msg("success")
	}
}

func Benchmark_Json_Use_ExtendHook_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson)
	l.ExtendDefaultHook(newDfltHook())
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			ExtendHook(
				&logger.ImmutableString{"name", "jack"},
				&logger.ImmutableInt{"age", 18},
				newDfltHook(),
			).
			Msg("success")
	}
}

func Benchmark_Json_Format(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson).
		ExtendDefaultHook(
			newDfltHook(),
			logger.ImmutString("name", "jack"),
			logger.ImmutInt("age", 18),
		)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			Printf("success: %s", "ok")
	}
}

func Benchmark_Json_Format_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatJson).
		ExtendDefaultHook(
			newDfltHook(),
			logger.ImmutString("name", "jack"),
			logger.ImmutInt("age", 18),
		)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			Printf("success: %s", "ok")
	}
}
