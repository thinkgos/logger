package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type ImmutableField[T any] struct {
	Value T
}

type ImmutableNamedField[T any] struct {
	Key   string
	Value T
}

type ImmutableErr ImmutableField[error]

func (m ImmutableErr) DoHook(ctx context.Context) Field {
	return zap.Error(m.Value)
}

type ImmutableErrors ImmutableNamedField[[]error]

func (m ImmutableErrors) DoHook(ctx context.Context) Field {
	return zap.Errors(m.Key, m.Value)
}

type ImmutableNamedError ImmutableNamedField[error]

func (m ImmutableNamedError) DoHook(ctx context.Context) Field {
	return zap.NamedError(m.Key, m.Value)
}

type ImmutableBinary ImmutableNamedField[[]byte]

func (m ImmutableBinary) DoHook(ctx context.Context) Field {
	return zap.Binary(m.Key, m.Value)
}

type ImmutableBool ImmutableNamedField[bool]

func (m *ImmutableBool) DoHook(ctx context.Context) Field {
	return zap.Bool(m.Key, m.Value)
}

type ImmutableBoolp ImmutableNamedField[*bool]

func (m *ImmutableBoolp) DoHook(ctx context.Context) Field {
	return zap.Boolp(m.Key, m.Value)
}

type ImmutableByteString ImmutableNamedField[[]byte]

func (m *ImmutableByteString) DoHook(ctx context.Context) Field {
	return zap.ByteString(m.Key, m.Value)
}

type ImmutableComplex128 ImmutableNamedField[complex128]

func (m *ImmutableComplex128) DoHook(ctx context.Context) Field {
	return zap.Complex128(m.Key, m.Value)
}

type ImmutableComplex128p ImmutableNamedField[*complex128]

func (m *ImmutableComplex128p) DoHook(ctx context.Context) Field {
	return zap.Complex128p(m.Key, m.Value)
}

type ImmutableComplex64 ImmutableNamedField[complex64]

func (m *ImmutableComplex64) DoHook(ctx context.Context) Field {
	return zap.Complex64(m.Key, m.Value)
}

type ImmutableComplex64p ImmutableNamedField[*complex64]

func (m *ImmutableComplex64p) DoHook(ctx context.Context) Field {
	return zap.Complex64p(m.Key, m.Value)
}

type ImmutableFloat64 ImmutableNamedField[float64]

func (m *ImmutableFloat64) DoHook(ctx context.Context) Field {
	return zap.Float64(m.Key, m.Value)
}

type ImmutableFloat64p ImmutableNamedField[*float64]

func (m *ImmutableFloat64p) DoHook(ctx context.Context) Field {
	return zap.Float64p(m.Key, m.Value)
}

type ImmutableFloat32 ImmutableNamedField[float32]

func (m *ImmutableFloat32) DoHook(ctx context.Context) Field {
	return zap.Float32(m.Key, m.Value)
}

type ImmutableFloat32p ImmutableNamedField[*float32]

func (m *ImmutableFloat32p) DoHook(ctx context.Context) Field {
	return zap.Float32p(m.Key, m.Value)
}

type ImmutableInt ImmutableNamedField[int]

func (m *ImmutableInt) DoHook(ctx context.Context) Field {
	return zap.Int(m.Key, m.Value)
}

type ImmutableIntp ImmutableNamedField[*int]

func (m *ImmutableIntp) DoHook(ctx context.Context) Field {
	return zap.Intp(m.Key, m.Value)
}

type ImmutableInt64 ImmutableNamedField[int64]

func (m *ImmutableInt64) DoHook(ctx context.Context) Field {
	return zap.Int64(m.Key, m.Value)
}

type ImmutableInt64p ImmutableNamedField[*int64]

func (m *ImmutableInt64p) DoHook(ctx context.Context) Field {
	return zap.Int64p(m.Key, m.Value)
}

type ImmutableInt32 ImmutableNamedField[int32]

func (m *ImmutableInt32) DoHook(ctx context.Context) Field {
	return zap.Int32(m.Key, m.Value)
}

type ImmutableInt32p ImmutableNamedField[*int32]

func (m *ImmutableInt32p) DoHook(ctx context.Context) Field {
	return zap.Int32p(m.Key, m.Value)
}

type ImmutableInt16 ImmutableNamedField[int16]

func (m *ImmutableInt16) DoHook(ctx context.Context) Field {
	return zap.Int16(m.Key, m.Value)
}

type ImmutableInt16p ImmutableNamedField[*int16]

func (m *ImmutableInt16p) DoHook(ctx context.Context) Field {
	return zap.Int16p(m.Key, m.Value)
}

type ImmutableInt8 ImmutableNamedField[int8]

func (m *ImmutableInt8) DoHook(ctx context.Context) Field {
	return zap.Int8(m.Key, m.Value)
}

type ImmutableInt8p ImmutableNamedField[*int8]

func (m *ImmutableInt8p) DoHook(ctx context.Context) Field {
	return zap.Int8p(m.Key, m.Value)
}

type ImmutableUint ImmutableNamedField[uint]

