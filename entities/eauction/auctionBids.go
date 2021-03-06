package eauction

import (
	"errors"
	"sync"

	"github.com/emirpasic/gods/trees/binaryheap"
)

type AuctionBids struct {
	mu sync.Mutex

	auctionId int
	bids      *binaryheap.Heap
}

func bidComp(b1, b2 interface{}) int {
	bid1, ok1 := b1.(*Bid)
	bid2, ok2 := b2.(*Bid)
	if !ok1 && ok2 {
		return 0
	}
	if bid1.Price() < bid2.Price() {
		return 1
	}
	return -1
}

func NewAuctionBids(auctionId int) *AuctionBids {
	return &AuctionBids{
		auctionId: auctionId,
		bids:      binaryheap.NewWith(bidComp),
	}
}

func (aucbids *AuctionBids) GetHighestBid() *Bid {
	if top, ok := aucbids.bids.Peek(); !ok {
		return nil
	} else if bid, ok := top.(*Bid); ok {
		return bid
	} else {
		return nil
	}
}

func (aucbids *AuctionBids) ValidateBid(bid *Bid) error {
	topbid := aucbids.GetHighestBid()
	if topbid != nil {
		if bid.Price() <= topbid.Price() {
			return errors.New("bid not high enough")
		}
		if bid.UserId() == topbid.UserId() {
			return errors.New("user is already highest bidder")
		}
	}
	return nil
}

func (aucbids *AuctionBids) AddBid(bid *Bid) error {
	aucbids.mu.Lock()
	defer aucbids.mu.Unlock()
	if err := aucbids.ValidateBid(bid); err != nil {
		return err
	}
	aucbids.bids.Push(bid)
	return nil
}

func (aucbids *AuctionBids) GetList() []*Bid {
	var bids []*Bid
	var tmpbids []interface{}
	for !aucbids.bids.Empty() {
		b, _ := aucbids.bids.Pop()
		tmpbids = append(tmpbids, b)
		bid, _ := b.(*Bid)
		bids = append(bids, bid)
	}
	aucbids.bids.Push(tmpbids...)
	return bids
}
