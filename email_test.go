package customerimporter

import (
	"testing"
)

func TestNewEmail(t *testing.T) {
	t.Run("Should NOT return error when given email is valid", func(t *testing.T) {
		_, err := NewEmail("juanitovalderrama@gmail.com")

		if err != nil {
			t.Errorf("Email should be valid, cause: %v", err)
		}
	})

	t.Run("Should return error when given email has WRONG format", func(t *testing.T) {
		_, err := NewEmail("juanitovalderrama@gmail,com")

		if err == nil {
			t.Error("Email should not be valid")
		}
	})
}

func TestDomain(t *testing.T) {
	t.Run("Should NOT return error when email has valid domain", func(t *testing.T) {
		email, err := NewEmail("juanitovalderrama@gmail.com")

		if err != nil {
			t.Errorf("Email should be valid, cause: %v", err)
		}

		_, err = email.Domain()
		if err != nil {
			t.Errorf("Unable to get domain, cause: %v", err)
		}
	})

	t.Run("Should return error when email has NOT valid domain", func(t *testing.T) {
		// Create instace of email skipping the validation
		var email Email = "juanitovalderrama@dd-qwe---&gmail.com"

		_, err := email.Domain()
		if err == nil {
			t.Errorf("Domain should not be valid, cause: %v", err)
		}
	})
}
