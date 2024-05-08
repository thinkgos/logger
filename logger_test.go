package logger_test

import (
	"context"
	"testing"

	"github.com/things-go/logger"
)

func init() {
	l, lv := logger.New(logger.WithLevel(logger.DebugLevel.String()))
	logger.ReplaceGlobals(logger.NewLoggerWith(l, lv))
	logger.SetDefaultValuer(
		logger.Caller(3),
		func(ctx context.Context) logger.Field {
			return logger.String("deft_key1", "deft_val1")
		},
	)
}

func Test_LoggerNormal(t *testing.T) {
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
	logger.DPanic("DPanic")
}

func Test_LoggerFormater(t *testing.T) {
	logger.Debugf("Debugf: %s", "debug")
	logger.Infof("Infof: %s", "info")
	logger.Warnf("Warnf: %s", "warn")
	logger.Errorf("Errorf: %s", "error")
	logger.DPanicf("DPanicf: %s", "dPanic")
}

func Test_LoggerKeyValue(t *testing.T) {
	logger.Debugw("Debugw", "Debugw", "w")
	logger.Infow("Infow", "Infow", "w")
	logger.Warnw("Warnw", "Warnw", "w")
	logger.Infow("Infow", "Infow", "w")
	logger.Errorw("Errorw", "Errorw", "w")
	logger.DPanicw("DPanicw", "DPanicw", "w")
}

func TestPanic(t *testing.T) {
	shouldPanic(t, func() {
		logger.Panic("Panic")
	})
	shouldPanic(t, func() {
		logger.Panicf("Panicf: %s", "panic")
	})
	shouldPanic(t, func() {
		logger.Panicw("Panicw: %s", "panic", "w")
	})
}

func Test_LoggerWith(t *testing.T) {
	logger.With(
		logger.String("string", "bb"),
		logger.Int16("int16", 100),
	).
		Debug("debug with")
}

func Test_LoggerNamed(t *testing.T) {
	logger.Named("another").Debug("debug named")
}
func Test_Logger_ZapLogger(t *testing.T) {
	logger.Logger().Debug("desugar")
}

func Test_LoggerNamespace(t *testing.T) {
	logger.Logger().With(logger.Namespace("aaaa")).With(logger.String("xx", "yy"), logger.String("aa", "bb")).Debug("with namespace")

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
	logger.WithValuer(ctxValuer).
		DebugContext(ctx, "with context")
}

func Test_Logger_Caller(t *testing.T) {
	log := logger.NewLogger(logger.WithLevel(logger.DebugLevel.String()))

	log.Error("error")
	log.Debug("debug")

	log.SetCallerLevel(logger.DebugLevel)

	log.Error("error")
	log.Debug("debug")
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
