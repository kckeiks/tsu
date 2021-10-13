package main

import (
	"os"

	"github.com/kckeiks/tsu/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
