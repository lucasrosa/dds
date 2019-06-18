package domain

import (
	"testing"
)

type domainRepository struct{}

func NewMemoryDomainRepository() SecondaryPort {
	return &domainRepository{}
}

func (r *domainRepository) FindAllKeywords() ([]string, error) {
	keywords := []string{
		"itau",
		"unibanco",
	}

	return keywords, nil
}

func TestCalculateDeceptiveScore(t *testing.T) {
	domainRepo := NewMemoryDomainRepository()

	domainService := NewDomainService(domainRepo)
	
	d := Domain{Name: "itau"}
	expected := 100
	got, _ := domainService.CalculateDeceptiveScore(&d)

	if got != expected {
		t.Error("Expected",expected,"got",got)
	}

}