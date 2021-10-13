package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var unicodeCmdExample = `
  Unicode code points may omit a prefix or include "U+"

  uc convert -u U+4EACU+4EAC U+4EAC
  output: 
  京京
  京

  uc convert 京
  output: 
  U+4EAC 
`

var (
	inputCodePointConvertCmd bool
	removeSpaceConvertCmd    bool
	convertCmd               = &cobra.Command{
		Use:     "convert [<args>]",
		Short:   "Convert string to/from Unicode code points",
		Example: unicodeCmdExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    runConvertCmd,
	}
)

func init() {
	convertCmd.Flags().BoolVarP(&removeSpaceConvertCmd, "remove-space", "", false, "removes space between Unicode code points")
	convertCmd.Flags().BoolVarP(&inputCodePointConvertCmd, "unicode", "u", false, "input is a sequence of Unicode code points")
	rootCmd.AddCommand(convertCmd)
}

func runConvertCmd(cmd *cobra.Command, args []string) error {
	if inputCodePointConvertCmd {
		for _, codePointSequence := range args {
			// Split returns empty string as first element because arg has prefix U+
			codePointSequence = strings.TrimSpace(codePointSequence)
			if codePointSequence == "" {
				return emptyStrError
			}
			if !strings.HasPrefix(codePointSequence, "U+") {
				return invalidArgType
			}
			for _, codePoint := range strings.Split(codePointSequence, "U+")[1:] {
				i, err := strconv.ParseUint(codePoint, 16, 32)
				if err != nil {
					return err
				}
				fmt.Printf("%c", i)
			}
			fmt.Printf("\n")
		}
	} else {
		for _, str := range args {
			// each arg is converted to a sequence of code points
			space := " "
			if removeSpaceConvertCmd {
				space = ""
			}
			for len(str) > 0 {
				r, size := utf8.DecodeRuneInString(str)
				fmt.Printf("%U%s", r, space)
				str = str[size:]
			}
			fmt.Printf("\n")
		}
	}
	return nil
}
