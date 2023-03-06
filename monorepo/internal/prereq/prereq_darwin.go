//go:build darwin

package prereq

import (
	"github.com/idlebot/monorepo/monorepo/global"
)

func checkPrerequisites() error {
	brewPath, err := exec.LookPath("brew")
	if err != nil {
		return fmt.Errorf("brew not found")
	}

	formulas, err := listInstalledFormulas()
	if err != nil {
		return err
	}

	if err := formulas.IsInstalled("git"); err != nil {
		return err
	}

	return nil
}

func checkBrewFormul(name string) error {
	brewPath, err := exec.LookPath("brew")
	if err != nil {
		return fmt.Errorf("brew not found")
	}

	output, err := shell.ExecuteOutput(brewPath, "-qq", "list", name)
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

// isFormulaInstalled returns true if a formula or cask is already installed
func isFormulaInstalled(formula string) bool {
	formulas, err := listInstalledFormulae()
	if err != nil {
		return false
	}
	_, isInstalled := formulas[formula]
	return isInstalled
}

type FormulaVersion struct {
	Stable string `json:"stable"`
}

type Formula struct {
	Name     string         `json:"name"`
	FullName string         `json:"full_name"`
	Version  FormulaVersion `json:"versions"`
}

type Formulas map[string]string

func (f Formulas) IsInstalled(formula string) bool {
	_, isInstalled := f[formula]
	return isInstalled
}

// listInstalledFormulas list all installed formulas. We named the latin
// plural Formulae just to be consistent with Homebrew vocabulary
func listInstalledFormulas() (Formulas, error) {
	brewPath, err := Path()
	if err != nil {
		return nil, err
	}

	output, err := shell.ExecuteOutput(brewPath, "info", "--installed", "--json")
	if err != nil {
		return nil, fmt.Errorf("brew info --installed --json: %w", err)
	}

	formulas := []*Formula{}
	if err := json.Unmarshal([]byte(output), &formulas); err != nil {
		return nil, fmt.Errorf("brew info --installed --json: %w", err)
	}

	installed := Formulas{}
	for _, formula := range formulas {
		installed[formula.Name] = formula.Version.Stable
	}

	return installed, nil
}
