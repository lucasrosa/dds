package domain

import "testing"

func TestTolaratedDistance(t *testing.T) {
	cases:= []struct{
		needleLength int
		expected int
	}{{
		needleLength: 0,
		expected: 0,
	},{
		needleLength: 3,
		expected: 0,
	},{
		needleLength: 4,
		expected: 1,
	},{
		needleLength: 8,
		expected: 1,
	},{
		needleLength: 9,
		expected: 2,
	},{
		needleLength: 16,
		expected: 2,
	},{
		needleLength: 17,
		expected: 3,
	}}

	for i, c := range cases {
		got := tolaratedDistance(c.needleLength)
		if got != c.expected {
			t.Error("Expected", c.expected, "got", got, "in case", i)
		}
	}
}

func TestCalculateLevenshteinWhenNeedleIsBiggerThanHaystack(t *testing.T) {
	haystack := "visa"
	needle := "visao"
	expected := false

	got := calculateLevenshtein(haystack, needle)
	if got != expected {
		t.Error("Expected", expected, "got",got, "in test case")
	}
}

func TestCalculateLevenshtein(t *testing.T) {
	cases := []struct{
		needle string
		haystack string
		expected bool
	}{{
		needle: "bb",
		haystack: "bb",
		expected: true,
	},{
		needle: "tim",
		haystack: "tim",
		expected: true,
	},{
		needle: "visa",
		haystack: "visa",
		expected: true,
	},{
		needle: "bb",
		haystack: "atendimentobb",
		expected: true,
	},
	{
		needle: "bb",
		haystack: "atendimentobx",
		expected: false,
	},
	{
		needle: "visa",
		haystack: "melhorvizacredito",
		expected: true,
	},
	{
		needle: "tim",
		haystack: "melhortimcredito",
		expected: true,
	},
	{
		needle: "tim",
		haystack: "melhorbimcredito",
		expected: false,
	},
	{
		needle: "coca-cola",
		haystack: "toca-tola",
		expected: true,
	}}

	for i, c := range cases {
		got := calculateLevenshtein(c.haystack, c.needle)
		if got != c.expected {
			t.Error("Expected", c.expected, "got",got, "in test case",i)
		}
	}
}