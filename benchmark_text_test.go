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
		l.InfoxContext(
			ctx,
			"success",
			logger.String("name", "jack"),
			logger.Int("age", 18),
			dfltCtx(ctx),
		)
	}
}

func Benchmark_Text_Logger_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.SetDefaultValuer(dfltCtx)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.InfoxContext(
			ctx,
			"success",
			logger.String("name", "jack"),
			logger.Int("age", 18),
		)
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

func Benchmark_Text_KeyValuePair(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.InfowContext(ctx,
			"success",
			"name", "jack",
			"age", 18,
			dfltCtx(ctx),
		)
	}
}

func Benchmark_Text_KeyValuePairFields(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.InfowContext(ctx,
			"success",
			logger.String("name", "jack"),
			logger.Int("age", 18),
			dfltCtx(ctx),
		)
	}
}

func Benchmark_Text_KeyValuePairFields_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.SetDefaultValuer(dfltCtx)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.InfowContext(ctx,
			"success",
			logger.String("name", "jack"),
			logger.Int("age", 18),
		)
	}
}

func Benchmark_Text_KeyValuePairFields_Use_WithFields(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.With(
			logger.String("name", "jack"),
			logger.Int("age", 18),
			dfltCtx(ctx),
		).InfowContext(ctx, "success")
	}
}

func Benchmark_Text_KeyValuePairFields_Use_WithFields_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.SetDefaultValuer(dfltCtx)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.With(
			logger.String("name", "jack"),
			logger.Int("age", 18),
		).InfowContext(ctx, "success")
	}
}

func Benchmark_Text_KeyValuePairFields_Use_WithValuer(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.WithValuer(
			logger.ImmutString("name", "jack"),
			logger.ImmutInt("age", 18),
			dfltCtx,
		).InfowContext(ctx, "success")
	}
}

func Benchmark_Text_KeyValuePairFields_Use_WithValuer_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.SetDefaultValuer(dfltCtx)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.WithValuer(
			logger.ImmutString("name", "jack"),
			logger.ImmutInt("age", 18),
		).InfowContext(ctx, "success")
	}
}

func Benchmark_Text_Format(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.WithValuer(
			func(ctx context.Context) logger.Field {
				return logger.String("name", "jack")
			},
			func(ctx context.Context) logger.Field {
				return logger.Int("age", 18)
			},
			dfltCtx,
		).InfofContext(ctx,
			"success",
		)
	}
}

func Benchmark_Text_Format_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	l := newDiscardLogger(logger.FormatConsole)
	l.SetDefaultValuer(dfltCtx)
	ctx := context.Background()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l.WithValuer(
			logger.ImmutString("name", "jack"),
			logger.ImmutInt("age", 18),
		).InfofContext(ctx,
			"success",
		)
	}
}
