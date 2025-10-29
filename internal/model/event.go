package model

// Event represents a events event
type Event struct {
	ID        string
	Domain    string
	URL       string
	Title     string
	Content   string
	CreatedAt int64

	UserID string
}
