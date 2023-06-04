package prettyterm

import "fmt"

type Color string

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

type BGColor string

const (
	BGBlack         = "\033[40m"
	BGRed           = "\033[41m"
	BGGreen         = "\033[42m"
	BGYellow        = "\033[43m"
	BGBlue          = "\033[44m"
	BGMagenta       = "\033[45m"
	BGCyan          = "\033[46m"
	BGWhite         = "\033[47m"
	BGBrightBlack   = "\033[40;1m"
	BGBrightRed     = "\033[41;1m"
	BGBrightGreen   = "\033[42;1m"
	BGBrightYellow  = "\033[43;1m"
	BGBrightBlue    = "\033[44;1m"
	BGBrightMagenta = "\033[45;1m"
	BGBrightCyan    = "\033[46;1m"
	BGBrightWhite   = "\033[47;1m"
)

type Decoration string

const (
	Underlined = "\033[4m"
	Reversed   = "\033[7m"
)

func ClearTerminal() {
	fmt.Print("\033[2J")
}
func SetCursor(x int, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}
func SetColor(c Color) {
	fmt.Print(c)
}
func SetBGColor(c BGColor) {
	fmt.Print(c)
}
func SetDecoration(d Decoration) {
	fmt.Print(d)
}
func ResetTerminal() {
	fmt.Print("\033[0m")
}
