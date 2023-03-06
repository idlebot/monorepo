//go:build linux && amd64

package prereq

import (
	"github.com/idlebot/monorepo/monorepo/global"
)

func setPlatformInfo() {
	global.Architecture = "x86_64"
	global.OS = "Linux"
}
