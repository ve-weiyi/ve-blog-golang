package tmpl

const NotEditMark = `
`

const Header = NotEditMark + `
package {{.Package}}

import(	
	{{range .ImportPkgPaths}}{{.}}` + "\n" + `{{end}}
)
`

const KeyApi = "api"
const KeyRouter = "router"
const KeyController = "controller"
const KeyService = "service"
const KeyRepository = "repository"
const KeyModel = "model"
