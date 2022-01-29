package interfaces

import "github.com/vkstack/bidding/entities"

type IAucUC interface {
	CreateUser(name string) (*entities.User, error)
	CreateAuction(userId int, article entities.Article, minprice, maxprice float64, finalization string) (*entities.Auction, error)
}
