package core

import (
	"os"
	"os/user"
	"testing"
)

var (
	defaultAppDataPath string
	homeDir            string
	testAppPath        string
)

func Init() {

}

func homeDir(t *testing.T) string {
	user, err := user.Current()
	if err != nil {
		t.Error(err)
	}
	if user.HomeDir == "" {
		t.Error("Cant get home dir")
	}
	return user.HomeDir
}

// Test default directory == homedir + app data path
// Test database and bucket creation. Validate with successful transaction
// Test add, update and delete snippet functions
