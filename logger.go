package logger

import (
	"context"
	"fmt"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log wrap zap logger
// - methods named after the log level or ending in "Context" for log.Print-style logging
// - methods ending in "w" or "wContext" for loosely-typed structured logging
// - methods ending in "f" or "fContext" for log.Printf-style logging
// - methods ending in "x" or "xContext" for structured logging
type Log struct {
	log   *zap.Logger
	level AtomicLevel
	hook  []Valuer
	// for caller
	callerCore *CallerCore
}

// NewLoggerWith new logger with zap logger and atomic level
func NewLoggerWith(logger *zap.Logger, lv AtomicLevel) *Log {
	return &Log{
		log:        logger,
		level:      lv,
		hook:       nil,
		callerCore: NewCallerCore(),
	}
}

// NewLogger new logger
// 默认配置:
//
// Level: 日志等级, 默认warn
// Format: 编码格式, 默认json
// EncodeLevel: 编码器类型, 默认LowercaseLevelEncoder
// Adapter: 默认输出适合器, 默认console`
// Stack: 是否使能栈调试输出, 默认false
// Path: 日志保存路径, 默认当前路径
// Writer: 当adapter有附带custom时, 如果为writer为空, 将使用os.Stdout
// EncoderConfig: 如果配置该项,则 EncodeLevel 将被覆盖
//
// 文件日志切割配置(启用file时生效)
// Filename 空字符使用默认, 默认<processname>-lumberjack.log
// MaxSize 每个日志文件最大尺寸(MB), 默认100MB
// MaxAge 日志文件保存天数, 默认0 不删除
// MaxBackups 日志文件保存备份数, 默认0 都保存
// LocalTime 是否格式化时间戳, 默认UTC时间
// Compress 是否使用gzip压缩文件, 采用默认不压缩
//
// Caller相关
// callerLevel caller日志级别, 默认warn
// callerSkip caller设置跳过深度, 默认0
// callerSkipPackages caller设置跳过的包名, 默认空
func NewLogger(opts ...Option) *Log { return NewLoggerWith(New(opts...)) }

// AddCallerSkip add the number of callers skipped by caller annotation.
func (l *Log) AddCallerSkip(callerSkip int) *Log {
	l.callerCore.AddSkip(callerSkip)
	return l
}

// AddCallerSkipPackage add the caller skip package.
func (l *Log) AddCallerSkipPackage(vs ...string) *Log {
	l.callerCore.AddSkipPackage(vs...)
	return l
}

// SetCallerLevel set the caller level.
func (l *Log) SetCallerLevel(lv Level) *Log {
	l.callerCore.SetLevel(lv)
	return l
}

// UseExternalCallerLevel use external caller level, which controller by user.
func (l *Log) UseExternalCallerLevel(lvl AtomicLevel) *Log {
	l.callerCore.UseExternalLevel(lvl)
	return l
}

// UnderlyingCallerLevel get underlying caller level.
func (l *Log) UnderlyingCallerLevel() AtomicLevel {
	return l.callerCore.UnderlyingLevel()
}

// SetCaller set the caller function.
func (l *Log) SetCaller(f func(depth int, skipPackages ...string) Field) *Log {
	if f != nil {
		l.callerCore.Caller = f
	}
	return l
}

// SetLevelWithText alters the logging level.
// ParseAtomicLevel set the logging level based on a lowercase or all-caps ASCII
// representation of the log level.
// If the provided ASCII representation is
// invalid an error is returned.
// see zapcore.Level
func (l *Log) SetLevelWithText(text string) error {
	lv, err := zapcore.ParseLevel(text)
	if err != nil {
		return err
	}
	l.level.SetLevel(lv)
	return nil
}

// SetLevel alters the logging level.
func (l *Log) SetLevel(lv Level) *Log {
	l.level.SetLevel(lv)
	return l
}

// SetDefaultValuer set default Valuer function, which hold always until you call XXXContext.
func (l *Log) SetDefaultValuer(fs ...Valuer) *Log {
	fn := make([]Valuer, 0, len(fs)+len(l.hook))
	fn = append(fn, l.hook...)
	fn = append(fn, fs...)
	l.hook = fn
	return l
}

// GetLevel returns the minimum enabled log level.
func (l *Log) GetLevel() Level { return l.level.Level() }

// Enabled returns true if the given level is at or above this level.
func (l *Log) Enabled(lvl Level) bool { return l.level.Enabled(lvl) }

// V returns true if the given level is at or above this level.
// same as Enabled
func (l *Log) V(lvl Level) bool { return l.level.Enabled(lvl) }

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func (l *Log) Sugar() *zap.SugaredLogger { return l.log.Sugar() }

// Logger return internal logger
func (l *Log) Logger() *zap.Logger { return l.log }

// WithValuer with Valuer function.
func (l *Log) WithValuer(fs ...Valuer) *Log {
	fn := make([]Valuer, 0, len(fs)+len(l.hook))
	fn = append(fn, l.hook...)
	fn = append(fn, fs...)
	return &Log{
		log:        l.log,
		level:      l.level,
		hook:       fn,
		callerCore: l.callerCore,
	}
}

// WithNewValuer return log with new Valuer function without default Valuer.
func (l *Log) WithNewValuer(fs ...Valuer) *Log {
	return &Log{
		log:        l.log,
		level:      l.level,
		hook:       fs,
		callerCore: l.callerCore,
	}
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (l *Log) With(fields ...Field) *Log {
	return &Log{
		log:        l.log.With(fields...),
		level:      l.level,
		hook:       l.hook,
		callerCore: l.callerCore,
	}
}

// Named adds a sub-scope to the logger's name. See Log.Named for details.
func (l *Log) Named(name string) *Log {
	return &Log{
		log:        l.log.Named(name),
		level:      l.level,
		hook:       l.hook,
		callerCore: l.callerCore,
	}
}

// Sync flushes any buffered log entries.
func (l *Log) Sync() error {
	return l.log.Sync()
}

func (l *Log) Log(ctx context.Context, level Level, args ...any) {
	l.Logx(ctx, level, sprintMessage(args...))
}

func (l *Log) Logf(ctx context.Context, level Level, template string, args ...any) {
	l.Logx(ctx, level, fmt.Sprintf(template, args...))
}

func (l *Log) Logw(ctx context.Context, level Level, msg string, keysAndValues ...any) {
	if !l.level.Enabled(level) {
		return
	}
	fc := PoolGet()
	defer PoolPut(fc)
	if l.callerCore.Enabled(level) {
		fc.Fields = append(fc.Fields, l.callerCore.Caller(l.callerCore.Skip, l.callerCore.SkipPackages...))
	}
	for _, f := range l.hook {
		fc.Fields = append(fc.Fields, f(ctx))
	}
	fc.Fields = l.appendSweetenFields(ctx, fc.Fields, keysAndValues)
	l.log.Log(level, msg, fc.Fields...)
}

func (l *Log) Logx(ctx context.Context, level Level, msg string, fields ...Field) {
	if !l.level.Enabled(level) {
		return
	}
	if needCaller := l.callerCore.Enabled(level); needCaller || len(l.hook) > 0 {
		fc := PoolGet()
		defer PoolPut(fc)
		if needCaller {
			fc.Fields = append(fc.Fields, l.callerCore.Caller(l.callerCore.Skip, l.callerCore.SkipPackages...))
		}
		for _, f := range l.hook {
			fc.Fields = append(fc.Fields, f(ctx))
		}
		fc.Fields = append(fc.Fields, fields...)
		l.log.Log(level, msg, fc.Fields...)
	} else {
		l.log.Log(level, msg, fields...)
	}
}

// ****** named after the log level or ending in "Context" for log.Print-style logging

// Debug (see DebugContext)
func (l *Log) Debug(args ...any) {
	l.DebugContext(context.Background(), args...)
}

// DebugContext uses fmt.Sprint to construct and log a message.
func (l *Log) DebugContext(ctx context.Context, args ...any) {
	l.Log(ctx, DebugLevel, args...)
}

// Info see InfoContext
func (l *Log) Info(args ...any) {
	l.InfoContext(context.Background(), args...)
}

// InfoContext uses fmt.Sprint to construct and log a message.
func (l *Log) InfoContext(ctx context.Context, args ...any) {
	l.Log(ctx, InfoLevel, args...)
}

// Warn see WarnContext
func (l *Log) Warn(args ...any) {
	l.WarnContext(context.Background(), args...)
}

// WarnContext uses fmt.Sprint to construct and log a message.
func (l *Log) WarnContext(ctx context.Context, args ...any) {
	l.Log(ctx, WarnLevel, args...)
}

// Error see ErrorContext
func (l *Log) Error(args ...any) {
	l.ErrorContext(context.Background(), args...)
}

// ErrorContext uses fmt.Sprint to construct and log a message.
func (l *Log) ErrorContext(ctx context.Context, args ...any) {
	l.Log(ctx, ErrorLevel, args...)
}

// DPanic see DPanicContext
func (l *Log) DPanic(args ...any) {
	l.DPanicContext(context.Background(), args...)
}

// DPanicContext uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (see DPanicLevel for details.)
func (l *Log) DPanicContext(ctx context.Context, args ...any) {
	l.Log(ctx, DPanicLevel, args...)
}

// Panic see PanicContext
func (l *Log) Panic(args ...any) {
	l.PanicContext(context.Background(), args...)
}

// PanicContext uses fmt.Sprint to to construct and log a message, then panics.
func (l *Log) PanicContext(ctx context.Context, args ...any) {
	l.Log(ctx, PanicLevel, args...)
}

// Fatal see FatalContext
func (l *Log) Fatal(args ...any) {
	l.FatalContext(context.Background(), args...)
}

// FatalContext uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Log) FatalContext(ctx context.Context, args ...any) {
	l.Log(ctx, FatalLevel, args...)
}

// ****** ending in "f" or "fContext" for log.Printf-style logging

// Debugf see DebugfContext
func (l *Log) Debugf(template string, args ...any) {
	l.DebugfContext(context.Background(), template, args...)
}

// DebugfContext uses fmt.Sprintf to log a templated message.
func (l *Log) DebugfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, DebugLevel, template, args...)
}

