package logger

import (
	"context"
	"fmt"
	"sync"
)

var eventPool = &sync.Pool{
	New: func() any {
		return &Event{
			fields: make([]Field, 0, 32),
		}
	},
}

// getEvent
func getEvent() *Event {
	e := eventPool.Get().(*Event)
	return e.reset()
}

// putEvent adds x to the pool.
func putEvent(e *Event) {
	if e == nil {
		return
	}
	eventPool.Put(e.reset())
}

// Event represents a log event.
// It is instanced by one of the level method of Logger and finalized by the Msg, Print, Printf method.
type Event struct {
	log    *Log
	level  Level
	fields []Field
	ctx    context.Context
}

func (e *Event) reset() *Event {
	if e == nil {
		return e
	}
	e.fields = e.fields[:0]
	e.ctx = context.Background()
	return e
}

func (e *Event) msg(msg string) {
	defer putEvent(e)
	if needCaller := e.log.callerCore.Enabled(e.level); needCaller {
		e.fields = append(e.fields, e.log.callerCore.Caller(e.log.callerCore.Skip, e.log.callerCore.SkipPackages...))
	}
	for _, h := range e.log.hooks {
		h.RunHook(e)
	}
	e.log.log.Log(e.level, msg, e.fields...)
}

// Context returns the context of the event.
func (e *Event) Context() context.Context { return e.ctx }

// Level returns the level of the event.
func (e *Event) Level() Level { return e.level }

// NOTICE: once this method is called, the *Event should be disposed.
func (e *Event) Print(args ...any) {
	if e == nil {
		return
	}
	e.msg(sprintMessage(args...))
}

// NOTICE: once this method is called, the *Event should be disposed.
func (e *Event) Printf(template string, args ...any) {
	if e == nil {
		return
	}
	e.msg(sprintfMessage(template, args...))
}

// NOTICE: once this method is called, the *Event should be disposed.
func (e *Event) Msg(msg string) {
	if e == nil {
		return
	}
	e.msg(msg)
}

// WithContext adds the Go Context to the *Event context.
// The context is not rendered in the output message, but is available to hooks calls.
// A typical use case is to extract tracing information from the Go Ctx.
func (e *Event) WithContext(ctx context.Context) *Event {
	if e == nil {
		return e
	}
	e.ctx = ctx
	return e
}

func (e *Event) With(fields ...Field) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, fields...)
	return e
}

// Hook hook immediately.
func (e *Event) Hook(hs ...Hook) *Event {
	if e == nil {
		return e
	}
	for _, f := range hs {
		f.RunHook(e)
	}
	return e
}

// HookFunc hook function immediately.
func (e *Event) HookFunc(hs ...HookFunc) *Event {
	if e == nil {
		return e
	}
	for _, f := range hs {
		f(e)
	}
	return e
}

// HookField hook field immediately.
func (e *Event) HookField(hs ...HookField) *Event {
	if e == nil {
		return e
	}
	for _, f := range hs {
		f.RunHook(e)
	}
	return e
}

// HookIf hook conditionally immediately.
func (e *Event) HookIf(b bool, hs ...Hook) *Event {
	if b {
		e.Hook(hs...)
	}
	return e
}

// HookFuncIf hook conditionally immediately.
func (e *Event) HookFuncIf(b bool, hs ...HookFunc) *Event {
	if b {
		e.HookFunc(hs...)
	}
	return e
}

// HookFieldIf hook conditionally immediately.
func (e *Event) HookFieldIf(b bool, hs ...HookField) *Event {
	if b {
		e.HookField(hs...)
	}
	return e
}

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

// sprintfMessage format with fmt.Sprintf.
func sprintfMessage(template string, args ...any) string {
	if len(args) == 0 {
		return template
	}
	return fmt.Sprintf(template, args...)
}
