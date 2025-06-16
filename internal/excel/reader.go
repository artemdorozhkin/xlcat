package excel

import "github.com/xuri/excelize/v2"

func OpenExcelFile(filename string) (*excelize.File, error) {
	return excelize.OpenFile(filename)
}
