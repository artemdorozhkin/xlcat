package excel

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var ExcelExtensions []string = []string{
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

		if slices.Contains(ExcelExtensions, itemExt) {
			excelFiles = append(excelFiles, item.Name())
		}
	}

	return excelFiles, nil
}
