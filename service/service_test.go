package service

import "testing"

func TestGetLongURL(t *testing.T) {
	mockStore := &MockStore{
		urls: make(map[string]string),
	}

	s := NewService(mockStore)

	err := s.store.SaveURL("test1", "http://example.com/test1")
	if err != nil {
		t.Fatalf("failed to save mock URL: %v", err)
	}

	longURL, err := s.store.GetURL("test1")

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
	if longURL != "http://example.com/test1" {
		t.Errorf("expected 'http://example.com/test1', but got '%s'", longURL)
	}
}
