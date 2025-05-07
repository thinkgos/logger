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
			hooks:     make([]Hook, 0, 32),
			fields:    make([]zap.Field, 0, 32),
			tmpFields: make([]zap.Field, 0, 64),
		}
	},
}

func getEvent() *Event {
	e := eventPool.Get().(*Event)
	return e.reset()
}

// PoolPut adds x to the pool.
// NOTE: See PoolGet.
func putEvent(e *Event) {
	if e == nil {
		return
	}
	eventPool.Put(e.reset())
}

type Event struct {
	log       *Log
	level     Level
	hooks     []Hook
	fields    []Field
	tmpFields []Field
	ctx       context.Context
}

func (e *Event) reset() *Event {
	if e == nil {
		return e
	}
	e.hooks = e.hooks[:0]
	e.fields = e.fields[:0]
	e.tmpFields = e.tmpFields[:0]
	e.ctx = context.Background()
	return e
}

func (e *Event) WithContext(ctx context.Context) *Event {
	if e == nil {
		return e
	}
	e.ctx = ctx
	return e
}

func (e *Event) ExtendHook(hs ...Hook) *Event {
	if e == nil {
		return e
	}
	e.hooks = append(e.hooks, hs...)
	return e
}

func (e *Event) ExtendHookFunc(hs ...HookFunc) *Event {
	if e == nil {
		return e
	}
	for _, f := range hs {
		e.hooks = append(e.hooks, f)
	}
	return e
}

func (e *Event) WithNewHook(hs ...Hook) *Event {
	if e == nil {
		return e
	}
	e.hooks = e.hooks[:0]
	e.hooks = append(e.hooks, hs...)
	return e
}

func (e *Event) msg(msg string) {
	defer putEvent(e)
	if needCaller := e.log.callerCore.Enabled(e.level); needCaller || len(e.hooks) > 0 {
		if needCaller {
			e.tmpFields = append(e.tmpFields, e.log.callerCore.Caller(e.log.callerCore.Skip, e.log.callerCore.SkipPackages...))
		}
		for _, h := range e.hooks {
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

// NOTICE: Not Recommended
func (e *Event) Pairs(keysAndValues ...any) *Event {
	if e == nil {
		return e
	}
	e.fields = e.log.appendSweetenFields(e.ctx, e.fields, keysAndValues)
	return e
}

func (e *Event) With(fields ...Field) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, fields...)
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
