package command

import (
	"fmt"
	"termsnippet/core"

	"gopkg.in/urfave/cli.v1"
)

func init() {
	InstallCommand(listSnippetCommand())
}

const (
	listSnippetTemplate = "%s: %s\n\n"
)

func listSnippetCommand() cli.Command {
	cmd := cli.Command{
		Name:        "list",
		Description: "list all saved snippets",

		Action: func(c *cli.Context) error {
			snippets, err := core.AllSnippets()
			if err != nil {
				return fmt.Errorf("Error getting saved snippets: %s", err)
			}
			for _, snippet := range snippets {
				fmt.Printf(listSnippetTemplate, snippet.Title, snippet.Description)
			}
			return nil
		},
	}
	return cmd
}
