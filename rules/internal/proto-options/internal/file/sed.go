package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/idlebot/monorepo/rules/internal/proto-options/internal/console"
)

type LineOperation interface {
	Execute(line string) []string
	Eof() []string
}

func SedString(text string, operations []LineOperation) (string, error) {
	reader := strings.NewReader(text)
	outputBuffer := &bytes.Buffer{}
	writer := bufio.NewWriter(outputBuffer)

	err := process(reader, writer, operations)
	if err != nil {
		return "", err
	}

	err = writer.Flush()
	if err != nil {
		return "", err
	}

	return outputBuffer.String(), nil
}

// Sed provides functionality similar to the sed utility
func Sed(inputFile, outputFile string, operations []LineOperation) error {
	// preserve file stat so we can applied the same permissions
	// on the edited file
	inputFileStat, err := os.Stat(inputFile)
	if err != nil {
		return err
	}

	err = processFile(inputFile, outputFile, operations)
	if err != nil {
		return err
	}

	err = os.Chmod(outputFile, inputFileStat.Mode())
	if err != nil {
		return fmt.Errorf("unable to apply file permissions to edited file '%s': %w", outputFile, err)
	}

	return nil
}

// SedTemplate takes an input string as template and apply operations saving it to the output file
func SedTemplate(template string, outputFile string, operations []LineOperation) error {
	reader := strings.NewReader(template)

	writer, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("unable to create file '%s': %w", outputFile, err)
	}
	defer writer.Close()

	return process(reader, writer, operations)
}

func processFile(inputFile, outputFile string, operations []LineOperation) error {
	reader, err := os.OpenFile(inputFile, os.O_RDONLY, 0o644)
	if err != nil {
		return fmt.Errorf("unable to open file '%s': %w", inputFile, err)
	}
	defer reader.Close()

	if _, err := os.Stat(outputFile); err == nil {
		err = os.Remove(outputFile)
		if err != nil {
			return fmt.Errorf("unable to overwrite output file '%s': %w", outputFile, err)
		}
	}

	writer, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("unable to create file '%s': %w", outputFile, err)
	}
	defer writer.Close()

	return process(reader, writer, operations)
}

