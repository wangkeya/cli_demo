/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"cli_demo/cmd"

	_ "cli_demo/cmd/get"
	_ "cli_demo/cmd/version"
)

func main() {
	cmd.Execute()
}
