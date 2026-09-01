package static

import (
	"embed"
	"net/http"
)

//go:embed *
var FS embed.FS

func Handler() http.Handler {
	return http.FileServer(http.FS(FS))
}
