package cli

import (
	"gopkg.in/urfave/cli.v1"
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
