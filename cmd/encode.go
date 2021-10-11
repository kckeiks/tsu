package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
	"bytes"
	"strconv"
)

var removeSpaceEncodeCmd bool
var resultInHexEncodeCmd bool
var prefixEncodeCmd string
var inputCodePointEncodeCmd bool
var encodeCmd = &cobra.Command{
	Use:   "encode [<args>]",
	Short: "Encode string using UTF-8",
	Long: "Convert a string to a sequence of UTF-8 encoded values",
	Args: cobra.MinimumNArgs(1),
	RunE: runEncodeCmd,																																											
}

func init() {
	encodeCmd.Flags().BoolVarP(&removeSpaceEncodeCmd, "remove-space", "", false, "removes space in between each digit")
	encodeCmd.Flags().BoolVarP(&inputCodePointEncodeCmd, "unicode", "u", false, "input is a sequence of Unicode code points")
	encodeCmd.Flags().BoolVarP(&resultInHexEncodeCmd, "hex", "x", false, "return result in hex")
	encodeCmd.Flags().StringVarP(&prefixEncodeCmd, "prefix", "", "", "add prefix to every two hex digits")
	rootCmd.AddCommand(encodeCmd)
}

func runEncodeCmd(cmd *cobra.Command, args []string) error {
	if inputCodePointEncodeCmd {
		return handleCodePointInput(args)
	} else {
		return handleStringInput(args)
	}
}

func handleStringInput(args []string) error {
	for _, str := range args {
		for len(str) > 0 {
			r, size := utf8.DecodeRuneInString(str)
			strUTF8CmdPrint(r)
			str = str[size:]
		}
		fmt.Printf("\n")																																		
	}
	return nil
}

func handleCodePointInput(args []string) error {
	result := bytes.NewBuffer([]byte{})
	// utf8 uses up to 4 bytes
	result.Grow(len(args)*4)
	buf := [4]byte{}
	for _, str := range args {
		// input is sequence of Unicode code points
		if str[:2] == "U+" {
			str = str[2:]
		}
		codepoint, err := strconv.ParseUint(str, 16, 32)
		if err != nil {
			return err
		}
		n := utf8.EncodeRune(buf[:], rune(codepoint))
		result.Write(buf[:n])																												
	}
	cpUTF8CmdPrint(result.Bytes())
	return nil
}

func cpUTF8CmdPrint(buff []byte) {
	space := " "
	if removeSpaceEncodeCmd {
		space = ""
	}
	for _, b := range buff {
		if resultInHexEncodeCmd {
			fmt.Printf("%s%X%s", prefixEncodeCmd, b, space)
		} else {
			fmt.Printf("%d%s", b, space)
		}
	}
	fmt.Printf("\n")
}

func strUTF8CmdPrint(r rune) {
	space := " "
	if removeSpaceEncodeCmd {
		space = ""
	}
	result := make([]byte, 4)
	n := utf8.EncodeRune(result, r)
	for _, b := range result[:n] {
		if resultInHexEncodeCmd {
			fmt.Printf("%s%x%s", prefixEncodeCmd, b, space)
		} else {
			fmt.Printf("%d%s", b, space)
		}
	}
}
