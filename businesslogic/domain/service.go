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
	score := 0
	for _, keyword := range keywords {
		currentScore := calculateScore(keyword, domain.Name)

		// These checks allows the algorithm to look at all keywords
		// until it finds the one that has the highest score
		if currentScore == 100 {
			return 100, nil
		} else if currentScore > score {
			score = currentScore
		}
	}

	return score, nil
}

func calculateScore(keyword Keyword, domainName string) int {
	if contains(domainName, keyword.Word) {
		if keyword.KeywordType == "brand" {
			return 100
		}
		return 90
	} else if calculateLevenshtein(domainName, keyword.Word) {
		if keyword.KeywordType == "brand" {
			return 80
		}
		return 70
	}

	return 0
}