package shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/idlebot/monorepo/monorepo/internal/console"
)

func ExecuteCmd(cmd *exec.Cmd) (int, error) {

	console.Verbosef("Executing: %s", cmd.Args)

	err := cmd.Start()
	if err != nil {
		return -1, fmt.Errorf("execute: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return cmd.ProcessState.ExitCode(), fmt.Errorf("execute: %w", err)
	}

	return cmd.ProcessState.ExitCode(), nil
}

func Execute(path string, args ...string) (int, error) {
	cmdArgs := make([]string, 0, len(args)+1)
	cmdArgs = append(cmdArgs, path)
	cmdArgs = append(cmdArgs, args...)

	cmd := &exec.Cmd{
		Path:   path,
		Args:   cmdArgs,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	return ExecuteCmd(cmd)
}

func ExecuteOutput(path string, args ...string) (string, error) {
	cmd := exec.Command(path, args...)
	buffer := bytes.NewBuffer(make([]byte, 0, 128))
	cmd.Stdout = buffer

	code, err := ExecuteCmd(cmd)
	if err != nil {
		return "", fmt.Errorf("error '%v' executing command: %w", code, err)
	}

	return strings.Trim(buffer.String(), "\n"), nil
}

func ExecuteShellScript(script string, args ...string) (int, error) {

	bashPath, err := exec.LookPath(GetCurrent())
	if err != nil {
		return -1, fmt.Errorf("bash is not installed, unable to execute script '%s'", script)
	}

	cmdArgs := make([]string, 0, len(args)+2)
	cmdArgs = append(cmdArgs, bashPath, script)
	cmdArgs = append(cmdArgs, args...)

	cmd := &exec.Cmd{
		Path:   bashPath,
		Args:   cmdArgs,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	return ExecuteCmd(cmd)
}
