package core

import (
	"fmt"
	"os/user"

	"github.com/boltdb/bolt"
)

const (
	appDataRelativePath = ".termsnippet/data"
	appDataFileMode     = 0600
	bucketName          = "snippets"
)

func getHomeDirectory() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.HomeDir, nil
}

func defaultAppDataDir() (string, error) {
	home, err := getHomeDirectory()
	if err != nil {
		return "", err
	}
	return home + "/" + appDataRelativePath, nil
}

type SnippetDatabase struct {
	DB       *bolt.DB
	DataPath string
}

func NewSnippetDatabase(dataPath string) (*SnippetDatabase, error) {
	var err error

	if dataPath == "" {
		return nil, fmt.Errorf("Database path required for initialization!")
	}

	sd := &SnippetDatabase{DataPath: dataPath}

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

func (sd *SnippetDatabase) GetSnippet(title string) (Snippet, error) {

	var s Snippet
	err := sd.DB.Update(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("data error:\nunable to find data. either it has been corrupted or removed unexpectedly")
		}

		snippet := b.Get([]byte(title))
		s, err = UnmarshalSnippet(snippet)
		return err
	})

	return s, err
}

func (sd *SnippetDatabase) AddSnippet(s Snippet) error {

	name, snippetBlob, err := MarshalSnippetForStorage(s)
	if err != nil {
		return err
	}

	err = sd.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("data error:\nunable to find data. either it has been corrupted or removed unexpectedly")
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
			return fmt.Errorf("data error:\nunable to find data. either it has been corrupted or removed unexpectedly")
		}

		err := b.Delete([]byte(name))
		return err
	})

	return err
}

func (sd *SnippetDatabase) AllSnippets() ([]Snippet, error) {
	var snippets []Snippet
	err := sd.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		stats := bucket.Stats()
		snippets = make([]Snippet, 0, stats.KeyN)
		cursor := bucket.Cursor()

		n, v := cursor.First()
		for v != nil && n != nil {
			snippet, e := UnmarshalSnippet(v)
			if e != nil {
				snippets = nil
				return fmt.Errorf("Unable to retrieve all Snippets: %s", e)
			}
			snippets = append(snippets, snippet)
			n, v = cursor.Next()
		}
		return nil
	})
	return snippets, err
}
