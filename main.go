package main

import (
	"github.com/dasuken/sshctl/cmd"
	"github.com/dasuken/sshctl/pkg/sshctl"
	"log"
	"os"
)

func main() {
	app := cmd.MakeApp("sshctl", "make ssh alias", "1.0.0")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(sshctl.ShowError(err.Error()))
	}
}