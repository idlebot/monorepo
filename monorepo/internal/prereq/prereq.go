package prereq

import (
	"fmt"

	"github.com/idlebot/monorepo/monorepo/internal/console"
)

func SetPlatformInfo() {
	setPlatformInfo()
}

func CheckPrerequisites() error {
	console.Verbose("Checking prerequisites")

	err := checkPrerequisites()
	if err != nil {
		return fmt.Errorf("check prerequisites: %w", err)
	}

	return nil
}
