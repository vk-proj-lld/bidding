package uc

import (
	"errors"
	"time"

	"github.com/vkstack/bidding/entities"
	"github.com/vkstack/bidding/interfaces"
	"github.com/vkstack/bidding/interfaces/notifiers"
)

type auctionUseCase struct {
	repo interfaces.IAucRepo

	aucchan       chan *entities.Auction
	outputchannel notifiers.IOut
}

func NewAuctionUseCase(repo interfaces.IAucRepo, out notifiers.IOut) interfaces.IAucUC {
	aucuc := auctionUseCase{
		repo:    repo,
		aucchan: make(chan *entities.Auction),
	}
	go aucuc.InitAuctionsListner()
	return &aucuc
}

func (aucuc *auctionUseCase) CreateUser(name string) (*entities.User, error) {
	user := entities.NewUser(name)
	if err := aucuc.repo.SaveUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

/*
	error is returned in following 4 cases
		case 1. finalizatiotime is incorrect
		case 2. finalizatiotime is invalid
*/
func (aucuc *auctionUseCase) CreateAuction(userId int, article entities.Article, minprice, maxprice float64, finalization string) (*entities.Auction, error) {
	finalizationTime, err := time.Parse("2006-01-02 15:04:05", finalization)
	if err != nil {
		return nil, err
	}
	currettime := time.Now()
	if currettime.After(finalizationTime) {
		return nil, errors.New("could not create past date auction")
	}
	auction := entities.NewAuction(userId, &article, currettime, finalizationTime, minprice, maxprice)
	if err := aucuc.repo.SaveAuction(auction); err != nil {
		return nil, err
	}
	go aucuc.RunInBackground(auction)
	return auction, nil
}

/*
	error is returned in following 4 cases
		case 1. Auction is closed
		case 2. priceNot in range
		case 3. current leader has higher bidding
		case 4: user is already highest bidder
*/
func (aucuc *auctionUseCase) PlaceBid(userId, auctionId int, price float64) (bool, error) {
	auction := aucuc.repo.GetAuction(auctionId)
	auctionbids := aucuc.repo.GetAuctionBids(auctionId)

	if auction == nil || auctionbids == nil {
		return false, errors.New("invalid auction, non-existing")
	}
	bid := entities.NewBid(auctionId, userId, price, time.Now())

	if err := auction.ValidateBid(bid); err != nil {
		return false, err
	}
	if err := auctionbids.AddBid(bid); err != nil {
		return false, err
	}
	return true, nil
}

func (aucuc *auctionUseCase) GetWinner(auctionId int) (*entities.Bid, *entities.User, error) {
	return nil, nil, nil
}

func (aucuc *auctionUseCase) GetLeaderBoard(auctionId int) []*entities.Bid {
	return nil
}

func (aucuc *auctionUseCase) RunInBackground(auction *entities.Auction) {
	d := auction.FinalizationTime().Sub(auction.CreationTime())
	ticker := time.NewTicker(d)
	<-ticker.C
	aucuc.aucchan <- auction
}

func (aucuc *auctionUseCase) InitAuctionsListner() {
	for auction := range aucuc.aucchan {
		auctionBids := aucuc.repo.GetAuctionBids(auction.Id())
		winningbid := auctionBids.GetHighestBid()
		winner := aucuc.repo.GetUser(winningbid.UserId())
		aucuc.outputchannel.Write(
			"following auction is closed and finalized",
			auction,
			auction.CalculateProfit(winningbid),
			winner,
		)
	}
}
