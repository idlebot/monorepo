package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseServiceName(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		validateParseServiceName(t, "example", &serviceName{
			Name: "example",
		})
	})

	t.Run("example.v1_alpha", func(t *testing.T) {
		validateParseServiceName(t, "example.v1_alpha", &serviceName{
			Name:    "example",
			Version: "v1_alpha",
		})
	})

	t.Run("example.v2_beta", func(t *testing.T) {
		validateParseServiceName(t, "example.v2_beta", &serviceName{
			Name:    "example",
			Version: "v2_beta",
		})
	})

	t.Run("example.some.other.package.v3", func(t *testing.T) {
		validateParseServiceName(t, "example.some.other.package.v3", &serviceName{
			Name:    "package",
			Version: "v3",
		})
	})

	t.Run("example.invalid.package.name.", func(t *testing.T) {
		validateParseServiceName(t, "example.invalid.package.name.", nil)
	})
}

func validateParseServiceName(t *testing.T, input string, expected *serviceName) {
	actual := parseServiceName(input)

	if expected == nil {
		assert.Nil(t, actual, "expected nil but got a non-nil result")
	} else {
		assert.NotNil(t, actual, "expected a non-nil result but got nil")
		assert.Equal(t, expected.Name, actual.Name, "unexpected service name")
		assert.Equal(t, expected.Version, actual.Version, "unexpected version")
	}
}
