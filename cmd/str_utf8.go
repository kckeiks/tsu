package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

var strUTF8Cmd = &cobra.Command{
	Use:   "utf8 string...",
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
	result := make([]byte, 4)
	n := utf8.EncodeRune(result, r)
	for _, b := range result[:n] {
		if resultInHex {
			fmt.Printf("%s%x%s", prefix, b, space)
		} else {
			fmt.Printf("%d%s", b, space)
		}
	}
}
