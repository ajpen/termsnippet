package core

import (
	"os"
	"os/user"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	defaultAppDataPath string
	testAppPath        = "./test.db"
	testSnippet        = Snippet{
		Title:       "Test Snippet",
		Body:        "randome stuff asdsadasdsda",
		Description: "Stuff",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testSnippet2 = Snippet{
		Title:       "Test_Snippet2",
		Body:        "Updaasdasds stuff asdsadasdsda",
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

func TestDatabaseCreation(t *testing.T) {
	var err error
	_, err = NewSnippetDatabase(testAppPath)
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
	defer os.Remove(testAppPath)

	err = db.AddSnippet(testSnippet)
	if err != nil {
		t.Error("Failed to add Snippet, ", err)
	}

	_, err = db.GetSnippet(testSnippet.Title)
	if err != nil {
		t.Error("Failed to retrieve Snippet. It was not added properly, ", err)
	}
}

func TestAllSnippet(t *testing.T) {
	var passed = false
	db, err := NewSnippetDatabase(testAppPath)
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(testAppPath)

	err = db.AddSnippet(testSnippet)
	if err != nil {
		t.Error("Failed to add Snippet, ", err)
	}

	err = db.AddSnippet(testSnippet2)
	if err != nil {
		t.Error("Failed to add Snippet, ", err)
	}

	snippets, e := db.AllSnippets()
	if e != nil {
		t.Error("Failed to retrieve all snippets, ", e)
	}
	var count = 0
	for _, s := range snippets {
		if s.Title == testSnippet.Title || s.Title == testSnippet2.Title {
			passed = true
		} else {
			t.Errorf("Unexpected snippet %+v", s)
			passed = false
		}
		count++
	}
	if count != 2 {
		t.Errorf("snippet count is invalid: expected 2 got %d", count)
	}
	if !passed {
		t.Fail()
	}
}

func TestUpdateSnippet(t *testing.T) {

	db, err := NewSnippetDatabase(testAppPath)
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(testAppPath)

	err = db.AddSnippet(testSnippet)
	if err != nil {
		t.Error("Failed to add Snippet, ", err)
	}

	var snippet Snippet

	err = db.UpdateSnippet(testUpdateSnippet)
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
	db, err := NewSnippetDatabase(testAppPath)
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(testAppPath)

	err = db.AddSnippet(testUpdateSnippet)
	if err != nil {
		t.Error("Failed to add Snippet, ", err)
	}

	err = db.DeleteSnippet(testUpdateSnippet.Title)

	if err != nil {
		t.Error("Unable to delete snippet, ", err)
	}

	_, err = db.GetSnippet(testSnippet.Title)
	if err == nil {
		t.Error("Deleted snippet was retrieved. Snippet was not properly deleted ")
	}
}
