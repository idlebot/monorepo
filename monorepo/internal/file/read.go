package file

import (
	"fmt"
	"io/ioutil"
)

func Read(filename string) (string, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("unable to read file '%s': %w", filename, err)
	}
	return string(buffer), nil
}
