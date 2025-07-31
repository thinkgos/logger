package logger

import "context"

// OnLevel starts a new message with customize level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnLevel(level Level) *Event {
	if !l.Enabled(level) {
		return nil
	}
	e := getEvent()
	e.log = l
	e.level = level
	e.ctx = context.Background()
	return e
}

// OnLevelContext starts a new message with customize level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnLevelContext(ctx context.Context, level Level) *Event {
	return l.OnLevel(level).WithContext(ctx)
}

// OnDebug starts a new message with [DebugLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnDebug() *Event {
	return l.OnLevel(DebugLevel)
}

// Debug starts a new message with [DebugLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnDebugContext(ctx context.Context) *Event {
	return l.OnLevel(DebugLevel).WithContext(ctx)
}

// OnInfo starts a new message with [InfoLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnInfo() *Event {
	return l.OnLevel(InfoLevel)
}

// OnInfoContext starts a new message with [InfoLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnInfoContext(ctx context.Context) *Event {
	return l.OnLevel(InfoLevel).WithContext(ctx)
}

// OnWarn starts a new message with [WarnLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnWarn() *Event {
	return l.OnLevel(WarnLevel)
}

// OnWarnContext starts a new message with [WarnLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnWarnContext(ctx context.Context) *Event {
	return l.OnLevel(WarnLevel).WithContext(ctx)
}

// OnError starts a new message with [ErrorLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnError() *Event {
	return l.OnLevel(ErrorLevel)
}

// OnErrorContext starts a new message with [ErrorLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnErrorContext(ctx context.Context) *Event {
	return l.OnLevel(ErrorLevel).WithContext(ctx)
}

// OnDPanic starts a new message with [DPanicLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnDPanic() *Event {
	return l.OnLevel(DPanicLevel)
}

// OnDPanicContext starts a new message with [DPanicLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnDPanicContext(ctx context.Context) *Event {
	return l.OnLevel(DPanicLevel).WithContext(ctx)
}

// OnPanic starts a new message with [PanicLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnPanic() *Event {
	return l.OnLevel(PanicLevel)
}

// OnPanicContext starts a new message with [PanicLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnPanicContext(ctx context.Context) *Event {
	return l.OnLevel(PanicLevel).WithContext(ctx)
}

// OnFatal starts a new message with [FatalLevel] level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnFatal() *Event {
	return l.OnLevel(FatalLevel)
}

// OnFatalContext starts a new message with [FatalLevel] level, and adds the Go Context to the *Event context.
//
// You must call Msg on the returned event in order to send the event.
func (l *Log) OnFatalContext(ctx context.Context) *Event {
	return l.OnLevel(FatalLevel).WithContext(ctx)
}

func (l *Log) Debug(args ...any) {
	l.OnDebug().Print(args...)
}

func (l *Log) Info(args ...any) {
	l.OnInfo().Print(args...)
}

func (l *Log) Warn(args ...any) {
	l.OnWarn().Print(args...)
}

func (l *Log) Error(args ...any) {
	l.OnError().Print(args...)
}

func (l *Log) Panic(args ...any) {
	l.OnPanic().Print(args...)
}

func (l *Log) DPanic(args ...any) {
	l.OnDPanic().Print(args...)
}

func (l *Log) Fatal(args ...any) {
	l.OnFatal().Print(args...)
}

func (l *Log) Debugf(template string, args ...any) {
	l.OnDebug().Printf(template, args...)
}

func (l *Log) Infof(template string, args ...any) {
	l.OnInfo().Printf(template, args...)
}

func (l *Log) Warnf(template string, args ...any) {
	l.OnWarn().Printf(template, args...)
}

func (l *Log) Errorf(template string, args ...any) {
	l.OnError().Printf(template, args...)
}

func (l *Log) Panicf(template string, args ...any) {
	l.OnPanic().Printf(template, args...)
}

func (l *Log) DPanicf(template string, args ...any) {
	l.OnDPanic().Printf(template, args...)
}

func (l *Log) Fatalf(template string, args ...any) {
	l.OnFatal().Printf(template, args...)
}
