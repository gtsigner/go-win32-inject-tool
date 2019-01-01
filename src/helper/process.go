package helper

import (
	"github.com/JamesHovious/w32"
	"unsafe"
)

type Process struct {
	ProcessID int
	Name      string
	Exe       string
}

//获取进程的名字
func GetProcessesByName(name string) *Process {
	handle := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0)
	if handle == w32.ERROR_INVALID_HANDLE {
		return nil
	}
	defer w32.CloseHandle(handle)
	var entry w32.MODULEENTRY32
	entry.Size = uint32(unsafe.Sizeof(entry))

	success := w32.Module32First(handle, &entry)
	if !success {
		return nil
	}
	var process Process
	//比较
	if name == w32.UTF16PtrToString(&entry.SzModule[0]) {
		process.Name = w32.UTF16PtrToString(&entry.SzModule[0])
		process.ProcessID = int(entry.ProcessID)
		process.Exe = w32.UTF16PtrToString(&entry.SzExePath[0])
		return &process
	}

	//定义一个实体类型
	for true {
		success = w32.Module32Next(handle, &entry)
		if !success {
			return nil
		}
		processName := w32.UTF16PtrToString(&entry.SzModule[0])
		if name == processName {
			process.Name = processName
			process.ProcessID = int(entry.ProcessID)
			process.Exe = w32.UTF16PtrToString(&entry.SzExePath[0])
			return &process
		}
		break
	}
	return nil
}
