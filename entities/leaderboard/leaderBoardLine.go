package leaderboard

import (
	"fmt"

	"github.com/vkstack/bidding/entities"
	"github.com/vkstack/bidding/entities/eauction"
)

type LeaderboardLine struct {
	lineId int
	bid    *eauction.Bid
	user   *entities.User
}

func NewLine(lineId int, bid eauction.Bid, user entities.User) *LeaderboardLine {
	return &LeaderboardLine{
		lineId: lineId,
		bid:    &bid,
		user:   &user,
	}
}

func (lbline *LeaderboardLine) String() string {
	return fmt.Sprintf("(%d) Bid price: %.2f by (%d)%s", lbline.lineId, lbline.bid.Price(), lbline.bid.UserId(), lbline.user.Name())
}
