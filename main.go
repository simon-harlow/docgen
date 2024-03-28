//go:generate go run generate-asset.go

package main

import (
	"github.com/simon-harlow/docgen/cmd"
)

func main() {
	cmd.Execute()
}
