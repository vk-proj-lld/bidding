package main

import (
	"github.com/vkstack/bidding/business/repo"
	"github.com/vkstack/bidding/business/uc"
)

// func
func main() {
	auctionusecase := uc.NewAuctionUseCase(
		//dependency injection
		repo.NewAuctionRepo(),
		uc.NewConsoleOutPutUsecase(),
	)
	// uc := business.GetAuctionUseCase()
	// auction := uc.CreateAuction("article-1", "antique article", []string{}, 1000.0, 1000000, time.Now().Add(5*time.Second))

	// u1 := entity.User{UserID: 1}
	// u2 := entity.User{UserID: 2}
	// u3 := entity.User{UserID: 3}
	// auction.AddBid(entity.CreateBid(&u2, auction.AuctionID, 2000))
	// auction.AddBid(entity.CreateBid(&u1, auction.AuctionID, 1000))
	// auction.AddBid(entity.CreateBid(&u3, auction.AuctionID, 3000))
	// auction.AddBid(entity.CreateBid(&u1, auction.AuctionID, 4000))
	// auction.AddBid(entity.CreateBid(&u1, auction.AuctionID, 5000))
	// auction.AddBid(entity.CreateBid(&u2, auction.AuctionID, 6000))
	// time.Sleep(time.Minute * 5)
	// auction.AddBid(entity.CreateBid(&u3, auction.AuctionID, 7000))
	// var wg sync.WaitGroup
	// wg.Add(1)
}
