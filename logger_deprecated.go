package logger

import (
	"context"
)

// SetDefaultValuer set default Valuer function, which hold always until you call XXXContext.
//
// Deprecated: As of 0.4.0, use [Log.SetDefaultHookFunc]/[Log.SetDefaultHook] instead.
func (l *Log) SetDefaultValuer(vs ...Valuer) *Log {
	return l.SetDefaultHookFunc(vs...)
}

// WithValuer return new Log with extend Valuer.
//
// Deprecated: As of 0.4.0, use [Log.ExtendHookFunc]/[Log.ExtendHook] instead.
func (l *Log) WithValuer(vs ...Valuer) *Log {
	return l.ExtendHookFunc(vs...)
}

// WithNewValuer return new log with new Valuer without default Valuer.
//
// Deprecated: As of 0.4.0, use [Log.WithNewHookFunc]/[Log.WithNewHook] instead.
func (l *Log) WithNewValuer(vs ...Valuer) *Log {
	return l.WithNewHookFunc(vs...)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).Print(args...) instead.
func (l *Log) Log(ctx context.Context, level Level, args ...any) {
	l.OnLevelContext(ctx, level).Print(args...)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).Printf(template, args...) instead.
func (l *Log) Logf(ctx context.Context, level Level, template string, args ...any) {
	l.OnLevelContext(ctx, level).Printf(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) Logw(ctx context.Context, level Level, msg string, keysAndValues ...any) {
	l.OnLevelContext(ctx, level).Pairs(keysAndValues...).Msg(msg)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).With(fields...).Msg(msg) instead.
func (l *Log) Logx(ctx context.Context, level Level, msg string, fields ...Field) {
	l.OnLevelContext(ctx, level).With(fields...).Msg(msg)
}

// ****** named after the log level or ending in "Context" for log.Print-style logging

// Debug (see DebugContext)
//
// Deprecated: As of 0.4.0, use Log.OnDebug().Print(args...) instead.
func (l *Log) Debug(args ...any) {
	l.OnDebug().Print(args...)
}

// DebugContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).Print(args...) instead.
func (l *Log) DebugContext(ctx context.Context, args ...any) {
	l.OnDebugContext(ctx).Print(args...)
}

// Info see InfoContext
//
// Deprecated: As of 0.4.0, use Log.OnInfo().Print(args...) instead.
func (l *Log) Info(args ...any) {
	l.OnInfo().Print(args...)
}

// InfoContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).Print(args...) instead.
func (l *Log) InfoContext(ctx context.Context, args ...any) {
	l.OnInfoContext(ctx).Print(args...)
}

// Warn see WarnContext
//
// Deprecated: As of 0.4.0, use Log.OnWarn().Print(args...) instead.
func (l *Log) Warn(args ...any) {
	l.OnWarn().Print(args...)
}

// WarnContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).Print(args...) instead.
func (l *Log) WarnContext(ctx context.Context, args ...any) {
	l.OnWarnContext(ctx).Print(args...)
}

// Error see ErrorContext
//
// Deprecated: As of 0.4.0, use Log.Error().Print(args...) instead.
func (l *Log) Error(args ...any) {
	l.OnError().Print(args...)
}

// ErrorContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).Print(args...) instead.
func (l *Log) ErrorContext(ctx context.Context, args ...any) {
	l.OnErrorContext(ctx).Print(args...)
}

// DPanic see DPanicContext
//
// Deprecated: As of 0.4.0, use Log.DPanic().Print(args...) instead.
func (l *Log) DPanic(args ...any) {
	l.OnDPanic().Print(args...)
}

// DPanicContext uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (see DPanicLevel for details.)
//
// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).Print(args...) instead.
func (l *Log) DPanicContext(ctx context.Context, args ...any) {
	l.OnDPanicContext(ctx).Print(args...)
}

// Panic see PanicContext
//
// Deprecated: As of 0.4.0, use Log.OnPanic().Print(args...) instead.
func (l *Log) Panic(args ...any) {
	l.OnPanic().Print(args...)
}

// PanicContext uses fmt.Sprint to to construct and log a message, then panics.
//
// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).Print(args...) instead.
func (l *Log) PanicContext(ctx context.Context, args ...any) {
	l.OnPanicContext(ctx).Print(args...)
}

// Fatal see FatalContext
//
// Deprecated: As of 0.4.0, use Log.OnFatal().Print(args...) instead.
func (l *Log) Fatal(args ...any) {
	l.OnFatal().Print(args...)
}

// FatalContext uses fmt.Sprint to construct and log a message, then calls os.Exit.
//
// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).Print(args...) instead.
func (l *Log) FatalContext(ctx context.Context, args ...any) {
	l.OnFatalContext(ctx).Print(args...)
}

// ****** ending in "f" or "fContext" for log.Printf-style logging

// Debugf see DebugfContext
//
// Deprecated: As of 0.4.0, use Log.OnDebug().Printf(template, args...) instead.
func (l *Log) Debugf(template string, args ...any) {
	l.OnDebug().Printf(template, args...)
}

// DebugfContext uses fmt.Sprintf to log a templated message.
//
// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).Printf(template, args...) instead.
func (l *Log) DebugfContext(ctx context.Context, template string, args ...any) {
	l.OnDebugContext(ctx).Printf(template, args...)
}

// Infof see InfofContext
//
// Deprecated: As of 0.4.0, use Log.OnInfo().Printf(template, args...) instead.
func (l *Log) Infof(template string, args ...any) {
	l.OnInfo().Printf(template, args...)
}

// InfofContext uses fmt.Sprintf to log a templated message.
//
// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).Printf(template, args...) instead.
func (l *Log) InfofContext(ctx context.Context, template string, args ...any) {
	l.OnInfoContext(ctx).Printf(template, args...)
}

// Warnf see WarnfContext
//
// Deprecated: As of 0.4.0, use Log.OnWarn().Printf(template, args...) instead.
func (l *Log) Warnf(template string, args ...any) {
	l.OnWarn().Printf(template, args...)
}

// WarnfContext uses fmt.Sprintf to log a templated message.
//
// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).Printf(template, args...) instead.
func (l *Log) WarnfContext(ctx context.Context, template string, args ...any) {
	l.OnWarnContext(ctx).Printf(template, args...)
}

// Errorf see ErrorfContext
//
// Deprecated: As of 0.4.0, use Log.OnError().Printf(template, args...) instead.
func (l *Log) Errorf(template string, args ...any) {
	l.OnError().Printf(template, args...)
}

// ErrorfContext uses fmt.Sprintf to log a templated message.
//
// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).Printf(template, args...) instead.
func (l *Log) ErrorfContext(ctx context.Context, template string, args ...any) {
	l.OnErrorContext(ctx).Printf(template, args...)
}

// DPanicf see DPanicfContext
//
// Deprecated: As of 0.4.0, use Log.OnDPanic().Printf(template, args...) instead.
func (l *Log) DPanicf(template string, args ...any) {
	l.OnDPanic().Printf(template, args...)
}

// DPanicfContext uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (see DPanicLevel for details.)
//
// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).Printf(template, args...) instead.
func (l *Log) DPanicfContext(ctx context.Context, template string, args ...any) {
	l.OnDPanicContext(ctx).Printf(template, args...)
}

// Panicf see PanicfContext
//
// Deprecated: As of 0.4.0, use Log.OnPanic().Printf(template, args...) instead.
func (l *Log) Panicf(template string, args ...any) {
	l.OnPanic().Printf(template, args...)
}

// PanicfContext uses fmt.Sprintf to log a templated message, then panics.
//
// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).Printf(template, args...) instead.
func (l *Log) PanicfContext(ctx context.Context, template string, args ...any) {
	l.OnPanicContext(ctx).Printf(template, args...)
}

// Fatalf see FatalfContext
//
// Deprecated: As of 0.4.0, use Log.nFatal().Printf(template, args...) instead.
func (l *Log) Fatalf(template string, args ...any) {
	l.OnFatal().Printf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
//
// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).Printf(template, args...) instead.
func (l *Log) FatalfContext(ctx context.Context, template string, args ...any) {
	l.OnFatalContext(ctx).Printf(template, args...)
}

// ****** ending in "w" or "wContext" for loosely-typed structured logging

// Debugw see DebugwContext
//
// Deprecated: As of 0.4.0, use Log.OnDebug().Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) Debugw(msg string, keysAndValues ...any) {
	l.OnDebug().Pairs(keysAndValues...).Msg(msg)
}

// DebugwContext logs a message with some additional context. The variadic key-value or Field
// pairs or Field are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(fields).Debug(msg)
//
// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) DebugwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.OnDebugContext(ctx).Pairs(keysAndValues...).Msg(msg)
}

