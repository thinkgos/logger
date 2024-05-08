package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Valuer is returns a log value.
type Valuer func(ctx context.Context) Field

/**************************** immutable Valuer ******************************************/

func wrapperField(field Field) Valuer {
	return func(ctx context.Context) Field {
		return field
	}
}

func ImmutErr(val error) Valuer                    { return wrapperField(zap.Error(val)) }
func ImmutErrors(key string, val []error) Valuer   { return wrapperField(zap.Errors(key, val)) }
func ImmutNamedError(key string, val error) Valuer { return wrapperField(zap.NamedError(key, val)) }

func ImmutBinary(key string, v []byte) Valuer            { return wrapperField(zap.Binary(key, v)) }
func ImmutBool(key string, v bool) Valuer                { return wrapperField(zap.Bool(key, v)) }
func ImmutBoolp(key string, v *bool) Valuer              { return wrapperField(zap.Boolp(key, v)) }
func ImmutByteString(key string, v []byte) Valuer        { return wrapperField(zap.ByteString(key, v)) }
func ImmutComplex128(key string, v complex128) Valuer    { return wrapperField(zap.Complex128(key, v)) }
func ImmutComplex128p(key string, v *complex128) Valuer  { return wrapperField(zap.Complex128p(key, v)) }
func ImmutComplex64(key string, v complex64) Valuer      { return wrapperField(zap.Complex64(key, v)) }
func ImmutComplex64p(key string, v *complex64) Valuer    { return wrapperField(zap.Complex64p(key, v)) }
func ImmutFloat64(key string, v float64) Valuer          { return wrapperField(zap.Float64(key, v)) }
func ImmutFloat64p(key string, v *float64) Valuer        { return wrapperField(zap.Float64p(key, v)) }
func ImmutFloat32(key string, v float32) Valuer          { return wrapperField(zap.Float32(key, v)) }
func ImmutFloat32p(key string, v *float32) Valuer        { return wrapperField(zap.Float32p(key, v)) }
func ImmutInt(key string, v int) Valuer                  { return wrapperField(zap.Int(key, v)) }
func ImmutIntp(key string, v *int) Valuer                { return wrapperField(zap.Intp(key, v)) }
func ImmutInt64(key string, v int64) Valuer              { return wrapperField(zap.Int64(key, v)) }
func ImmutInt64p(key string, v *int64) Valuer            { return wrapperField(zap.Int64p(key, v)) }
func ImmutInt32(key string, v int32) Valuer              { return wrapperField(zap.Int32(key, v)) }
func ImmutInt32p(key string, v *int32) Valuer            { return wrapperField(zap.Int32p(key, v)) }
func ImmutInt16(key string, v int16) Valuer              { return wrapperField(zap.Int16(key, v)) }
func ImmutInt16p(key string, v *int16) Valuer            { return wrapperField(zap.Int16p(key, v)) }
func ImmutInt8(key string, v int8) Valuer                { return wrapperField(zap.Int8(key, v)) }
func ImmutInt8p(key string, v *int8) Valuer              { return wrapperField(zap.Int8p(key, v)) }
func ImmutUint(key string, v uint) Valuer                { return wrapperField(zap.Uint(key, v)) }
func ImmutUintp(key string, v *uint) Valuer              { return wrapperField(zap.Uintp(key, v)) }
func ImmutUint64(key string, v uint64) Valuer            { return wrapperField(zap.Uint64(key, v)) }
func ImmutUint64p(key string, v *uint64) Valuer          { return wrapperField(zap.Uint64p(key, v)) }
func ImmutUint32(key string, v uint32) Valuer            { return wrapperField(zap.Uint32(key, v)) }
func ImmutUint32p(key string, v *uint32) Valuer          { return wrapperField(zap.Uint32p(key, v)) }
func ImmutUint16(key string, v uint16) Valuer            { return wrapperField(zap.Uint16(key, v)) }
func ImmutUint16p(key string, v *uint16) Valuer          { return wrapperField(zap.Uint16p(key, v)) }
func ImmutUint8(key string, v uint8) Valuer              { return wrapperField(zap.Uint8(key, v)) }
func ImmutUint8p(key string, v *uint8) Valuer            { return wrapperField(zap.Uint8p(key, v)) }
func ImmutString(key string, v string) Valuer            { return wrapperField(zap.String(key, v)) }
func ImmutStringp(key string, v *string) Valuer          { return wrapperField(zap.Stringp(key, v)) }
func ImmutUintptr(key string, v uintptr) Valuer          { return wrapperField(zap.Uintptr(key, v)) }
func ImmutUintptrp(key string, v *uintptr) Valuer        { return wrapperField(zap.Uintptrp(key, v)) }
func ImmutReflect(key string, v any) Valuer              { return wrapperField(zap.Reflect(key, v)) }
func ImmutNamespace(key string) Valuer                   { return wrapperField(zap.Namespace(key)) }
func ImmutStringer(key string, v fmt.Stringer) Valuer    { return wrapperField(zap.Stringer(key, v)) }
func ImmutTime(key string, v time.Time) Valuer           { return wrapperField(zap.Time(key, v)) }
func ImmutTimep(key string, v *time.Time) Valuer         { return wrapperField(zap.Timep(key, v)) }
func ImmutStack(key string) Valuer                       { return wrapperField(zap.Stack(key)) }
func ImmutStackSkip(key string, skip int) Valuer         { return wrapperField(zap.StackSkip(key, skip)) }
func ImmutDuration(key string, v time.Duration) Valuer   { return wrapperField(zap.Duration(key, v)) }
func ImmutDurationp(key string, v *time.Duration) Valuer { return wrapperField(zap.Durationp(key, v)) }
func ImmutObject(key string, val ObjectMarshaler) Valuer { return wrapperField(zap.Object(key, val)) }
func ImmutInline(val ObjectMarshaler) Valuer             { return wrapperField(zap.Inline(val)) }
func ImmutDict(key string, val ...Field) Valuer          { return wrapperField(zap.Dict(key, val...)) }
func ImmutAny(key string, v any) Valuer                  { return wrapperField(zap.Any(key, v)) }

