package console

import (
	"fmt"

	"github.com/idlebot/monorepo/monorepo/global"
)

func Info(message string) {
	if global.Quiet {
		return
	}
	fmt.Println(ToolNamePrefix, message)
}

func Infof(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	Info(message)
}
