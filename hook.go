package logger

import (
	"context"
)

// Hook defines an interface to a log hook.
type Hook interface {
	RunHook(e *Event)
}

type HookFunc func(e *Event)

// RunHook implements the Hook interface.
func (f HookFunc) RunHook(e *Event) {
	f(e)
}

// HookField is an adaptor to allow the use of an ordinary function as a Hook.
type HookField func(context.Context) Field

// RunHook implements the Hook interface.
func (f HookField) RunHook(e *Event) {
	e.Fields(f(e.Context()))
}
