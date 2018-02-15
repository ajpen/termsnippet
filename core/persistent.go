package core

import (
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
	appDataPath string
)

func init() {
	home, err := getHomeDirectory()
	if err != nil {
		panic(err)
	}
	appDataPath = home + "/" + appDataRelativePath
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
	sd := &SnippetDatabase{DataPath: dataPath}
	var err error

	sd.DB, err = bolt.Open(dataPath, appDataFileMode, nil)
	if err != nil {
		panic(err)
	}

	err = sd.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return sd, nil
}

func (sd *SnippetDatabase) AddSnippet(s Snippet) error {

	name, snippetBlob, err := MarshalSnippetForStorage(s)
	if err != nil {
		return err
	}

	err = sd.DB.Update(func(tx *bolt.Tx) error {

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

func (sd *SnippetDatabase) UpdateSnippet(s Snippet) error {
	return sd.AddSnippet(s)
}

func (sd *SnippetDatabase) DeleteSnippet(name string) error {
	err := sd.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("Data error:\nUnable to find data. Either it has been corrupted or removed unexpectedly.")
		}

		err := b.Delete([]byte(name))
		return err
	})

	return err
}
