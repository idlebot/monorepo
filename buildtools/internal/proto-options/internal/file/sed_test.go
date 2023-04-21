package file

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const backtick = "`"

func containsLine(t *testing.T, line string, contents []string) bool {
	found := false
	for _, contentLine := range contents {
		if contentLine == line {
			found = true
			break
		}
	}
	return assert.True(t, found)
}

func TestProcess(t *testing.T) {
	inputContents := `
	if [ -z "$SSH_AUTH_SOCK" ] ; then
    eval ` + backtick + "ssh-agent -s" + backtick + `
    ssh-add
fi

# Start Docker daemon automatically when logging in if not running.'
RUNNING=` + backtick + "ps aux | grep dockerd | grep -v grep" + backtick + `
if [ -z "$RUNNING" ]; then
    sudo dockerd > /dev/null 2>&1 &
    disown
fi

PATH=$PATH:/home/marcelou/.truepay/gcp-common-scripts`

	t.Run("AppendIfNotFound", func(t *testing.T) {
		output, err := SedString(
			"",
			[]LineOperation{
				AppendIfNotFound("source ./hello"),
			},
		)
		assert.NoError(t, err)
		lines := strings.Split(output, "\n")

		containsLine(t, "source ./hello", lines)
	})

	t.Run("OperationsAreExecutedInOrder", func(t *testing.T) {
		output, err := SedString(
			inputContents,
			[]LineOperation{
				ReplaceAll("PATH", "path"),
				// will not found the line below as the operation above changes PATH to path
				DeleteLine("PATH=$PATH:/home/marcelou/.truepay/gcp-common-scripts"),
			},
		)
		assert.NoError(t, err)

		lines := strings.Split(output, "\n")

		containsLine(t, "path=$path:/home/marcelou/.truepay/gcp-common-scripts", lines)
	})
}

func TestReplaceAll(t *testing.T) {
	t.Run("Eof", func(t *testing.T) {
		op := ReplaceAll("hello", "world")
		result := op.Eof()
		assert.Empty(t, result)
	})

	t.Run("NotFound", func(t *testing.T) {
		op := ReplaceAll("hello", "world")
		result := op.Execute("nothing to see here")
		assert.Equal(t, "nothing to see here", result[0])
	})

	t.Run("Found", func(t *testing.T) {
		op := ReplaceAll("hello", "world")
		result := op.Execute("this is your hello")
		assert.Equal(t, "this is your world", result[0])
	})

	t.Run("MultipleFound", func(t *testing.T) {
		op := ReplaceAll("hello", "life")
		result := op.Execute("hello, hello!")
		assert.Equal(t, "life, life!", result[0])
	})
}

func TestDeleteLine(t *testing.T) {
	t.Run("Eof", func(t *testing.T) {
		op := DeleteLine("hello, world!")
		result := op.Eof()
		assert.Empty(t, result)
	})

	t.Run("NotFound", func(t *testing.T) {
		op := DeleteLine("hello, world!")
		result := op.Execute("nothing to see here")
		assert.Equal(t, "nothing to see here", result[0])
	})

	t.Run("Found", func(t *testing.T) {
		op := DeleteLine("hello, world!")
		result := op.Execute("hello, world!")
		assert.Empty(t, result)
	})
}

func TestAppendIfNotFound(t *testing.T) {
	t.Run("NotFound", func(t *testing.T) {
		op := AppendIfNotFound("source ./hello")
		op.Execute("# line here")
		op.Execute("# line there")
		result := op.Eof()
		assert.Equal(t, "source ./hello", result[0])
	})

	t.Run("Found", func(t *testing.T) {
		op := AppendIfNotFound("source ./hello")
		op.Execute("source ./hello")
		op.Execute("# line there")
		result := op.Eof()
		assert.Empty(t, result)
	})
}

