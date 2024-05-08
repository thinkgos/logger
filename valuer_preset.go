package logger

import (
	"context"
)

// Package returns a Valuer that returns an immutable Valuer which key is pkg
func Package(v string) Valuer {
	return ImmutString("pkg", v)
}

func App(v string) Valuer {
	return ImmutString("app", v)
}
func Component(v string) Valuer {
	return ImmutString("component", v)
}
func Module(v string) Valuer {
	return ImmutString("module", v)
}
func Unit(v string) Valuer {
	return ImmutString("unit", v)
}
func Kind(v string) Valuer {
	return ImmutString("kind", v)
}
func Type(v string) Valuer {
	return ImmutString("type", v)
}
func TraceId(f func(c context.Context) string) Valuer {
	return FromString("traceId", f)
}
func RequestId(f func(c context.Context) string) Valuer {
	return FromString("requestId", f)
}
func Source(f func(c context.Context) string) Valuer {
	return FromString("source", f)
}
