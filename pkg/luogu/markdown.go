package luogu

import "github.com/charmbracelet/glamour"

func renderMarkdown(in string, style string) (string, error) {
	if style != "" && style != "auto" {
		return glamour.Render(in, style)
	}

	r, err := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
	)
	if err != nil {
		return "", err
	}
	return r.Render(in)
}
