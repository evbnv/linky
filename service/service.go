package service

import (
	"crypto/rand"
	"linky/models"
	"log"
	"math/big"
)

const (
	alphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shortURLLength = 6
)

type Service struct {
	store models.Storer
}

func NewService(s models.Storer) *Service {
	return &Service{store: s}
}

func (s *Service) URLTransform(longURL string) string {
	var shortURL string
	shortURLBytes := make([]byte, shortURLLength)

	alphabetLength := big.NewInt(int64(len(alphabet)))

	for {
		for i := 0; i < shortURLLength; i++ {
			index, err := rand.Int(rand.Reader, alphabetLength)
			if err != nil {
				continue
			}
			shortURLBytes[i] = alphabet[index.Int64()]
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