// Infof see InfofContext
func (l *Log) Infof(template string, args ...any) {
	l.InfofContext(context.Background(), template, args...)
}

// InfofContext uses fmt.Sprintf to log a templated message.
func (l *Log) InfofContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, InfoLevel, template, args...)
}

// Warnf see WarnfContext
func (l *Log) Warnf(template string, args ...any) {
	l.WarnfContext(context.Background(), template, args...)
}

// WarnfContext uses fmt.Sprintf to log a templated message.
func (l *Log) WarnfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, WarnLevel, template, args...)
}

// Errorf see ErrorfContext
func (l *Log) Errorf(template string, args ...any) {
	l.ErrorfContext(context.Background(), template, args...)
}

// ErrorfContext uses fmt.Sprintf to log a templated message.
func (l *Log) ErrorfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, ErrorLevel, template, args...)
}

// DPanicf see DPanicfContext
func (l *Log) DPanicf(template string, args ...any) {
	l.DPanicfContext(context.Background(), template, args...)
}

// DPanicfContext uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (see DPanicLevel for details.)
func (l *Log) DPanicfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, DPanicLevel, template, args...)
}

// Panicf see PanicfContext
func (l *Log) Panicf(template string, args ...any) {
	l.PanicfContext(context.Background(), template, args...)
}

