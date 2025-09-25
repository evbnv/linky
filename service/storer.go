package service

type Storer interface {
	GetURL(shortURL string) (string, error)
	SaveURL(shortURL, longURL string) error
}

type Service struct {
	store Storer
}

func NewService(store Storer) *Service {
	return &Service{store: store}
}