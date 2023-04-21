package generator

import (
	"fmt"
	"strings"

	"github.com/idlebot/monorepo/rules/internal/proto-options/internal/file"
)

func buildOperations(settings Settings) []file.LineOperation {
	packageInfo := []string{
		fmt.Sprintf("package %s;", settings.PackageName()),
		"",
		fmt.Sprintf(`option go_package = "%s";`, settings.GoPackage()),
		fmt.Sprintf(`option java_multiple_files = %v;`, settings.JavaMultipleFiles()),
		fmt.Sprintf(`option java_outer_classname = "%s";`, settings.JavaOuterClassname()),
		fmt.Sprintf(`option java_package = "%s";`, settings.JavaPackage()),
		fmt.Sprintf(`option csharp_namespace = "%s";`, settings.CSharpNamespace()),
		fmt.Sprintf(`option php_namespace = "%s";`, settings.PHPNamespace()),
		fmt.Sprintf(`option ruby_package = "%s";`, settings.RubyPackage()),
		fmt.Sprintf(`option objc_class_prefix = "%s";`, settings.ObjcClassPrefix()),
	}

	operations := []file.LineOperation{
		file.RegExpDelete(`option\s+go_package\s*[=]`),
		file.RegExpDelete(`option\s+java_multiple_files\s*[=]`),
		file.RegExpDelete(`option\s+java_outer_classname\s*[=]`),
		file.RegExpDelete(`option\s+java_package\s*[=]`),
		file.RegExpDelete(`option\s+csharp_namespace\s*[=]`),
		file.RegExpDelete(`option\s+php_namespace\s*[=]`),
		file.RegExpDelete(`option\s+ruby_package\s*[=]`),
		file.RegExpDelete(`option\s+objc_class_prefix\s*[=]`),
		file.RegExpReplace(`\s*package\s+([a-zA-Z][\w.]*)\s*\;`, strings.Join(packageInfo, "\n")),
	}

	return operations
}

func Execute(settings Settings) error {
	operations := buildOperations(settings)

	return file.Sed(settings.InputFile(), settings.OutputFile(), operations)
}
