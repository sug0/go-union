package tagged

import (
    "reflect"

    union "github.com/sug0/go-union"
)

// The type that is currently in use
// by the union.
type Kind int

// A tagged Union type.
type Union struct {
    kind  Kind
    union union.Union
}

// Returns a zeroed Union described by 'aStruct',
// which is a struct where all its fields are
// UnionCaster implementers.
func Uninit(aStruct interface{}) (*Union, error) {
    s := reflect.ValueOf(aStruct)

    // aStruct needs to be, well, a struct
    if s.Kind() != reflect.Struct {
        return nil, union.ErrInvalidStruct
    }

    numFields := s.NumField()

    if numFields == 0 {
        return nil, union.ErrInvalidStruct
    }

    var maxSize uintptr

    for i := 0; i < numFields; i++ {
        caster, ok := s.Field(i).Interface().(UnionCaster)
        if !ok {
            return nil, union.ErrNotCaster
        }
        if size := caster.UnionSize(); size > maxSize {
            maxSize = size
        }
    }

    return &Union{union: union.WithCapacity(maxSize)}, nil
}

// The analogous of Uninit(), but with an initialized value.
func InitWith(aStruct interface{}, initUnion func(*Union)) (*Union, error) {
    u, err := Uninit(aStruct)
    if err != nil {
        return nil, err
    }
    initUnion(u)
    return u, nil
}

// Returns the type of the Union's value that is currently in use.
func (u *Union) Kind() Kind {
    return Kind(u.kind)
}

// Sets the type of the Union's value that is currently in use.
func (u *Union) ChangeKind(to Kind) {
    u.kind = to
}

// Returns the raw data in the Union.
func (u *Union) Data() union.Union {
    return u.union
}

// Casts this Union into a new data type.
func (u *Union) CastTo(caster UnionCaster) {
    u.Data().CastTo(caster)
    u.ChangeKind(caster.UnionKind())
}