func (m *ImmutableUint) DoHook(ctx context.Context) Field {
	return zap.Uint(m.Key, m.Value)
}

type ImmutableUintp ImmutableNamedField[*uint]

func (m *ImmutableUintp) DoHook(ctx context.Context) Field {
	return zap.Uintp(m.Key, m.Value)
}

type ImmutableUint64 ImmutableNamedField[uint64]

func (m *ImmutableUint64) DoHook(ctx context.Context) Field {
	return zap.Uint64(m.Key, m.Value)
}

type ImmutableUint64p ImmutableNamedField[*uint64]

func (m *ImmutableUint64p) DoHook(ctx context.Context) Field {
	return zap.Uint64p(m.Key, m.Value)
}

type ImmutableUint32 ImmutableNamedField[uint32]

func (m *ImmutableUint32) DoHook(ctx context.Context) Field {
	return zap.Uint32(m.Key, m.Value)
}

type ImmutableUint32p ImmutableNamedField[*uint32]

func (m *ImmutableUint32p) DoHook(ctx context.Context) Field {
	return zap.Uint32p(m.Key, m.Value)
}

type ImmutableUint16 ImmutableNamedField[uint16]

func (m *ImmutableUint16) DoHook(ctx context.Context) Field {
	return zap.Uint16(m.Key, m.Value)
}

type ImmutableUint16p ImmutableNamedField[*uint16]

func (m *ImmutableUint16p) DoHook(ctx context.Context) Field {
	return zap.Uint16p(m.Key, m.Value)
}

type ImmutableUint8 ImmutableNamedField[uint8]

func (m *ImmutableUint8) DoHook(ctx context.Context) Field {
	return zap.Uint8(m.Key, m.Value)
}

type ImmutableUint8p ImmutableNamedField[*uint8]

func (m *ImmutableUint8p) DoHook(ctx context.Context) Field {
	return zap.Uint8p(m.Key, m.Value)
}

type ImmutableString ImmutableNamedField[string]

func (m *ImmutableString) DoHook(ctx context.Context) Field {
	return zap.String(m.Key, m.Value)
}

type ImmutableStringp ImmutableNamedField[*string]

func (m *ImmutableStringp) DoHook(ctx context.Context) Field {
	return zap.Stringp(m.Key, m.Value)
}

type ImmutableUintptr ImmutableNamedField[uintptr]

func (m *ImmutableUintptr) DoHook(ctx context.Context) Field {
	return zap.Uintptr(m.Key, m.Value)
}

type ImmutableUintptrp ImmutableNamedField[*uintptr]

func (m *ImmutableUintptrp) DoHook(ctx context.Context) Field {
	return zap.Uintptrp(m.Key, m.Value)
}

type ImmutableReflect ImmutableNamedField[any]

func (m *ImmutableReflect) DoHook(ctx context.Context) Field {
	return zap.Reflect(m.Key, m.Value)
}

type ImmutableNamespace ImmutableField[string]

func (m *ImmutableNamespace) DoHook(ctx context.Context) Field {
	return zap.Namespace(m.Value)
}

type ImmutableStringer ImmutableNamedField[fmt.Stringer]

func (m *ImmutableStringer) DoHook(ctx context.Context) Field {
	return zap.Stringer(m.Key, m.Value)
}

type ImmutableTime ImmutableNamedField[time.Time]

func (m *ImmutableTime) DoHook(ctx context.Context) Field {
	return zap.Time(m.Key, m.Value)
}

type ImmutableTimep ImmutableNamedField[*time.Time]

func (m *ImmutableTimep) DoHook(ctx context.Context) Field {
	return zap.Timep(m.Key, m.Value)
}

type ImmutableStack struct {
	Key string
}

func (m *ImmutableStack) DoHook(ctx context.Context) Field {
	return zap.Stack(m.Key)
}

type ImmutableStackSkip struct {
	Key  string
	Skip int
}

func (m *ImmutableStackSkip) DoHook(ctx context.Context) Field {
	return zap.StackSkip(m.Key, m.Skip)
}

type ImmutableDuration ImmutableNamedField[time.Duration]

func (m *ImmutableDuration) DoHook(ctx context.Context) Field {
	return zap.Duration(m.Key, m.Value)
}

type ImmutableDurationp ImmutableNamedField[*time.Duration]

func (m *ImmutableDurationp) DoHook(ctx context.Context) Field {
	return zap.Durationp(m.Key, m.Value)
}

type ImmutableObject ImmutableNamedField[ObjectMarshaler]

func (m *ImmutableObject) DoHook(ctx context.Context) Field {
	return zap.Object(m.Key, m.Value)
}

type ImmutableInline ImmutableField[ObjectMarshaler]

func (m *ImmutableInline) DoHook(ctx context.Context) Field {
	return zap.Inline(m.Value)
}

type ImmutableDict ImmutableNamedField[[]Field]

func (m *ImmutableDict) DoHook(ctx context.Context) Field {
	return zap.Dict(m.Key, m.Value...)
}

type ImmutableAny ImmutableNamedField[any]

func (m *ImmutableAny) DoHook(ctx context.Context) Field {
	return zap.Any(m.Key, m.Value)
}
