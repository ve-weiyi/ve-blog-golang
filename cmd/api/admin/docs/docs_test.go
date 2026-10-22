package docs

import (
	_ "embed"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

//go:embed admin.json
var docs string

func Test_Load(t *testing.T) {

	doc, err := loads.Analyzed(json.RawMessage(docs), "")
	if err != nil {
		t.Fatalf("load spec failed: %v", err)
	}

	sp := doc.Spec()
	if len(sp.Paths.Paths) == 0 {
		t.Fatal("no paths in spec")
	}

	if len(getRoutes(sp)) == 0 {
		t.Fatal("no routes in spec")
	}

	//var out map[string]map[string]any
	//jsonconv.AnyToAny(sp.Paths.Paths, &out)
	//t.Log(debugx.Dump(out))
}

func getRoutes(sp *spec.Swagger) map[string]map[string]*spec.Operation {
	// map[path][method] -> operation
	routes := make(map[string]map[string]*spec.Operation)

	for k, v := range sp.Paths.Paths {
		if routes[k] == nil {
			routes[k] = make(map[string]*spec.Operation)
		}

		if v.Get != nil {
			routes[k][http.MethodGet] = v.Get
		}

		if v.Put != nil {
			routes[k][http.MethodPut] = v.Put
		}

		if v.Post != nil {
			routes[k][http.MethodPost] = v.Post
		}

		if v.Delete != nil {
			routes[k][http.MethodDelete] = v.Delete
		}

		if v.Options != nil {
			routes[k][http.MethodOptions] = v.Options
		}

		if v.Head != nil {
			routes[k][http.MethodHead] = v.Head
		}
	}

	return routes
}
