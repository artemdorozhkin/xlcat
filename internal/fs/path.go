package fs

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
