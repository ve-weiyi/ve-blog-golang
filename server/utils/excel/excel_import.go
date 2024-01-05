package excel

import (
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

type ExcelImporter interface {
	OpenFile(fileName string) error
	OpenReader(r io.Reader) error
	SetActiveSheet(sheetName string) error
	GetAllRowValue() ([][]string, error)
	GetRowValue(rowIndex int) ([]string, error)
	GetAllRowCount() int
}

type ExcelImportImpl struct {
	activeSheetName string
	File            *excelize.File
}

func NewExcelImporter() *ExcelImportImpl {
	f := excelize.NewFile()
	return &ExcelImportImpl{
		File: f,
	}
}

func (s *ExcelImportImpl) OpenFile(fileName string) error {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return err
	}

	s.File = f
	return nil
}

func (s *ExcelImportImpl) OpenReader(r io.Reader) error {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	s.File = f
	return nil
}

func (s *ExcelImportImpl) SetActiveSheet(sheetName string) error {

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

func (s *ExcelImportImpl) GetRowValue(rowIndex int) ([]string, error) {
	rows, err := s.File.GetRows(s.activeSheetName)
	if err != nil {
		return nil, err
	}

	if rowIndex >= len(rows) {
		return nil, fmt.Errorf("rowIndex out of range")
	}

	return rows[rowIndex], nil
}

func (s *ExcelImportImpl) GetAllRowValue() ([][]string, error) {
	rows, err := s.File.GetRows(s.activeSheetName)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *ExcelImportImpl) GetAllRowCount() int {
	rows, err := s.File.GetRows(s.activeSheetName)
	if err != nil {
		return 0
	}

	return len(rows)
}
