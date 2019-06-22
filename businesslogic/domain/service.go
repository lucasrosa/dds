package domain

type port struct {
	repo SecondaryPort
}
// NewDomainService receives a Secondary Port of domain and insantiates a Primary Port
func NewDomainService(repo SecondaryPort) PrimaryPort {
	return &port{
		repo,
	}
}
// CalculateDeceptiveScore calculates the deceptiveness score of a domain name
// If it is a brand name the score is bigger than when it is a word from the dictionary   
func (p *port) CalculateDeceptiveScore(domain *Domain) (int, error) {
	keywords, err := p.repo.FindAllKeywords()

	if err != nil {
		return 0, err
	}

	for _, keyword := range keywords {
		if contains(domain.Name, keyword.Word) {
			if keyword.KeywordType == "brand" {
				return 100, nil
			}
			return 90, nil
		} else if calculateLevenshtein(domain.Name, keyword.Word) {
			if keyword.KeywordType == "brand" {
				return 80, nil
			}
			return 70, nil
		}
	}
	return 0, nil
}
