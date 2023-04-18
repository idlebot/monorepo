package console

import (
	"fmt"
)

func Verbose(message string) {
	if quiet {
		return
	}

	if verbose {
		fmt.Println(message)
	}
}

func Verbosef(format string, a ...any) {
	if verbose {
		message := fmt.Sprintf(format, a...)
		Verbose(message)
	}
}
