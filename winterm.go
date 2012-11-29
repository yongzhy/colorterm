//+build windows

package colorterm

import (
	"bufio"
	"os"
)

const (
	COLOR_BLACK   = 0
	COLOR_BLUE    = 1
	COLOR_GREEN   = 2
	COLOR_CYAN    = 3
	COLOR_RED     = 4
	COLOR_PURPLE  = 5
	COLOR_BROWN   = 6
	COLOR_LGREY   = 7
	COLOR_DGRAY   = 8
	COLOR_LBLUE   = 9
	COLOR_LGREEN  = 10
	COLOR_LCYAN   = 11
	COLOR_LRED    = 12
	COLOR_LPURPLE = 13
	COLOR_YELLOW  = 14
	COLOR_WHITE   = 15
	COLOR_MAX     = COLOR_WHITE + 1
)

type ColorTerm struct {
	stdout  _HANDLE
	orgAttr uint16
}

func NewColorTerminal() *ColorTerm {
	term := new(ColorTerm)
	term.stdout = term._GetStdOutHandle()
	term.orgAttr = term._GetTextAttr()
	return term
}

func (term *ColorTerm) _GetStdOutHandle() _HANDLE {
	return _GetStdHandle(_STD_OUTPUT_HANDLE)
}

func (term *ColorTerm) _FlushStdOut() {
	stdout := bufio.NewWriter(os.Stdout)
	stdout.Flush()
}

func (term *ColorTerm) _GetTextAttr() uint16 {
	return _GetConsoleScreenBufferInfo(term.stdout).WAttributes
}

func (term *ColorTerm) _SetTextAttr(attr uint16) {
	_SetConsoleTextAttribute(term.stdout, attr)
}

func (term *ColorTerm) Reset() {
	term._SetTextAttr(term.orgAttr)
}

func (term *ColorTerm) ClearScreen() {
	topleft := _COORD{0, 0}
	csbi := _GetConsoleScreenBufferInfo(term.stdout)
	var size uint32 = uint32(csbi.DwSize.X) * uint32(csbi.DwSize.Y)

	_FillConsoleOutputCharacter(term.stdout,
		0x20, size, topleft)

	_FillConsoleOutputAttribute(term.stdout,
		csbi.WAttributes, size, topleft)

	_SetConsoleCursorPosition(term.stdout, topleft)
}

func (term *ColorTerm) Columns() uint16 {
	csbi := _GetConsoleScreenBufferInfo(term.stdout)
	return csbi.DwSize.X
}

func (term *ColorTerm) Lines() uint16 {
	csbi := _GetConsoleScreenBufferInfo(term.stdout)
	return csbi.DwSize.Y
}

func (term *ColorTerm) SetTitle(title string) {
	_SetConsoleTitle(title)
}

func (term *ColorTerm) SetPosition(x, y uint16) {
	_SetConsoleCursorPosition(term.stdout, _COORD{x, y})
}

func (term *ColorTerm) SetColor(fg, bg uint16) {

}

func (term *ColorTerm) SetBgColor(bg uint16) {
	if bg < COLOR_MAX {
		current := term._GetTextAttr()
		fg := current & 0x000F
		term._SetTextAttr(fg + (bg << 4))
	}
}

func (term *ColorTerm) SetTextColor(fg uint16) {
	if fg < COLOR_MAX {
		current := term._GetTextAttr()
		bg := current & 0x00F0
		term._SetTextAttr(fg + bg)
	}
}
