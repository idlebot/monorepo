//go:build linux

package prereq

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/idlebot/monorepo/monorepo/internal/console"
	"github.com/idlebot/monorepo/monorepo/internal/shell"
)

func checkPrerequisites() error {
	err := checkAptDependency("build-essential")
	if err != nil {
		return err
	}

	err = checkAptDependency("git")
	if err != nil {
		return err
	}

	wsl, err := isWSL()
	if err != nil {
		return err
	}

	if wsl {
		// Check Windows Subsystem for Linux dependencies
		err = checkWSLPrerequisites()
		if err != nil {
			return err
		}
	}

	return nil
}

func checkAptDependency(name string) error {
	aptPath, err := exec.LookPath("apt")
	if err != nil {
		return fmt.Errorf("apt not found")
	}

	output, err := shell.ExecuteOutput(aptPath, "-qq", "list", name)
	if err != nil {
		return fmt.Errorf("apt -qq list %s: %w", name, err)
	}

	if !strings.Contains(output, "[installed]") {
		console.Infof("Please install '%s'.", name)
		console.Infof("'sudo apt install %s'", name)
		return fmt.Errorf("'%s' not installed", name)
	}

	return nil
}

// isWSL returns true if running under Windows Subsystem for Linux
func isWSL() (bool, error) {
	unamePath, err := exec.LookPath("uname")
	if err != nil {
		return false, fmt.Errorf("uname not found")
	}

	output, err := shell.ExecuteOutput(unamePath, "-a")
	if err != nil {
		return false, fmt.Errorf("uname -a: %w", err)
	}

	return strings.Contains(output, "microsoft-standard-WSL2"), nil
}

func checkWSLPrerequisites() error {
	_, err := exec.LookPath("www-browser")
	if err != nil {
		console.Info("Please install Windows Subsystem for Linux Utilities (wslu).")
		console.Info("'sudo apt install wslu'")
		return fmt.Errorf("wslu not installed: %w", err)
	}
	return nil
}
