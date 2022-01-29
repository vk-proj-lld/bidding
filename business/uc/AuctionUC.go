package uc

import (
	"errors"
	"sync"
	"time"

	"github.com/vkstack/bidding/entities"
	"github.com/vkstack/bidding/entities/earticle"
	"github.com/vkstack/bidding/entities/eauction"
	"github.com/vkstack/bidding/entities/leaderboard"
	"github.com/vkstack/bidding/interfaces"
	"github.com/vkstack/bidding/interfaces/istrategy"
	"github.com/vkstack/bidding/interfaces/notifiers"
)

type auctionUseCase struct {
	repo interfaces.IAucRepo
	wg   *sync.WaitGroup

	aucchan       chan *eauction.Auction
	outputchannel notifiers.IOut
}

func NewAuctionUseCase(wg *sync.WaitGroup, repo interfaces.IAucRepo, out notifiers.IOut) interfaces.IAucUC {
	aucuc := auctionUseCase{
		wg:            wg,
		repo:          repo,
		aucchan:       make(chan *eauction.Auction),
		outputchannel: out,
	}
	go aucuc.initAuctionsListner()
	return &aucuc
}

func (aucuc *auctionUseCase) CreateUser(name string, strategy istrategy.IStrategy) (*entities.User, error) {
	user := entities.NewUser(name, strategy)
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
func (aucuc *auctionUseCase) CreateAuction(userId int, article earticle.Article, minprice, maxprice float64, finalization string) (*eauction.Auction, error) {
	finalizationTime, err := time.ParseInLocation("2006-01-02 15:04:05", finalization, time.Local)
	if err != nil {
		return nil, err
	}
	currettime := time.Now()
	if currettime.After(finalizationTime) {
		return nil, errors.New("could not create past date auction")
	}
	auction := eauction.NewAuction(userId, &article, currettime, finalizationTime, minprice, maxprice)
	if err := aucuc.repo.SaveAuction(auction); err != nil {
		return nil, err
	}
	aucuc.wg.Add(1)
	go aucuc.runInBackground(auction)
	return auction, nil
}

/*
	error is returned in following 4 cases
		case 1. Auction is closed
		case 2. priceNot in range
		case 3. current leader has higher bidding
		case 4: user is already highest bidder
*/
func (aucuc *auctionUseCase) PlaceBid(userId, auctionId int, price float64) (*eauction.Bid, error) {
	auction := aucuc.repo.GetAuction(auctionId)
	auctionbids := aucuc.repo.GetAuctionBids(auctionId)

	if auction == nil || auctionbids == nil {
		return nil, errors.New("invalid auction, non-existing")
	}
	bid := eauction.NewBid(auctionId, userId, price, time.Now())

	if err := auction.ValidateBid(bid); err != nil {
		return nil, err
	}
	if err := auctionbids.AddBid(bid); err != nil {
		return nil, err
	}
	return bid, nil
}

func (aucuc *auctionUseCase) GetWinner(auctionId int) (*eauction.Auction, *entities.User, error) {
	return nil, nil, nil
}

func (aucuc *auctionUseCase) GetLeaderBoard(auctionId int) []*eauction.Auction {
	return nil
}

func (aucuc *auctionUseCase) runInBackground(auction *eauction.Auction) {
	d := auction.FinalizationTime().Sub(auction.CreationTime())
	ticker := time.NewTicker(d)
	<-ticker.C
	aucuc.aucchan <- auction
}

func (aucuc *auctionUseCase) newLineList(bids []*eauction.Bid) (lines []*leaderboard.LeaderboardLine) {
	for idx, bid := range bids {
		user := aucuc.repo.GetUser(bid.UserId())
		lines = append(lines, leaderboard.NewLine(idx+1, *bid, *user))
	}
	return lines
}

func (aucuc *auctionUseCase) initAuctionsListner() {
	for auction := range aucuc.aucchan {
		auctionBids := aucuc.repo.GetAuctionBids(auction.Id())
		winningbid := auctionBids.GetHighestBid()
		bids := auctionBids.GetList()
		var outputs = []interface{}{
			auction,
			"is closed and finalized",
			auction.CalculateProfit(winningbid),
		}
		for _, line := range aucuc.newLineList(bids) {
			outputs = append(outputs, line)
		}
		aucuc.outputchannel.Write(outputs...)
		aucuc.wg.Done()
	}
}