// Infow see InfowContext
//
// Deprecated: As of 0.4.0, use Log.OnInfo().Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) Infow(msg string, keysAndValues ...any) {
	l.OnInfo().Pairs(keysAndValues...).Msg(msg)
}

// InfowContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
//
// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) InfowContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.OnInfoContext(ctx).Pairs(keysAndValues...).Msg(msg)
}

// Warnw see WarnwContext
//
// Deprecated: As of 0.4.0, use Log.OnWarn().Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) Warnw(msg string, keysAndValues ...any) {
	l.OnWarn().Pairs(keysAndValues...).Msg(msg)
}

// WarnwContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
//
// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) WarnwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.OnWarnContext(ctx).Pairs(keysAndValues...).Msg(msg)
}

// Errorw see ErrorwContext
//
// Deprecated: As of 0.4.0, use Log.OnError().Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) Errorw(msg string, keysAndValues ...any) {
	l.OnError().Pairs(keysAndValues...).Msg(msg)
}

// ErrorwContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
//
// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) ErrorwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.OnErrorContext(ctx).Pairs(keysAndValues...).Msg(msg)
}

// DPanicw see DPanicwContext
//
// Deprecated: As of 0.4.0, use Log.OnDPanic().Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) DPanicw(msg string, keysAndValues ...any) {
	l.OnDPanic().Pairs(keysAndValues...).Msg(msg)
}

