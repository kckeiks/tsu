package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

var inputIsString bool
var unicodeCmd = &cobra.Command{
	Use:   "unicode",
	Short: "Convert string to/from Unicode code points",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, str := range args {
			for len(str) > 0 {
				r, size := utf8.DecodeRuneInString(str)
				unicodeCmdPrint(r)
				str = str[size:]
			}
			fmt.Printf("\n")																																		
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
