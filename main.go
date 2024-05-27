package colorado

import "fmt"

// ANSI escape codes for colors
const (
	Reset  = "\033[0m"
	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// Colorado function applies the specified color to the input text
func Colorado(text string, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}
