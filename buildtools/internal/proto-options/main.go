package main

import (
	"os"

	"github.com/idlebot/monorepo/rules/internal/proto-options/internal/cmd"
	"github.com/idlebot/monorepo/rules/internal/proto-options/internal/console"
	"github.com/idlebot/monorepo/rules/internal/proto-options/internal/generator"
)

func main() {
	settings, err := cmd.ParseParameters()
	if err != nil {
		console.Error(err)
		os.Exit(1)
	}

	console.Initialize(settings)

	err = generator.Execute(settings)
	if err != nil {
		console.Error(err)
		os.Exit(1)
	}
}
