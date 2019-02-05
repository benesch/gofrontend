// dir_darwin.go -- For Darwin, which uses its own large file convention
// for readdir_r.

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package os

import (
	"syscall"
	_ "unsafe" // for go:linkname
)

//go:linkname libc_readdir_r readdir_r$INODE64
func libc_readdir_r(*syscall.DIR, *syscall.Dirent, **syscall.Dirent) syscall.Errno
