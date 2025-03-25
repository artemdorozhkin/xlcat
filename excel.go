package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/xuri/excelize/v2"
)

var excelExtensions []string = []string{
	".XLAM",
	".XLSM",
	".XLSX",
	".XLTM",
	".XLTX",
}

func FindExcelFiles(dir string) ([]string, error) {
	items, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("ERR: Cant read directory: %s", dir)
	}

	excelFiles := []string{}

	for _, item := range items {
		if item.IsDir() {
			continue
		}

		itemExt := strings.ToUpper(filepath.Ext((item.Name())))

		if slices.Contains(excelExtensions, itemExt) {
			excelFiles = append(excelFiles, item.Name())
		}
	}

	return excelFiles, nil
}

func PrintFilesList(excelFiles []string) {
	excelFilesCount := len(excelFiles)

	if excelFilesCount == 1 {
		fmt.Printf("Found %d excel file:\n", len(excelFiles))
	} else {
		fmt.Printf("Found %d excel files:\n", len(excelFiles))
	}
	for _, p := range excelFiles {
		fmt.Printf("  %s\n", p)
	}
}

func OpenExcelFile(filename string) (*excelize.File, error) {
	return excelize.OpenFile(filename)
}

func PrintSheetInfo(rows [][]string, opts *Options) {
	table := tablewriter.NewWriter(os.Stdout)
	var sb strings.Builder
	sb.WriteString("First ")
	sb.WriteString(strconv.Itoa(opts.Rows))
	sb.WriteString(" rows of the sheet '")
	sb.WriteString(opts.Sheet)
	sb.WriteString("'")
	table.SetCaption(true, sb.String())
	table.SetColWidth(10)
	colsCount := 0
	for i, row := range rows {
		if len(row) > colsCount {
			colsCount = len(row)
		}
		if i >= opts.Rows {
			footer := make([]string, colsCount)
			footer[0] = "---"
			table.SetFooter(footer)
			break
		}
		table.Append(row)
	}
	table.Render()
}

func PrintSheetsInfo(sheets []string, f *excelize.File, opts *Options) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sheet Name", "Rows Count", "Cols Count"})
	for i := range len(sheets) {
		rows, err := f.GetRows(sheets[i])
		if err != nil {
			return fmt.Errorf("ERR: failed to read rows of sheet[%s]: %v", sheets[i], err)
		}

		rowsCount := len(rows)
		colsCount := 0
		if len(rows) > 0 {
			for _, row := range rows {
				if len(row) > colsCount {
					colsCount = len(row)
				}
			}
		}
		table.Append([]string{sheets[i], strconv.Itoa(rowsCount), strconv.Itoa(colsCount)})
	}
	table.Render()
	return nil
}
