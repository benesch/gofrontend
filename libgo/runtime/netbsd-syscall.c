// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The shims in this file account for NetBSD's approach to symbol
// versioning.
//
// For example, take sigaction. It is defined in signal.h like so:
//
//    int sigaction(...) __asm__ ("__sigaction14")
//
// C code that includes signal.h will transparently link against the
// modern "__sigaction14" symbol. But the Go runtime directly declares
// "//extern sigaction" and thus will incorrectly link against the old
// unsuffixed "sigaction" symbol.
//
// So the sigaction shim in this file creates a symbol named "sigaction"
// that shadows the sigaction symbol in libc and delegates to the correct
// "__sigaction14" symbol.
//
// Importantly, by relying on the system headers to select the correct
// version, this approach avoids ever mentioning a specific symbol
// version like "__sigaction14", so there is no need to update these
// shims for new versions of NetBSD.

#include <lwp.h>
#include <signal.h>
#include <sys/event.h>
#include <sys/time.h>

int sigaction_shim(int sig, const struct sigaction * restrict act,
  struct sigaction * restrict oact) __asm__ ("sigaction");

int
sigaction_shim(int sig, const struct sigaction * restrict act,
  struct sigaction * restrict oact)
{
  return sigaction(sig, act, oact);
}

int sigaltstack_shim(const stack_t * restrict ss, stack_t * restrict oss)
  __asm__ ("sigaltstack");

int
sigaltstack_shim(const stack_t * restrict ss, stack_t * restrict oss)
{
  return sigaltstack(ss, oss);
}

int sigprocmask_shim(int how, const sigset_t * restrict set,
  sigset_t * restrict oset)
  __asm__ ("sigprocmask");

int
sigprocmask_shim(int how, const sigset_t * restrict set, sigset_t * restrict oset) {
  return sigprocmask(how, set, oset);
}

int sigemptyset_shim(sigset_t *set)
  __asm__ ("sigemptyset");

int
sigemptyset_shim(sigset_t *set)
{
  return sigemptyset(set);
}

int sigfillset_shim(sigset_t *set)
  __asm__ ("sigfillset");

int
sigfillset_shim(sigset_t *set)
{
  return sigfillset(set);
}

int sigaddset_shim(sigset_t *set, int signo)
  __asm__ ("sigaddset");

int
sigaddset_shim(sigset_t *set, int signo)
{
  return sigaddset(set, signo);
}

int sigdelset_shim(sigset_t *set, int signo)
  __asm__ ("sigdelset");

int
sigdelset_shim(sigset_t *set, int signo)
{
  return sigdelset(set, signo);
}

int sigblock_shim(int mask)
  __asm__ ("sigblock");

int
sigblock_shim(int mask)
{
  return sigblock(mask);
}

int sigmask_shim(int mask)
  __asm__ ("sigmask");

int
sigmask_shim(int mask)
{
  return sigmask(mask);
}

int kevent_shim(int kq, const struct kevent *changelist, size_t nchanges,
  struct kevent *eventlist, size_t nevents,
  const struct timespec *timeout)
  __asm__ ("kevent");

int
kevent_shim(int kq, const struct kevent *changelist, size_t nchanges,
	    struct kevent *eventlist, size_t nevents, const struct timespec *timeout)
{
  return kevent(kq, changelist, nchanges, eventlist, nevents, timeout);
}

int _lwp_park_shim(clockid_t clock_id, int flags, struct timespec *ts,
  lwpid_t unpark, const void *hint, const void *unparkhint)
  __asm__ ("_lwp_park");

int
_lwp_park_shim(clockid_t clock_id, int flags, struct timespec *ts,
	       lwpid_t unpark, const void *hint, const void *unparkhint)
{
  return _lwp_park(clock_id, flags, ts, unpark, hint, unparkhint);
}

int setitimer_shim(int which, const struct itimerval * restrict value,
  struct itimerval * restrict ovalue)
  __asm__ ("setitimer");

int
setitimer_shim(int which, const struct itimerval * restrict value,
	       struct itimerval * restrict ovalue)
{
  return setitimer(which, value, ovalue);
}
