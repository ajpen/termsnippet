package cli

import (
	"fmt"

	"github.com/ajpen/termsnippet/core"

	"github.com/atotto/clipboard"
	"gopkg.in/urfave/cli.v1"
)

func init() {
	InstallCommand(clipSnippetCommand())
}

func clipSnippetCommand() cli.Command {
	return cli.Command{
		Name:        "clip",
		Description: "copy snippet body to clipboard",
		ArgsUsage:   "Title (required) - Title of snippet to view",
		Action:      clipSnippetAction,
	}

}

func clipSnippetAction(c *cli.Context) error {
	if c.NArg() <= 0 {
		return ErrArgumentMissing("Title")
	}
	title := c.Args()[0]
	snippet, err := core.GetSnippet(title)
	if err != nil {
		return ErrDoesNotExist(title)
	}
	err = clipboard.WriteAll(snippet.Body)
	if err != nil {
		return fmt.Errorf("Unable to copy snippet body to clipboard: %s", err)
	}
	return nil
}
