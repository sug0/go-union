package union

// A type capable of being casted to from a Union.
type UnionCaster interface {
    CastUnion(u Union) error
}