func TestRegExpReplace(t *testing.T) {
	t.Run("Eof", func(t *testing.T) {
		op := RegExpReplace(`"version"([ \t]*):([ \t]*)"\d+.\d+.\d+"`, `"version"$1:$2"1.0.0"`)
		result := op.Eof()
		assert.Empty(t, result)
	})

	t.Run("NotFound", func(t *testing.T) {
		op := RegExpReplace(`"version"([ \t]*):([ \t]*)"\d+.\d+.\d+"`, `"version"$1:$2"1.0.42"`)
		result := op.Execute("# line here")
		assert.Equal(t, "# line here", result[0])
	})

	t.Run("Found", func(t *testing.T) {
		op := RegExpReplace(`"version"([ \t]*):([ \t]*)"\d+.\d+.\d+"`, `"version"$1:$2"1.0.42"`)
		result := op.Execute(`"version":"0.0.1"`)
		assert.Equal(t, `"version":"1.0.42"`, result[0])
	})

	t.Run("ReplacePreserveWhatIsAround", func(t *testing.T) {
		op := RegExpReplace(`"version"([ \t]*):([ \t]*)"\d+.\d+.\d+"`, `"version"$1:$2"1.0.42"`)
		result := op.Execute(`  "version": "0.0.1" `)
		assert.Equal(t, `  "version": "1.0.42" `, result[0])
		result = op.Execute(`// "version" : "0.33.1",`)
		assert.Equal(t, `// "version" : "1.0.42",`, result[0])
	})

	t.Run("ReplaceTerraformModuleVersion", func(t *testing.T) {
		op := RegExpReplace(`"git@github\.com:truepay\/tf-truepay-modules.git\/\/modules\/(.*)\?ref=v\d+\.\d+\.\d+"`,
			`"git@github.com:truepay/tf-truepay-modules.git//modules/$1?ref=v1.0.42"`,
		)
		result := op.Execute(`  source       = "git@github.com:truepay/tf-truepay-modules.git//modules/domain_name?ref=v0.165.0"`)
		assert.Equal(t, `  source       = "git@github.com:truepay/tf-truepay-modules.git//modules/domain_name?ref=v1.0.42"`, result[0])
	})
}

func TestRegExpDelete(t *testing.T) {
	t.Run("Eof", func(t *testing.T) {
		op := RegExpDelete(`"version"([ \t]*):([ \t]*)"\d+.\d+.\d+"`)
		result := op.Eof()
		assert.Empty(t, result)
	})

	t.Run("NotFound", func(t *testing.T) {
		op := RegExpDelete(`"version"([ \t]*):([ \t]*)"\d+.\d+.\d+"`)
		result := op.Execute("# line here")
		assert.Equal(t, "# line here", result[0])
	})

	t.Run("Found", func(t *testing.T) {
		op := RegExpDelete(`"version"([ \t]*):([ \t]*)"\d+.\d+.\d+"`)
		result := op.Execute(`"version":"0.0.1"`)
		assert.Empty(t, result)
	})

	t.Run("FoundSubString", func(t *testing.T) {
		op := RegExpDelete(`\d+.\d+.\d+"`)
		result := op.Execute(`"version":"0.0.1"`)
		assert.Empty(t, result)
	})
}

func TestReplaceVariables(t *testing.T) {
	t.Run("Eof", func(t *testing.T) {
		op := ReplaceVariables(map[string]string{
			"greeting": "hello",
			"planet":   "earth",
		})
		result := op.Eof()
		assert.Empty(t, result)
	})

	t.Run("NotFound", func(t *testing.T) {
		op := ReplaceVariables(map[string]string{
			"greeting": "hello",
			"planet":   "earth",
		})
		result := op.Execute("# line here")
		assert.Equal(t, "# line here", result[0])
	})

	t.Run("FoundOneVariableBeginningOfLine", func(t *testing.T) {
		op := ReplaceVariables(map[string]string{
			"greeting": "hello",
			"planet":   "earth",
		})
		result := op.Execute("${greeting}, world!")
		assert.Equal(t, "hello, world!", result[0])
	})

	t.Run("FoundOneVariableMiddleOfLine", func(t *testing.T) {
		op := ReplaceVariables(map[string]string{
			"greeting": "hello",
			"planet":   "earth",
		})
		result := op.Execute("hello, ${planet}!")
		assert.Equal(t, "hello, earth!", result[0])
	})

	t.Run("ReplaceAll", func(t *testing.T) {
		op := ReplaceVariables(map[string]string{
			"greeting": "hello",
			"planet":   "earth",
		})
		result := op.Execute("${greeting}, ${planet}! ${planet} is a nice planet.")
		assert.Equal(t, "hello, earth! earth is a nice planet.", result[0])
	})

	t.Run("$$Escapes", func(t *testing.T) {
		op := ReplaceVariables(map[string]string{
			"greeting": "hello",
			"planet":   "earth",
		})
		result := op.Execute("$${greeting}, $${planet}!")
		assert.Equal(t, "$${greeting}, $${planet}!", result[0])
	})
}
