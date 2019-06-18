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
	return 1, nil
}
