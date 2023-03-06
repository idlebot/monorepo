//go:build darwin && amd64

package prereq

import (
	"github.com/idlebot/monorepo/monorepo/global"
)

func setPlatformInfo() {
	global.Architecture = "x86_64"
	global.OS = "Darwin"
}
