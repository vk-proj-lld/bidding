package entities

import (
	"fmt"
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

var bidcounter uint32

func NewBid(auctionId, userId int, price float64, btime time.Time) *Bid {
	return &Bid{
		aucitonId: auctionId,
		userId:    userId,
		bidId:     int(atomic.AddUint32(&bidcounter, 1)),

		btime: btime,
		price: price,
	}
}

func (b *Bid) String() string {
	return fmt.Sprintf("Bid of price:%.2f placed by %d in auction %d", b.price, b.userId, b.aucitonId)
}

func (b *Bid) Price() float64 { return b.price }

func (b *Bid) UserId() int { return b.userId }
