package plate

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// 生成目录，文件名，模板内容，填充数据 data
type PlateMeta struct {
	Key          string
	AutoCodePath string //生成的代码路径  blog/api/article.go
	Replace      bool   //是否替换文件

	TemplateString string //模版文件内容
	//TemplatePath   string      //模版文件路径   tpl/api.go.tpl
	Data interface{} //填充内容
}

func (tempMeta *PlateMeta) CreateTempFile() error {
	if !tempMeta.Replace {
		if fileExist(tempMeta.AutoCodePath) {
			return errors.New(fmt.Sprintf("目标文件已存在:%s\n", tempMeta.AutoCodePath))
		}
	}

	//创建文件夹
	err := os.MkdirAll(filepath.Dir(tempMeta.AutoCodePath), 0755)
	if err != nil {
		return err
	}
	//创建.go文件
	f, err := os.Create(tempMeta.AutoCodePath)
	if err != nil {
		return err
	}
	defer f.Close()

	//解析模板
	temp, err := tempMeta.getTemplate()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = temp.Execute(&buf, tempMeta.Data)
	if err != nil {
		return err
	}

	if path.Ext(tempMeta.AutoCodePath) == ".go" {
		err = output(tempMeta.AutoCodePath, buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		os.WriteFile(tempMeta.AutoCodePath, buf.Bytes(), 0640)
	}

	return nil
}

func (tempMeta *PlateMeta) RollBack() error {
	if !tempMeta.Replace {
		return nil
	}

	err := deLFile(tempMeta.AutoCodePath)
	if err != nil {
		return err
	}
	return nil
}

func (tempMeta *PlateMeta) MoveTempFile(movePath string) error {
	//判断目标文件是否都可以移动
	if movePath != "" {
		if fileExist(movePath) {
			return errors.New(fmt.Sprintf("目标文件已存在:%s\n", movePath))
		}

		if err := fileMove(tempMeta.AutoCodePath, movePath); err != nil {
			return err
		}
		log.Println("file move success:", movePath)
	}
	return nil
}

func (tempMeta *PlateMeta) getTemplate() (*template.Template, error) {
	//if tempMeta.TemplatePath != "" {
	//	//解析模板
	//	temp, err := template.ParseFiles(tempMeta.TemplatePath)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return temp, nil
	//}

	if tempMeta.TemplateString != "" {
		//解析模板
		temp, err := template.New("temp").Parse(tempMeta.TemplateString)
		if err != nil {
			return nil, err
		}
		return temp, nil
	}

	return nil, errors.New("TemplateString or TemplatePath all null ")
}
