package console

import (
	"fmt"
)

func Info(message string) {
	if quiet {
		return
	}
	fmt.Println(toolNamePrefix, message)
}

func Infof(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	Info(message)
}
