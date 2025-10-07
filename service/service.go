package service

import (
	"linky/models"
	"log"
	"math/rand"
)

type Service struct {
	store models.Storer
}

func NewService(s models.Storer) *Service {
	return &Service{store: s}
}

func (s *Service) URLTransform(longURL string) string {
	alphabet := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var shortURL string
	shortURLBytes := make([]byte, 6)

	for {
		for i := range shortURLBytes {
			index := rand.Intn(len(alphabet))
			shortURLBytes[i] = alphabet[index]
		}
		shortURL = string(shortURLBytes)
		err := s.store.SaveURL(shortURL, longURL)
		if err != nil {
			log.Println(err)
			// there is collision
			log.Println("There is collision")
			continue
		}
		break
	}
	log.Println("Short URL created")
	return shortURL
}

func (s *Service) GetLongURL(shortURL string) (string, error) {
	longURL, err := s.store.GetURL(shortURL)
	return longURL, err
}
