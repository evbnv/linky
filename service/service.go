package service

import (
	"crypto/sha256"
	"encoding/base64"
	"linky/models"
	"log"
	"math/rand"
	"strconv"
)

type Service struct {
	store models.Storer
}

func NewService(s models.Storer) *Service {
	return &Service{store: s}
}

func (s *Service) URLTransform(longURL string) string {
	var shortURL string

	for {
		h := sha256.New()
		h.Write([]byte(longURL))
		shortURL = base64.URLEncoding.EncodeToString([]byte(h.Sum(nil)[:8]))

		err := s.store.SaveURL(shortURL, longURL)
		if err != nil {
			log.Println(err)
			// there is collision
			log.Println("There is collision")
			longURL = longURL + ":" + strconv.Itoa(rand.Int())
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
