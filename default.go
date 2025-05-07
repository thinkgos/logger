package logger

import (
	"context"

	"go.uber.org/zap"
)

var defaultLogger = NewLoggerWith(zap.NewNop(), zap.NewAtomicLevel())

// ReplaceGlobals replaces the global Log only once.
func ReplaceGlobals(logger *Log) { defaultLogger = logger }

// UnderlyingLogger underlying global logger.
func UnderlyingLogger() *Log { return defaultLogger }

// AddCallerSkip add the number of callers skipped by caller annotation.
func AddCallerSkip(callerSkip int) *Log {
	defaultLogger.AddCallerSkip(callerSkip)
	return defaultLogger
}

// AddCallerSkipPackage add the caller skip package.
func AddCallerSkipPackage(vs ...string) *Log {
	defaultLogger.AddCallerSkipPackage(vs...)
	return defaultLogger
}

// SetCallerLevel set the caller level.
func SetCallerLevel(lv Level) *Log {
	defaultLogger.SetCallerLevel(lv)
	return defaultLogger
}

// UseExternalCallerLevel use external caller level, which controller by user.
func UseExternalCallerLevel(lvl AtomicLevel) *Log {
	defaultLogger.UseExternalCallerLevel(lvl)
	return defaultLogger
}

// UnderlyingCallerLevel get underlying caller level.
func UnderlyingCallerLevel() AtomicLevel {
	return defaultLogger.UnderlyingCallerLevel()
}

// SetLevelWithText alters the logging level.
// ParseAtomicLevel set the logging level based on a lowercase or all-caps ASCII
// representation of the log level.
// If the provided ASCII representation is
// invalid an error is returned.
func SetLevelWithText(text string) error { return defaultLogger.SetLevelWithText(text) }

// SetLevel alters the logging level.
func SetLevel(lv Level) *Log { return defaultLogger.SetLevel(lv) }

// GetLevel returns the minimum enabled log level.
func GetLevel() Level { return defaultLogger.GetLevel() }

// Enabled returns true if the given level is at or above this level.
func Enabled(lvl Level) bool { return defaultLogger.Enabled(lvl) }

// V returns true if the given level is at or above this level.
// same as Enabled
func V(lvl Level) bool { return defaultLogger.V(lvl) }

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func Sugar() *zap.SugaredLogger { return defaultLogger.Sugar() }

// Logger return internal logger
func Logger() *zap.Logger { return defaultLogger.Logger() }

// SetDefaultHook set default hook, which hold always until you call XXXContext.
func SetDefaultHook(hs ...Hook) *Log {
	return defaultLogger.SetDefaultHook(hs...)
}

// SetDefaultHookFunc set default hook, which hold always until you call XXXContext.
func SetDefaultHookFunc(hs ...HookFunc) *Log {
	return defaultLogger.SetDefaultHookFunc(hs...)
}

// ExtendHook return new Log with extend Hook.
func ExtendHook(hs ...Hook) *Log {
	return defaultLogger.ExtendHook(hs...)
}

// ExtendHookFunc return new Log with extend Hook.
func ExtendHookFunc(hs ...HookFunc) *Log {
	return defaultLogger.ExtendHookFunc(hs...)
}

// WithNewHook return new log with new hook without default hook.
func WithNewHook(hs ...Hook) *Log {
	return defaultLogger.WithNewHook(hs...)
}

// WithNewHookFunc return new log with new hook func without default hook.
func WithNewHookFunc(hs ...HookFunc) *Log {
	return defaultLogger.WithNewHookFunc(hs...)
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//
//	 sugaredLogger.With(
//	   "hello", "world",
//	   "failure", errors.New("oh no"),
//	   "count", 42,
//	   "user", User{Name: "alice"},
//	)
//
// is the equivalent of
//
//	unsugared.With(
//	  String("hello", "world"),
//	  String("failure", "oh no"),
//	  Stack(),
//	  Int("count", 42),
//	  Object("user", User{Name: "alice"}),
//	)
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func With(fields ...Field) *Log { return defaultLogger.With(fields...) }

// Named adds a sub-scope to the logger's name. See Log.Named for details.
func Named(name string) *Log { return defaultLogger.Named(name) }

// Sync flushes any buffered log entries.
func Sync() error { return defaultLogger.Sync() }

func OnLevel(level Level) *Event {
	return defaultLogger.OnLevel(level)
}

func OnLevelContext(ctx context.Context, level Level) *Event {
	return defaultLogger.OnLevel(level).WithContext(ctx)
}

func OnDebug() *Event {
	return defaultLogger.OnLevel(DebugLevel)
}

func OnDebugContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(DebugLevel).WithContext(ctx)
}
func OnInfo() *Event {
	return defaultLogger.OnLevel(InfoLevel)
}
func OnInfoContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(InfoLevel).WithContext(ctx)
}
func OnWarn() *Event {
	return defaultLogger.OnLevel(WarnLevel)
}
func OnWarnContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(WarnLevel).WithContext(ctx)
}
func OnError() *Event {
	return defaultLogger.OnLevel(ErrorLevel)
}
func OnErrorContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(ErrorLevel).WithContext(ctx)
}
func OnDPanic() *Event {
	return defaultLogger.OnLevel(DPanicLevel)
}
func OnDPanicContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(DPanicLevel).WithContext(ctx)
}
func OnPanic() *Event {
	return defaultLogger.OnLevel(PanicLevel)
}
func OnPanicContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(PanicLevel).WithContext(ctx)
}
func OnFatal() *Event {
	return defaultLogger.OnLevel(FatalLevel)
}
func OnFatalContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(FatalLevel).WithContext(ctx)
}
