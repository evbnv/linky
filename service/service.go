package service

import (
	"crypto/rand"
	"linky/models"
	"log/slog"
	"math/big"
	"unicode/utf8"
)

const (
	alphabet       = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
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
	shortURLRunes := make([]rune, shortURLLength)
	alphabetRunes := []rune(alphabet)
	alphabetLength := big.NewInt(int64(utf8.RuneCountInString(alphabet)))

	for {
		for i := 0; i < shortURLLength; i++ {
			index, err := rand.Int(rand.Reader, alphabetLength)
			if err != nil {
				continue
			}
			shortURLRunes[i] = alphabetRunes[index.Int64()]
		}
		shortURL = string(shortURLRunes)
		err := s.store.SaveURL(shortURL, longURL)
		if err != nil {
			slog.Warn("Collision detected during short URL generation",
				"event", "collision_detected", "shortURL", shortURL, "longURL", longURL, "error", err)
			continue
		}
		break
	}
	slog.Info("Short URL succesfully generated", "event", "url_created", "shortURL", shortURL, "longURL", longURL)
	return shortURL
}

func (s *Service) GetLongURL(shortURL string) (string, error) {
	longURL, err := s.store.GetURL(shortURL)
	return longURL, err
}
