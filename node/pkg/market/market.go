package market

import "github.com/odysseyhack/socialtec/node/pkg/store"

// Market defines the intrface to be implemented by market implementations
type Market interface {
	AddOffer(offer Offer) error
	GetAvailableOffers() ([]Offer, error)
	DeleteOffer(offerID int64) error
	ShowIntrest(offerID int64) error
	RemoveIntrest(offerID int64) error
	SendFormChain(ifOwner int64, ifGood int64, thenGood int64, thenReceiver int64) error
	ReferInterest(offerID int64, referID int64) error
	ListenBroadcasts(store store.Store)
}

// Offer defines the structure of an Offer
type Offer struct {
	Node     int64  `json:"node"`
	Name     string `json:"name"`
	Details  string `json:"details"`
	ImageURL string `json:"image_url"`
	ID       int64  `json:"id"`
}

// DefaultMarket is the singleton to be used
var DefaultMarket Market