// DPanicwContext logs a message with some additional context. In development, the
// logger then panics. (see DPanicLevel for details.) The variadic key-value
// pairs or Field are treated as they are in With.
//
// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) DPanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.OnDPanicContext(ctx).Pairs(keysAndValues...).Msg(msg)
}

// Panicw see PanicwContext
//
// Deprecated: As of 0.4.0, use Log.OnPanic().Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) Panicw(msg string, keysAndValues ...any) {
	l.OnPanic().Pairs(keysAndValues...).Msg(msg)
}

// PanicwContext logs a message with some additional context, then panics. The
// variadic key-value pairs or Field are treated as they are in With.
//
// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) PanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.OnPanicContext(ctx).Pairs(keysAndValues...).Msg(msg)
}

// Deprecated: As of 0.4.0, use Log.OnFatal().Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) Fatalw(msg string, keysAndValues ...any) {
	l.OnFatal().Pairs(keysAndValues...).Msg(msg)
}

// FatalwContext logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs or Field are treated as they are in With.
//
// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func (l *Log) FatalwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.OnFatalContext(ctx).Pairs(keysAndValues...).Msg(msg)
}

// ****** ending in "x" or "xContext" for structured logging

// Debug (see DebugContext)
//
// Deprecated: As of 0.4.0, use Log.OnDebug().With(fields...).Msg(msg) instead.
func (l *Log) Debugx(msg string, fields ...Field) {
	l.OnDebug().With(fields...).Msg(msg)
}

// DebugContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).With(fields...).Msg(msg) instead.
func (l *Log) DebugxContext(ctx context.Context, msg string, fields ...Field) {
	l.OnDebugContext(ctx).With(fields...).Msg(msg)
}

// Info see InfoContext
//
// Deprecated: As of 0.4.0, use Log.OnInfo().With(fields...).Msg(msg) instead.
func (l *Log) Infox(msg string, fields ...Field) {
	l.OnInfo().With(fields...).Msg(msg)
}

// InfoContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).With(fields...).Msg(msg) instead.
func (l *Log) InfoxContext(ctx context.Context, msg string, fields ...Field) {
	l.OnInfoContext(ctx).With(fields...).Msg(msg)
}

// Warn see WarnContext
//
// Deprecated: As of 0.4.0, use Log.OnWarn().With(fields...).Msg(msg) instead.
func (l *Log) Warnx(msg string, fields ...Field) {
	l.OnWarn().With(fields...).Msg(msg)
}

// WarnContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).With(fields...).Msg(msg) instead.
func (l *Log) WarnxContext(ctx context.Context, msg string, fields ...Field) {
	l.OnWarnContext(ctx).With(fields...).Msg(msg)
}

// Error see ErrorContext
//
// Deprecated: As of 0.4.0, use Log.OnError().With(fields...).Msg(msg) instead.
func (l *Log) Errorx(msg string, fields ...Field) {
	l.OnError().With(fields...).Msg(msg)
}

// ErrorContext uses fmt.Sprint to construct and log a message.
//
// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).With(fields...).Msg(msg) instead.
func (l *Log) ErrorxContext(ctx context.Context, msg string, fields ...Field) {
	l.OnErrorContext(ctx).With(fields...).Msg(msg)
}

// DPanic see DPanicContext
//
// Deprecated: As of 0.4.0, use Log.OnDPanic().With(fields...).Msg(msg) instead.
func (l *Log) DPanicx(msg string, fields ...Field) {
	l.OnDPanic().With(fields...).Msg(msg)
}

// DPanicContext uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (see DPanicLevel for details.)
//
// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).With(fields...).Msg(msg) instead.
func (l *Log) DPanicxContext(ctx context.Context, msg string, fields ...Field) {
	l.OnDPanicContext(ctx).With(fields...).Msg(msg)
}

// Panic see PanicContext
//
// Deprecated: As of 0.4.0, use Log.OnPanic().With(fields...).Msg(msg) instead.
func (l *Log) Panicx(msg string, fields ...Field) {
	l.OnPanic().With(fields...).Msg(msg)
}

// PanicContext uses fmt.Sprint to to construct and log a message, then panics.
//
// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).With(fields...).Msg(msg) instead.
func (l *Log) PanicxContext(ctx context.Context, msg string, fields ...Field) {
	l.OnPanicContext(ctx).With(fields...).Msg(msg)
}

// Fatal see FatalContext
//
// Deprecated: As of 0.4.0, use Log.OnFatal().With(fields...).Msg(msg) instead.
func (l *Log) Fatalx(msg string, fields ...Field) {
	l.OnFatal().With(fields...).Msg(msg)
}

// FatalContext uses fmt.Sprint to construct and log a message, then calls os.Exit.
//
// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).With(fields...).Msg(msg) instead.
func (l *Log) FatalxContext(ctx context.Context, msg string, fields ...Field) {
	l.OnFatalContext(ctx).With(fields...).Msg(msg)
}
