# go-union

This package implements C-like unions in Go, by messing with `unsafe`
code.

The API to work with them is somewhat ergonomic, utilizing
Go's interfaces to cast around the different types.

## Who would use this package?

A madman. Use the `unsafe` package directly, honestly.

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

    tagged "github.com/sug0/go-union/tagged"
    types "github.com/sug0/go-union/types"
)

func main() {
    // This struct will describe our union;
    // the fields need to be exported, because
    // the package typechecks the values at runtime
    // with reflection...
    myEnum := struct {
        Float32 *types.Float32
        Int32   *types.Int32
    }{}

    // Initialize the union with a value.
    intOrFloat, err := tagged.InitWith(myEnum, func(u *tagged.Union) {
        var f types.Float32
        u.CastTo(&f)
        *f.Ptr = 343.12
    })
    if err != nil {
        panic(err)
    }

    // Debug contents.
    switch intOrFloat.Kind() {
    default:
        panic("We're not a float32!")
    case types.KindFloat32:
        fmt.Println("I'm a float32!")
        fmt.Printf("%#v\n", intOrFloat)
    }

    // Update content.
    var i types.Int32
    intOrFloat.CastTo(&i)
    *i.Ptr = 45

    // Debug contents.
    switch intOrFloat.Kind() {
    default:
        panic("We're not an int32!")
    case types.KindInt32:
        fmt.Println("I'm an int32!")
        fmt.Printf("%#v\n", intOrFloat)
    }
}
```
