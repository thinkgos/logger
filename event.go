package logger

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

var eventPool = &sync.Pool{
	New: func() any {
		return &Event{
			fields:    make([]zap.Field, 0, 32),
			tmpFields: make([]zap.Field, 0, 64),
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
	log       *Log
	level     Level
	fields    []Field
	tmpFields []Field
	ctx       context.Context
}

func (e *Event) reset() *Event {
	if e == nil {
		return e
	}
	e.fields = e.fields[:0]
	e.tmpFields = e.tmpFields[:0]
	e.ctx = context.Background()
	return e
}

func (e *Event) msg(msg string) {
	defer putEvent(e)
	if needCaller := e.log.callerCore.Enabled(e.level); needCaller || len(e.log.hooks) > 0 {
		if needCaller {
			e.tmpFields = append(e.tmpFields, e.log.callerCore.Caller(e.log.callerCore.Skip, e.log.callerCore.SkipPackages...))
		}
		for _, h := range e.log.hooks {
			e.tmpFields = append(e.tmpFields, h.DoHook(e.ctx))
		}
		e.tmpFields = append(e.tmpFields, e.fields...)
		e.log.log.Log(e.level, msg, e.tmpFields...)
	} else {
		e.log.log.Log(e.level, msg, e.fields...)
	}
}

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

// DoHookFunc do hook func immediately.
func (e *Event) DoHookFunc(hs ...HookFunc) *Event {
	if e == nil {
		return e
	}
	for _, f := range hs {
		e.fields = append(e.fields, f.DoHook(e.ctx))
	}
	return e
}

// DoHookFunc do hook immediately.
func (e *Event) DoHook(hs ...Hook) *Event {
	if e == nil {
		return e
	}
	for _, f := range hs {
		e.fields = append(e.fields, f.DoHook(e.ctx))
	}
	return e
}

func (e *Event) Configure(f func(e *Event)) *Event {
	if e == nil {
		return e
	}
	f(e)
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
