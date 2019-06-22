package domain

import "regexp"

func containsNumbers(word string) bool {
	re := regexp.MustCompile("[0-9]+")
	if len(re.FindAllString(word, 1)) > 0 {
		return true
	}

	return false
}
