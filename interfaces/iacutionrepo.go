package interfaces

import "github.com/vkstack/bidding/entities"

type IAucRepo interface {
	SaveUser(user *entities.User) error
	SaveAuction(auction *entities.Auction) error
	GetAuction(auctionId int) *entities.Auction
	GetAuctionBids(auctionId int) *entities.AuctionBids
	GetUser(userId int) *entities.User
}
