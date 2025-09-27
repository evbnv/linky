package service

import "errors"

type MockStore struct {
	urls map[string]string
	SaveCalls int
}

func NewMockStore() *MockStore {
	return &MockStore{
		urls: make(map[string]string),
	}
}

func (m *MockStore) SaveURL(shortURL, longURL string) error {
	m.SaveCalls++
	if _, ok := m.urls[shortURL]; ok {
		return errors.New("collision")
	}
	m.urls[shortURL] = longURL
	return nil
}

func (m *MockStore) GetURL(shortURL string) (string, error) {
	longURL, ok := m.urls[shortURL]
	if !ok {
		return "", errors.New("not found 404")
	}
	return longURL, nil
}
