package command

import (
	"gopkg.in/urfave/cli.v1"
)

var (
	//Commands list of all defined commands
	Commands = make([]cli.Command, 0, 6)
)
