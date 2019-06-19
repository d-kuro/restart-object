package main

import (
	"os"

	"github.com/d-kuro/restart-object/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
