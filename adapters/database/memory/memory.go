package domainadaptermemory

import (
	"github.com/lucasrosa/edds/businesslogic/domain"
)

type domainRepository struct{}

func NewMemoryDomainRepository() domain.SecondaryPort {
	return &domainRepository{}
}

func (r *domainRepository) FindAllKeywords() ([]string, error) {
	keywords := []string{
		"itau",
		"unibanco",
	}

	return keywords, nil
}
