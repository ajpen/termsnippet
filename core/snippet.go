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

func UnmarshalSnippet(blob []byte) (Snippet, error) {

	s := &Snippet{}
	err := json.Unmarshal(blob, s)
	return *s, err
}

func NewSnippet(title, description, body string) *Snippet {
	now := time.Now()
	s := new(Snippet)
	s.Title = title
	s.Body = body
	s.Description = description
	s.CreatedAt = now
	return s
}
