package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type MutableFixedNamedField[T any] struct {
	Fc func(context.Context) T
}

type MutableNamedField[T any] struct {
	Key string
	Fc  func(context.Context) T
}

type MutableErr MutableFixedNamedField[error]

func (m *MutableErr) DoHook(ctx context.Context) Field {
	return zap.Error(m.Fc(ctx))
}

type MutableErrors MutableNamedField[[]error]

func (m *MutableErrors) DoHook(ctx context.Context) Field {
	return zap.Errors(m.Key, m.Fc(ctx))
}

type MutableNamedError MutableNamedField[error]

func (m *MutableNamedError) DoHook(ctx context.Context) Field {
	return zap.NamedError(m.Key, m.Fc(ctx))
}

type MutableBinary MutableNamedField[[]byte]

func (m *MutableBinary) DoHook(ctx context.Context) Field {
	return zap.Binary(m.Key, m.Fc(ctx))
}

type MutableBool MutableNamedField[bool]

func (m *MutableBool) DoHook(ctx context.Context) Field {
	return zap.Bool(m.Key, m.Fc(ctx))
}

type MutableBoolp MutableNamedField[*bool]

func (m *MutableBoolp) DoHook(ctx context.Context) Field {
	return zap.Boolp(m.Key, m.Fc(ctx))
}

type MutableByteString MutableNamedField[[]byte]

func (m *MutableByteString) DoHook(ctx context.Context) Field {
	return zap.ByteString(m.Key, m.Fc(ctx))
}

type MutableComplex128 MutableNamedField[complex128]

func (m *MutableComplex128) DoHook(ctx context.Context) Field {
	return zap.Complex128(m.Key, m.Fc(ctx))
}

type MutableComplex128p MutableNamedField[*complex128]

func (m *MutableComplex128p) DoHook(ctx context.Context) Field {
	return zap.Complex128p(m.Key, m.Fc(ctx))
}

type MutableComplex64 MutableNamedField[complex64]

func (m *MutableComplex64) DoHook(ctx context.Context) Field {
	return zap.Complex64(m.Key, m.Fc(ctx))
}

type MutableComplex64p MutableNamedField[*complex64]

func (m *MutableComplex64p) DoHook(ctx context.Context) Field {
	return zap.Complex64p(m.Key, m.Fc(ctx))
}

type MutableFloat64 MutableNamedField[float64]

func (m *MutableFloat64) DoHook(ctx context.Context) Field {
	return zap.Float64(m.Key, m.Fc(ctx))
}

type MutableFloat64p MutableNamedField[*float64]

func (m *MutableFloat64p) DoHook(ctx context.Context) Field {
	return zap.Float64p(m.Key, m.Fc(ctx))
}

type MutableFloat32 MutableNamedField[float32]

func (m *MutableFloat32) DoHook(ctx context.Context) Field {
	return zap.Float32(m.Key, m.Fc(ctx))
}

type MutableFloat32p MutableNamedField[*float32]

func (m *MutableFloat32p) DoHook(ctx context.Context) Field {
	return zap.Float32p(m.Key, m.Fc(ctx))
}

type MutableInt MutableNamedField[int]

func (m *MutableInt) DoHook(ctx context.Context) Field {
	return zap.Int(m.Key, m.Fc(ctx))
}

type MutableIntp MutableNamedField[*int]

func (m *MutableIntp) DoHook(ctx context.Context) Field {
	return zap.Intp(m.Key, m.Fc(ctx))
}

type MutableInt64 MutableNamedField[int64]

func (m *MutableInt64) DoHook(ctx context.Context) Field {
	return zap.Int64(m.Key, m.Fc(ctx))
}

type MutableInt64p MutableNamedField[*int64]

func (m *MutableInt64p) DoHook(ctx context.Context) Field {
	return zap.Int64p(m.Key, m.Fc(ctx))
}

type MutableInt32 MutableNamedField[int32]

func (m *MutableInt32) DoHook(ctx context.Context) Field {
	return zap.Int32(m.Key, m.Fc(ctx))
}

type MutableInt32p MutableNamedField[*int32]

func (m *MutableInt32p) DoHook(ctx context.Context) Field {
	return zap.Int32p(m.Key, m.Fc(ctx))
}

type MutableInt16 MutableNamedField[int16]

func (m *MutableInt16) DoHook(ctx context.Context) Field {
	return zap.Int16(m.Key, m.Fc(ctx))
}

type MutableInt16p MutableNamedField[*int16]

func (m *MutableInt16p) DoHook(ctx context.Context) Field {
	return zap.Int16p(m.Key, m.Fc(ctx))
}

type MutableInt8 MutableNamedField[int8]

func (m *MutableInt8) DoHook(ctx context.Context) Field {
	return zap.Int8(m.Key, m.Fc(ctx))
}

type MutableInt8p MutableNamedField[*int8]

