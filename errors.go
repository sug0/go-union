package union

import "errors"

// Union related errors.
var (
    UnionTooSmallErr = errors.New("union: not enough space in union")
)
