package parsex

import (
	"os"
	"path"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func ParseAPI(filename string) (out *spec.ApiSpec, err error) {
	if path.IsAbs(filename) {
		return parser.Parse(filename)
	}

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f := path.Join(dir, filename)
	return parser.Parse(f)
}
