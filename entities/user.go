package entities

import (
	"fmt"
	"sync/atomic"

	"github.com/vkstack/bidding/interfaces/istrategy"
)

var usercounter uint32

type User struct {
	id              int
	name            string
	biddingStrategy istrategy.IStrategy
}

func NewUser(name string, strategy istrategy.IStrategy) *User {
	return &User{
		name:            name,
		id:              int(atomic.AddUint32(&usercounter, 1)),
		biddingStrategy: strategy,
	}
}

func (u *User) Id() int { return u.id }

func (u *User) Name() string { return u.name }

func (u *User) String() string { return fmt.Sprintf("User (%d): %s", u.id, u.name) }