// PanicfContext uses fmt.Sprintf to log a templated message, then panics.
func (l *Log) PanicfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, PanicLevel, template, args...)
}

// Fatalf see FatalfContext
func (l *Log) Fatalf(template string, args ...any) {
	l.FatalfContext(context.Background(), template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Log) FatalfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, FatalLevel, template, args...)
}

// ****** ending in "w" or "wContext" for loosely-typed structured logging

// Debugw see DebugwContext
func (l *Log) Debugw(msg string, keysAndValues ...any) {
	l.DebugwContext(context.Background(), msg, keysAndValues...)
}

// DebugwContext logs a message with some additional context. The variadic key-value or Field
// pairs or Field are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(fields).Debug(msg)
func (l *Log) DebugwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, DebugLevel, msg, keysAndValues...)
}

// Infow see InfowContext
func (l *Log) Infow(msg string, keysAndValues ...any) {
	l.InfowContext(context.Background(), msg, keysAndValues...)
}

// InfowContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) InfowContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, InfoLevel, msg, keysAndValues...)
}

// Warnw see WarnwContext
func (l *Log) Warnw(msg string, keysAndValues ...any) {
	l.WarnwContext(context.Background(), msg, keysAndValues...)
}

// WarnwContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) WarnwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, WarnLevel, msg, keysAndValues...)
}

// Errorw see ErrorwContext
func (l *Log) Errorw(msg string, keysAndValues ...any) {
	l.ErrorwContext(context.Background(), msg, keysAndValues...)
}

// ErrorwContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) ErrorwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, ErrorLevel, msg, keysAndValues...)
}

