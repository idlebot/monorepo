package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var outputDir = "testdata/output"

func getInputFile(name string) string {
	dir, _ := os.Getwd()
	inputFile := path.Join(dir, "testdata", name)
	return inputFile
}

func getOutputFile(name string) string {
	outputFile := path.Join(outputDir, name)
	return outputFile
}

func TestNewSettings(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {
		inputFile := getInputFile("input.proto")
		outputFile := getOutputFile("long_output_name.proto")
		args := []string{
			"-input", inputFile,
			"-output", outputFile,
			"-repository", "github.com/hello/world",
			"-objc_class_prefix", "FOO",
		}

		settings, err := NewSettings(args)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, inputFile, settings.InputFile(), "unexpected input file")
		assert.Equal(t, outputFile, settings.OutputFile(), "unexpected output file")
		assert.Equal(t, "github.com/hello/world", settings.RepositoryName(), "unexpected github repository")
		assert.Equal(t, "input.v1", settings.PackageName(), "unexpected package name")
		assert.Equal(t, "github.com/hello/world/input/v1;inputv1", settings.GoPackage(), "unexpected go package name")
		assert.Equal(t, "Input.V1", settings.CSharpNamespace(), "unexpected c# namespace")
		assert.Equal(t, "com.github.hello.world.input.v1", settings.JavaPackage(), "unexpected java package name")
		assert.Equal(t, "LongOutputNameProto", settings.JavaOuterClassname(), "unexpected java outer classname")
		assert.Equal(t, `Input\\V1`, settings.PHPNamespace(), "unexpected PHP namespace")
		assert.Equal(t, "Input::V1", settings.RubyPackage(), "unexpected Ruby package name")
	})

	t.Run("missing_input_file", func(t *testing.T) {
		args := []string{
			"-output", "output.proto",
			"-repository", "github.com/hello/world",
			"-objc_class_prefix", "FOO",
		}

		_, err := NewSettings(args)
		assert.Error(t, err, "expected an error due to missing input file")
	})

	t.Run("missing_output_file", func(t *testing.T) {
		args := []string{
			"-input", getInputFile("input.proto"),
			"-repository", "github.com/hello/world",
			"-objc_class_prefix", "FOO",
		}

		_, err := NewSettings(args)
		assert.Error(t, err, "expected an error due to missing output file")
	})

	t.Run("missing_repository_name", func(t *testing.T) {
		args := []string{
			"-input", getInputFile("input.proto"),
			"-output", getOutputFile("output.proto"),
			"-objc_class_prefix", "FOO",
		}

		_, err := NewSettings(args)
		assert.Error(t, err, "expected an error due to missing github repository")
	})

	t.Run("missing_objc_class_prefix", func(t *testing.T) {
	})

	t.Run("invalid_input_file", func(t *testing.T) {
		args := []string{
			"-input", "missing.proto",
			"-output", getOutputFile("output.proto"),
			"-repository", "github.com/hello/world",
			"-objc_class_prefix", "FOO",
		}

		_, err := NewSettings(args)
		assert.Error(t, err)
	})
}
