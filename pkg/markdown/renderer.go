package markdown

import (
	"os"

	"github.com/charmbracelet/glamour"
	"golang.org/x/term"
)

func Render(in string, style string) (string, error) {
	if style != "" && style != "auto" {
		return glamour.Render(in, style)
	}

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return glamour.Render(in, "notty")
	}

	// detect background color and pick either the default dark or light theme
	r, err := glamour.NewTermRenderer(glamour.WithAutoStyle())
	if err != nil {
		return "", err
	}
	return r.Render(in)
}
