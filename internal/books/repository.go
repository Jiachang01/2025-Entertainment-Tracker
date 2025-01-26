package books

// Book represents a book in the tracker
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Repository handles book storage
type Repository struct {
	books []Book
}

// NewRepository initializes a new book repository
func NewRepository() *Repository {
	return &Repository{
		books: []Book{
			{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
		},
	}
}

// GetBooks returns all books
func (r *Repository) GetBooks() []Book {
	return r.books
}

// AddBook adds a new book
func (r *Repository) AddBook(book Book) {
	book.ID = len(r.books) + 1
	r.books = append(r.books, book)
}
