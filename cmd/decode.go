package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
	"bytes"
	"strconv"
)

var outputCodePoint bool
var inputHex bool
var decodeCmd = &cobra.Command{
	Use:   "decode [<args>]",
	Short: "Decode UTF-8 sequence to string",
	Long: "Convert a sequence of UTF-8 encoded values to a string",
	Args: cobra.MinimumNArgs(1),
	RunE: handleDecodeCmd,																																											
}

func init() {
	decodeCmd.Flags().BoolVarP(&outputCodePoint, "output-unicode", "u", false, "output will be a sequence of Unicode code points")
	decodeCmd.Flags().BoolVarP(&inputHex, "input-hex", "x", false, "input will be sequence of two hex digits")
	rootCmd.AddCommand(decodeCmd)
}

func handleDecodeCmd(cmd *cobra.Command, args []string) error {
	var b bytes.Buffer
	base := 10
	if inputHex {
		base = 16
	}
	for _, value := range args {
		codepoint,err := strconv.ParseUint(value, base, 8)
		if err != nil {
			return err
		}
		b.WriteByte(byte(codepoint))
	}
	if outputCodePoint {
		for rawb := b.Bytes(); len(rawb) > 0; {
			r, size := utf8.DecodeRune(rawb)
			fmt.Printf("%U ", r)
			rawb = rawb[size:]
		}
		fmt.Printf("\n")
	} else {
		fmt.Println(string(b.Bytes()))
	}
	return nil
}