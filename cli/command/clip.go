package command

import (
	"fmt"
	"termsnippet/core"

	"github.com/atotto/clipboard"
	"gopkg.in/urfave/cli.v1"
)

func init() {
	Commands = append(Commands, clipSnippetCommand())
}

func clipSnippetCommand() cli.Command {
	cmd := cli.Command{
		Name:        "clip",
		Description: "copy snippet body to clipboard",
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
			err = clipboard.WriteAll(snippet.Body)
			if err != nil {
				return fmt.Errorf("Unable to copy snippet body to clipboard: %s", err)
			}
			return nil
		},
	}
	return cmd
}
