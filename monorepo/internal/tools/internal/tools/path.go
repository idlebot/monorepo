package tools

import (
	"fmt"
	"os/exec"
)

func Path(toolName string) (string, error) {
	path, err := exec.LookPath(toolName)
	if err != nil {
		return "", fmt.Errorf(
			"unable to locate '%s' in path: %w",
			toolName,
			err,
		)
	}
	return path, nil
}
