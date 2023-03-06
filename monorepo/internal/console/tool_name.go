package console

import (
	"github.com/idlebot/monorepo/monorepo/global"
)

var (
	ToolNamePrefix string
)

func init() {
	ToolNamePrefix = SprintfColor(Yellow, "%s:", global.ToolName)
}
