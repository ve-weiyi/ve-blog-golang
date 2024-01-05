package excel

import (
	"testing"
)

func TestExport(t *testing.T) {
	var err error
	e := NewExcelExporter()

	err = e.NewActiveSheet("Sheet1")
	t.Log(err)

	err = e.SetSheetTitle([]string{"姓名", "年龄", "性别"})
	t.Log(err)

	err = e.AddRowValue([]interface{}{"张三", 18, "男"})
	t.Log(err)

	err = e.AddRowValue([]interface{}{"李四", 18, "男"})
	t.Log(err)

	err = e.NewActiveSheet("Sheet2")
	t.Log(err)

	err = e.SetSheetTitle([]string{"姓名", "年龄", "性别"})
	t.Log(err)

	err = e.AddRowValue([]interface{}{"张三", 18, "男"})
	t.Log(err)

	err = e.AddRowValue([]interface{}{"李四", 18, "男"})
	t.Log(err)

	err = e.ExportToFile("test.xlsx")
	if err != nil {
		t.Log(err)
	}
}

func TestImport(t *testing.T) {
	var err error
	i := NewExcelImporter()

	err = i.OpenFile("test.xlsx")
	t.Log(err)

	err = i.SetActiveSheet("Sheet1")
	t.Log(err)

	row, err := i.GetRowValue(0)
	t.Log(row, err)

	row, err = i.GetRowValue(1)
	t.Log(row, err)

	row, err = i.GetRowValue(2)
	t.Log(row, err)

	err = i.SetActiveSheet("Sheet2")
	t.Log(err)

	row, err = i.GetRowValue(1)
	t.Log(row, err)

	row, err = i.GetRowValue(2)
	t.Log(row, err)
}
