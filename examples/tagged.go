package main

import (
    "fmt"
    "reflect"

    tagged "github.com/sug0/go-union/tagged"
    types "github.com/sug0/go-union/types"
)

// these fields need to be exported!
type MyEnum struct {
    Float32 *types.Float32
    Int32   *types.Int32
}

func main() {
    // this will return a float32's zero value
    intOrFloat, err := tagged.Zero(0, reflect.ValueOf(MyEnum{}))
    if err != nil {
        panic(err)
    }

    var f types.Float32

    intOrFloat.Cast(&f)
    *f.Ptr = 343.12

    switch intOrFloat.Kind() {
    case types.KindFloat32:
        fmt.Println("I'm a float32!")
    case types.KindInt32:
        fmt.Println("I'm an int32!")
    }

    fmt.Printf("%#v\n", intOrFloat)
}
