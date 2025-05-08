package logger

import (
	"context"
)

// Hook defines an interface to a log hook.
type Hook interface {
	DoHook(ctx context.Context) Field
}

// HookFunc is an adaptor to allow the use of an ordinary function as a Hook.
type HookFunc func(context.Context) Field

// DoHook implements the Hook interface.
func (hf HookFunc) DoHook(ctx context.Context) Field { return hf(ctx) }
