package main

import (
	"errors"
	"fmt"
	"os"
)

func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		description := fmt.Sprintf("ERR: Argument must be a path: %s", path)
		return false, errors.New(description)
	}
	return fi.IsDir(), nil
}

func PrintErrln(a ...any) {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	a = append([]any{Red + fmt.Sprint(a[0])}, a[:len(a)-1]...)
	a = append(a, Reset)
	fmt.Fprintln(os.Stderr, a...)
}

func PrintErrf(format string, a ...any) {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	a = append([]any{Red + fmt.Sprint(a[0])}, a[:len(a)-1]...)
	a = append(a, Reset)
	fmt.Fprintf(os.Stderr, format, a...)
}
