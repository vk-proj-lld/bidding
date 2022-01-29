package entities

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

type Auction struct {
	article *Article

	aucitonID,
	sellerID int

	minprice,
	maxprice float64

	creationTime,
	finalizationTime time.Time
	closed bool
}

var auctioncounter uint32 = 1

func NewAuction(sellerID int, article *Article, creationTime, finalizationTime time.Time, minprice, maxprice float64) *Auction {
	return &Auction{
		aucitonID:        int(atomic.AddUint32(&auctioncounter, 1)),
		sellerID:         sellerID,
		article:          article,
		finalizationTime: finalizationTime,
		minprice:         minprice,
		maxprice:         maxprice,
		closed:           false,
		creationTime:     creationTime,
	}
}

func (auc *Auction) Id() int { return auc.aucitonID }

func (auc *Auction) String() string {
	return fmt.Sprintf("AuctionId (%d): \b%v", auc.aucitonID, auc.article)
}

func (auc *Auction) ValidateBid(bid *Bid) error {
	if bid.price < auc.minprice || bid.price > auc.maxprice {
		return errors.New("price not in range")
	}
	if auc.closed || bid.btime.After(auc.finalizationTime) {
		return errors.New("bids is closed")
	}
	return nil
}

func (auc *Auction) FinalizationTime() time.Time { return auc.finalizationTime }
func (auc *Auction) CreationTime() time.Time     { return auc.creationTime }

func (auc *Auction) Close() {
	auc.closed = true
}

func (auc *Auction) CalculateProfit(bid *Bid) interface{} {
	// changeable for later
	return fmt.Sprintf("Overall Profit made on auction with the Bid is: %.2f", bid.price)
}
