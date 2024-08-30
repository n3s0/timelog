package tl

import (
	"github.com/spf13/cobra"
)

var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "Handles project I/O",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(entryCmd)
}
