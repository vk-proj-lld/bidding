package repo

import (
	"github.com/vkstack/bidding/entities"
)

type auctionRepo struct {
	users       map[int]*entities.User
	auctions    map[int]*entities.Auction
	auctionBids map[int]*entities.AuctionBids
}

func NewAuctionRepo() *auctionRepo {
	return &auctionRepo{
		users: make(map[int]*entities.User),

		auctions:    make(map[int]*entities.Auction),
		auctionBids: make(map[int]*entities.AuctionBids),
	}
}

func (repo *auctionRepo) SaveUser(user *entities.User) error {
	repo.users[user.Id()] = user
	return nil
}

func (repo *auctionRepo) SaveAuction(auction *entities.Auction) error {
	repo.auctions[auction.Id()] = auction
	repo.auctionBids[auction.Id()] = entities.NewAuctionBids(auction.Id())
	return nil
}

func (repo *auctionRepo) GetAuction(auctionId int) *entities.Auction {
	return repo.auctions[auctionId]
}

func (repo *auctionRepo) GetAuctionBids(auctionId int) *entities.AuctionBids {
	return repo.auctionBids[auctionId]
}

func (repo *auctionRepo) GetUser(userId int) *entities.User {
	return repo.users[userId]
}
