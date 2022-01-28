package entity

import (
	"time"
)

type Bid struct {
	user         *User
	aucitonID    int
	bidID        int
	biddingPrice float64
	time         time.Time
}

func CreateBid(user *User, auctionID int, price float64) *Bid {
	return &Bid{user: user, aucitonID: auctionID, time: time.Now(), biddingPrice: price, bidID: getID()}
}
