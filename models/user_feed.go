package models

type UserFeed struct {
	Items []UserFeedItem `json:"items"`
}

type UserFeedItem struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	SourceName string `json:"source_name"`
	SourceURL  string `json:"source_url"`
}
