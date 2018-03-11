# Term Snippet
Copy useful snippets of code for later, right in the terminal

## Installation

- Run `go get github.com/ajpen/termsnippet`. 

Termsnippet should then be installed as `termsnippet` in `/$GOPATH/bin`.

Note: be sure to set `$EDITOR`; `termsnippet` relies on that environment variable to know what editor it should use when editing snippets or creating snippets.


## Usage

You can view all support commands with `termsnippet -h`. `termsnippet {command} -h` will display gelp information for a specific command. 

`termsnippet new {title}` creates a new snippet with the name `title`. This will open a text editor, annowing you to type (or paste) your code snippet. You can add `-c` to instruct `termsnippet` to use the code on the clipboard for the snippet. To include an optional description, use `-d` followed by the description.

`termsnippet edit {title}` opens a text editor allowing you to edit the snippet. You use `-d` to change the description.

`termsnippet clip {title}` copies the code snippet to clipboard.

`termsnippet list` prints displays all saved snippets. Using this command along with `less` gives a better experience.

`termsnippet view {title}` prints the name, description and contents of the snippet.

`termsnippet delete {title}` permanently removes an existing snippet. Use with care.
