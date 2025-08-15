package os

import (
	"os"
	"runtime"
	"syscall"
)

func ProcessExists(process *os.Process) (bool, error) {
	if runtime.GOOS == "windows" {
		_, err := os.FindProcess(process.Pid)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		err := process.Signal(syscall.Signal(0))
		if err == nil {
			return true, nil
		}
		if err.Error() == "os: process already finished" {
			return false, nil
		}
		errno, ok := err.(syscall.Errno)
		if !ok {
			return false, err
		}
		switch errno {
		case syscall.ESRCH:
			return false, nil
		case syscall.EPERM:
			return true, nil
		}
		return false, err
	}
}
