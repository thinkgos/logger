package main

import (
	"context"

	"github.com/things-go/logger"
)

type ctxKey struct{}

func ExampleHook(ctx context.Context) logger.Field {
	s, ok := ctx.Value(ctxKey{}).(string)
	if !ok {
		return logger.Skip()
	}
	return logger.String("ctx_key1", s)
}
func TmpHook(ctx context.Context) logger.Field {
	return logger.String("tmp_key1", "tmp_val1")
}

func main() {
	l := logger.NewLogger(
		logger.WithLevel(logger.DebugLevel.String()),
		logger.WithFormat(logger.FormatJson),
	).
		SetCallerLevel(logger.WarnLevel)
	logger.ReplaceGlobals(l)
	logger.SetDefaultValuer(ExampleHook)

	ctx := context.WithValue(context.Background(), ctxKey{}, "ctx_val1")

	logger.WithValuer(TmpHook).DebugContext(ctx, "Debug1")
	logger.WithNewValuer(TmpHook).DebugContext(ctx, "Debug2")

	logger.SetLevel(logger.WarnLevel)
	logger.DebugContext(ctx, "Debug3")

	err := logger.SetLevelWithText(logger.DebugLevel.String())
	_ = err
	logger.DebugContext(ctx, "Debug4")

	// 这里改成warn等级方便测试
	logger.SetLevel(logger.WarnLevel)
	if logger.Enabled(logger.InfoLevel) { // 提前过滤
		logger.InfoContext(ctx, "Info1") // 不执行
	}
	if logger.V(logger.InfoLevel) { // 提前过滤
		logger.InfoContext(ctx, "Info2") // 不执行
	}
	logger.SetLevel(logger.DebugLevel)

	logger.Named("叫个名字").DebugContext(ctx, " Debug5")

	logger.With(
		logger.String("name", "jack"),
		logger.Int("age", 18),
	).InfoContext(ctx, " Debug6")

	// `log.Print`风格的日志
	logger.DebugContext(ctx, "Debug")
	logger.InfoContext(ctx, "Info")
	logger.WarnContext(ctx, "Warn")
	logger.ErrorContext(ctx, "Error")
	logger.DPanicContext(ctx, "DPanic")

	// `log.Printf`风格的日志
	logger.DebugfContext(ctx, "Debugf: %s", "debug")
	logger.InfofContext(ctx, "Infof: %s", "info")
	logger.WarnfContext(ctx, "Warnf: %s", "warn")
	logger.ErrorfContext(ctx, "Errorf: %s", "error")
	logger.DPanicfContext(ctx, "DPanicf: %s", "dPanic")

	// 松散键值对风格的日志
	logger.DebugwContext(ctx, "Debugw", "k1", "v1", logger.String("k2", "v2"))
	logger.InfowContext(ctx, "Infow", "k1", "v1", logger.String("k2", "v2"))
	logger.WarnwContext(ctx, "Warnw", "k1", "v1", logger.String("k2", "v2"))
	logger.InfowContext(ctx, "Infow", "k1", "v1", logger.String("k2", "v2"))
	logger.ErrorwContext(ctx, "Errorw", "k1", "v1", logger.String("k2", "v2"))
	logger.DPanicwContext(ctx, "DPanicw", "k1", "v1", logger.String("k2", "v2"))

	// 纯结构型的日志
	logger.DebugxContext(ctx, "Debugx", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.InfoxContext(ctx, "Infox", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.WarnxContext(ctx, "Warnx", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.InfoxContext(ctx, "Infox", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.ErrorxContext(ctx, "Errorx", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.DPanicxContext(ctx, "DPanicx", logger.String("k1", "v1"), logger.String("k2", "v2"))
}
