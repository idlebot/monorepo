package file

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Write(outputFile, contents string) error {
	reader := strings.NewReader(contents)
	writer, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("unable to create file '%s': %w", outputFile, err)
	}
	defer writer.Close()
	_, err = io.Copy(writer, reader)
	if err != nil {
		return fmt.Errorf("unable to write file '%s': %w", outputFile, err)
	}

	return nil
}
