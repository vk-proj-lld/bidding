package repo

import (
	"github.com/vkstack/bidding/entities"
	"github.com/vkstack/bidding/entities/eauction"
	"github.com/vkstack/bidding/interfaces"
)

type auctionRepo struct {
	users       map[int]*entities.User
	auctions    map[int]*eauction.Auction
	auctionBids map[int]*eauction.AuctionBids
}

func NewAuctionRepo() interfaces.IAucRepo {
	return &auctionRepo{
		users: make(map[int]*entities.User),

		auctions:    make(map[int]*eauction.Auction),
		auctionBids: make(map[int]*eauction.AuctionBids),
	}
}

func (repo *auctionRepo) SaveUser(user *entities.User) error {
	repo.users[user.Id()] = user
	return nil
}

func (repo *auctionRepo) SaveAuction(auction *eauction.Auction) error {
	repo.auctions[auction.Id()] = auction
	repo.auctionBids[auction.Id()] = eauction.NewAuctionBids(auction.Id())
	return nil
}

func (repo *auctionRepo) GetAuction(auctionId int) *eauction.Auction {
	return repo.auctions[auctionId]
}

func (repo *auctionRepo) GetAuctionBids(auctionId int) *eauction.AuctionBids {
	return repo.auctionBids[auctionId]
}

func (repo *auctionRepo) GetUser(userId int) *entities.User {
	return repo.users[userId]
}
