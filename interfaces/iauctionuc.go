package interfaces

import (
	"github.com/vkstack/bidding/entities"
	"github.com/vkstack/bidding/entities/earticle"
	"github.com/vkstack/bidding/entities/eauction"
	"github.com/vkstack/bidding/interfaces/istrategy"
)

type IAucUC interface {
	CreateUser(name string, strategy istrategy.IStrategy) (*entities.User, error)
	CreateAuction(userId int, article earticle.Article, minprice, maxprice float64, finalization string) (*eauction.Auction, error)
	PlaceBid(userId, auctionId int, price float64) (*eauction.Bid, error)
}
