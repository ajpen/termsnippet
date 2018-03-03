package command

import (
	"termsnippet/core"
	"termsnippet/util"

	"gopkg.in/urfave/cli.v1"
)

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

		Action: func(c *cli.Context) *cli.ExitError {
			var body, title, desc string
			var err error
			var snippet core.Snippet

			if c.NArg() <= 0 {
				return cli.NewExitError("Snippet title argument missing", 4)
			}

			title = c.Args()[0]
			snippet, err = core.GetSnippet(title)
			if err != nil {
				return cli.NewExitError("Snippet "+title+" does not exist.", 5)
			}
			body, err = util.OpenInEditor(snippet.Body)
			if err != nil {
				return cli.NewExitError("Unable to open snippet body in editor: "+err.Error(), 94)
			}
			desc = c.String("description")

			err = core.EditSnippet(title, body)
			if err != nil {
				return cli.NewExitError("Unable to edit snippet body: "+err.Error(), 94)
			}
			if desc != "" {
				err = core.ChangeSnippetDescription(title, desc)
				if err != nil {
					return cli.NewExitError("Unable to edit snippet description: "+err.Error(), 94)
				}
			}
			return nil
		},
	}
	return cmd
}
