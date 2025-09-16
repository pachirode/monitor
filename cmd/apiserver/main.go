package main

import (
	"os"

	"github.com/pachirode/monitor/cmd/apiserver/app"
)

func main() {
	command := app.NewMonitorCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
