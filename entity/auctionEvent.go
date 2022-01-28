package entity

import (
	"errors"
	"time"

	"github.com/vkstack/bidding/interfaces/notifiers"
)

type auctionsEvent struct {
	output   notifiers.IOutPut
	auctions map[int]*Auction
}

func CreteAuctionEvent(channel notifiers.IOutPut) *auctionsEvent {
	return &auctionsEvent{output: channel, auctions: make(map[int]*Auction)}
}

func (event *auctionsEvent) ListAuctions() (auctionlist []*Auction) {
	for _, auc := range event.auctions {
		auctionlist = append(auctionlist, auc)
	}
	return auctionlist
}

func (event *auctionsEvent) GetAuction(auctionID int) (*Auction, error) {
	if auction, ok := event.auctions[auctionID]; ok {
		return auction, nil
	}
	return nil, errors.New("invalid auction id")
}

func (event *auctionsEvent) CreateAuction(name, description string, attributes interface{}, startPrice, endPrice float64, finaltime time.Time) *Auction {
	auction := CreateAuction(name, description, attributes, startPrice, endPrice, finaltime)
	event.auctions[auction.AuctionID] = auction
	auction.output = event.output
	return auction
}
