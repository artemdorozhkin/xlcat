package cli

import (
	"fmt"
	"strings"
)

func PrintUsage() {
	var usage strings.Builder
	usage.WriteString("USAGE:\n")
	usage.WriteString("  xlcat <file.xls[x|m]> [options]\n\n")
	usage.WriteString("OPTIONS:\n")
	usage.WriteString("  -r | --rows\tRows to preview\n")
	usage.WriteString("  -s | --sheet\tSheet to preview\n")
	fmt.Println(usage.String())
}
