// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The shims in this file account for NetBSD's approach to symbol
// versioning.
//
// See the comment in the analogous file in the runtime package,
// runtime/netbsd-syscall.c, for details about the approach.

#include <dirent.h>
#include <pwd.h>
#include <sys/mount.h>
#include <sys/param.h>
#include <sys/resource.h>
#include <sys/select.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <sys/time.h>
#include <sys/times.h>
#include <sys/wait.h>
#include <time.h>
#include <utime.h>

int mknod_shim(const char *path, mode_t mode, dev_t dev)
  __asm__ ("mknod");

int
mknod_shim(const char *path, mode_t mode, dev_t dev)
{
  return mknod(path, mode, dev);
}

int stat_shim(const char *path, struct stat *sb)
  __asm__ ("stat");

int
stat_shim(const char *path, struct stat *sb)
{
  return stat(path, sb);
}

int lstat_shim(const char *path, struct stat *sb)
  __asm__ ("lstat");

int
lstat_shim(const char *path, struct stat *sb)
{
  return lstat(path, sb);
}

int fstat_shim(int fd, struct stat *sb)
  __asm__ ("fstat");

int
fstat_shim(int fd, struct stat *sb)
{
  return fstat(fd, sb);
}

int readdir_r_shim(DIR * restrict dirp, struct dirent * restrict entry,
  struct dirent ** restrict result)
  __asm__ ("readdir_r");

int readdir_r_shim(DIR * restrict dirp, struct dirent * restrict entry,
		   struct dirent ** restrict result)
{
  return readdir_r(dirp, entry, result);
}

clock_t times_shim(struct tms *tp)
  __asm__ ("times");

clock_t
times_shim(struct tms *tp)
{
  return times(tp);
}

time_t time_shim(time_t *tloc)
  __asm__ ("time");

time_t
time_shim(time_t *tloc)
{
  return time(tloc);
}

int utime_shim(const char *file, const struct utimbuf *timep)
  __asm__ ("utime");

int
utime_shim(const char *file, const struct utimbuf *timep)
{
  return utime(file, timep);
}

int utimes_shim(const char *path, const struct timeval times[2])
  __asm__ ("utimes");

int
utimes_shim(const char *path, const struct timeval times[2])
{
  return utimes(path, times);
}

int gettimeofday_shim(struct timeval * restrict tp, void * restrict tzp)
  __asm__ ("gettimeofday");

int
gettimeofday_shim(struct timeval * restrict tp, void * restrict tzp)
{
  return gettimeofday(tp, tzp);
}

int settimeofday_shim(const struct timeval * restrict tp,
  const void * restrict tzp)
  __asm__ ("settimeofday");

int
settimeofday_shim(const struct timeval * restrict tp, const void * restrict tzp)
{
  return settimeofday(tp, tzp);
}

int clock_nanosleep_shim(clockid_t clock_id, int flags,
  const struct timespec *rqtp, struct timespec *rmtp)
  __asm__ ("clock_nanosleep");

int
clock_nanosleep_shim(clockid_t clock_id, int flags, const struct timespec *rqtp,
		     struct timespec *rmtp)
{
  return clock_nanosleep(clock_id, flags, rqtp, rmtp);
}

int nanosleep_shim(const struct timespec *rqtp, struct timespec *rmtp)
  __asm__ ("nanosleep");

int
nanosleep_shim(const struct timespec *rqtp, struct timespec *rmtp)
{
  return nanosleep(rqtp, rmtp);
}

int getrusage_shim(int who, struct rusage *rusage)
  __asm__ ("getrusage");

int
getrusage_shim(int who, struct rusage *rusage)
{
  return getrusage(who, rusage);
}

int getpwnam_r_shim(const char *name, struct passwd *pw, char *buffer,
  size_t buflen, struct passwd **result)
  __asm__ ("getpwnam_r");

int
getpwnam_r_shim(const char *name, struct passwd *pw, char *buffer,
		size_t buflen, struct passwd **result)
{
  return getpwnam_r(name, pw, buffer, buflen, result);
}

int getpwuid_r_shim(uid_t uid, struct passwd *pw, char *buffer, size_t buflen,
  struct passwd **result)
  __asm__ ("getpwuid_r");

int
getpwuid_r_shim(uid_t uid, struct passwd *pw, char *buffer, size_t buflen,
		struct passwd **result)
{
  return getpwuid_r(uid, pw, buffer, buflen, result);
}

int socket_shim(int domain, int type, int protocol)
  __asm__ ("socket");

int socket_shim(int domain, int type, int protocol)
{
  return socket(domain, type, protocol);
}

int select_shim(int nfds, fd_set * restrict readfds, fd_set * restrict writefds,
  fd_set * restrict exceptfds, struct timeval * restrict timeout)
  __asm__ ("select");

int select_shim(int nfds, fd_set * restrict readfds, fd_set * restrict writefds,
		fd_set * restrict exceptfds, struct timeval * restrict timeout)
{
  return select(nfds, readfds, writefds, exceptfds, timeout);
}

pid_t wait4_shim(pid_t wpid, int *status, int options, struct rusage *rusage)
  __asm__ ("wait4");

pid_t
wait4_shim(pid_t wpid, int *status, int options, struct rusage *rusage)
{
  return wait4(wpid, status, options, rusage);
}

int mount_shim(const char *type, const char *dir, int flags, void *data,
  size_t data_len)
  __asm__ ("mount");

int
mount_shim(const char *type, const char *dir, int flags, void *data,
	   size_t data_len)
{
  return mount(type, dir, flags, data, data_len);
}
