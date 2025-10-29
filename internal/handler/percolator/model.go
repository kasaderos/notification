package percolator

type Event struct {
	ID      string `json:"id"`
	Domain  string `json:"domain"`
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
