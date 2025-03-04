package model

import (
	"fmt"
	"os"
	"path"

	"github.com/ve-weiyi/ve-blog-golang/tools/cmd/model/helper"

	"github.com/ve-weiyi/ve-blog-golang/kit/quickstart/invent"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/files"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

type modelConfig struct {
	SqlFile string
	TplFile string
	OutPath string
	NameAs  string
}

func generateModel(tables []*helper.Table, conf modelConfig) error {
	t := files.ToAbs(conf.TplFile)
	o := conf.OutPath
	n := conf.NameAs

	var metas []invent.TemplateMeta

	tpl, err := os.ReadFile(t)
	if err != nil {
		return err
	}

	for _, table := range tables {
		data := helper.ConvertTableToData(table)

		meta := invent.TemplateMeta{
			Key:            "",
			Mode:           invent.ModeCreateOrReplace,
			CodeOutPath:    path.Join(o, fmt.Sprintf(n, table.Name)),
			TemplateString: string(tpl),
			FunMap: map[string]any{
				"funcFieldsKey": func(fs []*helper.ModelField) string {
					var name string
					for _, ff := range fs {
						name += ff.Name
					}
					return name
				},
				"funcFieldsKeyVar": func(fs []*helper.ModelField) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						tp := ff.Type
						if name != "" {
							name += ", "
						}
						name += fmt.Sprintf("%s %s", v, tp)
					}
					return name
				},
				"funcFieldsKeyCond": func(fs []*helper.ModelField) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						if name != "" {
							name += " and "
						}
						name += fmt.Sprintf("`%s` = ?", v)
					}
					return name
				},
				"funcFieldsKeyCondVar": func(fs []*helper.ModelField) string {
					var name string
					for _, ff := range fs {
						v := jsonconv.Case2Snake(ff.Name)
						if name != "" {
							name += ", "
						}
						name += v
					}
					return name
				},
			},
			Data: data,
		}
		metas = append(metas, meta)
	}

	for _, m := range metas {
		err := m.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
