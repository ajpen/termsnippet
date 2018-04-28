package cli

import (
	"fmt"

	"github.com/ajpen/termsnippet/core"
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
	return cli.Command{
		Name:        "view",
		Description: "view contents of snippet",
		ArgsUsage:   "Title (required) - Title of snippet to view",

		Action: viewSnippetAction,
	}

}

func viewSnippetAction(c *cli.Context) error {
	if c.NArg() <= 0 {
		return ErrArgumentMissing("Title")
	}
	title := c.Args()[0]
	snippet, err := core.GetSnippet(title)
	if err != nil {
		return ErrDoesNotExist(title)
	}
	fmt.Printf(snippetDisplayTemplate, snippet.Title, snippet.Description, snippet.Body)
	return nil
}
