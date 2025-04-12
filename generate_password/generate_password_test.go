package generate_password

import (
	"testing"
)

// TestGeneratePassword проверяет корректность длины и допустимость символов.
func TestGeneratePassword(t *testing.T) {
	// Задаем тестовые длины
	lengths := []int{0, 1, 5, 10, 32, 64}

	allowed := make(map[rune]bool)
	for _, c := range letters {
		allowed[c] = true
	}

	for _, length := range lengths {
		pwd, err := GeneratePassword(length)
		if err != nil {
			t.Errorf("Неожиданная ошибка при генерации пароля длины %d: %v", length, err)
		}

		if len(pwd) != length {
			t.Errorf("Ожидалась длина %d, получена длина %d для пароля '%s'", length, len(pwd), pwd)
		}

		for _, char := range pwd {
			if !allowed[char] {
				t.Errorf("Недопустимый символ '%c' в пароле '%s'", char, pwd)
			}
		}
	}
}

// TestRandomness проверяет, что последующие вызовы функции не генерируют идентичные пароли.
func TestRandomness(t *testing.T) {
	pwd1, err := GeneratePassword(16)
	if err != nil {
		t.Errorf("Ошибка при генерации пароля: %v", err)
	}
	pwd2, err := GeneratePassword(16)
	if err != nil {
		t.Errorf("Ошибка при генерации пароля: %v", err)
	}

	if pwd1 == pwd2 {
		t.Errorf("Два последовательных вызова функции вернули одинаковый пароль: %s", pwd1)
	}
}
