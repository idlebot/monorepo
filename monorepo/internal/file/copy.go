package file

import (
	"fmt"
	"io"
	"os"
)

func Copy(source, target string) error {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("'%s' is not a regular file", source)
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("unable to open source file '%s': %w", source, err)
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("unable to create destination file '%s': %w", target, err)
	}
	defer targetFile.Close()
	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return fmt.Errorf("unable to copy file contents '%s' to '%s': %w", source, target, err)
	}

	err = os.Chmod(target, sourceFileStat.Mode())
	if err != nil {
		return fmt.Errorf("unable to copy file permissions to destination file '%s': %w", target, err)
	}

	return nil
}
