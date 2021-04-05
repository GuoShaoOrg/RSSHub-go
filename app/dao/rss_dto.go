package dao

type RSSFeed struct {
	Title       string
	Link        string
	Description string
	Author      string
	Created     string
	ImageUrl    string
	Items       []RSSItem
	Tag         []string
}

type RSSItem struct {
	Title       string
	Link        string
	Description string
	Author      string
	Created     string
}
