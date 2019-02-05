// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

type mOS struct {
	initialized bool
	mutex       pthreadmutex
	cond        pthreadcond
	count       int
}

func unimplemented(name string) {
	println(name, "not implemented")
	*(*int)(unsafe.Pointer(uintptr(1231))) = 1231
}

//go:nosplit
func semacreate(mp *m) {
	if mp.mos.initialized {
		return
	}
	mp.mos.initialized = true
	if err := pthread_mutex_init(&mp.mos.mutex, nil); err != 0 {
		throw("pthread_mutex_init")
	}
	if err := pthread_cond_init(&mp.mos.cond, nil); err != 0 {
		throw("pthread_cond_init")
	}
}

//go:nosplit
func semasleep(ns int64) int32 {
	var start int64
	if ns >= 0 {
		start = nanotime()
	}
	mp := getg().m
	pthread_mutex_lock(&mp.mos.mutex)
	for {
		if mp.mos.count > 0 {
			mp.mos.count--
			pthread_mutex_unlock(&mp.mos.mutex)
			return 0
		}
		if ns >= 0 {
			spent := nanotime() - start
			if spent >= ns {
				pthread_mutex_unlock(&mp.mos.mutex)
				return -1
			}
			var t timespec
			t.set_nsec(int32(ns - spent))
			err := pthread_cond_timedwait_relative_np(&mp.mos.cond, &mp.mos.mutex, &t)
			if err == _ETIMEDOUT {
				pthread_mutex_unlock(&mp.mos.mutex)
				return -1
			}
		} else {
			pthread_cond_wait(&mp.mos.cond, &mp.mos.mutex)
		}
	}
}

//go:nosplit
func semawakeup(mp *m) {
	pthread_mutex_lock(&mp.mos.mutex)
	mp.mos.count++
	if mp.mos.count > 0 {
		pthread_cond_signal(&mp.mos.cond)
	}
	pthread_mutex_unlock(&mp.mos.mutex)
}
