package command

import (
	"fmt"
	"termsnippet/core"
	"termsnippet/util"

	"gopkg.in/urfave/cli.v1"
)

func init() {
	InstallCommand(editSnippetCommand())
}

func editSnippetCommand() cli.Command {
	cmd := cli.Command{
		Name:        "edit",
		Description: "Edit an existing code snippet",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "description, d",
				Usage: "update code snippet description",
				Value: "",
			},
		},
		ArgsUsage: "Title (required) - Title of the snippet to update",

		Action: func(c *cli.Context) error {
			var body, title, desc string
			var err error
			var snippet core.Snippet

			if c.NArg() <= 0 {
				return fmt.Errorf("Snippet title argument missing")
			}

			title = c.Args()[0]
			snippet, err = core.GetSnippet(title)
			if err != nil {
				return fmt.Errorf("Snippet %s does not exist", title)
			}
			body, err = util.OpenInEditor(snippet.Body)
			if err != nil {
				return fmt.Errorf("Unable to open snippet body in editor: %s", err)
			}
			desc = c.String("description")

			err = core.EditSnippet(title, body)
			if err != nil {
				return fmt.Errorf("Unable to edit snippet body: %s", err)
			}
			if desc != "" {
				err = core.ChangeSnippetDescription(title, desc)
				if err != nil {
					return fmt.Errorf("Unable to edit snippet description: %s", err)
				}
			}
			return nil
		},
	}
	return cmd
}
