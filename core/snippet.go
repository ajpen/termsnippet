package core

import (
	"encoding/json"
	"time"
)

type Snippet struct {
	Title       string
	Body        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func MarshalSnippetForStorage(s Snippet) ([]byte, []byte, error) {

	name := []byte(s.Title)
	snippetBlob, err := json.Marshal(s)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	return name, snippetBlob, err
}
