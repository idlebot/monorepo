package global

import (
	"os"
	"strings"
)

const (
	ToolName = "monorepo"
)

var (
	Command string = strings.Join(os.Args[1:], " ")

	Version      string
	OS           string
	Architecture string
)
