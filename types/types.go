// THIS PACKAGE IS AUTO GENERATED -- DO NOT EDIT!

package types

import (
    "unsafe"

    union "github.com/sug0/go-union"
    tagged "github.com/sug0/go-union/tagged"
)

const (
    KindUndefined = iota
    KindBool
    KindString
    KindInt
    KindInt8
    KindInt16
    KindInt32
    KindInt64
    KindUint
    KindUint8
    KindUint16
    KindUint32
    KindUint64
    KindUintptr
    KindByte
    KindRune
    KindFloat32
    KindFloat64
    KindComplex64
    KindComplex128
    KindIota
)

type Bool struct {
    Ptr *bool
}

func (*Bool) UnionKind() tagged.Kind {
    return KindBool
}

func (p *Bool) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Bool) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*bool)(u.Pointer())
    return nil
}

type String struct {
    Ptr *string
}

func (*String) UnionKind() tagged.Kind {
    return KindString
}

func (p *String) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *String) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*string)(u.Pointer())
    return nil
}

type Int struct {
    Ptr *int
}

func (*Int) UnionKind() tagged.Kind {
    return KindInt
}

func (p *Int) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Int) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*int)(u.Pointer())
    return nil
}

type Int8 struct {
    Ptr *int8
}

func (*Int8) UnionKind() tagged.Kind {
    return KindInt8
}

func (p *Int8) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Int8) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*int8)(u.Pointer())
    return nil
}

type Int16 struct {
    Ptr *int16
}

func (*Int16) UnionKind() tagged.Kind {
    return KindInt16
}

func (p *Int16) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Int16) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*int16)(u.Pointer())
    return nil
}

type Int32 struct {
    Ptr *int32
}

func (*Int32) UnionKind() tagged.Kind {
    return KindInt32
}

func (p *Int32) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Int32) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*int32)(u.Pointer())
    return nil
}

type Int64 struct {
    Ptr *int64
}

func (*Int64) UnionKind() tagged.Kind {
    return KindInt64
}

func (p *Int64) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Int64) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*int64)(u.Pointer())
    return nil
}

type Uint struct {
    Ptr *uint
}

func (*Uint) UnionKind() tagged.Kind {
    return KindUint
}

func (p *Uint) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Uint) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*uint)(u.Pointer())
    return nil
}

type Uint8 struct {
    Ptr *uint8
}

func (*Uint8) UnionKind() tagged.Kind {
    return KindUint8
}

func (p *Uint8) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Uint8) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*uint8)(u.Pointer())
    return nil
}

type Uint16 struct {
    Ptr *uint16
}

func (*Uint16) UnionKind() tagged.Kind {
    return KindUint16
}

func (p *Uint16) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Uint16) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*uint16)(u.Pointer())
    return nil
}

type Uint32 struct {
    Ptr *uint32
}

func (*Uint32) UnionKind() tagged.Kind {
    return KindUint32
}

func (p *Uint32) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Uint32) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*uint32)(u.Pointer())
    return nil
}

type Uint64 struct {
    Ptr *uint64
}

func (*Uint64) UnionKind() tagged.Kind {
    return KindUint64
}

func (p *Uint64) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Uint64) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*uint64)(u.Pointer())
    return nil
}

type Uintptr struct {
    Ptr *uintptr
}

func (*Uintptr) UnionKind() tagged.Kind {
    return KindUintptr
}

func (p *Uintptr) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Uintptr) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*uintptr)(u.Pointer())
    return nil
}

type Byte struct {
    Ptr *byte
}

func (*Byte) UnionKind() tagged.Kind {
    return KindByte
}

func (p *Byte) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Byte) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*byte)(u.Pointer())
    return nil
}

type Rune struct {
    Ptr *rune
}

func (*Rune) UnionKind() tagged.Kind {
    return KindRune
}

func (p *Rune) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Rune) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*rune)(u.Pointer())
    return nil
}

type Float32 struct {
    Ptr *float32
}

func (*Float32) UnionKind() tagged.Kind {
    return KindFloat32
}

func (p *Float32) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Float32) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*float32)(u.Pointer())
    return nil
}

type Float64 struct {
    Ptr *float64
}

func (*Float64) UnionKind() tagged.Kind {
    return KindFloat64
}

func (p *Float64) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Float64) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*float64)(u.Pointer())
    return nil
}

type Complex64 struct {
    Ptr *complex64
}

func (*Complex64) UnionKind() tagged.Kind {
    return KindComplex64
}

func (p *Complex64) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Complex64) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*complex64)(u.Pointer())
    return nil
}

type Complex128 struct {
    Ptr *complex128
}

func (*Complex128) UnionKind() tagged.Kind {
    return KindComplex128
}

func (p *Complex128) UnionSize() uintptr {
    return unsafe.Sizeof(*p.Ptr)
}

func (p *Complex128) CastUnion(u union.Union) error {
    if u.Cap() < p.UnionSize() {
        return union.ErrTooSmall
    }
    p.Ptr = (*complex128)(u.Pointer())
    return nil
}
