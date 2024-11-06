package knife4j

import (
	_ "embed"
	"net/http"
	"path/filepath"
	"strings"

	"golang.org/x/net/webdav"
)

//go:embed resource/doc.html
var index string

func Knife4jSwagHandler(prefix string) http.HandlerFunc {
	handler := &webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.Dir("resource"),
		LockSystem: webdav.NewMemLS(),
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		path := r.RequestURI

		path = strings.Replace(path, handler.Prefix, "", 1)
		if strings.Contains(path, "?") {
			path = path[:strings.Index(path, "?")]
		}

		switch filepath.Ext(path) {
		case ".html":
			r.Header.Set("Content-Type", "text/html; charset=utf-8")
		case ".css":
			r.Header.Set("Content-Type", "text/css; charset=utf-8")
		case ".js":
			r.Header.Set("Content-Type", "application/javascript")
		case ".png":
			r.Header.Set("Content-Type", "image/png")
		case ".json":
			r.Header.Set("Content-Type", "application/json; charset=utf-8")
		}

		switch path {
		case "index.html":
			w.Write([]byte(index))
		case "doc.html":
			w.Write([]byte(index))
		case "":
			http.Redirect(w, r, handler.Prefix+"index.html", http.StatusMovedPermanently)
		default:
			handler.ServeHTTP(w, r)
		}
	}
}
