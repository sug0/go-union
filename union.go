package union

import (
    "errors"
    "unsafe"
)

// The Union type, that will share the backing data between
// various different types.
type Union struct {
    data []byte
}

// A type capable of being casted to from a Union.
type UnionCaster interface {
    CastUnion(u Union) error
}

// Union related errors.
var (
    UnionTooSmallErr = errors.New("union: not enough space in union")
)

// Returns the (memory) capacity of this Union.
func (u Union) Cap() uintptr {
    return uintptr(cap(u.data))
}

// Return a pointer to the union data.
func (u Union) Pointer() unsafe.Pointer {
    return unsafe.Pointer(&u.data[0])
}

// Create a new Union from various types. Use union.Sizeof(type{})
// as the arguments.
func Of(t uintptr, types ...uintptr) Union {
    maxSize := t
    for _, ti := range types {
        if ti > maxSize {
            maxSize = ti
        }
    }
    return Union{data: make([]byte, maxSize)}
}
