package excel

import (
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelExporter interface {
	// 新建一个sheet页
	NewActiveSheet(sheetName string) error
	// 设置sheet页的标题
	SetSheetTitle(title []string) error
	// 设置sheet页的某一行的值
	SetRowValue(index int, rowData []interface{}) error
	// 在末尾添加一列
	AddRowValue(rowData []interface{}) error
	// 导出为文件
	ExportToFile(fileName string) error
	// 导出到指定位置
	Write(w io.Writer) error
}

type ExcelExportImpl struct {
	activeSheetName string
	File            *excelize.File
}

func NewExcelExporter() *ExcelExportImpl {
	f := excelize.NewFile()
	return &ExcelExportImpl{
		File: f,
	}
}

func (s *ExcelExportImpl) NewActiveSheet(sheetName string) error {

	s.activeSheetName = sheetName

	// Create a new sheet.
	index, err := s.File.NewSheet(s.activeSheetName)
	if err != nil {
		return err
	}
	// Set active sheet of the workbook.
	s.File.SetActiveSheet(index)

	return nil
}

func (s *ExcelExportImpl) SetSheetTitle(title []string) error {
	for i, v := range title {
		err := s.File.SetCellValue(s.activeSheetName, s.point(i, 1), v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ExcelExportImpl) SetRowValue(index int, rowData []interface{}) error {
	for i, v := range rowData {
		err := s.File.SetCellValue(s.activeSheetName, s.point(i, index), v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ExcelExportImpl) AddRowValue(rowData []interface{}) error {
	count := s.RowCount()
	return s.SetRowValue(count+1, rowData)
}

func (s *ExcelExportImpl) ExportToFile(fileName string) error {
	return s.File.SaveAs(fileName)
}

func (s *ExcelExportImpl) Write(w io.Writer) error {
	return s.File.Write(w)
}

func (s *ExcelExportImpl) RowCount() int {
	rows, err := s.File.GetRows(s.activeSheetName)
	if err != nil {
		return 0
	}
	return len(rows)
}

// 1,2,3 -> A1,B1,C1
func (s *ExcelExportImpl) point(charaID int, i int) string {
	row := strconv.Itoa(i)
	return s.charColIndex(charaID) + row
}

// 1,2,3 -> A,B,C
func (s *ExcelExportImpl) charColIndex(i int) string {
	if i > 26 {
		return "nil"
	}
	return string(rune(65 + i))
}
