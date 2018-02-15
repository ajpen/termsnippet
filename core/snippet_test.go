package core

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

var (
	title       = "something"
	body        = "something else"
	description = "another thing"
	createdAt   = time.Now()
	updatedAt   = time.Now()
)

func TestMarshalSnippetForStorage(t *testing.T) {
	s := Snippet{
		Title:       title,
		Body:        body,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	name, blob, err := MarshalSnippetForStorage(s)

	if err != nil {
		t.Error("Marshal failed: %s", err.Error())
	}

	if string(name) != title {
		t.Error("Mismatch in snippet name from marshal")
	}

	decodedBlob := &Snippet{}
	err = json.Unmarshal(blob, decodedBlob)

	if err != nil {
		t.Error("Unable to unmarshal snippet")
	}

	isEqual := cmp.Equal(*decodedBlob, s)

	if !isEqual {
		diff := cmp.Diff(*decodedBlob, s)
		t.Error("snippet blob is not the same as original snippet\nDiff: %s", diff)
	}

}
