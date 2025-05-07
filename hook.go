package logger

import (
	"context"
)

// Hook
type Hook interface {
	DoHook(ctx context.Context) Field
}

type HookFunc func(context.Context) Field

func (hf HookFunc) DoHook(ctx context.Context) Field { return hf(ctx) }
