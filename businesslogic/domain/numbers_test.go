package domain

import "testing"

func TestContainsNumbers(t *testing.T) {
	cases := []struct {
		word     string
		expected bool
	}{{
		word:     "hello",
		expected: false,
	}, {
		word:     "hello1",
		expected: true,
	}, {
		word:     "1hello",
		expected: true,
	}, {
		word:     "he1lo",
		expected: true,
	}, {
		word:     "hellodarknessmy0ldfriend",
		expected: true,
	}, {
		word:     "hellodarknessmyoldfriend",
		expected: false,
	}, {
		word:     "hellodarknessmyoldfriend2012",
		expected: true,
	}, {
		word:     "2018",
		expected: true,
	}}

	for i, c := range cases {
		got := containsNumbers(c.word)
		if got != c.expected {
			t.Error("Expected", c.expected, "got", got, "in case", i)
		}
	}
}