func process(reader io.Reader, writer io.Writer, operations []LineOperation) error {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		currentLines := []string{line}
		// we do this as a loop of currentLines as one of the line operations may
		// transform one line into multiple lines, so even though the scanner
		// only returns one line at a time, as operations execute, subsequent
		// operations need to be applied on all subsequent lines.
		for index := 0; index < len(currentLines); index++ {
			for _, operation := range operations {
				newLines := operation.Execute(currentLines[index])
				if len(newLines) == 0 {
					// remove current line and other operations are no longer relevant for the current line
					currentLines = append(currentLines[:index], currentLines[index+1:]...)
					break
				}

				// operation may return more than one line for a given current line, so
				// we replace the current line with the returned lines and move to the next
				// operation
				currentLines = append(currentLines[:index], append(newLines, currentLines[index+1:]...)...)
			}
		}

		err := writeLines(writer, currentLines)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("unable to scan file: %w", err)
	}

	for _, operation := range operations {
		newLines := operation.Eof()
		err := writeLines(writer, newLines)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeLines(writer io.Writer, lines []string) error {
	for _, currentLine := range lines {
		line := fmt.Sprintln(currentLine)
		_, err := writer.Write([]byte(line))
		if err != nil {
			return fmt.Errorf("unable to write line '%s': %w", line, err)
		}
	}
	return nil
}

type replaceAll struct {
	old string
	new string
}

func ReplaceAll(old, new string) LineOperation {
	return &replaceAll{
		old,
		new,
	}
}

func (l *replaceAll) Execute(line string) []string {
	newLine := strings.ReplaceAll(line, l.old, l.new)
	return []string{newLine}
}

func (l *replaceAll) Eof() []string {
	return []string{}
}

type deleteLine struct {
	line string
}

func DeleteLine(line string) LineOperation {
	return &deleteLine{
		line,
	}
}

func (l *deleteLine) Execute(line string) []string {
	if line == l.line {
		return []string{}
	}
	return []string{line}
}

func (l *deleteLine) Eof() []string {
	return []string{}
}

type appendIfNotFound struct {
	line  string
	found bool
}

func AppendIfNotFound(line string) LineOperation {
	return &appendIfNotFound{
		line:  line,
		found: false,
	}
}

func (l *appendIfNotFound) Execute(line string) []string {
	if line == l.line {
		l.found = true
	}
	return []string{line}
}

func (l *appendIfNotFound) Eof() []string {
	if !l.found {
		return []string{l.line}
	}
	return []string{}
}

type regExpReplace struct {
	searchRE *regexp.Regexp
	repl     string
}

func RegExpReplace(searchExpr, repl string) LineOperation {
	searchRE := regexp.MustCompile(searchExpr)
	return &regExpReplace{
		searchRE,
		repl,
	}
}

func (l *regExpReplace) Execute(line string) []string {
	result := l.searchRE.ReplaceAllString(line, l.repl)
	if result != line {
		console.Verbose(result)
		fmt.Println("found:", line)
		fmt.Println("result:", result)
	}

	return []string{result}
}

func (l *regExpReplace) Eof() []string {
	return []string{}
}

type regExpDelete struct {
	searchRE *regexp.Regexp
}

func RegExpDelete(searchExpr string) LineOperation {
	searchRE := regexp.MustCompile(searchExpr)
	return &regExpDelete{
		searchRE,
	}
}

func (l *regExpDelete) Execute(line string) []string {
	result := l.searchRE.FindString(line)
	if result == "" {
		// not found, we keep the line
		return []string{line}
	}
	fmt.Println("deleting:", line)
	return []string{}
}

func (l *regExpDelete) Eof() []string {
	return []string{}
}

type searchReplace struct {
	searchExpression *regexp.Regexp
	replaceValue     string
}
type replaceVariables struct {
	searchExpressions []searchReplace
}

var validVariableName = regexp.MustCompile(`^[a-z][a-z0-9\_]*[a-z0-9]$`)

// ReplaceVariables create a sed operation that replace variable
// values in a template. For example:
//
// Input template:
//
//	resource "google_project" "${project_resource_name}_project" {
//	  name = "${project_name}"
//	  project_id = "${project_name}-$${local.project_id}" // $$ escapes the $
//	}
//
// vars:
//
//	{
//		  "project_resource_name": "my_project",
//		  "project_name": 		   "my-project",
//	}
//
// Result:
//
//	resource "google_project" "my_project_project" {
//	  name = "my-project"
//	  project_id = "my-project-${local.project_id}" // $$ escapes the $
//	}
func ReplaceVariables(vars map[string]string) LineOperation {
	searchExpressions := make([]searchReplace, 0, len(vars))
	for name, value := range vars {
		if !validVariableName.MatchString(name) {
			panic(fmt.Sprintf("'%s' is not a valid replace variable name, must match regular expression `^[a-z][a-z0-9\\_]*[a-z0-9]$`", name))
		}
		searchExpression := searchReplace{
			searchExpression: regexp.MustCompile(fmt.Sprintf(`([^$])\${%s}`, name)),
			replaceValue:     "${1}" + value,
		}
		searchExpressions = append(searchExpressions, searchExpression)
	}
	return &replaceVariables{
		searchExpressions,
	}
}

func (l *replaceVariables) Execute(line string) []string {
	// a little bit of a hack so we don't have to deal with the beginning
	// of the line and $ being escaped
	paddedLine := " " + line
	for _, expr := range l.searchExpressions {
		paddedLine = expr.searchExpression.ReplaceAllString(paddedLine, expr.replaceValue)
	}
	return []string{paddedLine[1:]}
}

func (l *replaceVariables) Eof() []string {
	return []string{}
}
