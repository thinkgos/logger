package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

/**************************** immutable Valuer ******************************************/

func wrapperField(field Field) HookFunc {
	return func(ctx context.Context) Field {
		return field
	}
}

func ImmutErr(val error) HookFunc                  { return wrapperField(zap.Error(val)) }
func ImmutErrors(key string, val []error) HookFunc { return wrapperField(zap.Errors(key, val)) }
func ImmutNamedError(key string, val error) HookFunc {
	return wrapperField(zap.NamedError(key, val))
}

func ImmutBinary(key string, v []byte) HookFunc     { return wrapperField(zap.Binary(key, v)) }
func ImmutBool(key string, v bool) HookFunc         { return wrapperField(zap.Bool(key, v)) }
func ImmutBoolp(key string, v *bool) HookFunc       { return wrapperField(zap.Boolp(key, v)) }
func ImmutByteString(key string, v []byte) HookFunc { return wrapperField(zap.ByteString(key, v)) }
func ImmutComplex128(key string, v complex128) HookFunc {
	return wrapperField(zap.Complex128(key, v))
}
func ImmutComplex128p(key string, v *complex128) HookFunc {
	return wrapperField(zap.Complex128p(key, v))
}
func ImmutComplex64(key string, v complex64) HookFunc { return wrapperField(zap.Complex64(key, v)) }
func ImmutComplex64p(key string, v *complex64) HookFunc {
	return wrapperField(zap.Complex64p(key, v))
}
func ImmutFloat64(key string, v float64) HookFunc       { return wrapperField(zap.Float64(key, v)) }
func ImmutFloat64p(key string, v *float64) HookFunc     { return wrapperField(zap.Float64p(key, v)) }
func ImmutFloat32(key string, v float32) HookFunc       { return wrapperField(zap.Float32(key, v)) }
func ImmutFloat32p(key string, v *float32) HookFunc     { return wrapperField(zap.Float32p(key, v)) }
func ImmutInt(key string, v int) HookFunc               { return wrapperField(zap.Int(key, v)) }
func ImmutIntp(key string, v *int) HookFunc             { return wrapperField(zap.Intp(key, v)) }
func ImmutInt64(key string, v int64) HookFunc           { return wrapperField(zap.Int64(key, v)) }
func ImmutInt64p(key string, v *int64) HookFunc         { return wrapperField(zap.Int64p(key, v)) }
func ImmutInt32(key string, v int32) HookFunc           { return wrapperField(zap.Int32(key, v)) }
func ImmutInt32p(key string, v *int32) HookFunc         { return wrapperField(zap.Int32p(key, v)) }
func ImmutInt16(key string, v int16) HookFunc           { return wrapperField(zap.Int16(key, v)) }
func ImmutInt16p(key string, v *int16) HookFunc         { return wrapperField(zap.Int16p(key, v)) }
func ImmutInt8(key string, v int8) HookFunc             { return wrapperField(zap.Int8(key, v)) }
func ImmutInt8p(key string, v *int8) HookFunc           { return wrapperField(zap.Int8p(key, v)) }
func ImmutUint(key string, v uint) HookFunc             { return wrapperField(zap.Uint(key, v)) }
func ImmutUintp(key string, v *uint) HookFunc           { return wrapperField(zap.Uintp(key, v)) }
func ImmutUint64(key string, v uint64) HookFunc         { return wrapperField(zap.Uint64(key, v)) }
func ImmutUint64p(key string, v *uint64) HookFunc       { return wrapperField(zap.Uint64p(key, v)) }
func ImmutUint32(key string, v uint32) HookFunc         { return wrapperField(zap.Uint32(key, v)) }
func ImmutUint32p(key string, v *uint32) HookFunc       { return wrapperField(zap.Uint32p(key, v)) }
func ImmutUint16(key string, v uint16) HookFunc         { return wrapperField(zap.Uint16(key, v)) }
func ImmutUint16p(key string, v *uint16) HookFunc       { return wrapperField(zap.Uint16p(key, v)) }
func ImmutUint8(key string, v uint8) HookFunc           { return wrapperField(zap.Uint8(key, v)) }
func ImmutUint8p(key string, v *uint8) HookFunc         { return wrapperField(zap.Uint8p(key, v)) }
func ImmutString(key string, v string) HookFunc         { return wrapperField(zap.String(key, v)) }
func ImmutStringp(key string, v *string) HookFunc       { return wrapperField(zap.Stringp(key, v)) }
func ImmutUintptr(key string, v uintptr) HookFunc       { return wrapperField(zap.Uintptr(key, v)) }
func ImmutUintptrp(key string, v *uintptr) HookFunc     { return wrapperField(zap.Uintptrp(key, v)) }
func ImmutReflect(key string, v any) HookFunc           { return wrapperField(zap.Reflect(key, v)) }
func ImmutNamespace(key string) HookFunc                { return wrapperField(zap.Namespace(key)) }
func ImmutStringer(key string, v fmt.Stringer) HookFunc { return wrapperField(zap.Stringer(key, v)) }
func ImmutTime(key string, v time.Time) HookFunc        { return wrapperField(zap.Time(key, v)) }
func ImmutTimep(key string, v *time.Time) HookFunc      { return wrapperField(zap.Timep(key, v)) }
func ImmutStack(key string) HookFunc                    { return wrapperField(zap.Stack(key)) }
func ImmutStackSkip(key string, skip int) HookFunc      { return wrapperField(zap.StackSkip(key, skip)) }
func ImmutDuration(key string, v time.Duration) HookFunc {
	return wrapperField(zap.Duration(key, v))
}
func ImmutDurationp(key string, v *time.Duration) HookFunc {
	return wrapperField(zap.Durationp(key, v))
}
func ImmutObject(key string, val ObjectMarshaler) HookFunc {
	return wrapperField(zap.Object(key, val))
}
func ImmutInline(val ObjectMarshaler) HookFunc    { return wrapperField(zap.Inline(val)) }
func ImmutDict(key string, val ...Field) HookFunc { return wrapperField(zap.Dict(key, val...)) }
func ImmutAny(key string, v any) HookFunc         { return wrapperField(zap.Any(key, v)) }
