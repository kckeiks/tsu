package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
	"bytes"
	"strconv"
)

var outputIsCodePoint bool
var inputIsHex bool
var decodeCmd = &cobra.Command{
	Use:   "decode [<args>]",
	Short: "Decode UTF-8 sequence to string",
	Long: "Convert a sequence of UTF-8 encoded values to a string",
	Args: cobra.MinimumNArgs(1),
	Run: handle,																																											
}

func init() {
	decodeCmd.Flags().BoolVarP(&outputIsCodePoint, "output-unicode", "u", false, "output will be a sequence of Unicode code points")
	decodeCmd.Flags().BoolVarP(&inputIsHex, "input-hex", "x", false, "input will be sequence of two hex digits")
	rootCmd.AddCommand(decodeCmd)
}

func handle(cmd *cobra.Command, args []string) {
	var b bytes.Buffer
	base := 10
	if inputIsHex {
		base = 16
	}
	for _, value := range args {
		codepoint,err := strconv.ParseUint(value, base, 8)
		if err != nil {
			fmt.Println("Value is too big")
		}
		// since input is sequence of bytes
		err = b.WriteByte(byte(codepoint))
		if err != nil {
			fmt.Println("Something happened")
		}
	}
	if outputIsCodePoint {
		for rawb := b.Bytes(); len(rawb) > 0; {
			r, size := utf8.DecodeRune(rawb)
			fmt.Printf("%U ", r)
			rawb = rawb[size:]
		}
		fmt.Printf("\n")
	} else {
		fmt.Println(string(b.Bytes()))
	}
}