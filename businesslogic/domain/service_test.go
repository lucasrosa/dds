package domain

import (
	"testing"
)

type domainRepository struct{}

// newMemoryDomainRepository instantiates a secondary port for the tests
func newMemoryDomainRepository() SecondaryPort {
	return &domainRepository{}
}

func (r *domainRepository) FindAllKeywords() ([]Keyword, error) {
	keywords := []Keyword{
		{
			Word:        "itau",
			KeywordType: "brand",
		}, {
			Word:        "unibanco",
			KeywordType: "brand",
		}, {
			Word:        "bb",
			KeywordType: "brand",
		}, {
			Word:        "visa",
			KeywordType: "brand",
		}, {
			Word:        "bb",
			KeywordType: "brand",
		}, {
			Word:        "banco",
			KeywordType: "word",
		}, {
			Word:        "bancodobrasil",
			KeywordType: "brand",
		}, {
			Word:        "safra",
			KeywordType: "Word",
		}, {
			Word:        "bitcoin",
			KeywordType: "word",
		}, {
			Word:        "caixa",
			KeywordType: "word",
		}, {
			Word:        "santander",
			KeywordType: "brand",
		}, {
			Word:        "nubank",
			KeywordType: "brand",
		}}

	return keywords, nil
}

func TestCalculateDeceptiveScore(t *testing.T) {
	domainRepo := newMemoryDomainRepository()
	domainService := NewDomainService(domainRepo)

	cases := []struct {
		domain        Domain
		expectedScore int
	}{
		{domain: Domain{"itau"}, expectedScore: 100},
		{domain: Domain{"itamu"}, expectedScore: 80},
		{domain: Domain{"atendimentobb"}, expectedScore: 100},
		{domain: Domain{"hello2019"}, expectedScore: 50},
		{domain: Domain{"popap0la"}, expectedScore: 50},
	}

	for i, c := range cases {
		got, _ := domainService.CalculateDeceptiveScore(&c.domain)

		if got != c.expectedScore {
			t.Error("Expected", c.expectedScore, "got", got, "in case", i)
		}
	}
}


func BenchmarkCalculateScore(b *testing.B) {

	keyword := Keyword{Word: "cielo", KeywordType: "brand"}
	domainName := "atendimentocielo"
	for n:= 0; n < b.N; n++ {
		calculateScore(keyword, domainName)
	}
}