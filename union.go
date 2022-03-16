package union

import "unsafe"

// The Union type that will share the backing data between
// various different types.
type Union struct {
    data []byte
}

// Returns the (memory) capacity of this Union.
func (u Union) Cap() uintptr {
    return uintptr(len(u.data))
}

// Return a pointer to the Union data.
func (u Union) Pointer() unsafe.Pointer {
    return unsafe.Pointer(&u.data[0])
}

// Create a new Union from various types. Use unsafe.Sizeof(type{})
// as the arguments.
func Of(t uintptr, types ...uintptr) Union {
    maxSize := t
    for _, ti := range types {
        if ti > maxSize {
            maxSize = ti
        }
    }
    return WithCapacity(maxSize)
}

// Create a new Union with the given capacity in bytes.
func WithCapacity(capacity uintptr) Union {
    return Union{data: make([]byte, capacity)}
}
