package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

// subcommand:
// tsu utf8 unicode [\u782 \u123 ....]  // defaults: translate and return string
// tsu utf8 unicode -i [\u782 \u123 ....] // encode - return sequence in int
// tsu utf8 unicode -x [\u782 \u123 ....]  // encode - return sequence in hex

// subcommand:
// tsu utf8 hex 0x01 0x02 .... sequence of hex  // encode - defaults: return string
// tsu utf8 hex -u 0x01 0x02 .... sequence of hex  // encode - return sequence of unicode code popints

// Other command:
// tsu unicode [string.....]// translate and return sequence of unicode code popints
var resultInHex bool
var useHexPrefix bool

var utf8Cmd = &cobra.Command{
	Use:   "utf8",
	Short: "Translate string to/from sequence of UTF-8 encoded values",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, str := range args {
			for len(str) > 0 {
				r, size := utf8.DecodeRuneInString(str)
				printUTF8Value(r)
				str = str[size:]
			}
			fmt.Printf("\n")																																		
		}
	},																																												
}																																																																																																																																														

func printUTF8Value(r rune) {
	if resultInHex {
		prefix := ""
		if useHexPrefix {
			prefix = "0x"
		}
		fmt.Printf("%s%x ", prefix, r)
	} else {
		fmt.Printf("%d ", r)
	}
}

func init() {
	rootCmd.AddCommand(utf8Cmd)
	utf8Cmd.Flags().BoolVarP(&resultInHex, "hex", "x", false, "return result in hex")
	utf8Cmd.Flags().BoolVarP(&useHexPrefix, "hex-prefix", "p", false, "return result in hex with prefix \"0x\"")
}