// DPanicw see DPanicwContext
func (l *Log) DPanicw(msg string, keysAndValues ...any) {
	l.DPanicwContext(context.Background(), msg, keysAndValues...)
}

// DPanicwContext logs a message with some additional context. In development, the
// logger then panics. (see DPanicLevel for details.) The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) DPanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, DPanicLevel, msg, keysAndValues...)
}

// Panicw see PanicwContext
func (l *Log) Panicw(msg string, keysAndValues ...any) {
	l.PanicwContext(context.Background(), msg, keysAndValues...)
}

// PanicwContext logs a message with some additional context, then panics. The
// variadic key-value pairs or Field are treated as they are in With.
func (l *Log) PanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, PanicLevel, msg, keysAndValues...)
}

func (l *Log) Fatalw(msg string, keysAndValues ...any) {
	l.FatalwContext(context.Background(), msg, keysAndValues...)
}

// FatalwContext logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs or Field are treated as they are in With.
func (l *Log) FatalwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, FatalLevel, msg, keysAndValues...)
}

// ****** ending in "x" or "xContext" for structured logging

// Debug (see DebugContext)
func (l *Log) Debugx(msg string, fields ...Field) {
	l.DebugxContext(context.Background(), msg, fields...)
}

// DebugContext uses fmt.Sprint to construct and log a message.
func (l *Log) DebugxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, DebugLevel, msg, fields...)
}

// Info see InfoContext
func (l *Log) Infox(msg string, fields ...Field) {
	l.InfoxContext(context.Background(), msg, fields...)
}

// InfoContext uses fmt.Sprint to construct and log a message.
func (l *Log) InfoxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, InfoLevel, msg, fields...)
}

// Warn see WarnContext
func (l *Log) Warnx(msg string, fields ...Field) {
	l.WarnxContext(context.Background(), msg, fields...)
}

// WarnContext uses fmt.Sprint to construct and log a message.
func (l *Log) WarnxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, WarnLevel, msg, fields...)
}

// Error see ErrorContext
func (l *Log) Errorx(msg string, fields ...Field) {
	l.ErrorxContext(context.Background(), msg, fields...)
}

// ErrorContext uses fmt.Sprint to construct and log a message.
func (l *Log) ErrorxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, ErrorLevel, msg, fields...)
}

// DPanic see DPanicContext
func (l *Log) DPanicx(msg string, fields ...Field) {
	l.DPanicxContext(context.Background(), msg, fields...)
}

// DPanicContext uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (see DPanicLevel for details.)
func (l *Log) DPanicxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, DPanicLevel, msg, fields...)
}

// Panic see PanicContext
func (l *Log) Panicx(msg string, fields ...Field) {
	l.PanicxContext(context.Background(), msg, fields...)
}

// PanicContext uses fmt.Sprint to to construct and log a message, then panics.
func (l *Log) PanicxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, PanicLevel, msg, fields...)
}

// Fatal see FatalContext
func (l *Log) Fatalx(msg string, fields ...Field) {
	l.FatalxContext(context.Background(), msg, fields...)
}

// FatalContext uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Log) FatalxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, FatalLevel, msg, fields...)
}

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
	_multipleErrMsg     = "Multiple errors without a key."
)

// sprintMessage format with fmt.Sprint.
func sprintMessage(args ...any) string {
	if len(args) == 0 {
		return ""
	}
	if len(args) == 1 {
		if str, ok := args[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(args...)
}

// copy from zap(sugar.go)
func (l *Log) appendSweetenFields(ctx context.Context, fields []Field, keysAndValues []any) []Field {
	if len(keysAndValues) == 0 {
		return fields
	}

	var (
		invalid   invalidPairs
		seenError bool
	)

	for i := 0; i < len(keysAndValues); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := keysAndValues[i].(Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// If it is an error, consume it and move on.
		if err, ok := keysAndValues[i].(error); ok {
			if !seenError {
				seenError = true
				fields = append(fields, zap.Error(err))
			} else {
				l.Logx(ctx, ErrorLevel, _multipleErrMsg, zap.Error(err))
			}
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(keysAndValues)-1 {
			l.Logx(ctx, ErrorLevel, _oddNumberErrMsg, Any("ignored", keysAndValues[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := keysAndValues[i], keysAndValues[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(keysAndValues)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, Any(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		l.Logx(ctx, ErrorLevel, _nonStringKeyErrMsg, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value any
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	Any("key", p.key).AddTo(enc)
	Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}
