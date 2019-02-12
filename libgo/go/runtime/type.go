// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Runtime type representation.

package runtime

import "unsafe"

type _type struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	kind       uint8
	align      int8
	fieldAlign uint8
	_          uint8

	hashfn  func(unsafe.Pointer, uintptr) uintptr
	equalfn func(unsafe.Pointer, unsafe.Pointer) bool

	gcdata  *byte
	_string *string
	*uncommontype
	ptrToThis *_type
}

func (t *_type) string() string {
	return *t._string
}

// pkgpath returns the path of the package where t was defined, if
// available. This is not the same as the reflect package's PkgPath
// method, in that it returns the package path for struct and interface
// types, not just named types.
func (t *_type) pkgpath() string {
	if u := t.uncommontype; u != nil {
		if u.pkgPath == nil {
			return ""
		}
		return *u.pkgPath
	}
	return ""
}

// Return whether two type descriptors are equal.
// This is gccgo-specific, as gccgo, unlike gc, permits multiple
// independent descriptors for a single type.
func eqtype(t1, t2 *_type) bool {
	return t1 == t2
}

type method struct {
	name    *string
	pkgPath *string
	mtyp    *_type
	typ     *_type
	tfn     unsafe.Pointer
}

type uncommontype struct {
	name    *string
	pkgPath *string
	methods []method
}

type imethod struct {
	name    *string
	pkgPath *string
	typ     *_type
}

type interfacetype struct {
	typ     _type
	methods []imethod
}

type maptype struct {
	typ        _type
	key        *_type
	elem       *_type
	bucket     *_type // internal type representing a hash bucket
	keysize    uint8  // size of key slot
	valuesize  uint8  // size of value slot
	bucketsize uint16 // size of bucket
	flags      uint32
}

// Note: flag values must match those used in the TMAP case
// in ../cmd/compile/internal/gc/reflect.go:dtypesym.
func (mt *maptype) indirectkey() bool { // store ptr to key instead of key itself
	return mt.flags&1 != 0
}
func (mt *maptype) indirectvalue() bool { // store ptr to value instead of value itself
	return mt.flags&2 != 0
}
func (mt *maptype) reflexivekey() bool { // true if k==k for all keys
	return mt.flags&4 != 0
}
func (mt *maptype) needkeyupdate() bool { // true if we need to update key on an overwrite
	return mt.flags&8 != 0
}
func (mt *maptype) hashMightPanic() bool { // true if hash function might panic
	return mt.flags&16 != 0
}

type arraytype struct {
	typ   _type
	elem  *_type
	slice *_type
	len   uintptr
}

type chantype struct {
	typ  _type
	elem *_type
	dir  uintptr
}

type slicetype struct {
	typ  _type
	elem *_type
}

type functype struct {
	typ       _type
	dotdotdot bool
	in        []*_type
	out       []*_type
}

type ptrtype struct {
	typ  _type
	elem *_type
}

type structfield struct {
	name       *string // nil for embedded fields
	pkgPath    *string // nil for exported Names; otherwise import path
	typ        *_type  // type of field
	tag        *string // nil if no tag
	offsetAnon uintptr // byte offset of field<<1 | isAnonymous
}

func (f *structfield) offset() uintptr {
	return f.offsetAnon >> 1
}

func (f *structfield) anon() bool {
	return f.offsetAnon&1 != 0
}

type structtype struct {
	typ    _type
	fields []structfield
}
