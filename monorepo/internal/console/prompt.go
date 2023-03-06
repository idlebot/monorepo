package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/idlebot/monorepo/monorepo/internal/slices"
)

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("reading line from stdin: %w", err))
	}
	return ""
}

func printPrompt(prompt, defaultValue string) {
	fmt.Printf(
		"%s %s (default [%s]) ",
		ToolNamePrefix,
		prompt,
		TextColor(LightCyan, defaultValue),
	)
}

func Prompt(prompt string, defaultValue string) string {
	printPrompt(prompt, defaultValue)
	value := readLine()
	if value == "" {
		return defaultValue
	}
	return value
}

func PromptOptions(prompt string, defaultOption string, validOptions []string) string {
	for {
		printPrompt(prompt, defaultOption)
		value := readLine()
		if value == "" {
			return defaultOption
		}

		if slices.Contains(validOptions, value) {
			return value
		}

		Infof("'%s' is not a valid answer.", value)
	}
}

func PromptMenu(prompt string, defaultOption string, options []string) string {
	if len(options) < 1 {
		panic("PromptMenu: must have at least one option")
	}

	if defaultOption == "" {
		defaultOption = options[0]
	} else if !slices.Contains(options, defaultOption) {
		panic("PromptMenu: defaultOption is not a valid option")
	}

	fmt.Println()
	optionNumbers := make([]string, 0, len(options))
	for index, option := range options {
		optionNumber := strconv.Itoa(index + 1)
		optionNumbers = append(optionNumbers, optionNumber)
		optionNumberText := TextColor(LightCyan, optionNumber)
		fmt.Printf("[%s] %s\n", optionNumberText, option)
	}
	fmt.Println()

	defaultOptionIndex := slices.Index(options, defaultOption) + 1
	defaultOptionNumber := strconv.Itoa(defaultOptionIndex)
	optionNumber := PromptOptions(prompt, defaultOptionNumber, optionNumbers)
	index, err := strconv.Atoi(optionNumber)
	if err != nil {
		// should never happen as all valid options are a number
		panic(err)
	}
	option := options[index-1]
	return option
}
