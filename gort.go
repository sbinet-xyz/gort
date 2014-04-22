package main

import (
	"github.com/codegangsta/cli"
	"github.com/hybridgroup/gort/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gort"
	app.Version = VERSION
	app.Usage = "RobotOps by your command"
	app.Commands = []cli.Command{
		commands.Scan(),
	}
	app.Run(os.Args)
}