func (m *MutableInt8p) DoHook(ctx context.Context) Field {
	return zap.Int8p(m.Key, m.Fc(ctx))
}

type MutableUint MutableNamedField[uint]

func (m *MutableUint) DoHook(ctx context.Context) Field {
	return zap.Uint(m.Key, m.Fc(ctx))
}

type MutableUintp MutableNamedField[*uint]

func (m *MutableUintp) DoHook(ctx context.Context) Field {
	return zap.Uintp(m.Key, m.Fc(ctx))
}

type MutableUint64 MutableNamedField[uint64]

func (m *MutableUint64) DoHook(ctx context.Context) Field {
	return zap.Uint64(m.Key, m.Fc(ctx))
}

type MutableUint64p MutableNamedField[*uint64]

func (m *MutableUint64p) DoHook(ctx context.Context) Field {
	return zap.Uint64p(m.Key, m.Fc(ctx))
}

type MutableUint32 MutableNamedField[uint32]

func (m *MutableUint32) DoHook(ctx context.Context) Field {
	return zap.Uint32(m.Key, m.Fc(ctx))
}

type MutableUint32p MutableNamedField[*uint32]

func (m *MutableUint32p) DoHook(ctx context.Context) Field {
	return zap.Uint32p(m.Key, m.Fc(ctx))
}

type MutableUint16 MutableNamedField[uint16]

func (m *MutableUint16) DoHook(ctx context.Context) Field {
	return zap.Uint16(m.Key, m.Fc(ctx))
}

type MutableUint16p MutableNamedField[*uint16]

func (m *MutableUint16p) DoHook(ctx context.Context) Field {
	return zap.Uint16p(m.Key, m.Fc(ctx))
}

type MutableUint8 MutableNamedField[uint8]

func (m *MutableUint8) DoHook(ctx context.Context) Field {
	return zap.Uint8(m.Key, m.Fc(ctx))
}

type MutableUint8p MutableNamedField[*uint8]

func (m *MutableUint8p) DoHook(ctx context.Context) Field {
	return zap.Uint8p(m.Key, m.Fc(ctx))
}

type MutableString MutableNamedField[string]

func (m *MutableString) DoHook(ctx context.Context) Field {
	return zap.String(m.Key, m.Fc(ctx))
}

type MutableStringp MutableNamedField[*string]

func (m *MutableStringp) DoHook(ctx context.Context) Field {
	return zap.Stringp(m.Key, m.Fc(ctx))
}

type MutableUintptr MutableNamedField[uintptr]

func (m *MutableUintptr) DoHook(ctx context.Context) Field {
	return zap.Uintptr(m.Key, m.Fc(ctx))
}

type MutableUintptrp MutableNamedField[*uintptr]

func (m *MutableUintptrp) DoHook(ctx context.Context) Field {
	return zap.Uintptrp(m.Key, m.Fc(ctx))
}

type MutableReflect MutableNamedField[any]

func (m *MutableReflect) DoHook(ctx context.Context) Field {
	return zap.Reflect(m.Key, m.Fc(ctx))
}

type MutableNamespace MutableFixedNamedField[string]

func (m *MutableNamespace) DoHook(ctx context.Context) Field {
	return zap.Namespace(m.Fc(ctx))
}

type MutableStringer MutableNamedField[fmt.Stringer]

func (m *MutableStringer) DoHook(ctx context.Context) Field {
	return zap.Stringer(m.Key, m.Fc(ctx))
}

type MutableTime MutableNamedField[time.Time]

func (m *MutableTime) DoHook(ctx context.Context) Field {
	return zap.Time(m.Key, m.Fc(ctx))
}

type MutableTimep MutableNamedField[*time.Time]

func (m *MutableTimep) DoHook(ctx context.Context) Field {
	return zap.Timep(m.Key, m.Fc(ctx))
}

type MutableDuration MutableNamedField[time.Duration]

func (m *MutableDuration) DoHook(ctx context.Context) Field {
	return zap.Duration(m.Key, m.Fc(ctx))
}

type MutableDurationp MutableNamedField[*time.Duration]

func (m *MutableDurationp) DoHook(ctx context.Context) Field {
	return zap.Durationp(m.Key, m.Fc(ctx))
}

type MutableObject MutableNamedField[ObjectMarshaler]

func (m *MutableObject) DoHook(ctx context.Context) Field {
	return zap.Object(m.Key, m.Fc(ctx))
}

type MutableInline MutableFixedNamedField[ObjectMarshaler]

func (m *MutableInline) DoHook(ctx context.Context) Field {
	return zap.Inline(m.Fc(ctx))
}

type MutableDict MutableNamedField[[]Field]

func (m *MutableDict) DoHook(ctx context.Context) Field {
	return zap.Dict(m.Key, m.Fc(ctx)...)
}

type MutableAny MutableNamedField[any]

func (m *MutableAny) DoHook(ctx context.Context) Field {
	return zap.Any(m.Key, m.Fc(ctx))
}
