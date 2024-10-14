package swagger

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/net/webdav"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/static"
)

func RegisterHttpSwagHandler(server *rest.Server, prefix, swaggerFile string) {
	f, err := os.ReadFile(swaggerFile)
	if err != nil {
		log.Println(err)
		return
	}

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("%s%s", prefix, "docs/blog.json"),
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(f)
		},
	})

	server.AddRoutes(static.PrefixRoutes(prefix, func(w http.ResponseWriter, r *http.Request) {
		httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("%s%s", prefix, "docs/blog.json")), //The url pointing to API definition
		).ServeHTTP(w, r)
	}))
}

// prefix: "/api/v1/swagger/"
func RegisterKnife4jSwagHandler(server *rest.Server, prefix, swaggerFile string) {
	f, err := os.ReadFile(swaggerFile)
	if err != nil {
		log.Println(err)
		return
	}

	server.AddRoutes(static.PrefixRoutes(prefix, func(w http.ResponseWriter, r *http.Request) {
		Knife4jSwagHandler(prefix).ServeHTTP(w, r)
	}))

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("%s%s", prefix, "swagger-resources"),
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[
			  {
				"name": "2.X版本",
				"url": "/v2/api-docs",
				"swaggerVersion": "2.0",
				"location": "v2"
			  }
			]`))
		},
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("%s%s", prefix, "v2/api-docs"),
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(f)
		},
	})
}

//go:embed resources/doc.html
var index string

func Knife4jSwagHandler(prefix string) http.HandlerFunc {
	handler := &webdav.Handler{
		FileSystem: webdav.Dir(files.GetRuntimeRoot() + "/resources"),
		LockSystem: webdav.NewMemLS(),
	}
	handler.Prefix = prefix

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
