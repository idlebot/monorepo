package shell

import (
	"fmt"
	"os"
	"path"
)

func ShellPath() string {
	return os.Getenv("SHELL")
}

func GetCurrent() string {
	return path.Base(ShellPath())
}

func ProfileName() string {
	if GetCurrent() == "zsh" {
		return ".zprofile"
	}

	return ".profile"
}

func ProfilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to obtain user home directory: %w", err)
	}

	return path.Join(home, ProfileName()), nil
}
