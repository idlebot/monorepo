//go:build darwin && arm64

package prereq

import (
	"github.com/idlebot/monorepo/monorepo/global"
)

func setPlatformInfo() {
	global.Architecture = "arm64"
	global.OS = "Darwin"
}
