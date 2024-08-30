package tl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "timelog",
	Short: "Command line time tracker",
	Long:  "Command line time tracker for tasks and projects",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "timelog error: %+v\n", err)
		os.Exit(1)
	}
}
