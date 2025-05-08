package logger

import "context"

// SetDefaultValuer set default field function, which hold always until you call [Event.Msg]/[Event.Print]/[Event.Printf].
//
// Deprecated: As of 0.4.0, use [Log.SetDefaultHookFunc]/[Log.SetDefaultHook] instead.
func SetDefaultValuer(vs ...HookFunc) *Log {
	return defaultLogger.SetDefaultValuer(vs...)
}

// WithValuer with field function.
//
// Deprecated: As of 0.4.0, use [Log.ExtendHookFunc]/[Log.ExtendHook] instead.
func WithValuer(vs ...HookFunc) *Log {
	return defaultLogger.WithValuer(vs...)
}

// WithNewValuer return log with new Valuer function without default Valuer.
//
// Deprecated: As of 0.4.0, use [Log.WithNewHookFunc]/[Log.WithNewHook] instead.
func WithNewValuer(fs ...HookFunc) *Log {
	return defaultLogger.WithNewValuer(fs...)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).Print(args...) instead.
func Log2(ctx context.Context, level Level, args ...any) {
	defaultLogger.Log(ctx, level, args...)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).Printf(template, args...) instead.
func Logf(ctx context.Context, level Level, template string, args ...any) {
	defaultLogger.Logf(ctx, level, template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).Pairs(keysAndValues...).Msg(msg) instead.
func Logw(ctx context.Context, level Level, msg string, keysAndValues ...any) {
	defaultLogger.Logw(ctx, level, msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnLevelContext(ctx, level).With(fields...).Msg(msg) instead.
func Logx(ctx context.Context, level Level, msg string, fields ...Field) {
	defaultLogger.Logx(ctx, level, msg, fields...)
}

// ****** named after the log level or ending in "Context" for log.Print-style logging

// Deprecated: As of 0.4.0, use Log.OnDebug().Print(args...) instead.
func Debug(args ...any) {
	defaultLogger.Debug(args...)
}

// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).Print(args...) instead.
func DebugContext(ctx context.Context, args ...any) {
	defaultLogger.DebugContext(ctx, args...)
}

// Deprecated: As of 0.4.0, use Log.OnInfo().Print(args...) instead.
func Info(args ...any) {
	defaultLogger.Info(args...)
}

// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).Print(args...) instead.
func InfoContext(ctx context.Context, args ...any) {
	defaultLogger.InfoContext(ctx, args...)
}

// Warn see WarnContext
//
// Deprecated: As of 0.4.0, use Log.OnWarn().Print(args...) instead.
func Warn(args ...any) {
	defaultLogger.Warn(args...)
}

// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).Print(args...) instead.
func WarnContext(ctx context.Context, args ...any) {
	defaultLogger.WarnContext(ctx, args...)
}

// Deprecated: As of 0.4.0, use Log.Error().Print(args...) instead.
func Error(args ...any) {
	defaultLogger.Error(args...)
}

// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).Print(args...) instead.
func ErrorContext(ctx context.Context, args ...any) {
	defaultLogger.ErrorContext(ctx, args...)
}

// Deprecated: As of 0.4.0, use Log.DPanic().Print(args...) instead.
func DPanic(args ...any) {
	defaultLogger.DPanic(args...)
}

// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).Print(args...) instead.
func DPanicContext(ctx context.Context, args ...any) {
	defaultLogger.DPanicContext(ctx, args...)
}

// Deprecated: As of 0.4.0, use Log.OnPanic().Print(args...) instead.
func Panic(args ...any) {
	defaultLogger.Panic(args...)
}

// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).Print(args...) instead.
func PanicContext(ctx context.Context, args ...any) {
	defaultLogger.PanicContext(ctx, args...)
}

// Deprecated: As of 0.4.0, use Log.OnFatal().Print(args...) instead.
func Fatal(args ...any) {
	defaultLogger.Fatal(args...)
}

// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).Print(args...) instead.
func FatalContext(ctx context.Context, args ...any) {
	defaultLogger.FatalContext(ctx, args...)
}

// ****** ending in "f" or "fContext" for log.Printf-style logging

// Deprecated: As of 0.4.0, use Log.OnDebug().Printf(template, args...) instead.
func Debugf(template string, args ...any) {
	defaultLogger.Debugf(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).Printf(template, args...) instead.
func DebugfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.DebugfContext(ctx, template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnInfo().Printf(template, args...) instead.
func Infof(template string, args ...any) {
	defaultLogger.Infof(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).Printf(template, args...) instead.
func InfofContext(ctx context.Context, template string, args ...any) {
	defaultLogger.InfofContext(ctx, template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnWarn().Printf(template, args...) instead.
func Warnf(template string, args ...any) {
	defaultLogger.Warnf(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).Printf(template, args...) instead.
func WarnfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.WarnfContext(ctx, template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnError().Printf(template, args...) instead.
func Errorf(template string, args ...any) {
	defaultLogger.Errorf(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).Printf(template, args...) instead.
func ErrorfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.ErrorfContext(ctx, template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnDPanic().Printf(template, args...) instead.
func DPanicf(template string, args ...any) {
	defaultLogger.DPanicf(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).Printf(template, args...) instead.
func DPanicfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.DPanicfContext(ctx, template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnPanic().Printf(template, args...) instead.
func Panicf(template string, args ...any) {
	defaultLogger.Panicf(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).Printf(template, args...) instead.
func PanicfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.PanicfContext(ctx, template, args...)
}

// Deprecated: As of 0.4.0, use Log.nFatal().Printf(template, args...) instead.
func Fatalf(template string, args ...any) {
	defaultLogger.Fatalf(template, args...)
}

// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).Printf(template, args...) instead.
func FatalfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.FatalfContext(ctx, template, args...)
}

// ****** ending in "w" or "wContext" for loosely-typed structured logging

// Deprecated: As of 0.4.0, use Log.OnDebug().Pairs(keysAndValues...).Msg(msg) instead.
func Debugw(msg string, keysAndValues ...any) {
	defaultLogger.Debugw(msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func DebugwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.DebugwContext(ctx, msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnInfo().Pairs(keysAndValues...).Msg(msg) instead.
func Infow(msg string, keysAndValues ...any) {
	defaultLogger.Infow(msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func InfowContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.InfowContext(ctx, msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnWarn().Pairs(keysAndValues...).Msg(msg) instead.
func Warnw(msg string, keysAndValues ...any) {
	defaultLogger.Warnw(msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func WarnwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.WarnwContext(ctx, msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnError().Pairs(keysAndValues...).Msg(msg) instead.
func Errorw(msg string, keysAndValues ...any) {
	defaultLogger.Errorw(msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func ErrorwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.ErrorwContext(ctx, msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnDPanic().Pairs(keysAndValues...).Msg(msg) instead.
func DPanicw(msg string, keysAndValues ...any) {
	defaultLogger.DPanicw(msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func DPanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.DPanicwContext(ctx, msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnPanic().Pairs(keysAndValues...).Msg(msg) instead.
func Panicw(msg string, keysAndValues ...any) {
	defaultLogger.Panicw(msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func PanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.PanicwContext(ctx, msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnFatal().Pairs(keysAndValues...).Msg(msg) instead.
func Fatalw(msg string, keysAndValues ...any) {
	defaultLogger.Fatalw(msg, keysAndValues...)
}

// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).Pairs(keysAndValues...).Msg(msg) instead.
func FatalwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.FatalwContext(ctx, msg, keysAndValues...)
}

// ****** ending in "x" or "xContext" for structured logging

// Deprecated: As of 0.4.0, use Log.OnDebug().With(fields...).Msg(msg) instead.
func Debugx(msg string, fields ...Field) {
	defaultLogger.Debugx(msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnDebugContext(ctx).With(fields...).Msg(msg) instead.
func DebugxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.DebugxContext(ctx, msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnInfo().With(fields...).Msg(msg) instead.
func Infox(msg string, fields ...Field) {
	defaultLogger.Infox(msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnInfoContext(ctx).With(fields...).Msg(msg) instead.
func InfoxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.InfoxContext(ctx, msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnWarn().With(fields...).Msg(msg) instead.
func Warnx(msg string, fields ...Field) {
	defaultLogger.Warnx(msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnWarnContext(ctx).With(fields...).Msg(msg) instead.
func WarnxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.WarnxContext(ctx, msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnError().With(fields...).Msg(msg) instead.
func Errorx(msg string, fields ...Field) {
	defaultLogger.Errorx(msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnErrorContext(ctx).With(fields...).Msg(msg) instead.
func ErrorxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.ErrorxContext(ctx, msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnDPanic().With(fields...).Msg(msg) instead.
func DPanicx(msg string, fields ...Field) {
	defaultLogger.DPanicx(msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnDPanicContext(ctx).With(fields...).Msg(msg) instead.
func DPanicxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.DPanicxContext(ctx, msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnPanic().With(fields...).Msg(msg) instead.
func Panicx(msg string, fields ...Field) {
	defaultLogger.Panicx(msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnPanicContext(ctx).With(fields...).Msg(msg) instead.
func PanicxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.PanicxContext(ctx, msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnFatal().With(fields...).Msg(msg) instead.
func Fatalx(msg string, fields ...Field) {
	defaultLogger.Fatalx(msg, fields...)
}

// Deprecated: As of 0.4.0, use Log.OnFatalContext(ctx).With(fields...).Msg(msg) instead.
func FatalxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.FatalxContext(ctx, msg, fields...)
}
