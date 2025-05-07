package logger_test

import (
	"context"
	"testing"

	"github.com/thinkgos/logger"
)

func init() {
	l, lv := logger.New(logger.WithLevel(logger.DebugLevel.String()))
	logger.ReplaceGlobals(logger.NewLoggerWith(l, lv))
	logger.SetDefaultHookFunc(
		logger.Caller(3),
		func(ctx context.Context) logger.Field {
			return logger.String("deft_key1", "deft_val1")
		},
	)
}

func Test_LoggerNormal(t *testing.T) {
	logger.OnDebug().Msg("Debug")
	logger.OnInfo().Msg("Info")
	logger.OnWarn().Msg("Warn")
	logger.OnError().Msg("Error")
	logger.OnDPanic().Msg("DPanic")
}

func Test_LoggerFormater(t *testing.T) {
	logger.OnDebug().Printf("Debugf: %s", "debug")
	logger.OnInfo().Printf("Infof: %s", "info")
	logger.OnWarn().Printf("Warnf: %s", "warn")
	logger.OnError().Printf("Errorf: %s", "error")
	logger.OnDPanic().Printf("DPanicf: %s", "dpanic")
}

func Test_LoggerKeyValue(t *testing.T) {
	logger.OnDebug().Pairs("Debug", "pair").Msg("Debug")
	logger.OnInfo().Pairs("Info", "pair").Msg("Info")
	logger.OnWarn().Pairs("Warn", "pair").Msg("Warn")
	logger.OnInfo().Pairs("Info", "pair").Msg("Info")
	logger.OnError().Pairs("Error", "pair").Msg("Error")
	logger.OnDPanic().Pairs("DPanic", "pair").Msg("DPanic")
}

func TestPanic(t *testing.T) {
	shouldPanic(t, func() {
		logger.OnPanic().Msg("Panic")
	})
	shouldPanic(t, func() {
		logger.OnPanic().Printf("Panicf: %s", "panic")
	})
	shouldPanic(t, func() {
		logger.OnPanic().Pairs("panic", "pair").Msg("Panic: pair")
	})
}

func Test_LoggerWith(t *testing.T) {
	logger.OnDebug().
		With(
			logger.String("string", "bb"),
			logger.Int16("int16", 100),
		).
		Msg("debug with")
}

func Test_LoggerNamed(t *testing.T) {
	logger.Named("another").OnDebug().Msg("debug named")
}
func Test_Logger_ZapLogger(t *testing.T) {
	logger.Logger().Debug("desugar")
}

func Test_LoggerNamespace(t *testing.T) {
	logger.Logger().
		With(logger.Namespace("aaaa")).
		With(logger.String("xx", "yy"), logger.String("aa", "bb")).
		Debug("with namespace")

	_ = logger.Sync()
}

type ctxKey struct{}

func Test_Logger_Context(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "ctx_value")
	ctxValuer := func(ctx context.Context) logger.Field {
		s, ok := ctx.Value(ctxKey{}).(string)
		if !ok {
			return logger.Skip()
		}
		return logger.String("ctx_key", s)
	}
	logger.OnDebugContext(ctx).
		ExtendHookFunc(ctxValuer).
		Msg("with context")
}

func Test_Logger_Caller(t *testing.T) {
	log := logger.NewLogger(logger.WithLevel(logger.DebugLevel.String()))

	log.OnError().Msg("error")
	log.OnDebug().Msg("debug")

	log.SetCallerLevel(logger.DebugLevel)

	log.OnError().Msg("error")
	log.OnDebug().Msg("debug")
}

func shouldPanic(t *testing.T, f func()) {
	defer func() {
		e := recover()
		if e == nil {
			t.Errorf("should panic but not")
		}
	}()
	f()
}
