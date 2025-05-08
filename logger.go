package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log wrap zap logger
type Log struct {
	log   *zap.Logger
	level AtomicLevel
	hooks []Hook
	// for caller
	callerCore *CallerCore
}

// NewLoggerWith new logger with zap logger and atomic level
func NewLoggerWith(logger *zap.Logger, lv AtomicLevel) *Log {
	return &Log{
		log:        logger,
		level:      lv,
		hooks:      nil,
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

// SetNewCallerCore overwrite with new caller core
func (l *Log) SetNewCallerCore(c *CallerCore) *Log {
	if c != nil {
		l.callerCore = c
	}
	return l
}

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

// UnderlyingCallerLevel return underlying caller level.
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

// SetDefaultHook set default hook, which hold always until you call [Event.Msg]/[Event.Print]/[Event.Printf].
func (l *Log) SetDefaultHook(hs ...Hook) *Log {
	hooks := make([]Hook, len(l.hooks)+len(hs))
	copy(hooks, l.hooks)
	copy(hooks[len(l.hooks):], hs)
	l.hooks = hooks
	return l
}

// SetDefaultHookFunc set default hook, which hold always until you call [Event.Msg]/[Event.Print]/[Event.Printf].
func (l *Log) SetDefaultHookFunc(hs ...HookFunc) *Log {
	hooks := make([]Hook, len(l.hooks)+len(hs))
	copy(hooks, l.hooks)
	for i := range hs {
		hooks[len(l.hooks)+i] = hs[i]
	}
	l.hooks = hooks
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

// ExtendHook return new Log with extend Hook.
func (l *Log) ExtendHook(hs ...Hook) *Log {
	hooks := make([]Hook, len(l.hooks)+len(hs))
	copy(hooks, l.hooks)
	copy(hooks[len(l.hooks):], hs)
	return &Log{
		log:        l.log,
		level:      l.level,
		hooks:      hooks,
		callerCore: l.callerCore,
	}
}

// ExtendHookFunc return new Log with extend Hook.
func (l *Log) ExtendHookFunc(hs ...HookFunc) *Log {
	hooks := make([]Hook, len(l.hooks)+len(hs))
	copy(hooks, l.hooks)
	for i := range hs {
		hooks[len(l.hooks)+i] = hs[i]
	}
	return &Log{
		log:        l.log,
		level:      l.level,
		hooks:      hooks,
		callerCore: l.callerCore,
	}
}

// WithNewHook return new log with new hook without default hook.
func (l *Log) WithNewHook(hs ...Hook) *Log {
	hooks := make([]Hook, len(hs))
	copy(hooks, hs)
	return &Log{
		log:        l.log,
		level:      l.level,
		hooks:      hooks,
		callerCore: l.callerCore,
	}
}

// WithNewHookFunc return new log with new hook func without default hook.
func (l *Log) WithNewHookFunc(hs ...HookFunc) *Log {
	hooks := make([]Hook, len(hs))
	for i := range hs {
		hooks[i] = hs[i]
	}
	return &Log{
		log:        l.log,
		level:      l.level,
		hooks:      hooks,
		callerCore: l.callerCore,
	}
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
//
// NOTICE: if you do not need a new log, use [Event.With] instead.
func (l *Log) With(fields ...Field) *Log {
	return &Log{
		log:        l.log.With(fields...),
		level:      l.level,
		hooks:      l.hooks,
		callerCore: l.callerCore,
	}
}

// Named adds a sub-scope to the logger's name. See [Logger.Named] for details.
func (l *Log) Named(name string) *Log {
	return &Log{
		log:        l.log.Named(name),
		level:      l.level,
		hooks:      l.hooks,
		callerCore: l.callerCore,
	}
}

// Sync flushes any buffered log entries.
func (l *Log) Sync() error {
	return l.log.Sync()
}

func (l *Log) OnLevel(level Level) *Event {
	if !l.Enabled(level) {
		return nil
	}
	e := getEvent()
	e.log = l
	e.level = level
	e.hooks = append(e.hooks, l.hooks...)
	e.ctx = context.Background()
	return e
}

func (l *Log) OnLevelContext(ctx context.Context, level Level) *Event {
	return l.OnLevel(level).WithContext(ctx)
}

func (l *Log) OnDebug() *Event {
	return l.OnLevel(DebugLevel)
}

func (l *Log) OnDebugContext(ctx context.Context) *Event {
	return l.OnLevel(DebugLevel).WithContext(ctx)
}
func (l *Log) OnInfo() *Event {
	return l.OnLevel(InfoLevel)
}
func (l *Log) OnInfoContext(ctx context.Context) *Event {
	return l.OnLevel(InfoLevel).WithContext(ctx)
}
func (l *Log) OnWarn() *Event {
	return l.OnLevel(WarnLevel)
}
func (l *Log) OnWarnContext(ctx context.Context) *Event {
	return l.OnLevel(WarnLevel).WithContext(ctx)
}
func (l *Log) OnError() *Event {
	return l.OnLevel(ErrorLevel)
}
func (l *Log) OnErrorContext(ctx context.Context) *Event {
	return l.OnLevel(ErrorLevel).WithContext(ctx)
}
func (l *Log) OnDPanic() *Event {
	return l.OnLevel(DPanicLevel)
}
func (l *Log) OnDPanicContext(ctx context.Context) *Event {
	return l.OnLevel(DPanicLevel).WithContext(ctx)
}
func (l *Log) OnPanic() *Event {
	return l.OnLevel(PanicLevel)
}
func (l *Log) OnPanicContext(ctx context.Context) *Event {
	return l.OnLevel(PanicLevel).WithContext(ctx)
}
func (l *Log) OnFatal() *Event {
	return l.OnLevel(FatalLevel)
}
func (l *Log) OnFatalContext(ctx context.Context) *Event {
	return l.OnLevel(FatalLevel).WithContext(ctx)
}
