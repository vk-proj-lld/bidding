package interfaces

import (
	"time"

	"github.com/vkstack/bidding/entity"
)

type EventCreater interface {
	CreateAuction(name, description string, attributes interface{}, startPrice, endPrice float64, finaltime time.Time) *entity.Auction
}

type IAuctionUsecase interface {
	EventCreater
	Bid(user *entity.User, auctionID int, price float64) bool
}

type IAuctionEvent interface {
	EventCreater
	GetAuction(auctionID int) (*entity.Auction, error)
	ListAuctions() []*entity.Auction
}
