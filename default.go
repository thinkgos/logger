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

// SetNewCallerCore overwrite with new caller core
func SetNewCallerCore(c *CallerCore) *Log {
	return defaultLogger.SetNewCallerCore(c)
}

// AddCallerSkip add the number of callers skipped by caller annotation.
func AddCallerSkip(callerSkip int) *Log {
	return defaultLogger.AddCallerSkip(callerSkip)
}

// AddCallerSkipPackage add the caller skip package.
func AddCallerSkipPackage(vs ...string) *Log {
	return defaultLogger.AddCallerSkipPackage(vs...)
}

// SetCallerLevel set the caller level.
func SetCallerLevel(lv Level) *Log {
	return defaultLogger.SetCallerLevel(lv)
}

// UseExternalCallerLevel use external caller level, which controller by user.
func UseExternalCallerLevel(lvl AtomicLevel) *Log {
	return defaultLogger.UseExternalCallerLevel(lvl)
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

// ExtendDefaultHook set default hook, which hold always until you call [Event.Msg]/[Event.Print]/[Event.Printf].
func ExtendDefaultHook(hs ...Hook) *Log {
	return defaultLogger.ExtendDefaultHook(hs...)
}

// ExtendDefaultHookFunc set default hook, which hold always until you call [Event.Msg]/[Event.Print]/[Event.Printf].
func ExtendDefaultHookFunc(hs ...HookFunc) *Log {
	return defaultLogger.ExtendDefaultHookFunc(hs...)
}

// ExtendHook creates a child log with extend Hook.
func ExtendHook(hs ...Hook) *Log {
	return defaultLogger.ExtendHook(hs...)
}

// ExtendHookFunc creates a child log  with extend Hook.
func ExtendHookFunc(hs ...HookFunc) *Log {
	return defaultLogger.ExtendHookFunc(hs...)
}

// WithNewHook creates a child log with new hook without default hook.
func WithNewHook(hs ...Hook) *Log {
	return defaultLogger.WithNewHook(hs...)
}

// WithNewHookFunc creates a child log with new hook func without default hook.
func WithNewHookFunc(hs ...HookFunc) *Log {
	return defaultLogger.WithNewHookFunc(hs...)
}

// With creates a child log and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
//
// NOTICE: if you do not need a child log, use [Event.With] instead.
func With(fields ...Field) *Log { return defaultLogger.With(fields...) }

// Named adds a sub-scope to the logger's name. See Log.Named for details.
func Named(name string) *Log { return defaultLogger.Named(name) }

// Sync flushes any buffered log entries.
func Sync() error { return defaultLogger.Sync() }

// OnLevel starts a new message with customize level.
//
// You must call Msg on the returned event in order to send the event.

func OnLevel(level Level) *Event {
	return defaultLogger.OnLevel(level)
}

// OnLevelContext starts a new message with customize level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnLevelContext(ctx context.Context, level Level) *Event {
	return defaultLogger.OnLevel(level).WithContext(ctx)
}

// OnDebug starts a new message with [DebugLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func OnDebug() *Event {
	return defaultLogger.OnLevel(DebugLevel)
}

// Debug starts a new message with [DebugLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnDebugContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(DebugLevel).WithContext(ctx)
}

// OnInfo starts a new message with [InfoLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func OnInfo() *Event {
	return defaultLogger.OnLevel(InfoLevel)
}

// OnInfoContext starts a new message with [InfoLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnInfoContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(InfoLevel).WithContext(ctx)
}

// OnWarn starts a new message with [WarnLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func OnWarn() *Event {
	return defaultLogger.OnLevel(WarnLevel)
}

// OnWarnContext starts a new message with [WarnLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnWarnContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(WarnLevel).WithContext(ctx)
}

// OnError starts a new message with [ErrorLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func OnError() *Event {
	return defaultLogger.OnLevel(ErrorLevel)
}

// OnErrorContext starts a new message with [ErrorLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnErrorContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(ErrorLevel).WithContext(ctx)
}

// OnDPanic starts a new message with [DPanicLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func OnDPanic() *Event {
	return defaultLogger.OnLevel(DPanicLevel)
}

// OnDPanicContext starts a new message with [DPanicLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnDPanicContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(DPanicLevel).WithContext(ctx)
}

// OnPanic starts a new message with [PanicLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func OnPanic() *Event {
	return defaultLogger.OnLevel(PanicLevel)
}

// OnPanicContext starts a new message with [PanicLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnPanicContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(PanicLevel).WithContext(ctx)
}

// OnFatal starts a new message with [FatalLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func OnFatal() *Event {
	return defaultLogger.OnLevel(FatalLevel)
}

// OnFatalContext starts a new message with [FatalLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func OnFatalContext(ctx context.Context) *Event {
	return defaultLogger.OnLevel(FatalLevel).WithContext(ctx)
}
