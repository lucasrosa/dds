package domain

// PrimaryPort represents the actions that ‘drives’ the application logic.
// Conversations with the application domain are initiated by adapters,
// often as a direct consequence of some external stimulus such as an external
// system initiating a conversation.
type PrimaryPort interface {
	CalculateDeceptiveScore(domain *Domain) (int, error)
}

// SecondaryPort represents the actions that will be taken
// by the application domain to ‘drive’ adapters.
// In this case, this is always as a result of the application
// domain initiating the interaction, generally in order to
// communicate with an external system or device
type SecondaryPort interface {
	FindAllKeywords() ([]string, error)
}
