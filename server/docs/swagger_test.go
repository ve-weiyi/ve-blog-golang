package docs

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/swaggo/swag"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"log"
	"testing"
)

func TestSwagger(t *testing.T) {
	var SwaggerInfo = &swag.Spec{
		Version:          "1.0",
		Host:             "localhost:9999",
		BasePath:         "/api/v1",
		Schemes:          []string{},
		Title:            "Swagger Example API",
		Description:      "This is a sample server celler server.",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  docTemplate,
		LeftDelim:        "{{",
		RightDelim:       "}}",
	}

	data := &T{}
	jsoniter.UnmarshalFromString(SwaggerInfo.ReadDoc(), &data)
	log.Println(jsonconv.ObjectToJsonIndent(data.Paths))
	//log.Println(SwaggerInfo.ReadDoc())
}

type Parameter struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	In          string `json:"in"`
	Required    bool   `json:"required"`
	Schema      struct {
		Ref string `json:"$ref"`
	} `json:"schema"`
}

type Responses struct {
	Field1 struct {
		Description string `json:"description"`
		Schema      struct {
			AllOf []struct {
				Ref        string `json:"$ref,omitempty"`
				Type       string `json:"type,omitempty"`
				Properties struct {
					Data struct {
						Ref string `json:"$ref"`
					} `json:"data"`
				} `json:"properties,omitempty"`
			} `json:"allOf"`
		} `json:"schema"`
	} `json:"200"`
}

type Api map[string]interface{}

//type Api struct {
//	A    map[string]interface{} `json:"omitempty"`
//	Post struct {
//		//Security []struct {
//		//	ApiKeyAuth []interface{} `json:"ApiKeyAuth"`
//		//} `json:"security"`
//		//Consumes []string `json:"consumes"`
//		//Produces []string `json:"produces"`
//		//Tags     []string `json:"tags"`
//		//Summary  string   `json:"summary"`
//		//Parameters []Parameter `json:"parameters"`
//		//Responses  Responses   `json:"responses"`
//	} `json:"post"`
//}

type T struct {
	Schemes []interface{} `json:"schemes"`
	Swagger string        `json:"swagger"`
	Info    struct {
		Description    string `json:"description"`
		Title          string `json:"title"`
		TermsOfService string `json:"termsOfService"`
		Contact        struct {
			Name  string `json:"name"`
			Url   string `json:"url"`
			Email string `json:"email"`
		} `json:"contact"`
		License struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"license"`
		Version string `json:"version"`
	} `json:"info"`
	Host                string         `json:"host"`
	BasePath            string         `json:"basePath"`
	Paths               map[string]Api `json:"paths"`
	Definitions         struct{}       `json:"definitions"`
	SecurityDefinitions struct {
		ApiKeyAuth struct {
			Type string `json:"type"`
			Name string `json:"name"`
			In   string `json:"in"`
		} `json:"ApiKeyAuth"`
		BasicAuth struct {
			Type string `json:"type"`
		} `json:"BasicAuth"`
	} `json:"securityDefinitions"`
}
