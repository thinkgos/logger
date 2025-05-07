package logger_test

import (
	"context"
	"testing"

	"github.com/thinkgos/logger"
)

func Benchmark_Text_Logger_Deprecated(b *testing.B) {
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

func Benchmark_Text_Logger_Use_Hook_Deprecated(b *testing.B) {
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

func Benchmark_Text_KeyValuePair_Deprecated(b *testing.B) {
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

func Benchmark_Text_KeyValuePairFields_Deprecated(b *testing.B) {
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

func Benchmark_Text_KeyValuePairFields_Use_Hook_Deprecated(b *testing.B) {
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

func Benchmark_Text_KeyValuePairFields_Use_WithFields_Deprecated(b *testing.B) {
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

func Benchmark_Text_KeyValuePairFields_Use_WithFields_Hook_Deprecated(b *testing.B) {
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

func Benchmark_Text_KeyValuePairFields_Use_WithValuer_Deprecated(b *testing.B) {
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

func Benchmark_Text_KeyValuePairFields_Use_WithValuer_Hook_Deprecated(b *testing.B) {
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

func Benchmark_Text_Format_Deprecated(b *testing.B) {
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

func Benchmark_Text_Format_Use_Hook_Deprecated(b *testing.B) {
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
