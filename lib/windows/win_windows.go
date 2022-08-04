package windows

import (
	"fmt"
	"log"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/shirou/gopsutil/v3/process"
)

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong     // 结构大小
	cntUsage            ulong     // 此进程的引用计数
	th32ProcessID       ulong     // 进程id
	th32DefaultHeapID   ulong_ptr // 进程默认堆id
	th32ModuleID        ulong     // 进程模块id
	cntThreads          ulong     // 进程的线程数
	th32ParentProcessID ulong     // 父进程id
	pcPriClassBase      ulong     // 线程优先权
	dwFlags             ulong     // 保留
	szExeFile           [260]byte // 进程全名
}

type MODULEENTRY32 struct {
	dwSize        ulong     // 指定结构的长度，以字节为单位。在调用Module32First功能，设置这个成员SIZEOF（MODULEENTRY32）。如果你不初始化的dwSize，Module32First将失败。
	th32ModuleID  ulong     // 此成员已经不再被使用，通常被设置为1
	th32ProcessID ulong     // 正在检查的进程标识符。这个成员的内容，可以使用Win32 API的元素
	GlblcntUsage  ulong     // 全局模块的使用计数，即模块的总载入次数。通常这一项是没有意义的，被设置为0xFFFF。
	ProccntUsage  ulong     // 全局模块的使用计数（与GlblcntUsage相同）。通常这一项也是没有意义的，被设置为0xFFFF。
	modBaseAddr   byte      // 模块的基址，在其所属的进程范围内。
	modBaseSize   ulong     // 模块的大小，单位字节。
	hModule       ulong_ptr // 模块句柄
	szModule      [260]byte // 模块名称
	szExePath     [260]byte // NULL结尾的字符串，其中包含的位置，或模块的路径。 在VC++6.0中， _MAX_PATH的值为260。
}

func GetProcessList() []PROCESSENTRY32 {
	var processArray []PROCESSENTRY32
	/*
	   CreateToolhelp32Snapshot
	       指定快照中包含的系统内容，这个参数能够使用下列数值（常量）中的一个或多个。
	       TH32CS_INHERIT(0x80000000)      - 声明快照句柄是可继承的。
	       TH32CS_SNAPALL                  - 在快照中包含系统中所有的进程和线程。
	       TH32CS_SNAPHEAPLIST(0x00000001) - 在快照中包含在th32ProcessID中指定的进程的所有的堆。
	       TH32CS_SNAPMODULE(0x00000008)   - 在快照中包含在th32ProcessID中指定的进程的所有的模块。
	       TH32CS_SNAPPROCESS(0x00000002)  - 在快照中包含系统中所有的进程。
	       TH32CS_SNAPTHREAD(0x00000004)   - 在快照中包含系统中所有的线程。
	       H32CS_SNAPALL = (TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE)
	   th32ProcessID
	       指定将要快照的进程ID。如果该参数为0表示快照当前进程。该参数只有在设置了TH32CS_SNAPHEAPLIST或者TH32CS_SNAPMODULE后才有效，在其他情况下该参数被忽略，所有的进程都会被快照。
	*/
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	if int(pHandle) == -1 {
		return processArray
	}
	Process32Next := kernel32.NewProc("Process32Next")
	for {
		var proc PROCESSENTRY32
		proc.dwSize = ulong(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			// fmt.Println("ProcessName : " + string(proc.szExeFile[0:]))
			// fmt.Println("th32ModuleID : " + strconv.Itoa(int(proc.th32ModuleID)))
			// fmt.Println("ProcessID : " + strconv.Itoa(int(proc.th32ProcessID)))

			log.Println("进程ID:", th32ProcessID, " --- 进程名称 ==:", string(proc.szExeFile[:]))
			processArray = append(processArray, proc)
		} else {
			break
		}
	}
	CloseHandle := kernel32.NewProc("CloseHandle")
	_, _, _ = CloseHandle.Call(pHandle)

	return processArray
}

func GetProcess(pid int) (PROCESSENTRY32, error) {
	var targetProcess PROCESSENTRY32
	targetProcess = PROCESSENTRY32{
		dwSize: 0,
	}

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	if int(pHandle) == -1 {
		return targetProcess, fmt.Errorf("error:Can not find any proess.")
	}
	Process32Next := kernel32.NewProc("Process32Next")

	for {
		var proc PROCESSENTRY32
		proc.dwSize = ulong(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			if int(proc.th32ProcessID) == pid {
				targetProcess = proc
				// fmt.Println("ProcessName : " + string(proc.szExeFile[0:]))
				// fmt.Println("th32ModuleID : " + strconv.Itoa(int(proc.th32ModuleID)))
				// fmt.Println("ProcessID : " + strconv.Itoa(int(proc.th32ProcessID)))
				break
			}
		} else {
			break
		}
	}
	CloseHandle := kernel32.NewProc("CloseHandle")
	_, _, _ = CloseHandle.Call(pHandle)
	return targetProcess, nil
}

// 根据进程 id 计算 cpu 使用率
func pidProcess(pid int32) {
	for {
		p1, err := process.NewProcess(pid)
		if err != nil {
			log.Println("p1:", err)
			return
		}
		t1, err := p1.Times()
		if err != nil {
			log.Println("t1:", err)
			return
		}
		time.Sleep(time.Second * 1)

		p2, err := process.NewProcess(pid)
		if err != nil {
			log.Println("p2:", err)
			return
		}
		t2, err := p2.Times()
		if err != nil {
			log.Println("t2:", err)
			return
		}
		// CPU使用比例 = 在间隔时间内进程的CPU使用时间 / 计算机逻辑CPU数量
		cpu := (t2.Total() - t1.Total()) / 2 / float64(runtime.NumCPU()) * 3

		log.Println(cpu)
	}
}
