package interfaces

import (
	"github.com/vkstack/bidding/entities"
	"github.com/vkstack/bidding/entities/eauction"
)

type IAucRepo interface {
	SaveUser(user *entities.User) error
	SaveAuction(auction *eauction.Auction) error
	GetAuction(auctionId int) *eauction.Auction
	GetAuctionBids(auctionId int) *eauction.AuctionBids
	GetUser(userId int) *entities.User
}
