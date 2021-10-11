package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

var removeSpaceConvertCmd bool
var inputCodePointConvertCmd bool
var unicodeCmdExample = 
`
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

var convertCmd = &cobra.Command{
	Use:   "convert {-u code_point | string} ...",
	Short: "Convert string to/from Unicode code points",
	Example: unicodeCmdExample,
	Args: cobra.MinimumNArgs(1),
	RunE: runConvertCmd,																																												
}

func init() {
	convertCmd.Flags().BoolVarP(&removeSpaceConvertCmd, "remove-space", "", false, "removes space in between each digit")
	convertCmd.Flags().BoolVarP(&inputCodePointConvertCmd, "unicode", "u", false, "input is a sequence of Unicode code points")
	rootCmd.AddCommand(convertCmd)
}

func runConvertCmd(cmd *cobra.Command, args []string) error {
	for _, str := range args {
		if inputCodePointConvertCmd {
			// input is sequence of Unicode code points
			if str[:2] == "U+" {
				str = str[2:]
			}
			codepoint, err := strconv.ParseInt(str, 16, 32)
			if err != nil {
				return err
			}
			fmt.Println(string(codepoint))
		} else {
			// input is string
			for len(str) > 0 {
				r, size := utf8.DecodeRuneInString(str)
				unicodeCmdPrint(r)
				str = str[size:]
			}
			fmt.Printf("\n")	
		}																																	
	}
	return nil
}

func unicodeCmdPrint(r rune) {
	space := " "
	if removeSpaceConvertCmd {
		space = ""
	}
	if inputCodePointConvertCmd {
		fmt.Printf("%c%s", r, space)
	} else {
		fmt.Printf("%U%s", r, space)
	}
}
