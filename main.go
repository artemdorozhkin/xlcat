package main

import (
	"fmt"
	"os"
	"strings"

	"slices"
)

func main() {
	opts, err := ParseArgs()

	if err != nil {
		PrintUsage()
		PrintErrln(err)
		os.Exit(1)
	}

	if opts.Help {
		PrintUsage()
		return
	}

	if len(opts.Path) == 0 {
		PrintUsage()
		PrintErrln("ERR: File name is required")
		os.Exit(1)
	}

	pathIsDir, err := IsDir(opts.Path)

	if err != nil {
		PrintErrln(err)
		os.Exit(1)
	}

	if pathIsDir {
		excelFiles, err := FindExcelFiles(opts.Path)

		if err != nil {
			PrintErrln(err)
			os.Exit(1)
		}

		if len(excelFiles) == 0 {
			fmt.Printf("There are no supported excel files in directory: %s\n", opts.Path)
			fmt.Printf("Supported extensions: %s", strings.Join(excelExtensions, ", "))
			return
		}

		PrintFilesList(excelFiles)
		return
	}

	if len(opts.Sheet) == 0 && opts.Rows > 0 {
		PrintErrln("ERR: You need to specify sheet name")
		os.Exit(1)
	}

	f, err := OpenExcelFile(opts.Path)
	defer func() {
		if err := f.Close(); err != nil {
			PrintErrln("ERR:", err)
			os.Exit(1)
		}
	}()

	if err != nil {
		PrintErrf("ERR: Failed to open file '%s': %v", opts.Path, err)
		os.Exit(1)
	}

	sheets := f.GetSheetList()

	sheetSpecified := len(opts.Sheet) > 0
	if sheetSpecified && !slices.Contains(sheets, opts.Sheet) {
		PrintErrf("Sheet '%s' missing in file '%s'", opts.Sheet, opts.Path)
		os.Exit(1)
	}

	if sheetSpecified {
		if opts.Rows == 0 {
			opts.Rows = 5 // default rows limit
		}
		rows, err := f.GetRows(opts.Sheet)
		if err != nil {
			PrintErrf("Failed to read rows of sheet[%s]: %v", opts.Sheet, err)
			os.Exit(1)
		}

		PrintSheetInfo(rows, opts)
		return
	}

	if len(sheets) > 0 {
		if err := PrintSheetsInfo(sheets, f, opts); err != nil {
			PrintErrln(err)
			os.Exit(1)
		}
	}
}
