package books

// Service provides book-related operations
type Service struct {
	repo *Repository
}

// NewService creates a new book service
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetBooks retrieves all books from the repository
func (s *Service) GetBooks() []Book {
	return s.repo.GetBooks()
}

// AddBook adds a new book to the repository
func (s *Service) AddBook(title, author string) Book {
	book := Book{Title: title, Author: author}
	s.repo.AddBook(book)
	return book
}
