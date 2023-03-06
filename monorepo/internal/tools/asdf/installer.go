package asdf

import (
	"os/exec"

	"github.com/idlebot/monorepo/monorepo/internal/shell"
	"github.com/idlebot/monorepo/monorepo/internal/tools/installer"
)

var (
	defaultAsdfPlugins = []string{
		"golang",
		"buf",
		"protoc",
		"bazel",
	}
)

type asdfInstaller struct{}

var instance = &asdfInstaller{}

func AsdfInstaller() installer.Installer {
	return instance
}

func (i *asdfInstaller) Name() string {
	return "asdf"
}

func (i *asdfInstaller) IsInstalled() bool {
	_, err := exec.LookPath("asdf")
	return err == nil
}

func (i *asdfInstaller) Install() error {
	return nil
}

func (i *asdfInstaller) Update() error {
	return nil
}

func (i *asdfInstaller) ConfigureProfile(profile *shell.ProfileEditor) error {
	// asdfConfigurationScript := path.Join(homebrew.HomebrewRoot, "opt/asdf/libexec/asdf.sh")
	// err := profile.WriteLine(fmt.Sprintf(". %s", asdfConfigurationScript))
	// if err != nil {
	// 	return fmt.Errorf("configuring profile: %w", err)
	// }

	// // zsh shell completion is not working correctly
	// if shell.GetCurrent() == "bash" {
	// 	asdfCompletionScript := path.Join(homebrew.HomebrewRoot, "etc/bash_completion.d/asdf.bash")
	// 	err = profile.WriteLine(fmt.Sprintf(". %s", asdfCompletionScript))
	// 	if err != nil {
	// 		return fmt.Errorf("configuring profile: %w", err)
	// 	}
	// }

	return nil
}

func init() {
	installer.RegisterInstaller(instance)
}
