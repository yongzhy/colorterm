//+build windows

package colorterm

import (
	"syscall"
	"unsafe"
)

type (
	_HANDLE uintptr
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms682119.aspx
type _COORD struct {
	X, Y uint16
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms686311.aspx
type _SMALL_RECT struct {
	Left, Top, Right, Bottom uint16
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms682093.aspx
type _CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              _COORD
	DwCursorPosition    _COORD
	WAttributes         uint16
	SrWindow            _SMALL_RECT
	DwMaximumWindowSize _COORD
}

const (
	_STD_INPUT_HANDLE  = -10
	_STD_OUTPUT_HANDLE = -11
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetConsoleScreenBufferInfo = modkernel32.NewProc("GetConsoleScreenBufferInfo")
	procSetConsoleTextAttribute    = modkernel32.NewProc("SetConsoleTextAttribute")
	procSetConsoleTitle            = modkernel32.NewProc("SetConsoleTitleA")
	procSetConsoleCursorPosition   = modkernel32.NewProc("SetConsoleCursorPosition")
	procFillConsoleOutputCharacter = modkernel32.NewProc("FillConsoleOutputCharacterA")
	procFillConsoleOutputAttribute = modkernel32.NewProc("FillConsoleOutputAttribute")
	procGetStdHandle               = modkernel32.NewProc("GetStdHandle")
)

func _GetConsoleScreenBufferInfo(hConsoleOutput _HANDLE) *_CONSOLE_SCREEN_BUFFER_INFO {
	var csbi _CONSOLE_SCREEN_BUFFER_INFO
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}

func _SetConsoleTextAttribute(hConsoleOutput _HANDLE, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttributes))
	return ret != 0
}

func _SetConsoleTitle(title string) bool {
	b := []byte(title)
	ret, _, _ := procSetConsoleTitle.Call(uintptr(unsafe.Pointer(&b[0])))
	return ret != 0
}

func _SetConsoleCursorPosition(stdout _HANDLE, coord _COORD) bool {
	pCoord := uintptr(coord.X) + uintptr(coord.Y)<<16
	ret, _, _ := procSetConsoleCursorPosition.Call(
		uintptr(stdout),
		pCoord)
	return ret != 0
}

func _FillConsoleOutputCharacter(stdout _HANDLE, char uint8, length uint32, coord _COORD) uint32 {
	var written uint32

	pCoord := uintptr(coord.X) + uintptr(coord.Y)<<16
	ret, _, _ := procFillConsoleOutputCharacter.Call(
		uintptr(stdout),
		uintptr(char),
		uintptr(length),
		pCoord,
		uintptr(unsafe.Pointer(&written)))

	if ret == 0 {
		written = 0
	}
	return written
}

func _FillConsoleOutputAttribute(stdout _HANDLE, attr uint16, length uint32, coord _COORD) uint32 {
	var written uint32

	pCoord := uintptr(coord.X) + uintptr(coord.Y)<<16
	ret, _, _ := procFillConsoleOutputAttribute.Call(
		uintptr(stdout),
		uintptr(attr),
		uintptr(length),
		pCoord,
		uintptr(unsafe.Pointer(&written)))

	if ret == 0 {
		written = 0
	}

	return written
}

func _GetStdHandle(nStdHandle int32) _HANDLE {
	ret, _, _ := procGetStdHandle.Call(uintptr(nStdHandle))
	return _HANDLE(ret)
}
