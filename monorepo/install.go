package monorepo

import (
	"os"

	"github.com/idlebot/monorepo/monorepo/internal/console"
	"github.com/idlebot/monorepo/monorepo/internal/prereq"
)

func Install() {
	if err := prereq.CheckPrerequisites(); err != nil {
		console.Error(err)
		os.Exit(1)
	}
}
