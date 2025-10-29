package event

type Event struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Domain    string `json:"domain"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
}
