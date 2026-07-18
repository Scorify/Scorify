package main

import (
	"os"

	"github.com/scorify/scorify/pkg/cmd"
)

//go:generate go run -mod=mod github.com/scorify/generate

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
