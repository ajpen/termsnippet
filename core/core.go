package core

import (
	"fmt"
	"os"
)

const (
	defaultDbName = "data.db"
)

var (
	snippetStore *SnippetDatabase
)

func init() {
	defaultDataDir, err := defaultAppDataDir()
	if err != nil {
		panic(fmt.Errorf("Failed to initialize database: %s", err))
	}
	if _, err := os.Stat(defaultDataDir); os.IsNotExist(err) {
		err = os.MkdirAll(defaultDataDir, 0744)
		if err != nil {
			panic(fmt.Errorf("Failed to initialize database: %s", err))
		}
	}
	fullDataPath := defaultDataDir + "/" + defaultDbName
	sd, err := NewSnippetDatabase(fullDataPath)
	if err != nil {
		panic(fmt.Errorf("Failed to initialize database: %s", err))
	}
	snippetStore = sd
}

func GetSnippet(title string) (Snippet, error) {
	s, e := snippetStore.GetSnippet(title)
	return s, e
}

func AddSnippet(title, description, body string) error {
	s := NewSnippet(title, description, body)
	return snippetStore.AddSnippet(*s)
}

func RenameSnippet(oldTitle, newTitle string) error {
	s, err := GetSnippet(oldTitle)
	if err != nil {
		return fmt.Errorf("Unable to get snippet with title %s: %s", oldTitle, err)
	}

	s.Title = newTitle
	snippetStore.UpdateSnippet(s)
	return err
}

func ChangeSnippetDescription(title, newDescription string) error {
	s, err := GetSnippet(title)
	if err != nil {
		return fmt.Errorf("Unable to get snippet with title %s: %s", title, err)
	}

	s.Description = newDescription
	snippetStore.UpdateSnippet(s)
	return err
}

func EditSnippet(title, newBody string) error {
	s, err := GetSnippet(title)
	if err != nil {
		return fmt.Errorf("Unable to get snippet with title %s: %s", title, err)
	}

	s.Body = newBody
	snippetStore.UpdateSnippet(s)
	return err
}

func AllSnippets() ([]Snippet, error) {
	return snippetStore.AllSnippets()
}

func DeleteSnippet(title string) error {
	return snippetStore.DeleteSnippet(title)
}
