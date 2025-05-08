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
			dfltCtx(ctx),
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
			With(dfltCtx(ctx)).
			Msg("success")
	}
}

func Benchmark_Text_Logger_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.ExtendDefaultHookFunc(dfltCtx)
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
			dfltCtx(ctx),
		)
	}
}

func Benchmark_Text_Use_WithFields(b *testing.B) {
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
				dfltCtx(ctx),
			).
			Msg("success")
	}
}

func Benchmark_Text_Use_WithFields_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.ExtendDefaultHookFunc(dfltCtx)
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
	l := newDiscardLogger(logger.FormatConsole)
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

func Benchmark_Text_Use_ExtendHook_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.ExtendDefaultHookFunc(dfltCtx)
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

func Benchmark_Text_Format(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			ExtendHookFunc(
				func(ctx context.Context) logger.Field {
					return logger.String("name", "jack")
				},
				func(ctx context.Context) logger.Field {
					return logger.Int("age", 18)
				},
				dfltCtx,
			).
			Printf("success")
	}
}

func Benchmark_Text_Format_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.ExtendDefaultHookFunc(dfltCtx)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.OnInfoContext(ctx).
			ExtendHookFunc(
				logger.ImmutString("name", "jack"),
				logger.ImmutInt("age", 18),
			).
			Msg("success")
	}
}
