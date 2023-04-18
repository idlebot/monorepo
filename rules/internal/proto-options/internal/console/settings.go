package console

var (
	quiet          = false
	verbose        = false
	toolNamePrefix = ""
)

type Settings interface {
	Quiet() bool
	Verbose() bool
	ToolName() string
}

func Initialize(settings Settings) {
	quiet = settings.Quiet()
	verbose = settings.Verbose()
	toolName := settings.ToolName()
	toolNamePrefix = SprintfColor(Yellow, "%s:", toolName)
}
