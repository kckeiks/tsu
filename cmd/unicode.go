package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

var inputIsString bool
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

var unicodeCmd = &cobra.Command{
	Use:   "unicode {code_point | -s string} ...",
	Short: "Convert string to/from Unicode code points",
	Example: unicodeCmdExample,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, str := range args {
			if !inputIsString {
				// input is sequence of Unicode code points
				if str[:2] == "U+" {
					str = str[2:]
				}
				codepoint, _ := strconv.ParseInt(str, 16, 32)
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
	},																																												
}

func init() {
	unicodeCmd.PersistentFlags().BoolVarP(&removeSpace, "remove-space", "", false, "removes space in between each digit")
	unicodeCmd.Flags().BoolVarP(&inputIsString, "str", "s", false, "input is string")
	rootCmd.AddCommand(unicodeCmd)
}

func unicodeCmdPrint(r rune) {
	space := " "
	if removeSpace {
		space = ""
	}
	if inputIsString {
		fmt.Printf("%U%s", r, space)
	} else {
		fmt.Printf("%c%s", r, space)
	}
}
