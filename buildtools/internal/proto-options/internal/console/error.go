package console

import (
	"fmt"
	"os"
)

func formatError(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func Errorf(format string, a ...any) {
	Error(formatError(format, a...))
}

func Error(err error) {
	fmt.Fprintf(
		os.Stderr, `
%sError:%s %v

`,
		getColor(Red),
		getColor(Reset),
		err,
	)
}
