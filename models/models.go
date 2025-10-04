package models

type Storer interface {
	GetURL(shortURL string) (string, error)
	SaveURL(shortURL, longURL string) error
}

type Service struct {
	Store Storer
}

type ShortenRequest struct {
	URL string `json:"long_url"`
}
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func NewService(store Storer) *Service {
	return &Service{Store: store}
}
