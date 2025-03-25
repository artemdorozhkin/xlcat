package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Options struct {
	Path  string
	Sheet string
	Rows  int
	Help  bool
}

func ParseArgs() (*Options, error) {
	var opts Options

	args := os.Args
	if len(args) < 2 {
		err := "ERR: File name is required"
		return nil, errors.New(err)
	}

	for i := 1; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--rows", "-r":
			if i+1 < len(args) {
				opts.Rows, _ = strconv.Atoi(args[i+1])
				args = slices.Delete(args, i, i+2)
				i -= 1
			}
		case "--sheet", "-s":
			if i+1 < len(args) {
				opts.Sheet = args[i+1]
				args = slices.Delete(args, i, i+2)
				i -= 1
			}
		case "--help", "-h":
			opts.Help = true
			return &opts, nil
		default:
			if !strings.HasPrefix(arg, "--") && !strings.HasPrefix(arg, "-") {
				opts.Path = arg
				args = slices.Delete(args, i, i+1)
				i -= 1
			}
		}
	}
	return &opts, nil
}

func PrintUsage() {
	var usage strings.Builder
	usage.WriteString("USAGE:\n")
	usage.WriteString("  xlcat <file.xls[x|m]> [options]\n\n")
	usage.WriteString("OPTIONS:\n")
	usage.WriteString("  -r | --rows\tRows to preview\n")
	usage.WriteString("  -s | --sheet\tSheet to preview\n")
	usage.WriteString("  --sheets\tShows onle sheet names\n")
	fmt.Println(usage.String())
}
