package union

import "errors"

// Union related errors.
var (
    ErrTooSmall      = errors.New("union: not enough space in union")
    ErrInvalidStruct = errors.New("union: invalid struct")
    ErrInvalidIndex  = errors.New("union: invalid struct index")
    ErrNotCaster     = errors.New("union: not a tagged.UnionCaster")
)
