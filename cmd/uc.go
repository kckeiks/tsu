package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "uc",
	Short:        "Unicode Converter",
	SilenceUsage: true,
}

func Execute() error {
	return rootCmd.Execute()
}
