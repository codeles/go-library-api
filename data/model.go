package data

type Book struct {
	ID     int32  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{
	{1, "The Neverending Story", "Michael Ende"},
	{2, "The Fellowship of the Ring", "J.R.R. Tolkien"},
}
