package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/vkstack/bidding/business/repo"
	"github.com/vkstack/bidding/business/uc"
	"github.com/vkstack/bidding/entities"
	"github.com/vkstack/bidding/entities/earticle"
	"github.com/vkstack/bidding/entities/strategy"
)

var rangen = rand.New(rand.NewSource(time.Now().UnixMicro()))

// func
func main() {
	var wg sync.WaitGroup
	var out = uc.NewConsoleOutPutUsecase()
	auctionusecase := uc.NewAuctionUseCase(
		&wg,
		//dependency injection
		repo.NewAuctionRepo(),
		out,
	)
	// uc := business.GetAuctionUseCase()
	// auction := uc.CreateAuction("article-1", "antique article", []string{}, 1000.0, 1000000, time.Now().Add(5*time.Second))
	seller, err := auctionusecase.CreateUser("Himanshu", strategy.NewRandomStrategy(rangen))
	out.Write("creating user:", seller, err)
	auction, err := auctionusecase.CreateAuction(
		seller.Id(),
		*earticle.NewArticle("item-1", "precious vintage item"),
		1000.0,
		10000.0,
		time.Now().Local().Add(time.Second*10).Format("2006-01-02 15:04:05"),
	)
	out.Write("creating auction:", auction, err)

	u1, _ := auctionusecase.CreateUser("Divyesh", strategy.NewRandomStrategy(rangen))
	out.Write("creating user:", u1, err)
	u2, _ := auctionusecase.CreateUser("Vajahat", strategy.NewRandomStrategy(rangen))
	out.Write("creating user:", u2, err)
	u3, _ := auctionusecase.CreateUser("Priyanka", strategy.NewRandomStrategy(rangen))
	out.Write("creating user:", u3, err)
	u4, _ := auctionusecase.CreateUser("Muskan", strategy.NewRandomStrategy(rangen))
	out.Write("creating user:", u4, err)
	bidders := []*entities.User{u1, u2, u3, u4}
	go func() {
		slots := len(bidders)
		start := 500
		for !auction.IsClosed() {
			start = start + 10*rangen.Intn(50)
			u := bidders[rangen.Intn(slots)]
			b, err := auctionusecase.PlaceBid(u.Id(), auction.Id(), float64(start))
			out.Write("placing bid:", b, u, err)
			time.Sleep(time.Millisecond * 50 * time.Duration(rangen.Intn(10)))
		}
	}()
	wg.Wait()
	fmt.Println("All Auctions are closed.")
}
