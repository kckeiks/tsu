package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"unicode/utf8"
)

var unicodeCmdExample = `
  Unicode code points may omit a prefix or include "U+"

  uc convert -u U+4EAC
  output: 
  京

  uc convert -u 4EAC
  output: 
  京 

  uc convert 京
  output: 
  U+4EAC 
`

var (
	inputCodePointConvertCmd bool
	convertCmd               = &cobra.Command{
		Use:     "convert [<args>]",
		Short:   "Convert string to/from Unicode code points",
		Example: unicodeCmdExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    runConvertCmd,
	}
)

func init() {
	convertCmd.Flags().BoolVarP(&inputCodePointConvertCmd, "unicode", "u", false, "input is a sequence of Unicode code points")
	rootCmd.AddCommand(convertCmd)
}

func runConvertCmd(cmd *cobra.Command, args []string) error {
	for _, str := range args {
		if inputCodePointConvertCmd {
			// TODO: Check length
			if strings.HasPrefix(str, "U+") {
				str = str[2:]
			}
			codepoint, err := strconv.ParseInt(str, 16, 32)
			if err != nil {
				return err
			}
			fmt.Printf("%c", codepoint)
		} else {
			// input is string
			for len(str) > 0 {
				r, size := utf8.DecodeRuneInString(str)
				fmt.Printf("%U ", r)
				str = str[size:]
			}
		}
	}
	fmt.Printf("\n")
	return nil
}
