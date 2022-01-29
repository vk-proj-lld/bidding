package strategy

import (
	"math/rand"

	"github.com/vkstack/bidding/entities/eauction"
	"github.com/vkstack/bidding/interfaces/istrategy"
)

type ranstrategy struct {
	rangen *rand.Rand
}

func NewRandomStrategy(rangen *rand.Rand) istrategy.IStrategy {
	return &ranstrategy{
		rangen,
	}
}

func (ranst *ranstrategy) GetNextBid(maxbid *eauction.Bid, auction *eauction.Auction) float64 {
	return maxbid.Price() + float64(10*ranst.rangen.Intn(int(50)))
}
