package entities

import (
	"sync/atomic"
	"time"
)

type Bid struct {
	aucitonId,
	userId,
	bidId int

	price float64
	btime time.Time
}

var bidcounter uint32 = 1

func NewBid(auctionId, userId int, price float64, btime time.Time) *Bid {
	return &Bid{
		aucitonId: auctionId,
		userId:    userId,
		bidId:     int(atomic.AddUint32(&bidcounter, 1)),

		btime: btime,
		price: price,
	}
}

func (b *Bid) Price() float64 { return b.price }

func (b *Bid) UserId() int { return b.userId }
