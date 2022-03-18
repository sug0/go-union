package union

// Used to extract the size in
// bytes occupied by a particular type that
// implements this interface.
//
// It's needed because UnionCaster implementers
// will be architecture width sized (usually they'll
// be structs with a single pointer field).
type Sized interface {
    UnionSize() uintptr
}

// A type capable of being casted to from a Union.
type UnionCaster interface {
    Sized
    CastFromUnion(u Union) error
}
