package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ajpen/termsnippet/core"

	"gopkg.in/urfave/cli.v1"
)

func init() {
	InstallCommand(deleteSnippetCommand())
}

func deleteSnippetCommand() cli.Command {
	return cli.Command{
		Name:        "delete",
		Description: "delete snippet",

		Action: deleteSnippetAction,
	}
}

func deleteSnippetAction(c *cli.Context) error {
	var err error
	if c.NArg() <= 0 {
		return ErrArgumentMissing("Title")
	}
	title := c.Args()[0]
	snippet, err := core.GetSnippet(title)
	if err != nil {
		return ErrDoesNotExist(title)
	}
	fmt.Printf("Delete snippet %s? ", snippet.Title)
	decision := strings.ToLower(readLine())
	if decision == "y" || decision == "yes" {
		err = core.DeleteSnippet(title)
	}
	if err != nil {
		return fmt.Errorf("Unable to delete snippet: %s", err)
	}
	fmt.Printf("Snippet %s deleted\n", title)
	return nil
}

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
