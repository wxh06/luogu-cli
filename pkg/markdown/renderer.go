package markdown

import (
	"os"

	"github.com/charmbracelet/glamour"
	"golang.org/x/term"
)

func Render(in string, style string) (string, error) {
	// 源码（无样式）
	if style == "" {
		return in, nil
	}

	if style == "auto" && !term.IsTerminal(int(os.Stdout.Fd())) {
		return glamour.Render(in, "notty")
	}

	return glamour.Render(in, style)
}