/**************************** Dynamic Valuer ******************************************/

func FromErr(vf func(context.Context) error) Valuer {
	return func(ctx context.Context) Field {
		return zap.Error(vf(ctx))
	}
}
func FromErrors(key string, vf func(context.Context) []error) Valuer {
	return func(ctx context.Context) Field {
		return zap.Errors(key, vf(ctx))
	}
}
func FromNamedError(key string, vf func(context.Context) error) Valuer {
	return func(ctx context.Context) Field {
		return zap.NamedError(key, vf(ctx))
	}
}

func FromBinary(key string, vf func(context.Context) []byte) Valuer {
	return func(ctx context.Context) Field {
		return zap.Binary(key, vf(ctx))
	}
}
func FromBool(key string, vf func(context.Context) bool) Valuer {
	return func(ctx context.Context) Field {
		return zap.Bool(key, vf(ctx))
	}
}
func FromBoolp(key string, vf func(context.Context) *bool) Valuer {
	return func(ctx context.Context) Field {
		return zap.Boolp(key, vf(ctx))
	}
}
func FromByteString(key string, vf func(context.Context) []byte) Valuer {
	return func(ctx context.Context) Field {
		return zap.ByteString(key, vf(ctx))
	}
}
func FromComplex128(key string, vf func(context.Context) complex128) Valuer {
	return func(ctx context.Context) Field {
		return zap.Complex128(key, vf(ctx))
	}
}
func FromComplex128p(key string, vf func(context.Context) *complex128) Valuer {
	return func(ctx context.Context) Field {
		return zap.Complex128p(key, vf(ctx))
	}
}
func FromComplex64(key string, vf func(context.Context) complex64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Complex64(key, vf(ctx))
	}
}
func FromComplex64p(key string, vf func(context.Context) *complex64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Complex64p(key, vf(ctx))
	}
}
func FromFloat64(key string, vf func(context.Context) float64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Float64(key, vf(ctx))
	}
}
func FromFloat64p(key string, vf func(context.Context) *float64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Float64p(key, vf(ctx))
	}
}
func FromFloat32(key string, vf func(context.Context) float32) Valuer {
	return func(ctx context.Context) Field {
		return zap.Float32(key, vf(ctx))
	}
}
func FromFloat32p(key string, vf func(context.Context) *float32) Valuer {
	return func(ctx context.Context) Field {
		return zap.Float32p(key, vf(ctx))
	}
}
func FromInt(key string, vf func(context.Context) int) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int(key, vf(ctx))
	}
}
func FromIntp(key string, vf func(context.Context) *int) Valuer {
	return func(ctx context.Context) Field {
		return zap.Intp(key, vf(ctx))
	}
}
func FromInt64(key string, vf func(context.Context) int64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int64(key, vf(ctx))
	}
}
func FromInt64p(key string, vf func(context.Context) *int64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int64p(key, vf(ctx))
	}
}
func FromInt32(key string, vf func(context.Context) int32) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int32(key, vf(ctx))
	}
}
func FromInt32p(key string, vf func(context.Context) *int32) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int32p(key, vf(ctx))
	}
}
func FromInt16(key string, vf func(context.Context) int16) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int16(key, vf(ctx))
	}
}
func FromInt16p(key string, vf func(context.Context) *int16) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int16p(key, vf(ctx))
	}
}
func FromInt8(key string, vf func(context.Context) int8) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int8(key, vf(ctx))
	}
}
func FromInt8p(key string, vf func(context.Context) *int8) Valuer {
	return func(ctx context.Context) Field {
		return zap.Int8p(key, vf(ctx))
	}
}
func FromUint(key string, vf func(context.Context) uint) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint(key, vf(ctx))
	}
}
func FromUintp(key string, vf func(context.Context) *uint) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uintp(key, vf(ctx))
	}
}
func FromUint64(key string, vf func(context.Context) uint64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint64(key, vf(ctx))
	}
}
func FromUint64p(key string, vf func(context.Context) *uint64) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint64p(key, vf(ctx))
	}
}
func FromUint32(key string, vf func(context.Context) uint32) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint32(key, vf(ctx))
	}
}
func FromUint32p(key string, vf func(context.Context) *uint32) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint32p(key, vf(ctx))
	}
}
func FromUint16(key string, vf func(context.Context) uint16) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint16(key, vf(ctx))
	}
}
func FromUint16p(key string, vf func(context.Context) *uint16) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint16p(key, vf(ctx))
	}
}
func FromUint8(key string, vf func(context.Context) uint8) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint8(key, vf(ctx))
	}
}
func FromUint8p(key string, vf func(context.Context) *uint8) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uint8p(key, vf(ctx))
	}
}
func FromString(key string, vf func(context.Context) string) Valuer {
	return func(ctx context.Context) Field {
		return zap.String(key, vf(ctx))
	}
}
func FromStringp(key string, vf func(context.Context) *string) Valuer {
	return func(ctx context.Context) Field {
		return zap.Stringp(key, vf(ctx))
	}
}
func FromUintptr(key string, vf func(context.Context) uintptr) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uintptr(key, vf(ctx))
	}
}
func FromUintptrp(key string, vf func(context.Context) *uintptr) Valuer {
	return func(ctx context.Context) Field {
		return zap.Uintptrp(key, vf(ctx))
	}
}
func FromReflect(key string, vf func(context.Context) any) Valuer {
	return func(ctx context.Context) Field {
		return zap.Reflect(key, vf(ctx))
	}
}
func FromStringer(key string, vf func(context.Context) fmt.Stringer) Valuer {
	return func(ctx context.Context) Field {
		return zap.Stringer(key, vf(ctx))
	}
}
func FromTime(key string, vf func(context.Context) time.Time) Valuer {
	return func(ctx context.Context) Field {
		return zap.Time(key, vf(ctx))
	}
}
func FromTimep(key string, vf func(context.Context) *time.Time) Valuer {
	return func(ctx context.Context) Field {
		return zap.Timep(key, vf(ctx))
	}
}
func FromDuration(key string, vf func(context.Context) time.Duration) Valuer {
	return func(ctx context.Context) Field {
		return zap.Duration(key, vf(ctx))
	}
}
func FromDurationp(key string, vf func(context.Context) *time.Duration) Valuer {
	return func(ctx context.Context) Field {
		return zap.Durationp(key, vf(ctx))
	}
}
func FromAny(key string, vf func(context.Context) any) Valuer {
	return func(ctx context.Context) Field {
		return zap.Any(key, vf(ctx))
	}
}
