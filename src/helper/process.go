package helper

import (
	"github.com/JamesHovious/w32"
	"syscall"
	"unsafe"
)

type (
	HANDLE uintptr
	BOOL int32
)

var (
	modadvapi32               = syscall.NewLazyDLL("advapi32.dll")
)

func GetPrivileges() {
	var token syscall.Token
	handle, _ := syscall.GetCurrentProcess()
	//失败
	if nil != syscall.OpenProcessToken(handle, syscall.TOKEN_ALL_ACCESS, &token) {
		return
	}
	//	syscall.Syscall(procLookupPrivilegeValueW.Addr(),nil,)
}

type Process struct {
	ProcessID int
	Name      string
	Exe       string
}

//获取进程的名字
func GetProcessesByName(exeFile string) (*Process, string) {
	handle, _ := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if handle == 0 {
		return nil, "创建Snapshot失败"
	}
	defer syscall.CloseHandle(handle)

	//定义句柄存储
	var entry = syscall.ProcessEntry32{}
	entry.Size = uint32(unsafe.Sizeof(entry))
	var process Process
	//定义一个实体类型
	for true {
		if nil != syscall.Process32Next(handle, &entry) {
			break
		}
		//执行文件的名称
		_exeFile := w32.UTF16PtrToString(&entry.ExeFile[0])
		if exeFile == _exeFile {
			process.Name = _exeFile
			process.ProcessID = int(entry.ProcessID)
			process.Exe = _exeFile
			return &process, ""
		}

	}
	return nil, "未找到进程"
}
