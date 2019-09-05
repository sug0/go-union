package tagged

import (
    "reflect"
    "sync/atomic"

    union "github.com/sug0/go-union"
)

// The type that is currently in use
// by the union.
type Kind int32

// A tagged Union type.
type Union struct {
    kind  Kind
    union union.Union
}

func Zero(fieldNum int, aStruct reflect.Value) (Union, error) {
    // aStruct needs to be, well, a struct
    if aStruct.Kind() != reflect.Struct {
        return Union{}, union.ErrInvalidStruct
    }

    numFields := aStruct.NumField()

    if numFields == 0 {
        return Union{}, union.ErrInvalidStruct
    }
    if fieldNum < 0 || fieldNum >= numFields {
        return Union{}, union.ErrInvalidIndex
    }

    var maxSize uintptr

    for i := 0; i < numFields; i++ {
        field := aStruct.Field(i)
        if _, ok := field.Interface().(UnionKind); !ok {
            return Union{}, union.ErrNotKind
        }
        if size := field.Type().Size(); size > maxSize {
            maxSize = size
        }
    }

    k := aStruct.Field(fieldNum).Interface().(UnionKind)

    return Union{
        kind: k.UnionKind(),
        union: union.WithCapacity(maxSize),
    }, nil
}

func (u *Union) Kind() Kind {
    return Kind(atomic.LoadInt32((*int32)(&u.kind)))
}

func (u *Union) changeKind(to Kind) {
    atomic.StoreInt32((*int32)(&u.kind), int32(to))
}

func (u *Union) Cast(caster UnionCaster) {
    u.changeKind(caster.UnionKind())
    caster.CastUnion(u.union)
}
