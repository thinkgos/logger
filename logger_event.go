package logger

import "context"

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
