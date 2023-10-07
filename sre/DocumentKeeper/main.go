/*
Copyright © 2023 ritaCanavarro
*/
package main

import (
	"DocumentKeeper/cmd"
	"os"
)

func init() {
	os.Setenv("GOPATH", "/home/rita.canavarro/Documents/Community/OSC/Study/hiring-assignments/sre/DocumentKeeper")
}

func main() {
	cmd.Execute()
}
