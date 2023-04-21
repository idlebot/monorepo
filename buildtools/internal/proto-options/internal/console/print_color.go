package console

import (
	"fmt"
)

func TextColor(c Color, text string) string {
	return fmt.Sprintf("%s%s%s", getColor(c), text, getColor(Reset))
}

func SprintfColor(c Color, format string, a ...any) string {
	text := fmt.Sprintf(format, a...)
	return TextColor(c, text)
}

func PrintColor(color Color, a ...any) {
	fmt.Print(getColor(color))
	fmt.Print(a...)
	fmt.Print(getColor(Reset))
}

func PrintlnColor(color Color, a ...any) {
	fmt.Print(getColor(color))
	fmt.Println(a...)
	fmt.Print(getColor(Reset))
}
