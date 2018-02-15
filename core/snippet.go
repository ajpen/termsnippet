package core

import (
	"time"
)

type Snippet struct {
	Title       string
	Body        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
