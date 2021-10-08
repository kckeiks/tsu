package cmd

import (
	"fmt"
	"bytes"
	"strconv"
	"github.com/spf13/cobra"
	"unicode/utf8"
)

var cpUTF8Cmd = &cobra.Command{
	Use:   "utf8 code_point...",
	Short: "Convert a sequence of unicode code points to a sequence of UTF-8 encoded values",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := bytes.NewBuffer([]byte{})
		// utf8 uses up to 4 bytes
		result.Grow(len(args)*4)
		buf := [4]byte{}
		for _, str := range args {
			// input is sequence of Unicode code points
			if str[:2] == "U+" {
				str = str[2:]
			}
			codepoint, _ := strconv.ParseUint(str, 16, 32)
			n := utf8.EncodeRune(buf[:], rune(codepoint))
			result.Write(buf[:n])																												
		}
		cpUTF8CmdPrint(result.Bytes())	
	},																																											
}

func init() {
	cpUTF8Cmd.Flags().BoolVarP(&resultInHex, "hex", "x", false, "return result in hex")
	cpUTF8Cmd.Flags().StringVarP(&prefix, "prefix", "", "", "add prefix to every two hex digits")
	cpCmd.AddCommand(cpUTF8Cmd)
}

func cpUTF8CmdPrint(buff []byte) {
	space := " "
	if removeSpace {
		space = ""
	}
	for _, b := range buff {
		if resultInHex {
			fmt.Printf("%s%X%s", prefix, b, space)
		} else {
			fmt.Printf("%d%s", b, space)
		}
	}
	fmt.Printf("\n")
}