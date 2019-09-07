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

// Returns the zero value of the field 'fieldNum' of
// a Union described by 'aStruct', which is a struct
// where all its fields are UnionCaster implementers.
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
        caster, ok := aStruct.Field(i).Interface().(UnionCaster)
        if !ok {
            return Union{}, union.ErrNotCaster
        }
        if size := caster.UnionSize(); size > maxSize {
            maxSize = size
        }
    }

    k := aStruct.Field(fieldNum).Interface().(UnionKind)

    return Union{
        kind: k.UnionKind(),
        union: union.WithCapacity(maxSize),
    }, nil
}

// Returns the type of the Union's value that is currently in use.
func (u *Union) Kind() Kind {
    return Kind(atomic.LoadInt32((*int32)(&u.kind)))
}

func (u *Union) changeKind(to Kind) {
    atomic.StoreInt32((*int32)(&u.kind), int32(to))
}

// Changes the Union's value that is currently in use.
// Use is governed by the type that implements UnionCaster.
func (u *Union) Cast(caster UnionCaster) {
    u.changeKind(caster.UnionKind())
    caster.CastUnion(u.union)
}
