package main

import(
    "fmt"
    "runtime"
    "syscall"
    "unsafe"
)

var (
    winKernel32    = syscall.NewLazyDLL("kernel32.dll")
    winProcGetStdHandle = winKernel32.NewProc("GetStdHandle")
    winProcGetConsoleMode = winKernel32.NewProc("GetConsoleMode")
    winProcSetConsoleMode = winKernel32.NewProc("SetConsoleMode")
)

const (
    WIN_STD_OUTPUT_HANDLE = 4294967285
    WIN_ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
)

func enableVirtualTerminal() error {
    handle, _, _ := winProcGetStdHandle.Call(uintptr(WIN_STD_OUTPUT_HANDLE))
    if handle == uintptr(syscall.InvalidHandle) {
        return fmt.Errorf("failed to get stdout handle")
    }

    var mode uint32
    ret, _, _ := winProcGetConsoleMode.Call(handle, uintptr(unsafe.Pointer(&mode)))
    if ret == 0 {
        return fmt.Errorf("failed to get console mode")
    }

    mode |= WIN_ENABLE_VIRTUAL_TERMINAL_PROCESSING

    ret, _, _ = winProcSetConsoleMode.Call(handle, uintptr(mode))
    if ret == 0 {
        return fmt.Errorf("failed to set console mode")
    }

    return nil
}

func initTerminalOutput() {
	if runtime.GOOS == "windows" {
		enableVirtualTerminal()
	}
}