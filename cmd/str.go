package cmd

import (
	"github.com/spf13/cobra"
)

var strCmd = &cobra.Command{
	Use:   "str",
	Short: "Convert a user-given string",
}																																																																																																																																														

func init() {
	strCmd.PersistentFlags().BoolVarP(&removeSpace, "remove-space", "", false, "removes space in between each digit")
	rootCmd.AddCommand(strCmd)
}
