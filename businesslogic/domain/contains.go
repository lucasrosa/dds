package domain

import (
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

func contains(haystack, needle string) bool {
	t := language.Tag(language.BrazilianPortuguese)
	// IgnoreDiacritics is used to match things like á == a
	matcher := search.New(t, search.IgnoreDiacritics)

	start, _ := matcher.IndexString(haystack, needle)

	if start == -1 {
		return false
	}

	return true
}
