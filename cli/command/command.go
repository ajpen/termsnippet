package command

import (
	Cli "termsnippet/cli"

	"gopkg.in/urfave/cli.v1"
)

func InstallCommand(c cli.Command) {
	Cli.App.Commands = append(Cli.App.Commands, c)
}
