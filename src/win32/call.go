package win32

import (
	"syscall"
	"unsafe"
)

var (
	modkernel32, _            = syscall.LoadDLL("kernel32.dll")
	procVirtualAllocEx, _     = modkernel32.FindProc("VirtualAllocEx")
	procWriteProcessMemory, _ = modkernel32.FindProc("WriteProcessMemory")
	procCreateRemoteThread, _ = modkernel32.FindProc("CreateRemoteThread")
	procGetModuleHandleA, _   = modkernel32.FindProc("GetModuleHandleA")
	procGetProcAddress, _     = modkernel32.FindProc("GetProcAddress")
)

func GetModuleHandleA(name string) (r1 uintptr, err error) {
	bytes := []byte(name)
	r1, _, err = procGetModuleHandleA.Call(uintptr(unsafe.Pointer(&bytes[0])))
	err = syscall.GetLastError()
	return
}

func GetLoadLibraryAAddr() (uintptr, error) {
	handle, err := GetModuleHandleA("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	ptr, err := syscall.GetProcAddress(syscall.Handle(handle), "LoadLibraryA")
	err = syscall.GetLastError()
	return ptr, err
}

//分配虚拟内存
func VirtualAllocEx(hwnd syscall.Handle, lpaddress uint32, size uint32, tp uint32, tect uint32) (r1 uintptr, err error) {
	r1, _, _ = procVirtualAllocEx.Call(uintptr(hwnd), uintptr(lpaddress), uintptr(size), uintptr(tp), uintptr(tect), 0)
	err = syscall.GetLastError()
	return
}

func WriteProcessMemory(hwnd syscall.Handle, addr uint32, lpBuffer uintptr, nsize uint32, filewriten uint32) (r1 uintptr, err error) {
	r1, _, err = procWriteProcessMemory.Call(uintptr(hwnd), uintptr(addr), lpBuffer, uintptr(nsize), uintptr(filewriten), 0)
	err = syscall.GetLastError()
	return
}

//执行远程线程调用方法
func CreateRemoteThread(hwnd syscall.Handle, threadAttributes uint32, stackSize uint32, startAddress uintptr, parameter uintptr, creationFlags uint32, threadid uint32) (r1 uintptr, err error) {
	r1, _, err = procCreateRemoteThread.Call(uintptr(hwnd), uintptr(threadAttributes), uintptr(stackSize), uintptr(startAddress), uintptr(parameter), uintptr(creationFlags), uintptr(threadid))
	err = syscall.GetLastError()
	return
}
