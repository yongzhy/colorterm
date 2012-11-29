package main

import (
	"colorterm"
	"fmt"
)

func main() {
	term := colorterm.NewColorTerminal()

	fmt.Printf("Testing Set Title to 'HELLO WORLD'\n")
	term.SetTitle("HELLO WORLD!")

	fmt.Printf("Testing Clear Whole Screen\n")
	term.ClearScreen()
	fmt.Printf("\nThis Is the First line after clear screen\n")

	fmt.Printf("Testing Set Foreground Color to  Red\n")
	term.SetTextColor(colorterm.COLOR_RED)
	fmt.Printf("This line should be in read\n")

	fmt.Printf("Testing Set Background Color to  Green\n")
	term.SetBgColor(colorterm.COLOR_GREEN)
	fmt.Printf("This line background should be in Green\n")

	term.SetTextColor(colorterm.COLOR_GREEN)
	term.SetBgColor(colorterm.COLOR_BLACK)

	term.Reset()
}
