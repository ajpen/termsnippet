package command

import (
	"fmt"
	"termsnippet/core"

	"gopkg.in/urfave/cli.v1"
)

const (
	snippetDisplayTemplate = "\nTitle\n" +
		"%s \n\n" +
		"Description\n" +
		"%s \n\n" +
		"%s" +
		"\n"
)

func init() {
	InstallCommand(viewSnippetCommand())
}

func viewSnippetCommand() cli.Command {
	cmd := cli.Command{
		Name:        "view",
		Description: "view contents of snippet",
		ArgsUsage:   "Title (required) - Title of snippet to view",

		Action: func(c *cli.Context) error {
			if c.NArg() <= 0 {
				return fmt.Errorf("Snippet title argument missing")
			}
			title := c.Args()[0]
			snippet, err := core.GetSnippet(title)
			if err != nil {
				return fmt.Errorf("Snippet %s does not exist", title)
			}
			fmt.Printf(snippetDisplayTemplate, snippet.Title, snippet.Description, snippet.Body)
			return nil
		},
	}

	return cmd
}
