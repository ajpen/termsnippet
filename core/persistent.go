package core

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"os/user"
)

const (
	appDataRelativePath = ".termsnippet/data/data.db"
	appDataFileMode     = 0600
	bucketName          = "snippets"
)

var (
	appDataPath = nil
)

func init() {
	home, err := getHomeDirectory()
	if err != nil {
		panic(err)
	}
	appDataPath = home + '/' + appDataRelativePath
}

func getHomeDirectory() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.HomeDir, nil
}

type SnippetDatabase struct {
	DB       *bolt.DB
	DataPath string
}

func NewSnippetDatabase(dataPath string) (*SnippetDatabase, error) {
	fd := &SnippetDatabase{DataPath: dataPath}

	fd.DB, err = bolt.Open(dataPath, appDataFileMode, nil)
	if err != nil {
		panic(err)
	}

	err := fd.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fd, nil
}

func (sd *SnippetDatabase) AddSnippet(s Snippet) error {

	name, err := json.Marshal(s.Title)
	if err != nil {
		return err
	}

	snippetBlob, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err := sd.DB, Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("Data error:\nUnable to find data. Either it has been corrupted or removed unexpectedly.")
		}

		err := b.Put(name, snippetBlob)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
