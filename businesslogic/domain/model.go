package domain

// Domain struct represents the model of a domain
type Domain struct {
	Name string `json:"name"`
}

// Keyword is used to save the keyword itself and it's type. The type indicates
// whether the keyword is a brand name or a common word
type Keyword struct {
	Word string
	KeywordType string
}
