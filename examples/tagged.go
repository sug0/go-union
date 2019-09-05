package main

import (
    "fmt"
    "unsafe"
    "reflect"

    union "github.com/sug0/go-union"
    tagged "github.com/sug0/go-union/tagged"
)

type (
    Float32Ptr struct { ptr *float32 }
    Int32Ptr   struct { ptr *int32 }
)

const (
    KindFloat32 = iota
    KindInt32
)

// these fields need to be exported!
type MyEnum struct {
    Float32 *Float32Ptr
    Int32   *Int32Ptr
}

func main() {
    // this will return a float32's zero value
    intOrFloat, err := tagged.Zero(0, reflect.ValueOf(MyEnum{}))
    if err != nil {
        panic(err)
    }

    var f Float32Ptr

    intOrFloat.Cast(&f)
    *f.ptr = 343.12

    switch intOrFloat.Kind() {
    case KindFloat32:
        fmt.Println("I'm a float32!")
    case KindInt32:
        fmt.Println("I'm an int32!")
    }

    fmt.Printf("%#v\n", intOrFloat)
}

func (*Float32Ptr) UnionKind() tagged.Kind {
    return KindFloat32
}

func (f *Float32Ptr) CastUnion(u union.Union) error {
    if u.Cap() < unsafe.Sizeof(float32(0.0)) {
        return union.ErrTooSmall
    }
    f.ptr = (*float32)(u.Pointer())
    return nil
}

func (*Int32Ptr) UnionKind() tagged.Kind {
    return KindInt32
}

func (i *Int32Ptr) CastUnion(u union.Union) error {
    if u.Cap() < unsafe.Sizeof(int32(0)) {
        return union.ErrTooSmall
    }
    i.ptr = (*int32)(u.Pointer())
    return nil
}
