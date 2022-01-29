package entities

import "fmt"

type LeaderboardLine struct {
	lineId int
	bid    *Bid
	user   *User
}

func NewLine(lineId int, bid Bid, user User) *LeaderboardLine {
	return &LeaderboardLine{
		lineId: lineId,
		bid:    &bid,
		user:   &user,
	}
}

func (lbline *LeaderboardLine) String() string {
	return fmt.Sprintf("(%d) Bid price: %.2f by (%d)%s", lbline.lineId, lbline.bid.Price(), lbline.bid.UserId(), lbline.user.Name())
}
