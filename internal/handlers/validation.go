package handlers

import (
	"regexp"
	"strings"
)

func isEmpty(value string) bool {
	if strings.TrimSpace(value) == "" {
		return true
	}
	return false
}

func validEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}
