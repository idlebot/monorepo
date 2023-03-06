package git

import (
	"fmt"

	"github.com/idlebot/monorepo/monorepo/internal/shell"
	"github.com/idlebot/monorepo/monorepo/internal/tools/internal/tools"
)

func Path() (string, error) {
	return tools.Path("git")
}

func Clone(directory, repository string) error {
	gitPath, err := Path()
	if err != nil {
		return fmt.Errorf("git clone: %w", err)
	}
	_, err = shell.ExecuteOutput(gitPath, "clone", repository, directory)
	if err != nil {
		return fmt.Errorf("git clone: %w", err)
	}
	return nil
}
