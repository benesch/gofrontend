// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import "unsafe"

//sysnb raw_ptrace(request int, pid int, addr *byte, data *byte) (err Errno)
//ptrace(request _C_int, pid Pid_t, addr *byte, data *byte) _C_long

//sys	getdirent(fd int, buf []byte, basep *uintptr) (n int, err error)
//getdirent64(fd _C_int, buf *byte, nbyte Size_t, basep *uintptr) _C_int

func ReadDirent(fd int, buf []byte) (n int, err error) {
	// Final argument is (basep *uintptr) and the syscall doesn't take nil.
	// 64 bits should be enough. (32 bits isn't even on 386). Since the
	// actual system call is getdirentries64, 64 is a good guess.
	// TODO(rsc): Can we use a single global basep for all calls?
	basep := (*uintptr)(unsafe.Pointer(new(uint64)))
	return getdirent(fd, buf, basep)
}

//sys	Flock(fd int, how int) (err error)
//flock(fd _C_int, how _C_int) _C_int
