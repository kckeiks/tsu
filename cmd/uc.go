package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:          "uc",
		Short:        "Unicode Converter",
		SilenceUsage: true,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
