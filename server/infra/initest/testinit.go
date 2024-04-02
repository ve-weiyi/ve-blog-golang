package initest

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

// @title						Swagger Example API
// @version					0.0.1
// @description				This is a sample Server pets
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-token
// @BasePath					/
func Init(configPath ...string) {

}
