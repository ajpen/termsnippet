package main

import (
	"os"
	"termsnippet/cli"
)

func main() {
	app := cli.App
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
