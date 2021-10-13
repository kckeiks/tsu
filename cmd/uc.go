package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	invalidArgType = errors.New("invalid argument")
	emptyStrError  = errors.New("invalid empty arg")
)
var rootCmd = &cobra.Command{
	Use:          "uc",
	Short:        "Unicode Converter",
	SilenceUsage: true,
}

func Execute() error {
	return rootCmd.Execute()
}
