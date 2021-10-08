package cmd

import (
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Convert a Unicode code point to UTF encoded representation",
}																																																																																																																																														

func init() {
	cpCmd.PersistentFlags().BoolVarP(&removeSpace, "remove-space", "", false, "removes space in between each digit")
	rootCmd.AddCommand(cpCmd)
}

