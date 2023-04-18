package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/idlebot/monorepo/rules/internal/proto-options/internal/file"
)

const (
	toolName = "proto-options"

	verboseArg            = "verbose"
	quietArg              = "quiet"
	inputFileArg          = "input"
	outputFileArg         = "output"
	packageNameArg        = "package"
	repositoryNameArg     = "repository"
	goPackageArg          = "go_package"
	csharpNamespaceArg    = "csharp_namespace"
	javaPackageArg        = "java_package"
	javaOuterClassnameArg = "java_outer_classname"
	phpNamespaceArg       = "php_namespace"
	rubyPackageArg        = "ruby_package"
	objcClassPrefixArg    = "objc_class_prefix"
)

var (
	packageNameExp   = regexp.MustCompile(`package\s+(?P<PackageName>[a-zA-Z][\w.]*)`)
	packageNameIndex = packageNameExp.SubexpIndex("PackageName")
)

type Settings struct {
	toolName string

	quiet   bool
	verbose bool

	inputFile  string
	outputFile string

	packageName    string
	repositoryName string

	goPackage string

	csharpNamespace string

	javaPackage        string
	javaOuterClassname string

	phpNamespace string

	rubyPackage string

	objcClassPrefix string
}

func NewSettings(args []string) (*Settings, error) {
	settings := &Settings{
		toolName: "proto-options",
	}
	fs := flag.NewFlagSet(toolName, flag.ExitOnError)

	fs.BoolVar(&settings.verbose, verboseArg, false, "enable verbose logging (default: false)")
	fs.BoolVar(&settings.quiet, quietArg, true, "supress all output (default: true)")

	fs.StringVar(&settings.inputFile, inputFileArg, "", "Input .proto file (required).")
	fs.StringVar(&settings.outputFile, outputFileArg, "", "Output .proto file with options (required)")
	fs.StringVar(&settings.packageName, packageNameArg, "", "Protobuf package name; If empty, the package name will be inferred from the input file")
	fs.StringVar(&settings.repositoryName, repositoryNameArg, "", "Git repository name (required, e.g. 'github.com/hello/world'")
	fs.StringVar(&settings.goPackage, goPackageArg, "", "Go package name; If empty, the Go package name will be inferred from the Protobuf package name and Github repository name")
	fs.StringVar(&settings.csharpNamespace, csharpNamespaceArg, "", "C# namespace; If empty, the namespace will be inferred from the package name")
	fs.StringVar(&settings.javaPackage, javaPackageArg, "", "Java package name; If empty, the Java package name will be inferred from the Protobuf package name and Github repository name")
	fs.StringVar(&settings.javaOuterClassname, javaOuterClassnameArg, "", "Java outer classname; If empty, the Java package name will be inferred from the output file name")
	fs.StringVar(&settings.phpNamespace, phpNamespaceArg, "", "PHP namespace; If empty, the PHP namespace will be inferred from the Protobuf package name")
	fs.StringVar(&settings.rubyPackage, rubyPackageArg, "", "Ruby package; If empty, the Ruby package will be inferred from the Protobuf package name")
	fs.StringVar(&settings.objcClassPrefix, objcClassPrefixArg, "", "Objective C class prefix (required)")

	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}

	if settings.inputFile == "" {
		return nil, fmt.Errorf("-%s argument is required", inputFileArg)
	}

	if settings.outputFile == "" {
		return nil, fmt.Errorf("-%s argument is required", outputFileArg)
	}

	if settings.repositoryName == "" {
		return nil, fmt.Errorf("-%s argument is required", repositoryNameArg)
	}

	if settings.objcClassPrefix == "" {
		return nil, fmt.Errorf("-%s argument is required", objcClassPrefixArg)
	}

	if settings.packageName == "" {
		settings.packageName, err = getPackageName(settings.inputFile)
		if err != nil {
			return nil, err
		}
	}

	return settings, nil
}

func (s *Settings) ToolName() string {
	return s.toolName
}

