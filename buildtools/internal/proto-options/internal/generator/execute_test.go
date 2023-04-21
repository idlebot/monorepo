package generator

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/idlebot/monorepo/rules/internal/proto-options/internal/cmd"
)

const (
	// Set this to true to regerate testdata/*_expected.proto files.
	// Also remember to run using go test instead of bazel test, otherwise
	// it will generate the expected files in the bazel sandbox.
	REGENERATE_EXPECTED_TEST_OUTPUT = true
)

func TestRegenerateExpectedMustBeFalse(t *testing.T) {
	// we have this test case to ensure that we do not commit REGENERATE_EXPECTED_TEST_OUTPUT = true
	// this flag should only be used during development to regenerate the testdata/*_expected.proto files
	assert.False(t, REGENERATE_EXPECTED_TEST_OUTPUT, "REGENERATE_EXPECTED_TEST_OUTPUT must be false. It should only be set to update the testdata/expected_*.proto files.")
}

var outputDir string

func TestMain(m *testing.M) {
	if REGENERATE_EXPECTED_TEST_OUTPUT {
		dir, _ := os.Getwd()
		expectedDir := path.Join(dir, "testdata", "expected")
		err := os.MkdirAll(expectedDir, os.ModePerm)
		if err != nil {
			os.Exit(1)
		}
	}
	outputDir, err := ioutil.TempDir(os.Getenv("TEST_TEMPDIR"), "settings_test_*")
	if err != nil {
		os.Exit(1)
	}

	code := m.Run()

	os.RemoveAll(outputDir)
	os.Exit(code)
}

func getInputFile(name string) string {
	dir, _ := os.Getwd()
	inputFile := path.Join(dir, "testdata", name)
	return inputFile
}

func getOutputFile(name string) string {
	if REGENERATE_EXPECTED_TEST_OUTPUT {
		return getExpectedFile(name)
	}
	outputFile := path.Join(outputDir, name)
	return outputFile
}

func getExpectedFile(name string) string {
	dir, _ := os.Getwd()
	expectedFile := path.Join(dir, "testdata", "expected", name)

	return expectedFile
}

func compareFiles(file1, file2 string) (bool, error) {
	f1, err := os.Open(file1)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)

	for {
		hasLine1 := scanner1.Scan()
		hasLine2 := scanner2.Scan()

		if hasLine1 != hasLine2 {
			return false, nil
		}

		if !hasLine1 {
			break
		}

		if scanner1.Text() != scanner2.Text() {
			return false, nil
		}
	}

	return true, nil
}

func testExpectedProto(t *testing.T, file string) {
	inputFile := getInputFile(file)
	outputFile := getOutputFile(file)
	expectedFile := getExpectedFile(file)

	settings, err := cmd.NewSettings([]string{
		"-input", inputFile,
		"-output", outputFile,
		"-repository", "github.com/hello/world",
		"-objc_class_prefix", "HWM",
	})
	assert.NoError(t, err)

	err = Execute(settings)
	assert.NoError(t, err)

	areEqual, err := compareFiles(outputFile, expectedFile)
	assert.NoError(t, err)
	assert.True(t, areEqual)
}

func TestExecute(t *testing.T) {
	t.Run("complete_proto_with_options.proto", func(t *testing.T) {
		testExpectedProto(t, "complete_proto_with_options.proto")
	})

	t.Run("only_package_name_with_version.proto", func(t *testing.T) {
		testExpectedProto(t, "only_package_name_with_version.proto")
	})

	t.Run("only_package_name_without_version.proto", func(t *testing.T) {
		testExpectedProto(t, "only_package_name_without_version.proto")
	})

	t.Run("unformatted_proto.proto", func(t *testing.T) {
		testExpectedProto(t, "unformatted_proto.proto")
	})
}
