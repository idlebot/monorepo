package generator

// Settings is the interface that wraps the basic settings methods.

type Settings interface {
	InputFile() string
	OutputFile() string
	PackageName() string
	RepositoryName() string
	GoPackage() string
	CSharpNamespace() string
	JavaPackage() string
	JavaOuterClassname() string
	JavaMultipleFiles() bool
	PHPNamespace() string
	RubyPackage() string
	ObjcClassPrefix() string
}
