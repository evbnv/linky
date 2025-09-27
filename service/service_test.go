package service

import (
	"crypto/sha256"
	"encoding/base64"
	"linky/models"
	"strings"
	"testing"
)

func TestGetLongURL(t *testing.T) {
	mockStore := &MockStore{
		urls: make(map[string]string),
	}

	s := models.NewService(mockStore)

	err := s.Store.SaveURL("test1", "http://example.com/test1")
	if err != nil {
		t.Fatalf("failed to save mock URL: %v", err)
	}

	longURL, err := s.Store.GetURL("test1")

	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
	if longURL != "http://example.com/test1" {
		t.Errorf("expected 'http://example.com/test1', but got '%s'", longURL)
	}
}

func TestURLTransform_Success(t *testing.T) {
	mockStore := NewMockStore()
	s := NewService(mockStore)

	longURL := "https://example.com/test-success"

	shortURL := s.URLTransform(longURL)

	if mockStore.SaveCalls != 1 {
		t.Errorf("Ожидался 1 вызов SaveURL, получено %d. Коллизии не должно было быть.", mockStore.SaveCalls)
	}

	if shortURL == "" {
		t.Fatal("URLTransform вернул пустую строку")
	}

	savedLongURL, _ := mockStore.GetURL(shortURL)
	if savedLongURL != longURL {
		t.Errorf("Сохраненный длинный URL не соответствует исходному. Ожидалось: %s, Получено: %s", longURL, savedLongURL)
	}
}

func TestURLTransform_HandlesCollision(t *testing.T) {
	mockStore := NewMockStore()
	s := NewService(mockStore)

	initialLongURL := "https://example.com/test-collision"

	h := sha256.New()
	h.Write([]byte(initialLongURL))
	collisionKey := base64.URLEncoding.EncodeToString(h.Sum(nil)[:8])

	mockStore.urls[collisionKey] = "http://fake-collided-url.com"

	shortURL := s.URLTransform(initialLongURL)

	if mockStore.SaveCalls < 2 {
		t.Errorf("Ожидалось минимум 2 вызова SaveURL для разрешения коллизии, получено %d.", mockStore.SaveCalls)
	}

	if longURL, _ := mockStore.GetURL(collisionKey); longURL == initialLongURL {
		t.Error("Ошибка: Функция сохранила ключ, который должен был вызвать коллизию.")
	}

	finalLongURL, _ := mockStore.GetURL(shortURL)

	if !strings.Contains(finalLongURL, initialLongURL+":") {
		t.Errorf("Финально сохраненный URL не соответствует формату коллизии. Ожидалось, что он начнется с '%s:'", initialLongURL)
	}
}
