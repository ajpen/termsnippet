package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// OpenInEditor Opens a text editor and returns any text saved when the editor was opened.
func OpenInEditor(s string) (string, error) {
	var editor string
	var text []byte

	f, err := ioutil.TempFile("", "tsnppt")
	if err != nil {
		return "", fmt.Errorf("Unable to create temporary file for writing: %s", err)
	}

	defer os.Remove(f.Name())

	// populate file with contents of s
	_, err = f.WriteString(s)
	if err != nil {
		return "", fmt.Errorf("Unable to write contents to file: %s", err)
	}

	if e, ok := os.LookupEnv("EDITOR"); ok {
		editor = e
	}

	cmd := exec.Command(editor, f.Name())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		return "", fmt.Errorf("Cannot open editor for writing: %s", err)
	}

	text, err = ioutil.ReadFile(f.Name())

	if err != nil {
		return "", fmt.Errorf("Unable to obtain text from editor: %s", err)
	}
	return string(text), nil
}
