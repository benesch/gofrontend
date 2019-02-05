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

//sys	Fstat(fd int, stat *Stat_t) (err error)
//fstat$INODE64(fd _C_int, stat *Stat_t) _C_int

//sys	Ftruncate(fd int, length int64) (err error)
//ftruncate(fd _C_int, length Offset_t) _C_int

//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//getrlimit(resource _C_int, rlim *Rlimit) _C_int

//sys	Lstat(path string, stat *Stat_t) (err error)
//lstat$INODE64(path *byte, stat *Stat_t) _C_int

//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
//mmap(addr *byte, length Size_t, prot _C_int, flags _C_int, fd _C_int, offset Offset_t) *byte

//sys	Open(path string, mode int, perm uint32) (fd int, err error)
//__go_open(path *byte, mode _C_int, perm Mode_t) _C_int

//sys	Pread(fd int, p []byte, offset int64) (n int, err error)
//pread(fd _C_int, buf *byte, count Size_t, offset Offset_t) Ssize_t

//sys	Pwrite(fd int, p []byte, offset int64) (n int, err error)
//pwrite(fd _C_int, buf *byte, count Size_t, offset Offset_t) Ssize_t

//sys	Seek(fd int, offset int64, whence int) (off int64, err error)
//lseek$INODE64(fd _C_int, offset Offset_t, whence _C_int) Offset_t

//sysnb	Setrlimit(resource int, rlim *Rlimit) (err error)
//setrlimit(resource int, rlim *Rlimit) _C_int

//sys	Stat(path string, stat *Stat_t) (err error)
//stat$INODE64(path *byte, stat *Stat_t) _C_int

//sys	Truncate(path string, length int64) (err error)
//truncate(path *byte, length Offset_t) _C_int

// Make fcntl visible to go/internal/syscall/unix.
//go:linkname fcntl syscall.fcntl
