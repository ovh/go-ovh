package common

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Check checks e and panic if not nil
func Check(err error) {
	if err != nil {
		if Verbose {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// Exit func display an error message on stderr and exit 1
func Exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

// WrongUsage display a wrong usage error, shows the help and exit 1
func WrongUsage(cmd *cobra.Command) {
	fmt.Fprintln(os.Stderr, "Error: Wrong usage")
	cmd.Help()
	os.Exit(1)
}
