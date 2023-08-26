package util

import (
	"fmt"
	"os"
)

func InvalidArg(message string) {
	fmt.Fprintln(os.Stdout, message)
	os.Exit(1)
}

func ExecutionError(message string, err error) {
	fmt.Fprintln(os.Stderr, message, err)
	os.Exit(1)
}
