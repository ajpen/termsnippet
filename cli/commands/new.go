package command

import (
	"termsnippet/core"
	"termsnippet/util"

	"github.com/atotto/clipboard"
	"gopkg.in/urfave/cli.v1"
)

func newSnippetCommand() cli.Command {
	cmd := cli.Command{
		Name:        "new",
		Description: "Create a new code snippet",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "clip, c",
				Usage: "If set, the contents of the clipboard will be used as the snippet body",
			},
			cli.StringFlag{
				Name:  "description, d",
				Usage: "code snippet description",
				Value: "",
			},
		},
		ArgsUsage: "Title (required) - Sets the title of the code snippet",
		Action: func(c *cli.Context) *cli.ExitError {
			if c.NArg() <= 0 {
				return cli.NewExitError("Snippet title argument missing", 4)
			}
			var body, title, desc string
			var err error

			title = c.Args()[0]
			if c.Bool("clip") {
				body, err = clipboard.ReadAll()
				if err != nil {
					return cli.NewExitError("Unable to read from clipboard: "+err.Error(), 90)
				}
			} else {
				body, err = util.OpenInEditor("")
				if err != nil {
					return cli.NewExitError("Unable to read from text editor: "+err.Error(), 91)
				}
			}
			desc = c.String("description")
			err = core.AddSnippet(title, desc, body)
			if err != nil {
				return cli.NewExitError("Unable to save snippet: "+err.Error(), 20)
			}
			return nil
		},
	}
	return cmd
}