func (s *Settings) Quiet() bool {
	return s.quiet
}

func (s *Settings) Verbose() bool {
	return s.verbose
}

func (s *Settings) InputFile() string {
	return s.inputFile
}

func (s *Settings) OutputFile() string {
	return s.outputFile
}

func (s *Settings) PackageName() string {
	return s.packageName
}

func (s *Settings) GoPackage() string {
	if s.goPackage == "" {
		filename := path.Join(s.repositoryName, strings.ReplaceAll(s.packageName, ".", "/"))
		serviceName := parseServiceName(s.packageName)
		packageAlias := serviceName.Name + serviceName.Version
		s.goPackage = fmt.Sprintf("%s;%s", filename, packageAlias)
	}

	return s.goPackage
}

func (s *Settings) RepositoryName() string {
	return s.repositoryName
}

func (s *Settings) CSharpNamespace() string {
	if s.csharpNamespace == "" {
		s.csharpNamespace = snakeCaseToCamelCase(s.packageName)
	}
	return s.csharpNamespace
}

func (s *Settings) JavaPackage() string {
	if s.javaPackage == "" {
		repoParts := strings.Split(s.repositoryName, "/")
		domainParts := strings.Split(repoParts[0], ".")
		for i, j := 0, len(domainParts)-1; i < j; i, j = i+1, j-1 {
			domainParts[i], domainParts[j] = domainParts[j], domainParts[i]
		}
		s.javaPackage = strings.Join(append(append(domainParts, repoParts[1:]...), s.packageName), ".")
	}

	return s.javaPackage
}

func (s *Settings) JavaOuterClassname() string {
	if s.javaOuterClassname == "" {
		filename := path.Base(s.outputFile)
		baseName := filename[:len(filename)-len(path.Ext(filename))]
		s.javaOuterClassname = snakeCaseToCamelCase(baseName) + "Proto"
	}

	return s.javaOuterClassname
}

func (s *Settings) JavaMultipleFiles() bool {
	// we always force java_multiple_files to false to have a single faile that matches the
	// java_outer_classname + "Proto.Java", e.g. GreeterProto.java
	// This makes it easier to create Bazel rules so we know the name of the java generated file
	// during the analysis part of the build.
	return false
}

func (s *Settings) PHPNamespace() string {
	if s.phpNamespace == "" {
		s.phpNamespace = strings.ReplaceAll(snakeCaseToCamelCase(s.packageName), ".", `\\`)
	}

	return s.phpNamespace
}

func (s *Settings) RubyPackage() string {
	if s.rubyPackage == "" {
		s.rubyPackage = strings.ReplaceAll(snakeCaseToCamelCase(s.packageName), ".", "::")
	}

	return s.rubyPackage
}

func (s *Settings) ObjcClassPrefix() string {
	return s.objcClassPrefix
}

func snakeCaseToCamelCase(s string) string {
	// Split the string into words based on underscores
	words := strings.Split(s, "_")

	// Convert each word to camel case
	for i := range words {
		words[i] = strings.Title(words[i])
	}

	// Join the words together without any underscores
	return strings.Join(words, "")
}

func getPackageName(inputFile string) (string, error) {
	matches, err := file.Grep(inputFile, regexp.MustCompile(`^\s*package\s+([a-zA-Z][\w.]*)\s*;$`))
	if err != nil {
		return "", err
	}

	if len(matches) == 0 {
		return "", fmt.Errorf("unable to find package name from input file '%s'", inputFile)
	}

	if len(matches) > 1 {
		return "", fmt.Errorf("multiple package names found in input file '%s'", inputFile)
	}

	packageNameMatches := packageNameExp.FindStringSubmatch(matches[0])
	packageNameIndex := packageNameExp.SubexpIndex("PackageName")

	return packageNameMatches[packageNameIndex], nil
}

func ParseParameters() (*Settings, error) {
	return NewSettings(os.Args[1:])
}
