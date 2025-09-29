package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestGetLongURL(t *testing.T) {
	tests := []struct {
		name            string
		shortURL        string
		mockSetup       map[string]string
		expectedLongURL string
		expectedError   error
	}{
		{
			name:            "Успешное получение",
			shortURL:        "testkey1",
			mockSetup:       map[string]string{"testkey1": "https://example.com/original1"},
			expectedLongURL: "https://example.com/original1",
			expectedError:   nil,
		},
		{
			name:            "Ключ не найден (404)",
			shortURL:        "missingkey",
			mockSetup:       map[string]string{"existing": "http://present.com"},
			expectedLongURL: "",
			expectedError:   errors.New("not found 404"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := NewMockStore()
			for short, long := range tt.mockSetup {
				mockStore.urls[short] = long
			}

			s := NewService(mockStore)

			longURL, err := s.GetLongURL(tt.shortURL)

			if !errors.Is(err, tt.expectedError) && err != tt.expectedError && err.Error() != tt.expectedError.Error() {
				t.Fatalf("Ошибка не совпадает. Ожидалась: '%v', Получено: '%v'", tt.expectedError, err)
			}

			if longURL != tt.expectedLongURL {
				t.Errorf("Полученный URL не совпадает. Ожидалось: '%s', Получено: '%s'", tt.expectedLongURL, longURL)
			}
		})
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

	if !strings.Contains(finalLongURL, initialLongURL) {
		t.Errorf("Финально сохраненный URL не соответствует формату коллизии. Ожидалось, что он будет %s", initialLongURL)
	}
}
