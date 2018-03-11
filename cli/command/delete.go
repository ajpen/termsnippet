package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"termsnippet/core"

	"gopkg.in/urfave/cli.v1"
)

func init() {
	InstallCommand(deleteSnippetCommand())
}

func deleteSnippetCommand() cli.Command {
	cmd := cli.Command{
		Name:        "delete",
		Description: "delete snippet",

		Action: func(c *cli.Context) error {
			var err error
			if c.NArg() <= 0 {
				return fmt.Errorf("Snippet title argument missing")
			}
			title := c.Args()[0]
			snippet, err := core.GetSnippet(title)
			if err != nil {
				return fmt.Errorf("Snippet %s does not exist", title)
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
		},
	}
	return cmd
}

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
