package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

var removeSpaceConvertCmd bool
var inputIsStrConvertCmd bool
var unicodeCmdExample = 
`
  Unicode code points may omit a prefix or include "U+"

  uc unicode U+4EAC
  output: 
  京

  uc unicode 4EAC
  output: 
  京 

  uc unicode -s 京
  output: 
  U+4EAC 
`

var convertCmd = &cobra.Command{
	Use:   "convert {code_point | -s string} ...",
	Short: "Convert string to/from Unicode code points",
	Example: unicodeCmdExample,
	Args: cobra.MinimumNArgs(1),
	RunE: runConvertCmd,																																												
}

func init() {
	convertCmd.Flags().BoolVarP(&removeSpaceConvertCmd, "remove-space", "", false, "removes space in between each digit")
	convertCmd.Flags().BoolVarP(&inputIsStrConvertCmd, "str", "s", false, "input is string")
	rootCmd.AddCommand(convertCmd)
}

func runConvertCmd(cmd *cobra.Command, args []string) error {
	for _, str := range args {
		if !inputIsStrConvertCmd {
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
	if inputIsStrConvertCmd {
		fmt.Printf("%U%s", r, space)
	} else {
		fmt.Printf("%c%s", r, space)
	}
}
