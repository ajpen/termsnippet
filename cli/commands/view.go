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
		"%s"
)

func viewSnippetCommand() cli.Command {
	cmd := cli.Command{
		Name:        "view",
		Description: "view contents of snippet",
		ArgsUsage:   "Title (required) - Title of snippet to view",

		Action: func(c *cli.Context) *cli.ExitError {
			if c.NArg() <= 0 {
				return cli.NewExitError("Snippet title argument missing", 4)
			}
			title := c.Args()[0]
			snippet, err := core.GetSnippet(title)
			if err != nil {
				return cli.NewExitError("Snippet "+title+" does not exist.", 5)
			}
			fmt.Printf(snippetDisplayTemplate, snippet.Title, snippet.Description, snippet.Body)
			return nil
		},
	}

	return cmd
}
