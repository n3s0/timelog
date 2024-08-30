package tl

import (
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"proj"},
	Short:   "Handles project I/O",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
