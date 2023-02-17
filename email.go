package customerimporter

import (
	"fmt"
	"regexp"
)

type Email string

func NewEmail(e string) (*Email, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	isValid := re.MatchString(e)

	if !isValid {
		return nil, fmt.Errorf("Invalid email format: %s", e)
	}

	var email Email = Email(e)

	return &email, nil
}

func (e Email) Domain() (string, error) {
	re := regexp.MustCompile(`(?i)@([a-z0-9\-]+(\.[a-z0-9\-]+)*\.[a-z]{2,})$`)
	domain := re.FindStringSubmatch(string(e))

	if len(domain) != 3 {
		return "", fmt.Errorf("Unable to extract domain of: %s", e)
	}

	return domain[1], nil
}
