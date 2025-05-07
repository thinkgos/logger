package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Valuer is returns a log value.
//
// Deprecated: As of 0.4.0, use [HookFunc] instead.
type Valuer = HookFunc

/**************************** Dynamic Valuer ******************************************/

func FromErr(vf func(context.Context) error) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Error(vf(ctx))
	}
}
func FromErrors(key string, vf func(context.Context) []error) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Errors(key, vf(ctx))
	}
}
func FromNamedError(key string, vf func(context.Context) error) HookFunc {
	return func(ctx context.Context) Field {
		return zap.NamedError(key, vf(ctx))
	}
}

func FromBinary(key string, vf func(context.Context) []byte) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Binary(key, vf(ctx))
	}
}
func FromBool(key string, vf func(context.Context) bool) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Bool(key, vf(ctx))
	}
}
func FromBoolp(key string, vf func(context.Context) *bool) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Boolp(key, vf(ctx))
	}
}
func FromByteString(key string, vf func(context.Context) []byte) HookFunc {
	return func(ctx context.Context) Field {
		return zap.ByteString(key, vf(ctx))
	}
}
func FromComplex128(key string, vf func(context.Context) complex128) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Complex128(key, vf(ctx))
	}
}
func FromComplex128p(key string, vf func(context.Context) *complex128) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Complex128p(key, vf(ctx))
	}
}
func FromComplex64(key string, vf func(context.Context) complex64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Complex64(key, vf(ctx))
	}
}
func FromComplex64p(key string, vf func(context.Context) *complex64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Complex64p(key, vf(ctx))
	}
}
func FromFloat64(key string, vf func(context.Context) float64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Float64(key, vf(ctx))
	}
}
func FromFloat64p(key string, vf func(context.Context) *float64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Float64p(key, vf(ctx))
	}
}
func FromFloat32(key string, vf func(context.Context) float32) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Float32(key, vf(ctx))
	}
}
func FromFloat32p(key string, vf func(context.Context) *float32) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Float32p(key, vf(ctx))
	}
}
func FromInt(key string, vf func(context.Context) int) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int(key, vf(ctx))
	}
}
func FromIntp(key string, vf func(context.Context) *int) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Intp(key, vf(ctx))
	}
}
func FromInt64(key string, vf func(context.Context) int64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int64(key, vf(ctx))
	}
}
func FromInt64p(key string, vf func(context.Context) *int64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int64p(key, vf(ctx))
	}
}
func FromInt32(key string, vf func(context.Context) int32) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int32(key, vf(ctx))
	}
}
func FromInt32p(key string, vf func(context.Context) *int32) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int32p(key, vf(ctx))
	}
}
func FromInt16(key string, vf func(context.Context) int16) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int16(key, vf(ctx))
	}
}
func FromInt16p(key string, vf func(context.Context) *int16) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int16p(key, vf(ctx))
	}
}
func FromInt8(key string, vf func(context.Context) int8) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int8(key, vf(ctx))
	}
}
func FromInt8p(key string, vf func(context.Context) *int8) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Int8p(key, vf(ctx))
	}
}
func FromUint(key string, vf func(context.Context) uint) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint(key, vf(ctx))
	}
}
func FromUintp(key string, vf func(context.Context) *uint) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uintp(key, vf(ctx))
	}
}
func FromUint64(key string, vf func(context.Context) uint64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint64(key, vf(ctx))
	}
}
func FromUint64p(key string, vf func(context.Context) *uint64) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint64p(key, vf(ctx))
	}
}
func FromUint32(key string, vf func(context.Context) uint32) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint32(key, vf(ctx))
	}
}
func FromUint32p(key string, vf func(context.Context) *uint32) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint32p(key, vf(ctx))
	}
}
func FromUint16(key string, vf func(context.Context) uint16) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint16(key, vf(ctx))
	}
}
func FromUint16p(key string, vf func(context.Context) *uint16) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint16p(key, vf(ctx))
	}
}
func FromUint8(key string, vf func(context.Context) uint8) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint8(key, vf(ctx))
	}
}
func FromUint8p(key string, vf func(context.Context) *uint8) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uint8p(key, vf(ctx))
	}
}
func FromString(key string, vf func(context.Context) string) HookFunc {
	return func(ctx context.Context) Field {
		return zap.String(key, vf(ctx))
	}
}
func FromStringp(key string, vf func(context.Context) *string) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Stringp(key, vf(ctx))
	}
}
func FromUintptr(key string, vf func(context.Context) uintptr) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uintptr(key, vf(ctx))
	}
}
func FromUintptrp(key string, vf func(context.Context) *uintptr) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Uintptrp(key, vf(ctx))
	}
}
func FromReflect(key string, vf func(context.Context) any) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Reflect(key, vf(ctx))
	}
}
func FromStringer(key string, vf func(context.Context) fmt.Stringer) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Stringer(key, vf(ctx))
	}
}
func FromTime(key string, vf func(context.Context) time.Time) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Time(key, vf(ctx))
	}
}
func FromTimep(key string, vf func(context.Context) *time.Time) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Timep(key, vf(ctx))
	}
}
func FromDuration(key string, vf func(context.Context) time.Duration) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Duration(key, vf(ctx))
	}
}
func FromDurationp(key string, vf func(context.Context) *time.Duration) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Durationp(key, vf(ctx))
	}
}
func FromAny(key string, vf func(context.Context) any) HookFunc {
	return func(ctx context.Context) Field {
		return zap.Any(key, vf(ctx))
	}
}
