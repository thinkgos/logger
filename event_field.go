package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func (e *Event) Error(val error) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Error(val))
	return e
}
func (e *Event) Errors(key string, val []error) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Errors(key, val))
	return e
}
func (e *Event) NamedError(key string, val error) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.NamedError(key, val))
	return e
}
func (e *Event) Binary(key string, v []byte) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Binary(key, v))
	return e
}
func (e *Event) Bool(key string, v bool) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Bool(key, v))
	return e
}
func (e *Event) Boolp(key string, v *bool) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Boolp(key, v))
	return e
}
func (e *Event) ByteString(key string, v []byte) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.ByteString(key, v))
	return e
}
func (e *Event) Complex128(key string, v complex128) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Complex128(key, v))
	return e
}
func (e *Event) Complex128p(key string, v *complex128) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Complex128p(key, v))
	return e
}
func (e *Event) Complex64(key string, v complex64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Complex64(key, v))
	return e
}
func (e *Event) Complex64p(key string, v *complex64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Complex64p(key, v))
	return e
}
func (e *Event) Float64(key string, v float64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Float64(key, v))
	return e
}
func (e *Event) Float64p(key string, v *float64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Float64p(key, v))
	return e
}
func (e *Event) Float32(key string, v float32) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Float32(key, v))
	return e
}
func (e *Event) Float32p(key string, v *float32) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Float32p(key, v))
	return e
}
func (e *Event) Int(key string, v int) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int(key, v))
	return e
}
func (e *Event) Intp(key string, v *int) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Intp(key, v))
	return e
}
func (e *Event) Int64(key string, v int64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int64(key, v))
	return e
}
func (e *Event) Int64p(key string, v *int64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int64p(key, v))
	return e
}
func (e *Event) Int32(key string, v int32) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int32(key, v))
	return e
}
func (e *Event) Int32p(key string, v *int32) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int32p(key, v))
	return e
}
func (e *Event) Int16(key string, v int16) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int16(key, v))
	return e
}
func (e *Event) Int16p(key string, v *int16) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int16p(key, v))
	return e
}
func (e *Event) Int8(key string, v int8) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int8(key, v))
	return e
}
func (e *Event) Int8p(key string, v *int8) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Int8p(key, v))
	return e
}
func (e *Event) Uint(key string, v uint) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint(key, v))
	return e
}
func (e *Event) Uintp(key string, v *uint) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uintp(key, v))
	return e
}
func (e *Event) Uint64(key string, v uint64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint64(key, v))
	return e
}
func (e *Event) Uint64p(key string, v *uint64) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint64p(key, v))
	return e
}
func (e *Event) Uint32(key string, v uint32) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint32(key, v))
	return e
}
func (e *Event) Uint32p(key string, v *uint32) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint32p(key, v))
	return e
}
func (e *Event) Uint16(key string, v uint16) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint16(key, v))
	return e
}
func (e *Event) Uint16p(key string, v *uint16) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint16p(key, v))
	return e
}
func (e *Event) Uint8(key string, v uint8) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint8(key, v))
	return e
}
func (e *Event) Uint8p(key string, v *uint8) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uint8p(key, v))
	return e
}
func (e *Event) String(key string, v string) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.String(key, v))
	return e
}
func (e *Event) Stringp(key string, v *string) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Stringp(key, v))
	return e
}
func (e *Event) Uintptr(key string, v uintptr) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uintptr(key, v))
	return e
}
func (e *Event) Uintptrp(key string, v *uintptr) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Uintptrp(key, v))
	return e
}
func (e *Event) Reflect(key string, v any) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Reflect(key, v))
	return e
}
func (e *Event) Namespace(key string) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Namespace(key))
	return e
}
func (e *Event) Stringer(key string, v fmt.Stringer) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Stringer(key, v))
	return e
}
func (e *Event) Time(key string, v time.Time) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Time(key, v))
	return e
}
func (e *Event) Timep(key string, v *time.Time) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Timep(key, v))
	return e
}
func (e *Event) Stack(key string) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Stack(key))
	return e
}
func (e *Event) StackSkip(key string, skip int) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.StackSkip(key, skip))
	return e
}
func (e *Event) Duration(key string, v time.Duration) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Duration(key, v))
	return e
}
func (e *Event) Durationp(key string, v *time.Duration) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Durationp(key, v))
	return e
}
func (e *Event) Object(key string, val ObjectMarshaler) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Object(key, val))
	return e
}
func (e *Event) Inline(val ObjectMarshaler) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Inline(val))
	return e
}
func (e *Event) Dict(key string, val ...Field) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Dict(key, val...))
	return e
}
func (e *Event) Any(key string, v any) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Any(key, v))
	return e
}

func (e *Event) Array(key string, v ArrayMarshaler) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, zap.Array(key, v))
	return e
}

func (e *Event) Caller(depth int) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, DefaultCaller(depth))
	return e
}

func (e *Event) CallerFile(depth int) *Event {
	if e == nil {
		return e
	}
	e.fields = append(e.fields, DefaultCallerFile(depth))
	return e
}
