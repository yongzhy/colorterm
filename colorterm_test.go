package colorterm

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	term := NewColorTerminal()
	fmt.Printf("Testing Get Console Buffer Size\n")
	fmt.Printf("Buffer Columns=%d, Lines=%d\n", term.Columns(), term.Lines())

	fmt.Printf("Testing Set Title to 'HELLO WORLD'\n")
	term.SetTitle("HELLO WORLD!")

	fmt.Printf("Testing Clear Whole Screen\n")
	term.ClearScreen()
	fmt.Printf("\nThis Is the First line after clear screen\n")

	fmt.Printf("Testing Set Foreground Color to  Red\n")
	term.SetTextColor(COLOR_RED)
	fmt.Printf("This line should be in read\n")

	fmt.Printf("Testing Set Background Color to  Green\n")
	term.SetBgColor(COLOR_GREEN)
	fmt.Printf("This line background should be in Green\n")

	term.SetTextColor(COLOR_GREEN)
	term.SetBgColor(COLOR_BLACK)

	term.Reset()
}
