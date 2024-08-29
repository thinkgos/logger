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

// SetDefaultValuer set default field function, which hold always until you call XXXContext.
// suggest
func SetDefaultValuer(vs ...Valuer) *Log {
	return defaultLogger.SetDefaultValuer(vs...)
}

// WithValuer with field function.
func WithValuer(vs ...Valuer) *Log {
	return defaultLogger.WithValuer(vs...)
}

// WithNewValuer return log with new Valuer function without default Valuer.
func WithNewValuer(fs ...Valuer) *Log {
	return defaultLogger.WithNewValuer(fs...)
}

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func Sugar() *zap.SugaredLogger { return defaultLogger.Sugar() }

// Logger return internal logger
func Logger() *zap.Logger { return defaultLogger.Logger() }

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

// ****** named after the log level or ending in "Context" for log.Print-style logging

func Debug(args ...any) {
	defaultLogger.Debug(args...)
}
func DebugContext(ctx context.Context, args ...any) {
	defaultLogger.DebugContext(ctx, args...)
}
func Info(args ...any) {
	defaultLogger.Info(args...)
}
func InfoContext(ctx context.Context, args ...any) {
	defaultLogger.InfoContext(ctx, args...)
}
func Warn(args ...any) {
	defaultLogger.Warn(args...)
}
func WarnContext(ctx context.Context, args ...any) {
	defaultLogger.WarnContext(ctx, args...)
}
func Error(args ...any) {
	defaultLogger.Error(args...)
}
func ErrorContext(ctx context.Context, args ...any) {
	defaultLogger.ErrorContext(ctx, args...)
}
func DPanic(args ...any) {
	defaultLogger.DPanic(args...)
}
func DPanicContext(ctx context.Context, args ...any) {
	defaultLogger.DPanicContext(ctx, args...)
}
func Panic(args ...any) {
	defaultLogger.Panic(args...)
}
func PanicContext(ctx context.Context, args ...any) {
	defaultLogger.PanicContext(ctx, args...)
}
func Fatal(args ...any) {
	defaultLogger.Fatal(args...)
}
func FatalContext(ctx context.Context, args ...any) {
	defaultLogger.FatalContext(ctx, args...)
}

// ****** ending in "f" or "fContext" for log.Printf-style logging

func Debugf(template string, args ...any) {
	defaultLogger.Debugf(template, args...)
}
func DebugfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.DebugfContext(ctx, template, args...)
}
func Infof(template string, args ...any) {
	defaultLogger.Infof(template, args...)
}
func InfofContext(ctx context.Context, template string, args ...any) {
	defaultLogger.InfofContext(ctx, template, args...)
}
func Warnf(template string, args ...any) {
	defaultLogger.Warnf(template, args...)
}
func WarnfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.WarnfContext(ctx, template, args...)
}
func Errorf(template string, args ...any) {
	defaultLogger.Errorf(template, args...)
}
func ErrorfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.ErrorfContext(ctx, template, args...)
}
func DPanicf(template string, args ...any) {
	defaultLogger.DPanicf(template, args...)
}
func DPanicfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.DPanicfContext(ctx, template, args...)
}
func Panicf(template string, args ...any) {
	defaultLogger.Panicf(template, args...)
}
func PanicfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.PanicfContext(ctx, template, args...)
}
func Fatalf(template string, args ...any) {
	defaultLogger.Fatalf(template, args...)
}
func FatalfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.FatalfContext(ctx, template, args...)
}

// ****** ending in "w" or "wContext" for loosely-typed structured logging

func Debugw(msg string, keysAndValues ...any) {
	defaultLogger.Debugw(msg, keysAndValues...)
}
func DebugwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.DebugwContext(ctx, msg, keysAndValues...)
}
func Infow(msg string, keysAndValues ...any) {
	defaultLogger.Infow(msg, keysAndValues...)
}
func InfowContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.InfowContext(ctx, msg, keysAndValues...)
}
func Warnw(msg string, keysAndValues ...any) {
	defaultLogger.Warnw(msg, keysAndValues...)
}
func WarnwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.WarnwContext(ctx, msg, keysAndValues...)
}
func Errorw(msg string, keysAndValues ...any) {
	defaultLogger.Errorw(msg, keysAndValues...)
}
func ErrorwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.ErrorwContext(ctx, msg, keysAndValues...)
}
func DPanicw(msg string, keysAndValues ...any) {
	defaultLogger.DPanicw(msg, keysAndValues...)
}
func DPanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.DPanicwContext(ctx, msg, keysAndValues...)
}
func Panicw(msg string, keysAndValues ...any) {
	defaultLogger.Panicw(msg, keysAndValues...)
}
func PanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.PanicwContext(ctx, msg, keysAndValues...)
}
func Fatalw(msg string, keysAndValues ...any) {
	defaultLogger.Fatalw(msg, keysAndValues...)
}
func FatalwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.FatalwContext(ctx, msg, keysAndValues...)
}

// ****** ending in "x" or "xContext" for structured logging

func Debugx(msg string, fields ...Field) {
	defaultLogger.Debugx(msg, fields...)
}
func DebugxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.DebugxContext(ctx, msg, fields...)
}
func Infox(msg string, fields ...Field) {
	defaultLogger.Infox(msg, fields...)
}
func InfoxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.InfoxContext(ctx, msg, fields...)
}
func Warnx(msg string, fields ...Field) {
	defaultLogger.Warnx(msg, fields...)
}
func WarnxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.WarnxContext(ctx, msg, fields...)
}
func Errorx(msg string, fields ...Field) {
	defaultLogger.Errorx(msg, fields...)
}
func ErrorxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.ErrorxContext(ctx, msg, fields...)
}
func DPanicx(msg string, fields ...Field) {
	defaultLogger.DPanicx(msg, fields...)
}
func DPanicxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.DPanicxContext(ctx, msg, fields...)
}
func Panicx(msg string, fields ...Field) {
	defaultLogger.Panicx(msg, fields...)
}
func PanicxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.PanicxContext(ctx, msg, fields...)
}
func Fatalx(msg string, fields ...Field) {
	defaultLogger.Fatalx(msg, fields...)
}
func FatalxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.FatalxContext(ctx, msg, fields...)
}
