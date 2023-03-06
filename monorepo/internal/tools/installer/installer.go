package installer

import (
	"github.com/idlebot/monorepo/monorepo/internal/shell"
)

type Installer interface {
	Name() string
	IsInstalled() bool
	Install() error
	Update() error
	ConfigureProfile(profile *shell.ProfileEditor) error
}

var (
	installers []Installer = make([]Installer, 0, 4)
)

func RegisterInstaller(installer Installer) {
	installers = append(installers, installer)
}

func Installers() []Installer {
	return installers
}
