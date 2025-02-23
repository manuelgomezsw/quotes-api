package books

type Book struct {
	BookID   int    `json:"book_id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Synopsis string `json:"synopsis"`
	Source   string `json:"source"`
}
