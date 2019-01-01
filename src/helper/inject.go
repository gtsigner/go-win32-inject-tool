package helper

import (
	"dll_inject_to_wechat/src/win32"
	"fmt"
	"github.com/JamesHovious/w32"
	"syscall"
	"unsafe"
)

const (
	ProcessAllAccess = 0x2035711
)

func Inject(name string, dll string) string {
	//1.获取微信的进程ID
	process, err := GetProcessesByName(name)
	if err != "" {
		return err
	}

	//2.打开进程
	handle, ex := syscall.OpenProcess(uint32(ProcessAllAccess), false, uint32(process.ProcessID))
	if ex != nil {
		return "打开进程失败"
	}
	defer syscall.CloseHandle(handle)
	var dllLength = len(dll) + 1
	//3.分配虚拟内存，写入dll名字路径
	dllMemAddr, ex := win32.VirtualAllocEx(handle, 0, uint32(dllLength), 4096, 4)
	if ex != nil {
		return "分配内存失败"
	}
	bt := []byte(dll)
	//4.写入内存
	_, ex = win32.WriteProcessMemory(handle, uint32(dllMemAddr), uintptr(unsafe.Pointer(&bt[0])), uint32(dllLength), 0)
	if ex != nil {
		return "写入内存失败"
	}

	//5.测试一下读出内存
	bytes, _ := w32.ReadProcessMemory(w32.HANDLE(handle), uint32(dllMemAddr), uint(dllLength))
	fmt.Println("开始加载DLL：", string(bytes[:]))

	//5.远程执行
	loadAddr, ex := win32.GetLoadLibraryAAddr()
	println(loadAddr)
	if ex != nil {
		return "获取内核地址失败"
	}
	pch, ex := win32.CreateRemoteThread(handle, 0, 0, loadAddr, dllMemAddr, 0, 0)
	if ex != nil {
		println(ex)
		return "远程加载DLL失败:"
	}
	defer syscall.CloseHandle(syscall.Handle(pch))

	return "DLL注入成功"
}

func unject() {

}
