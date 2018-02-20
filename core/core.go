package core

import (
	"fmt"
)

var (
	SnippetStore *SnippetDatabase
)

func Init() {
	defaultPath, err := DefaultAppDataPath()
	if err != nil {
		panic(err)
	}
	SnippetStore, err = NewSnippetDatabase(defaultPath)
	if err != nil {
		panic(err)
	}
}

func GetSnippet(title string) (Snippet, error) {
	s, e := SnippetStore.GetSnippet(title)
	return s, e
}

func AddSnippet(title, description, body string) error {
	s := NewSnippet(title, description, body)
	return SnippetStore.AddSnippet(*s)
}

func RenameSnippet(oldTitle, newTitle string) error {
	s, err := GetSnippet(oldTitle)
	if err != nil {
		return fmt.Errorf("Unable to get snippet with title %s: %s", oldTitle, err)
	}

	s.Title = newTitle
	SnippetStore.UpdateSnippet(s)
	return err
}

func ChangeSnippetDescription(title, newDescription string) error {
	s, err := GetSnippet(title)
	if err != nil {
		return fmt.Errorf("Unable to get snippet with title %s: %s", title, err)
	}

	s.Description = newDescription
	SnippetStore.UpdateSnippet(s)
	return err
}

func EditSnippet(title, newBody string) error {
	s, err := GetSnippet(title)
	if err != nil {
		return fmt.Errorf("Unable to get snippet with title %s: %s", title, err)
	}

	s.Body = newBody
	SnippetStore.UpdateSnippet(s)
	return err
}
