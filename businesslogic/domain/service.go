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
func (p *port) CalculateDeceptiveScore(domain *Domain) (int, error) {
	keywords, err := p.repo.FindAllKeywords()

	if err != nil {
		return 0, err
	}

	for _, k := range keywords {
		if k == domain.Name {
			return 100, nil
		}
	}
	return 0, nil
}
