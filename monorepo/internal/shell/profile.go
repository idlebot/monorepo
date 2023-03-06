package shell

import (
	"fmt"
	"os"
	"path"

	"github.com/idlebot/monorepo/monorepo/global"
	"github.com/idlebot/monorepo/monorepo/internal/file"
)

type ProfileEditor struct {
	script *Script
}

func NewProfileEditor() (*ProfileEditor, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("get user home directory: %w", err)
	}

	startupConfigScript := path.Join(homeDir, ".config", global.ToolName, fmt.Sprintf("%s-startup", global.ToolName))

	script, err := NewScript(startupConfigScript)
	if err != nil {
		return nil, err
	}

	return &ProfileEditor{
		script,
	}, nil
}

func (p *ProfileEditor) Close() error {
	p.script.Close()

	profilePath, err := ProfilePath()
	if err != nil {
		return fmt.Errorf("unable to get profile path: %w", err)
	}

	_, err = os.Stat(profilePath)
	if err != nil {
		profile, err := os.Create(profilePath)
		if err != nil {
			return fmt.Errorf("unable to create '%s': %w", profilePath, err)
		}
		err = profile.Close()
		if err != nil {
			return fmt.Errorf("unable to create '%s': %w", profilePath, err)
		}
	}

	fmt.Println("editing:", profilePath)

	err = file.Sed(profilePath, []file.LineOperation{
		file.AppendIfNotFound(fmt.Sprintf(". %s", p.script.Path())),
	})
	if err != nil {
		return fmt.Errorf("unable edit file '%s': %w", profilePath, err)
	}
	p.script = nil

	return nil
}

func (p *ProfileEditor) WriteLine(text string) error {
	return p.script.WriteLine(text)
}
