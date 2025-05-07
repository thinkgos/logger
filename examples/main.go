package main

import (
	"context"

	"github.com/thinkgos/logger"
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
	logger.SetDefaultHookFunc(ExampleHook)

	ctx := context.WithValue(context.Background(), ctxKey{}, "ctx_val1")

	logger.OnDebugContext(ctx).ExtendHookFunc(TmpHook).Msg("Debug1")
	logger.OnDebugContext(ctx).WithNewHook(logger.HookFunc(TmpHook)).Msg("Debug2")

	logger.SetLevel(logger.WarnLevel)
	logger.OnDebugContext(ctx).Msg("Debug3")

	err := logger.SetLevelWithText(logger.DebugLevel.String())
	_ = err
	logger.OnDebugContext(ctx).Msg("Debug4")

	// 这里改成warn等级方便测试
	logger.SetLevel(logger.WarnLevel)
	if logger.Enabled(logger.InfoLevel) { // 提前过滤
		logger.OnInfoContext(ctx).Msg("Info1") // 不执行
	}
	if logger.V(logger.InfoLevel) { // 提前过滤
		logger.OnInfoContext(ctx).Msg("Info2") // 不执行
	}
	logger.SetLevel(logger.DebugLevel)

	logger.Named("叫个名字").OnDebugContext(ctx).Msg(" Debug5")

	logger.OnInfoContext(ctx).
		With(
			logger.String("name", "jack"),
			logger.Int("age", 18),
		).
		Msg(" Debug6")

	// `log.Print`风格的日志
	logger.OnDebugContext(ctx).Print("Debug")
	logger.OnInfoContext(ctx).Print("Info")
	logger.OnWarnContext(ctx).Print("Warn")
	logger.OnErrorContext(ctx).Print("Error")
	logger.OnDPanicContext(ctx).Print("DPanic")

	// `log.Printf`风格的日志
	logger.OnDebugContext(ctx).Printf("Debugf: %s", "debug")
	logger.OnInfoContext(ctx).Printf("Infof: %s", "info")
	logger.OnWarnContext(ctx).Printf("Warnf: %s", "warn")
	logger.OnErrorContext(ctx).Printf("Errorf: %s", "error")
	logger.OnDPanicContext(ctx).Printf("DPanicf: %s", "dpanic")

	// 松散键值对风格的日志
	logger.OnDebugContext(ctx).Pairs("k1", "v1", logger.String("k2", "v2")).Msg("Debug Pair")
	logger.OnInfoContext(ctx).Pairs("k1", "v1", logger.String("k2", "v2")).Msg("Info Pair")
	logger.OnWarnContext(ctx).Pairs("k1", "v1", logger.String("k2", "v2")).Msg("Warn Pair")
	logger.OnInfoContext(ctx).Pairs("k1", "v1", logger.String("k2", "v2")).Msg("Info Pair")
	logger.OnErrorContext(ctx).Pairs("k1", "v1", logger.String("k2", "v2")).Msg("Error Pair")
	logger.OnDPanicContext(ctx).Pairs("k1", "v1", logger.String("k2", "v2")).Msg("DPanic Pair")

	// 纯结构型的日志
	logger.OnDebugContext(ctx).String("k1", "v1").String("k2", "v2").Msg("Debug")
	logger.OnInfoContext(ctx).String("k1", "v1").String("k2", "v2").Msg("Info")
	logger.OnWarnContext(ctx).String("k1", "v1").String("k2", "v2").Msg("Warn")
	logger.OnInfoContext(ctx).String("k1", "v1").String("k2", "v2").Msg("Info")
	logger.OnErrorContext(ctx).String("k1", "v1").String("k2", "v2").Msg("Error")
	logger.OnDPanicContext(ctx).String("k1", "v1").String("k2", "v2").Msg("DPanic")
}
