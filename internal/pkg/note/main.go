package note

import "time"

// Note structure
type Note struct {
	ID        int64
	JournalID int64
	Title     string
	Body      string
	Mood      string
	Tags      []string
	CreatedAt time.Time
}

// Repository provideDAL
type Repository interface {
	Get(id int64) (*Note, error)
	GetByJournalID(id int64) ([]Note, error)
	Create(journalID int64, title, body, mood string, tags []string) (int64, error)
	Remove(id int64) error
}