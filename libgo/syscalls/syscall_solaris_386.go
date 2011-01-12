// syscall_solaris_386.go -- Solaris/x86 specific support

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import "unsafe"

const ARCH = "386"

// FIXME: ptrace(3C) has this, but exec.go expects the next.
//func libc_ptrace(request int, pid Pid_t, addr int, data int) int __asm__ ("ptrace")

func libc_ptrace(request int, pid Pid_t, addr uintptr, data *byte) int __asm__ ("ptrace")

var dummy *byte
const sizeofPtr uintptr = uintptr(unsafe.Sizeof(dummy))