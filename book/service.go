package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookInput) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookInput BookInput) (Book, error) {
	priceBook := bookInput.Price
	newRate := bookInput.Rating
	newBook := Book{
		Title:       bookInput.Title,
		Description: bookInput.Description,
		Price:       priceBook,
		Rating:      newRate,
	}
	book, err := s.repository.Create(newBook)
	return book, err
}
