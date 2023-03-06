package file

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func Sha256(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("sha256 file '%s': %w", filename, err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("sha256 file '%s': %w", filename, err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
