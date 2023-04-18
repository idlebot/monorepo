package console

type Color int

const (
	Reset Color = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	LightGray
	Gray
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

var (
	colors = map[Color]string{
		Reset:        "\033[0m",
		Red:          "\033[31m",
		Green:        "\033[32m",
		Yellow:       "\033[33m",
		Blue:         "\033[34m",
		Purple:       "\033[35m",
		Cyan:         "\033[36m",
		LightGray:    "\033[37m",
		Gray:         "\033[90m",
		LightRed:     "\033[91m",
		LightGreen:   "\033[92m",
		LightYellow:  "\033[93m",
		LightBlue:    "\033[94m",
		LightMagenta: "\033[95m",
		LightCyan:    "\033[96m",
		White:        "\033[97m",
	}
)

func getColor(c Color) string {
	return colors[c]
}
