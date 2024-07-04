package main

import (
	"github.com/scorify/scorify/pkg/cmd"
)

//go:generate go run -mod=mod github.com/scorify/generate

func main() {
	cmd.Execute()
}
