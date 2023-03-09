package file

import (
	"fmt"
	"os"
)

func Read(filename string) (string, error) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("unable to read file '%s': %w", filename, err)
	}
	return string(buffer), nil
}
