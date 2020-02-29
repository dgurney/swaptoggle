package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

func swaptoggle(path string, on bool) syscall.Errno {
	swap := []byte(path)
	operation := syscall.SYS_SWAPOFF
	if on {
		operation = syscall.SYS_SWAPON
	}
	_, _, err := syscall.Syscall(uintptr(operation), uintptr(unsafe.Pointer(&swap[0])), 0, 0)
	return err
}

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("This isn't Linux")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("You need to provide the path to your swap partition/file(s)!")
		return
	}

	// Linux supports multiple swap spaces
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Toggling swap space at", os.Args[i])
		err := swaptoggle(os.Args[i], false)
		if err != 0 {
			panic("swapoff says " + err.Error())
		}

		err = swaptoggle(os.Args[i], true)
		if err != 0 {
			panic("swapon says " + err.Error())
		}
	}
}
