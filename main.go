package main

import (
	"os"
	"termsnippet/cli/command"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "TermSnippet"
	app.Usage = "A command line based reusable code manager"
	app.Description = "TermSnippet is a command line based code snippet database for storing and accessing resuable code"
	app.Author = "Anfernee Jervis"
	app.Version = "0.0.1"
	app.Commands = command.Commands

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
