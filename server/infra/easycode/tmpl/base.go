package tmpl

const NotEditMark = `
`

const Header = NotEditMark + `
package {{.Package}}

import(	
	{{range .ImportPkgPaths}}{{.}}` + "\n" + `{{end}}
)
`

const KeyController = "controller"
const KeyService = "service"
const KeyRouter = "router"
const KeyRepository = "repository"
const KeyModel = "model"
