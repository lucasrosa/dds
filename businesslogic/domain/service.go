package domain

type port struct {
	repo SecondaryPort
}

func NewDomainService(repo SecondaryPort) PrimaryPort {
	return &port{
		repo,
	}
}

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
