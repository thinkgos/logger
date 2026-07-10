package logger_test

import (
	"context"
	"io"
	"testing"

	"github.com/thinkgos/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Benchmark_Text_NativeLogger(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(testNativeZapEncoderConfig),
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
			dfltHookField(ctx),
		)
	}
}

func Benchmark_Text_Logger(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			String("name", "jack").
			Int("age", 18).
			With(dfltHookField(ctx)).
			Msg("success")
	}
}

func Benchmark_Text_Logger_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.ExtendDefaultHookField(dfltHookField)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			String("name", "jack").
			Int("age", 18).
			Msg("success")
	}
}

func Benchmark_Text_NativeSugar(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(testNativeZapEncoderConfig),
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
			dfltHookField(ctx),
		)
	}
}

func Benchmark_Text_Use_With(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			With(
				logger.String("name", "jack"),
				logger.Int("age", 18),
				dfltHookField(ctx),
			).
			Msg("success")
	}
}

func Benchmark_Text_Use_With_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole).
		ExtendDefaultHookField(dfltHookField)
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

func Benchmark_Text_Use_ExtendHook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole).
		ExtendHook(
			&ImmutableMetadata{name: "jack", age: 18},
			&ImmutableDflt{},
		)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			Msg("success")
	}
}

func Benchmark_Text_Format(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole).
		ExtendDefaultHook(
			&ImmutableDflt{},
			&ImmutableMetadata{name: "jack", age: 18},
		)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			Printf("success: %s", "ok")
	}
}

func Benchmark_Text_Format_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole).
		ExtendDefaultHook(
			&ImmutableDflt{},
			&ImmutableMetadata{name: "jack", age: 18},
		)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			Printf("success: %s", "ok")
	}
}
