package market

// Market defines the intrface to be implemented by market implementations
type Market interface {
	ShowIntrest(offerID string) error
	GetAvailableOffers(pageNumber int) ([]Offer, error)
	RemoveIntrest(offerID string) error
	AddOffers([]Offer) error
}

// Offer defines the structure of an Offer
type Offer struct {
	Name     string `json:"name"`
	Details  string `json:"details"`
	ImageURL string `json:"image_url"`
	ID       string `json:"id"`
}
