package domainadaptermemory

import (
	"github.com/lucasrosa/edds/businesslogic/domain"
)

type domainRepository struct{}

func NewMemoryDomainRepository() domain.SecondaryPort {
	return &domainRepository{}
}

func (r *domainRepository) FindAllKeywords() ([]domain.Keyword, error) {
	keywords := []domain.Keyword{
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
