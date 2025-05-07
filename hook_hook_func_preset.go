package logger

import (
	"context"
)

// Package returns a Valuer that returns an immutable Valuer which key is pkg
func Package(v string) HookFunc {
	return ImmutString("pkg", v)
}
func App(v string) HookFunc {
	return ImmutString("app", v)
}
func Component(v string) HookFunc {
	return ImmutString("component", v)
}
func Module(v string) HookFunc {
	return ImmutString("module", v)
}
func Unit(v string) HookFunc {
	return ImmutString("unit", v)
}
func Kind(v string) HookFunc {
	return ImmutString("kind", v)
}
func Type(v string) HookFunc {
	return ImmutString("type", v)
}
func TraceId(f func(c context.Context) string) HookFunc {
	return FromString("traceId", f)
}
func SpanId(f func(c context.Context) string) HookFunc {
	return FromString("spanId", f)
}
func RequestId(f func(c context.Context) string) HookFunc {
	return FromString("requestId", f)
}
func Source(f func(c context.Context) string) HookFunc {
	return FromString("source", f)
}
