package business

import (
	"time"

	"github.com/vkstack/bidding/entity"
	"github.com/vkstack/bidding/interfaces"
)

type AuctionUseCase struct {
	event interfaces.IAuctionEvent
}

var aucUC AuctionUseCase

func GetAuctionUseCase() interfaces.IAuctionUsecase {
	aucUC.event = entity.CreteAuctionEvent(notifier)
	return &aucUC
}

func (aucUC *AuctionUseCase) Bid(user *entity.User, auctionID int, price float64) bool {
	if auction, err := aucUC.event.GetAuction(auctionID); err == nil {
		bid := entity.CreateBid(user, auctionID, price)
		auction.AddBid(bid)
		return true
	} else {
		//some error is to be logged
		return false
	}
}

func (aucUC *AuctionUseCase) CreateAuction(name, description string, attributes interface{}, startPrice, endPrice float64, finaltime time.Time) *entity.Auction {
	return aucUC.event.CreateAuction(name, description, attributes, startPrice, endPrice, finaltime)
}

func (aucUC *AuctionUseCase) ListAuctions() []*entity.Auction {
	return aucUC.event.ListAuctions()
}
