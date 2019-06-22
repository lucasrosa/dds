package domain

import (
	"github.com/agext/levenshtein" 
) 

func tolaratedDistance(needleLength int) int {
	distance := 3 // for words bigger than 16 characters

	if needleLength < 4 {
		distance = 0
	}else if needleLength >= 4 && needleLength <= 8 {
		distance = 1
	}  else if needleLength  > 8 && needleLength  <= 16 {
		distance = 2
	}
	
	return distance
}

// quando for marca, o score é maior do que quando for palavra do dicionario
func calculateLevenshtein(haystack, needle string) bool {
	nl := len(needle)
	hl := len(haystack)
	if nl > hl {
		return false
	}
	// The expected levenshtein distance is based on the size of the needle
	td := tolaratedDistance(nl)

	total := hl - nl
	for i := 0; i <= total; i++ {
		if levenshtein.Distance(needle, haystack[i:i+nl], nil) <= td {
			return true
		}
	}

	return false
}