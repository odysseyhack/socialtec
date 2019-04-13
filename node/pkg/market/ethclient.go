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
	keyString    string
	contractAddr common.Address
	session      *GivoSession
	account      accounts.Account
	conn         *ethclient.Client
	netURL       string
}

// NewMarket Creates a market instance
func NewMarket(netURL string, keyString string) Market {
	conn, err := ethclient.Dial(netURL)
	if err != nil {
		log.Fatalf("Something went wrong connecting err: %v", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Fatalf("Something went wrong creating ECDSA key: %v", err)
	}
	tc := bind.NewKeyedTransactor(key)
	addr, _, givo, err := DeployGivo(tc, conn)
	if err != nil {
		log.Fatalf("Something went wrong deploying with error err: %v", err)
	}

	session := GivoSession{
		Contract:     givo,
		TransactOpts: *tc,
		CallOpts:     bind.CallOpts{Context: context.Background()},
	}
	return &client{
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

//AddOffer adds an offer
func (c *client) AddOffer(offer Offer) error {
	fmt.Println("givo: ", c.contractAddr.String())
	tr, err := c.session.CreateOffer(offer.Name, offer.ImageURL, offer.Details)
	fmt.Println("Created offer", tr.Value(), err)
	return err
}

func (c *client) GetAvailableOffers(pageNumber int) ([]Offer, error) {
	var offers []Offer
	for i := 0; i < 1; i++ {
		offer, err := c.session.Offers(big.NewInt(0), big.NewInt(0))
		if err != nil {
			return nil, err
		}
		offers = append(offers, Offer{
			Node:     offer.Node.Hex(),
			Name:     offer.Name,
			Details:  offer.IpfsDetails,
			ImageURL: offer.IpfsImage,
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
	_, err := c.session.AddIntrest(big.NewInt(0), big.NewInt(offerID))
	if err != nil {
		return err
	}
	return nil
}
func (c *client) RemoveIntrest(offerID int64) error {
	_, err := c.session.DeleteIntrest(big.NewInt(0), big.NewInt(offerID))
	if err != nil {
		return err
	}
	return nil
}
