package templates

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed *.html
var FS embed.FS

func Execute(w io.Writer, name string, data interface{}) error {
	tmpl, err := template.ParseFS(FS, name)
	if err != nil {
		return fmt.Errorf("failed to parse embedded template %s: %w", name, err)
	}
	return tmpl.Execute(w, data)
}
