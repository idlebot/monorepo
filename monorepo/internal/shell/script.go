package shell

import (
	"fmt"
	"os"
)

type Script struct {
	scriptPath string
	file       *os.File
}

func NewScript(scriptPath string) (*Script, error) {
	script, err := os.Create(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("create '%s' script: %w", scriptPath, err)
	}

	return &Script{
		scriptPath,
		script,
	}, nil
}

func (s *Script) Path() string {
	return s.scriptPath
}

func (s *Script) File() *os.File {
	return s.file
}

func (s *Script) WriteLine(text string) error {
	_, err := s.file.WriteString(fmt.Sprintf("%s\n", text))
	if err != nil {
		return fmt.Errorf("write file '%s': %w", s.scriptPath, err)
	}
	return nil
}

func (s *Script) Close() error {
	err := s.file.Sync()
	if err != nil {
		return err
	}
	err = s.file.Close()
	if err != nil {
		return err
	}
	err = os.Chmod(s.scriptPath, 0744)
	if err != nil {
		return err
	}

	s.file = nil
	return nil
}
