package istrategy

import (
	"github.com/vkstack/bidding/entities/eauction"
)

type IStrategy interface {
	GetNextBid(maxbid *eauction.Bid, auction *eauction.Auction) float64
}
