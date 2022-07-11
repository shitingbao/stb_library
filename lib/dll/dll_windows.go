package dll

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 注意的是golang由于数据类型和c++的不一致
// 在需要传参的时候需要把所有的参数都转换成uintptr指针类型
// 而且转换的过程需要借助unsafe.Pointer指针

// int 转化
func IntPtr(n int) uintptr {
	return uintptr(n)
}

// string 转化
func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

// 执行 c++ 代码文件
// path 为 dll 文件地址
// example 两数相加
func addExample(a, b int, path string) {
	lib := syscall.NewLazyDLL(path)
	fmt.Println("dll:", lib.Name)
	add := lib.NewProc("sumtest")

	fmt.Println("+++++++NewProc:", add, "+++++++")

	ret, _, err := add.Call(uintptr(a), uintptr(b))
	if err != nil {
		fmt.Println("lib.dll运算结果为:", ret)
	}
}
