// +build linux,386

package colorterm

import (
	"fmt"
	"os"
	"strconv"
)

const (
	COLOR_BLACK   = 0
	COLOR_RED     = 1
	COLOR_GREEN   = 2
	COLOR_BROWN   = 3
	COLOR_BLUE    = 4
	COLOR_PURPLE  = 5
	COLOR_CYAN    = 6
	COLOR_LGREY   = 7
	COLOR_DGRAY   = 8
	COLOR_LRED    = 9
	COLOR_LGREEN  = 10
	COLOR_YELLOW  = 11
	COLOR_LBLUE   = 12
	COLOR_LPURPLE = 13
	COLOR_LCYAN   = 14
	COLOR_WHITE   = 15
	COLOR_MAX     = COLOR_WHITE + 1
)

var _Code = map[string]string{
	"escape":     "\033[",
	"reset":      "\033[0m",
	"bold":       "\033[01m",
	"clear":      "\033[2J",
	"clear_eol":  "\033[K",
	"gotoxy":     "\033[%d;%dH",
	"move_up":    "\033[%dA",
	"move_down":  "\033[%dB",
	"move_right": "\033[%dC",
	"move_left":  "\033[%dD",
	"save":       "\033[s",
	"restore":    "\033[u",
	"title":      "\033]1;\007\033]2;%s\007",
}

var _ColorFg = map[int]string{
	0:  "30m",
	1:  "31m",
	2:  "32m",
	3:  "33m",
	4:  "34m",
	5:  "35m",
	6:  "36m",
	7:  "37m",
	8:  "1;30m",
	9:  "1;31m",
	10: "1;32m",
	11: "1;33m",
	12: "1;34m",
	13: "1;35m",
	14: "1;36m",
	15: "1;37m",
}

var _ColorBg = map[int]string{
	0: "40m",
	1: "41m",
	2: "42m",
	3: "43m",
	4: "44m",
	5: "45m",
	6: "46m",
	7: "47m",
}

type ColorTerm struct {
	stdout *os.File
	stderr *os.File
}

func NewColorTerminal() *ColorTerm {
	term := new(ColorTerm)
	term.stdout = os.Stdout
	term.stderr = os.Stderr
	return term
}

func (term *ColorTerm) Reset() {
	term.stdout.WriteString(_Code["reset"])
}

func (term *ColorTerm) ClearScreen() {
	term.stdout.WriteString(_Code["clear"])
	term.stdout.Sync()
}

func (term *ColorTerm) Columns() uint16 {
	var cols uint16 = 0
	env := os.Getenv("COLUMNS")
	if env != "" {
		c, _ := strconv.ParseUint(env, 10, 64)
		cols = uint16(c)
	}
	return cols
}

func (term *ColorTerm) Lines() uint16 {
	var cols uint16 = 0
	env := os.Getenv("LINES")
	if env != "" {
		c, _ := strconv.ParseUint(env, 10, 64)
		cols = uint16(c)
	}
	return cols
}

func (term *ColorTerm) SetTitle(title string) {
	t := os.Getenv("TERM")
	if t == "xterm" || t == "Eterm" || t == "aterm" || t == "rxvt" || t == "xterm-color" {
		term.stderr.WriteString(fmt.Sprintf(_Code["title"], title))
	}
}

func (term *ColorTerm) SetPosition(x, y uint16) {
	term.stdout.WriteString(fmt.Sprintf(_Code["gotoxy"], x, y))
}

func (term *ColorTerm) SetColor(fg, bg uint16) {
	term.SetBgColor(bg)
	term.SetTextColor(fg)
}

func (term *ColorTerm) SetBgColor(bg uint16) {
	if bg < COLOR_MAX {
		term.stdout.WriteString(_Code["escape"] + _ColorBg[int(bg)%len(_ColorBg)])
		term.stdout.Sync()
	}
}

func (term *ColorTerm) SetTextColor(fg uint16) {
	if fg < COLOR_MAX {
		term.stdout.WriteString(_Code["escape"] + _ColorFg[int(fg)])
		term.stdout.Sync()
	}
}
