package main

import (
	"xlcat/internal/cli"
	"xlcat/internal/excel"
	path "xlcat/internal/fs"

	"fmt"
	"os"
	"strings"

	"slices"
)

func main() {
	opts, err := cli.ParseArgs()

	if err != nil {
		cli.PrintUsage()
		cli.PrintErrln(err)
		os.Exit(1)
	}

	if opts.Help {
		cli.PrintUsage()
		return
	}

	if len(opts.Path) == 0 {
		cli.PrintUsage()
		cli.PrintErrln("ERR: File name is required")
		os.Exit(1)
	}

	pathIsDir, err := path.IsDir(opts.Path)

	if err != nil {
		cli.PrintErrln(err)
		os.Exit(1)
	}

	if pathIsDir {
		excelFiles, err := excel.FindExcelFiles(opts.Path)

		if err != nil {
			cli.PrintErrln(err)
			os.Exit(1)
		}

		if len(excelFiles) == 0 {
			fmt.Printf("There are no supported excel files in directory: %s\n", opts.Path)
			fmt.Printf("Supported extensions: %s", strings.Join(excel.ExcelExtensions, ", "))
			return
		}

		excel.PrintFilesList(excelFiles)
		return
	}

	if len(opts.Sheet) == 0 && opts.Rows > 0 {
		cli.PrintErrln("ERR: You need to specify sheet name")
		os.Exit(1)
	}

	f, err := excel.OpenExcelFile(opts.Path)
	defer func() {
		if err := f.Close(); err != nil {
			cli.PrintErrln("ERR:", err)
			os.Exit(1)
		}
	}()

	if err != nil {
		cli.PrintErrf("ERR: Failed to open file '%s': %v", opts.Path, err)
		os.Exit(1)
	}

	sheets := f.GetSheetList()

	sheetSpecified := len(opts.Sheet) > 0
	if sheetSpecified && !slices.Contains(sheets, opts.Sheet) {
		cli.PrintErrf("Sheet '%s' missing in file '%s'", opts.Sheet, opts.Path)
		os.Exit(1)
	}

	if sheetSpecified {
		if opts.Rows == 0 {
			opts.Rows = 5 // default rows limit
		}
		rows, err := f.GetRows(opts.Sheet)
		if err != nil {
			cli.PrintErrf("Failed to read rows of sheet[%s]: %v", opts.Sheet, err)
			os.Exit(1)
		}

		excel.PrintSheetInfo(rows, opts)
		return
	}

	if len(sheets) > 0 {
		if err := excel.PrintSheetsInfo(sheets, f, opts); err != nil {
			cli.PrintErrln(err)
			os.Exit(1)
		}
	}
}
