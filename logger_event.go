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
	e.hooks = append(e.hooks, l.hooks...)
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
