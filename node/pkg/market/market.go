package market

// Market defines the intrface to be implemented by market implementations
type Market interface {
	AddOffer(offer Offer) error
	GetAvailableOffers(pageNumber int) ([]Offer, error)
	DeleteOffer(offerID int64) error
	ShowIntrest(offerID int64) error
	RemoveIntrest(offerID int64) error
}

// Offer defines the structure of an Offer
type Offer struct {
	Node     uint64 `json:"node"`
	Name     string `json:"name"`
	Details  string `json:"details"`
	ImageURL string `json:"image_url"`
	ID       int64  `json:"id"`
}

// DefaultMarket is the singleton to be used
var DefaultMarket Market
