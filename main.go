package main

import (
	"github.com/kckeiks/tsu/cmd"
  "os"
)
  
func main() {
  if err := cmd.Execute(); err != nil {
    os.Exit(1)
  }
}