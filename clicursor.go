package clicursor

import "fmt"

// Show the cursor
func Show() {
	fmt.Printf("\u001b[?25h")
}

// Hide the cursor
func Hide() {
	fmt.Printf("\u001b[?25l")
}
