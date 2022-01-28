package entity

import (
	"fmt"
	"sync"
	"time"

	"github.com/vkstack/bidding/interfaces/notifiers"
)

var lastAI int
var mu sync.Mutex

func getID() int {
	mu.Lock()
	defer mu.Unlock()
	lastAI += 1
	return lastAI
}

type Auction struct {
	AuctionID        int
	article          *Article
	finalIzationTime time.Time
	startPrice       float64
	endPrice         float64
	sellerID         int //seller
	// isClosed         bool //

	winningBid *Bid
	bids       []*Bid

	output notifiers.IOutPut
}

func (auction *Auction) run() {
	ticker := time.NewTicker(auction.finalIzationTime.Sub(time.Now()))
	go func() {
		<-ticker.C // 9:40
		// <-ticker.C// 10:40
		// <-ticker.C// 11:40
		// <-ticker.C// 12:40
		auction.output.PrintFinalWinner(auction.WinnerMessage())
	}()
}

func CreateAuction(name, description string, attributes interface{}, startPrice, endPrice float64, finaltime time.Time) *Auction {
	auc := Auction{
		article:          getNewarticle(name, description, attributes),
		startPrice:       startPrice,
		endPrice:         endPrice,
		finalIzationTime: finaltime,
		AuctionID:        getID(),
	}
	auc.run() //8:40
	return &auc
}

//adds valid bid
func (auction *Auction) AddBid(bid *Bid) bool {
	if auction.sellerID == bid.user.UserID || auction.finalIzationTime.Before(bid.time) {
		// if auction.finalIzationTime.Before(bid.time) {
		//bid.time.After(auction.finalIzationTime)
		return false
	}
	if bid.biddingPrice < auction.startPrice || auction.endPrice < bid.biddingPrice {
		fmt.Println(bid.biddingPrice < auction.startPrice, auction.endPrice < bid.biddingPrice)
		return false
	}
	auction.bids = append(auction.bids, bid)
	auction.evalWinner(bid)
	return true
}

func (auction *Auction) evalWinner(bid *Bid) {
	//race condition
	if auction.winningBid == nil || bid.biddingPrice > auction.winningBid.biddingPrice {
		auction.winningBid = bid
	}
}

func (auction *Auction) GetWinnerStats() (*Bid, float64) {
	return auction.winningBid, auction.calculateProfit()
}

func (auction *Auction) calculateProfit() float64 {
	//variable logic
	return auction.winningBid.biddingPrice
}

func (auction *Auction) WinnerMessage() string {
	return fmt.Sprintf(
		`
		---------
		WinnerID: %d
		BidPrice: %f
		Profit:%f
		---------

		`,
		auction.winningBid.user.UserID,
		auction.winningBid.biddingPrice,
		auction.calculateProfit(),
	)
}
