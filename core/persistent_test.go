package core

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"os/user"
	"testing"
	"time"
)

var (
	db                 *SnippetDatabase
	defaultAppDataPath string
	testAppPath        = "./test.db"
	testSnippet        = Snippet{
		Title:       "Test Snippet",
		Body:        "randome stuff asdsadasdsda",
		Description: "Stuff",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testUpdateSnippet = Snippet{
		Title:       "Test Snippet",
		Body:        "Updaasdasds stuff asdsadasdsda",
		Description: "Stuff",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
)

func homeDir() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}
	if user.HomeDir == "" {
		return ""
	}
	return user.HomeDir
}

func TestDefaultDir(t *testing.T) {

	defaultDir, err := DefaultAppDataPath()
	if err != nil {
		t.Error(err)
	}

	home := homeDir()
	expected := home + "/" + appDataRelativePath
	if defaultDir != expected {
		t.Errorf("DefaultAppDataPath does not return the expected string\nExpected: %s\nGot: %s ", expected, defaultDir)
	}
}

func TestDatabaseCreation(t *testing.T) {
	var err error
	db, err = NewSnippetDatabase(testAppPath)
	if err != nil {
		t.Error(err)
	}

	if err = os.Remove(testAppPath); os.IsNotExist(err) {
		t.Error("Database file was not successfully created")
	}
}

func TestGetAndAddSnippet(t *testing.T) {
	db, err := NewSnippetDatabase(testAppPath)
	if err != nil {
		t.Error(err)
	}

	err = db.AddSnippet(testSnippet)
	if err != nil {
		t.Error("Failed to add Snippet, ", err)
	}

	_, err = db.GetSnippet(testSnippet.Title)
	if err != nil {
		t.Error("Failed to retrieve Snippet. It was not added properly, ", err)
	}

}

func TestUpdateSnippet(t *testing.T) {

	var snippet Snippet
	err := db.UpdateSnippet(testUpdateSnippet)
	if err != nil {
		t.Error("Failed to update Snippet, ", err)
	}

	snippet, err = db.GetSnippet(testSnippet.Title)
	if err != nil {
		t.Error("Failed to retrieve Snippet. It was not added properly, ", err)
	}

	isEqual := cmp.Equal(snippet, testUpdateSnippet)

	if !isEqual {
		t.Error("Snippet was not updated properly.")
	}
}

func TestDeleteSnippet(t *testing.T) {

	err := db.DeleteSnippet(testUpdateSnippet.Title)

	if err != nil {
		t.Error("Unable to delete snippet, ", err)
	}

	_, err = db.GetSnippet(testSnippet.Title)
	if err == nil {
		t.Error("Deleted snippet was retrieved. Snippet was not properly deleted ")
	}
}
