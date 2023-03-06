package console

import (
	"fmt"

	"github.com/idlebot/monorepo/monorepo/global"
)

func Verbose(message string) {
	if global.Quiet {
		return
	}

	if global.Verbose {
		fmt.Println(message)
	}
}

func Verbosef(format string, a ...any) {
	if global.Verbose {
		message := fmt.Sprintf(format, a...)
		Verbose(message)
	}
}
