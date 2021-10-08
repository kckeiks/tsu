package cmd

import (
	"github.com/spf13/cobra"
)

var strCmd = &cobra.Command{
	Use:   "str",
	Short: "Translate a user given string",
}																																																																																																																																														

func init() {
	rootCmd.AddCommand(strCmd)
}
