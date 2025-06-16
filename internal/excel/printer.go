package excel

import (
	"xlcat/internal/cli"

	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/xuri/excelize/v2"
)

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

func PrintSheetInfo(rows [][]string, opts *cli.Options) {
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

func PrintSheetsInfo(sheets []string, f *excelize.File, opts *cli.Options) error {
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
