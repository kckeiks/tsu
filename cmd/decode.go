package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"unicode/utf8"
)

var decodeCmdExample = `
  uc decode 104 101 108 108 111 32 119 111 114 108 100
  output:
  hello world
`

var (
	outputCodePointDecodeCmd bool
	inputHexDecodeCmd        bool
	decodeCmd                = &cobra.Command{
		Use:     "decode [<args>]",
		Short:   "Decode UTF-8 encoded sequence",
		Long:    "Convert a sequence of UTF-8 encoded values to a string",
		Args:    cobra.MinimumNArgs(1),
		Example: decodeCmdExample,
		RunE:    handleDecodeCmd,
	}
)

func init() {
	decodeCmd.Flags().BoolVarP(&outputCodePointDecodeCmd, "output-unicode", "u", false, "output will be a sequence of Unicode code points")
	decodeCmd.Flags().BoolVarP(&inputHexDecodeCmd, "input-hex", "x", false, "input will be sequence of two hex digits")
	rootCmd.AddCommand(decodeCmd)
}

func handleDecodeCmd(cmd *cobra.Command, args []string) error {
	var buf bytes.Buffer
	base := 10
	if inputHexDecodeCmd {
		base = 16
	}
	for _, value := range args {
		codepoint, err := strconv.ParseUint(value, base, 8)
		if err != nil {
			return err
		}
		buf.WriteByte(byte(codepoint))
	}
	if outputCodePointDecodeCmd {
		for bytesBuff := buf.Bytes(); len(bytesBuff) > 0; {
			r, size := utf8.DecodeRune(bytesBuff)
			fmt.Printf("%U ", r)
			bytesBuff = bytesBuff[size:]
		}
		fmt.Printf("\n")
	} else {
		fmt.Println(string(buf.Bytes()))
	}
	return nil
}
