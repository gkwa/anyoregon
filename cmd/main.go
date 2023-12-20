package main

import (
	"os"

	"github.com/taylormonacelli/anyoregon"
)

func main() {
	code := anyoregon.Execute()
	os.Exit(code)
}
