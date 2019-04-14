package market

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type client struct {
	id           int64
	keyString    string
	contractAddr common.Address
	session      *GivoSession
	account      accounts.Account
	conn         *ethclient.Client
	netURL       string
}

// Deploy used for only deployments
func Deploy(tc *bind.TransactOpts, conn *ethclient.Client) string {
	addr, _, _, err := DeployGivo(tc, conn)
	if err != nil {
		log.Fatalf("Something went wrong deploying with error err: %v", err)
	}
	fmt.Println(addr.Hex())
	return addr.Hex()
}

// NewMarket Creates a market instance
func NewMarket(netURL string, keyString string, addrString string) Market {
	conn, err := ethclient.Dial(netURL)
	if err != nil {
		log.Fatalf("Something went wrong connecting err: %v", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Fatalf("Something went wrong creating ECDSA key: %v", err)
	}
	tc := bind.NewKeyedTransactor(key)

	addr := common.HexToAddress(addrString)
	givo, err := NewGivo(addr, conn)
	if err != nil {
		log.Fatalf("failed creating givo instance: %v", err)
	}

	session := GivoSession{
		Contract:     givo,
		TransactOpts: *tc,
		CallOpts:     bind.CallOpts{Context: context.Background()},
	}
	log.Println("Successfully connected to ethereum network")

	nodeID, _ := session.AddressToId(tc.From)

	return &client{
		id:           nodeID.Int64(),
		keyString:    keyString,
		contractAddr: addr,
		session:      &session,
		conn:         conn,
		netURL:       netURL,
	}
}

func (c *client) LoadAccout() {
	ks := keystore.NewKeyStore(
		"/tmp/eth/givo",
		keystore.LightScryptN,
		keystore.LightScryptP)
	account, err := ks.NewAccount("randompassphrase")
	fmt.Println(account, err)
	return
}

func (c *client) ID() int64 {
	if c.id != 0 {
		nodeID, _ := c.session.AddressToId(c.session.TransactOpts.From)
		c.id = nodeID.Int64()
	}
	return c.id
}

//AddOffer adds an offer
func (c *client) AddOffer(offer Offer) error {
	_, err := c.session.CreateOffer(offer.Name, offer.ImageURL, offer.Details)
	return err
}

func (c *client) GetAvailableOffers() ([]Offer, error) {
	var offers []Offer
	maxInt, _ := c.session.MaxAllGood()
	var max, lastIndex int64 = maxInt.Int64(), 0
	for i := max; i >= lastIndex; i-- {
		offer, err := c.session.AllGoods(big.NewInt(i))
		if err != nil {
			return nil, err
		}
		offers = append(offers, Offer{
			Node:     offer.Node.Int64(),
			Name:     offer.Name,
			Details:  offer.IpfsDetails,
			ImageURL: offer.IpfsImage,
			ID:       i,
		})
	}

	return offers, nil
}

func (c *client) DeleteOffer(offerID int64) error {
	_, err := c.session.DeleteOffer(big.NewInt(offerID))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) ShowIntrest(offerID int64) error {
	_, err := c.session.AddIntrest(big.NewInt(offerID))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) RemoveIntrest(offerID int64) error {
	_, err := c.session.DeleteIntrest(big.NewInt(offerID))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) ReferInterest(offerID int64, referID int64) error {
	_, err := c.session.ReferIntrest(big.NewInt(offerID), big.NewInt(referID))
	if err != nil {
		return err
	}
	return nil
}

// to, get_good_id,  give_owner, give_good_id
func (c *client) SendFormChain(ifOwner int64, ifGood int64, thenGood int64, thenReceiver int64) error {
	_, err := c.session.CycleFormed(
		big.NewInt(ifOwner),
		big.NewInt(ifGood),
		big.NewInt(thenGood),
		big.NewInt(thenReceiver),
	)
	if err != nil {
		return err
	}
	return nil
}

// to, get_good_id,  give_owner, give_good_id
func (c *client) Propagate(ifOwner int64, ifGood int64, ifReceiver int64, thenGood int64, thenReceiver int64) error {
	_, err := c.session.CyclePropagate(
		big.NewInt(ifOwner),
		big.NewInt(ifGood),
		big.NewInt(ifReceiver),
		big.NewInt(thenGood),
		big.NewInt(thenReceiver),
	)
	if err != nil {
		return err
	}
	return nil
}
