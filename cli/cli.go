package cli

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

// Errors
var (
	// ArgumentMissingf returns an error message indicating a missing argument.
	// argument: argument name
	ErrArgumentMissing = func(argument string) error {
		return fmt.Errorf("Missing required argument: %s", argument)
	}

	ErrDoesNotExist = func(lookingFor string) error {
		return fmt.Errorf("Not Found: %s", lookingFor)
	}
)

func init() {
	App.Name = "TermSnippet"
	App.Usage = "A command line based reusable code manager"
	App.Description = "TermSnippet is a command line based code snippet database for storing and accessing resuable code"
	App.Author = "Anfernee Jervis"
	App.Version = "0.0.1"
}

var (
	App = cli.NewApp()
)

func InstallCommand(c cli.Command) {
	App.Commands = append(App.Commands, c)
}
