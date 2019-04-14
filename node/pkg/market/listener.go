package market

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/odysseyhack/socialtec/node/pkg/store"
)

func (c *client) ListenBroadcasts(store store.Store) {
	referer := make(chan *GivoRefer, 100)
	chain := make(chan *GivoChained, 100)
	interested := make(chan *GivoIntrested, 100)
	notInterested := make(chan *GivoNotIntrested, 100)

	watchopts := &bind.WatchOpts{}

	c.session.Contract.WatchRefer(watchopts, referer)
	c.session.Contract.WatchChained(watchopts, chain)
	c.session.Contract.WatchIntrested(watchopts, interested)
	c.session.Contract.WatchNotIntrested(watchopts, notInterested)

	log.Println("Watching for all the nodes")

	for {
		select {
		case event := <-referer:
			c.handleReference(event)
		case event := <-chain:
			c.handleChained(event)
		case gInterest := <-interested:
			c.handleInterested(gInterest)
		case gNoInterest := <-notInterested:
			fmt.Println("no-interest", gNoInterest)
		case <-time.After(time.Second):
		}
	}
}

// Interest stores interest
type Interest struct {
	Who  int64
	What int64
}

func (c *client) checkIfILike(from int64) (int64, bool) {
	var offers []Offer
	if err := store.Default.Get("my_interests", &offers); err != nil {
		return 0, false
	}
	for _, offer := range offers {
		if offer.Node == from {
			return offer.ID, true
		}
	}
	return 0, false
}

func (c *client) handleInterested(event *GivoIntrested) {
	if c.ID() != event.To.Int64() {
		return
	}
	fmt.Printf("\ninterest from:%d to:%d for:%d nodeID: %d\n",
		event.From.Int64(),
		event.To.Int64(),
		event.GoodId.Int64(),
		c.ID(),
	)

	if offerID, ok := c.checkIfILike(event.From.Int64()); ok {
		if err := c.SendFormChain(
			event.From.Int64(),
			offerID,
			event.GoodId.Int64(),
			event.From.Int64(),
		); err != nil {
			log.Printf("failed sending chain for request. Err:%+v", err)
		}
		return
	}

	//  add to refs map
	refs := make(map[int64]Interest)
	if err := store.Default.Get("refs", &refs); err != nil && err != store.ErrorNotFound {
		log.Printf("error getting refs")
		return
	}
	refs[event.From.Int64()] = Interest{Who: event.From.Int64(), What: event.GoodId.Int64()}
	if err := store.Default.Set("refs", &refs); err != nil && err != store.ErrorNotFound {
		log.Printf("error setting refs")
		return
	}

	intrestedParties := make(map[int64]Interest)
	if err := store.Default.Get("intrestedParties", &intrestedParties); err != nil && err != store.ErrorNotFound {
		log.Printf("error getting interested parties", err)
		return
	}
	intrestedParties[event.From.Int64()] = Interest{Who: event.From.Int64(), What: event.GoodId.Int64()}
	if err := store.Default.Set("intrestedParties", &intrestedParties); err != nil && err != store.ErrorNotFound {
		log.Printf("error setting refs")
		return
	}

}

func (c *client) handleReference(event *GivoRefer) {
	if c.ID() != event.To.Int64() {
		return
	}

	fmt.Printf("referer from:%d to:%d for:%d reference:%d",
		event.From.Int64(),
		event.To.Int64(),
		event.GoodId.Int64(),
		event.ReferId.Int64(),
	)

	if offerID, ok := c.checkIfILike(event.ReferId.Int64()); ok {
		if err := c.SendFormChain(
			event.ReferId.Int64(),
			offerID,
			event.GoodId.Int64(),
			event.From.Int64(),
		); err != nil {
			log.Printf("failed sending chain for request. Err:%+v", err)
		}
		return
	}

	//  add to refs map
	refs := make(map[int64]Interest)
	if err := store.Default.Get("refs", &refs); err != nil && err != store.ErrorNotFound {
		log.Printf("error getting refs")
		return
	}
	refs[event.ReferId.Int64()] = Interest{Who: event.From.Int64(), What: event.GoodId.Int64()}
	if err := store.Default.Set("refs", &refs); err != nil && err != store.ErrorNotFound {
		log.Printf("error setting refs")
		return
	}

	intrestedParties := make(map[int64]Interest)
	if err := store.Default.Get("intrestedParties", &intrestedParties); err != nil && err != store.ErrorNotFound {
		log.Printf("error getting refs")
		return
	}
	intrestedParties[event.ReferId.Int64()] = Interest{Who: event.From.Int64(), What: event.GoodId.Int64()}
	if err := store.Default.Set("intrestedParties", &intrestedParties); err != nil && err != store.ErrorNotFound {
		log.Printf("error setting refs")
		return
	}

}

func (c *client) handleChained(event *GivoChained) {
	if c.ID() != event.ThenReceiver.Int64() {
		return
	}
	fmt.Println("chained event",
		event.IfOwner.Int64(),
		event.IfGood.Int64(),
		event.IfReceiver.Int64(),
		event.ThenGood.Int64(),
		event.ThenReceiver.Int64(),
	)

	if c.ID() == event.IfOwner.Int64() {
		log.Println("chain propagation over")
		return
	}

	// propagate
	refs := make(map[int64]Interest)
	if err := store.Default.Get("refs", &refs); err != nil && err != store.ErrorNotFound {
		log.Printf("error getting refs for propagate on handle chained")
		return
	}
	referred, ok := refs[event.IfOwner.Int64()]
	if !ok {
		log.Println("Bad cycle owner not in referred map", referred, c.ID(), event.IfOwner)
		return
	}

	if err := c.Propagate(
		event.IfOwner.Int64(),
		event.IfGood.Int64(),
		event.IfReceiver.Int64(),
		referred.What,
		referred.Who,
	); err != nil {
		log.Printf("failed propagating chain for request. Err:%+v", err)
	}
}
