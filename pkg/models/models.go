package models

import (
	"errors"
	"fmt"
	"time"
)

var ErrorNoRecord = errors.New("models: record not found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

func (s *Snippet) String() string {
	return fmt.Sprintf("id: %d, title: %s, conent: %s", s.ID, s.Title, s.Content)
}

type Article struct {
	ID      int
	Title   string
	Slug    string
	Created time.Time
	Expires time.Time
}

func (a *Article) String() string {
	return fmt.Sprintf("id: %d, title: %s, slug: %s", a.ID, a.Title, a.Slug)
}
