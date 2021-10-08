package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

var resultInHex bool
var prefix string
var strUTF8Cmd = &cobra.Command{
	Use:   "utf8",
	Short: "Convert a string to a sequence of UTF-8 encoded values",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, str := range args {
			for len(str) > 0 {
				r, size := utf8.DecodeRuneInString(str)
				strUTF8CmdPrint(r)
				str = str[size:]
			}
			fmt.Printf("\n")																																		
		}
	},																																											
}

func init() {
	strUTF8Cmd.Flags().BoolVarP(&resultInHex, "hex", "x", false, "return result in hex")
	strUTF8Cmd.Flags().StringVarP(&prefix, "prefix", "", "", "add prefix to every two hex digits")
	strCmd.AddCommand(strUTF8Cmd)
}

func strUTF8CmdPrint(r rune) {
	space := " "
	if removeSpace {
		space = ""
	}
	if resultInHex {
		fmt.Printf("%s%x%s", prefix, r, space)
	} else {
		fmt.Printf("%d%s", r, space)
	}
}
