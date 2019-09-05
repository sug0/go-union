package tagged

import (
    union "github.com/sug0/go-union"
)

// A type that belongs to a Union's typesystem.
type UnionKind interface {
    UnionKind() Kind
}

// A type capable of being casted to from a Union.
type UnionCaster interface {
    UnionKind
    union.UnionCaster
}
