package cli

import (
	"fmt"
	"os"
)

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
