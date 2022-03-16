PRIMITIVES = [
    'bool',
    'string',
    'int',
    'int8',
    'int16',
    'int32',
    'int64',
    'uint',
    'uint8',
    'uint16',
    'uint32',
    'uint64',
    'uintptr',
    'byte',
    'rune',
    'float32',
    'float64',
    'complex64',
    'complex128',
]

def typename(t):
    return t.title().replace(' ', '')

def typekind(t):
    return 'Kind' + t.title().replace(' ', '') 

def struct(t):
    print('type %s struct {' % typename(t))
    print('    Ptr *%s' % t)
    print('}')
    print('')

def union_kind(t):
    print('func (*%s) UnionKind() tagged.Kind {' % typename(t))
    print('    return %s' % typekind(t))
    print('}')
    print('')

def union_size(t):
    print('func (p *%s) UnionSize() uintptr {' % typename(t))
    print('    return unsafe.Sizeof(*p.Ptr)')
    print('}')
    print('')

def union_caster(t):
    print('func (p *%s) CastUnion(u union.Union) error {' % typename(t))
    print('    if u.Cap() < p.UnionSize() {')
    print('        return union.ErrTooSmall')
    print('    }')
    print('    p.Ptr = (*%s)(u.Pointer())' % t)
    print('    return nil')
    print('}')
    print('')

if __name__ == '__main__':
    # package header
    print('// THIS PACKAGE IS AUTO GENERATED -- DO NOT EDIT!')
    print('')
    print('package types')
    print('')
    print('import (')
    print('    "unsafe"')
    print('')
    print('    union "github.com/sug0/go-union"')
    print('    tagged "github.com/sug0/go-union/tagged"')
    print(')')
    print('')

    # constants
    print('const (')
    print('    KindUndefined = iota')

    for t in PRIMITIVES:
        print('    %s' % typekind(t))

    print('    KindIota')
    print(')')
    print('')

    # implementation of primitives
    for t in PRIMITIVES:
        struct(t)
        union_kind(t)
        union_size(t)
        union_caster(t)
