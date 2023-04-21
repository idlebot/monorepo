package file

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Grep(inputFile string, expr *regexp.Regexp) ([]string, error) {
	// preserve file stat so we can applied the same permissions
	// on the edited file
	_, err := os.Stat(inputFile)
	if err != nil {
		return nil, fmt.Errorf("grep '%s': %w", inputFile, err)
	}

	reader, err := os.OpenFile(inputFile, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("grep '%s': %w", inputFile, err)
	}
	defer reader.Close()

	matches := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, expr.FindAllString(line, -1)...)
	}

	return matches, nil
}
