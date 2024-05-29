package colorado

import "fmt"

// ANSI escape codes for colors
const (
	Reset          = "\033[0m"
	Black          = "\033[30m"
	Red            = "\033[31m"
	Green          = "\033[32m"
	Yellow         = "\033[33m"
	Blue           = "\033[34m"
	Purple         = "\033[35m"
	Cyan           = "\033[36m"
	White          = "\033[37m"
	BlackBg        = "\033[40m"
	RedBg          = "\033[41m"
	GreenBg        = "\033[42m"
	YellowBg       = "\033[43m"
	BlueBg         = "\033[44m"
	PurpleBg       = "\033[45m"
	CyanBg         = "\033[46m"
	WhiteBg        = "\033[47m"
	BrightBlackBg  = "\033[100m"
	BrightRedBg    = "\033[101m"
	BrightGreenBg  = "\033[102m"
	BrightYellowBg = "\033[103m"
	BrightBlueBg   = "\033[104m"
	BrightPurpleBg = "\033[105m"
	BrightCyanBg   = "\033[106m"
	BrightWhiteBg  = "\033[107m"
)

// ApplyColor applies the specified color to the input text
func Color(text, color, background string) string {
	if background == "" {
		return fmt.Sprintf("%s%s%s", color, text, Reset)
	}
	//if background color is specified
	return fmt.Sprintf("%s%s%s%s", color, background, text, Reset)
}
