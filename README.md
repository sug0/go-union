# go-union

This package implements C-like unions in Go, by messing with `unsafe`
code.

The API to work with them is somewhat ergonomic, utilizing
Go's interfaces to cast around the different types.

## API documentation

Find it at godoc:

* [Unsafe unions](https://godoc.org/github.com/sug0/go-union)
* [Tagged unions](https://godoc.org/github.com/sug0/go-union/tagged)
* [Auto generated wrappers over native Go types](https://godoc.org/github.com/sug0/go-union/types)

## Example usage

```go
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
```
