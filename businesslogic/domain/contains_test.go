package domain

import "testing"

func TestContains(t *testing.T) {
	cases := []struct {
		needle    string
		haystack  string
		expected bool
	}{{
		needle:    "bahia",
		haystack: "casasbahia",
		expected:       true,
	},
	{
		needle:    "bahía",
		haystack: "casasbahia",
		expected:       true,
	},
	{
		needle:    "magazine",
		haystack: "magazineluiza",
		expected:       true,
	},
	{
		needle:    "magazineluiza",
		haystack: "magazine",
		expected:       false,
	}}

	for i, c := range cases {
		got := contains(c.haystack, c.needle)

		if c.expected != got {
			t.Error("Expected", c.expected, "got", got, "in case", i)
		}
	}
}