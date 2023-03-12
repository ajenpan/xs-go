//go:build !windows
// +build !windows

package filelock

import (
	"io"
	"syscall"
)

// use fcntl POSIX locks for the most consistent behavior across platforms, and
// hopefully some campatibility over NFS and CIFS.

func RLockFile(fd uintptr, wait bool) error {
	flock := &syscall.Flock_t{
		Type:   syscall.F_RDLCK,
		Whence: int16(io.SeekStart),
		Start:  0,
		Len:    0,
	}
	cmd := syscall.F_SETLK
	if wait {
		cmd = syscall.F_SETLKW
	}
	return syscall.FcntlFlock(fd, cmd, flock)
}

func WLockFile(fd uintptr, wait bool) error {
	flock := &syscall.Flock_t{
		Type:   syscall.F_WRLCK,
		Whence: int16(io.SeekStart),
		Start:  0,
		Len:    0,
	}
	cmd := syscall.F_SETLK
	if wait {
		cmd = syscall.F_SETLKW
	}
	return syscall.FcntlFlock(fd, cmd, flock)
}

func UnLockFile(fd uintptr) error {
	flock := &syscall.Flock_t{
		Type:   syscall.F_UNLCK,
		Whence: int16(io.SeekStart),
		Start:  0,
		Len:    0,
	}
	return syscall.FcntlFlock(fd, syscall.F_SETLKW, flock)
}
