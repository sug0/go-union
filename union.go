package main

import (
    "fmt"
    "errors"
    "unsafe"
)

// The Union type, that will share the backing data between
// various different types.
type Union struct {
    data []byte
}

// A type capable of being decoded from a Union.
type UnionDecoder interface {
    DecodeUnion(u Union) error
}

// Returns the (memory) capacity of this Union.
func (u Union) Cap() uintptr {
    return uintptr(cap(u.data))
}

// Return a pointer to the union data.
func (u Union) Pointer() unsafe.Pointer {
    return unsafe.Pointer(&u.data[0])
}

// The union types we'll be working with...
type (
    Int    struct { value *int }
    String struct { value *string }
)

// Union related errors.
var (
    UnionTooSmallErr = errors.New("union: not enough space in union")
)

func (i *Int) DecodeUnion(u Union) error {
    if u.Cap() < unsafe.Sizeof(int(0)) {
        return UnionTooSmallErr
    }
    i.value = (*int)(u.Pointer())
    return nil
}

func (s *String) DecodeUnion(u Union) error {
    if u.Cap() < unsafe.Sizeof("") {
        return UnionTooSmallErr
    }
    s.value = (*string)(u.Pointer())
    return nil
}

func UnionOf(t uintptr, types ...uintptr) Union {
    maxSize := t
    for _, ti := range types {
        if ti > maxSize {
            maxSize = ti
        }
    }
    return Union{data: make([]byte, maxSize)}
}

func main() {
    stringOrInt := UnionOf(
        unsafe.Sizeof(int(0)),
        unsafe.Sizeof(string("")),
    )

    var (
        s String
        i Int
    )

    err := s.DecodeUnion(stringOrInt)
    if err != nil {
        panic(err)
    }
    *s.value = "cenas"

    fmt.Println(stringOrInt)

    err = i.DecodeUnion(stringOrInt)
    if err != nil {
        panic(err)
    }
    *i.value = 45

    fmt.Println(stringOrInt)
}
