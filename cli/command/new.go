package command

import (
	"fmt"
	"termsnippet/util"

	"github.com/ajpen/termsnippet/core"

	"github.com/atotto/clipboard"
	"gopkg.in/urfave/cli.v1"
)

func init() {
	InstallCommand(newSnippetCommand())
}

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
		Action: func(c *cli.Context) error {
			if c.NArg() <= 0 {
				return fmt.Errorf("Snippet title argument missing")
			}
			var body, title, desc string
			var err error

			title = c.Args()[0]
			if c.Bool("clip") {
				body, err = clipboard.ReadAll()
				if err != nil {
					return fmt.Errorf("Unable to read from clipboard: %s", err)
				}
			} else {
				body, err = util.OpenInEditor("")
				if err != nil {
					return fmt.Errorf("Unable to read from text editor: %s", err)
				}
			}
			desc = c.String("description")
			err = core.AddSnippet(title, desc, body)
			if err != nil {
				return fmt.Errorf("Unable to save snippet: %s", err)
			}
			return nil
		},
	}
	return cmd
}
