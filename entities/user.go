package entities

import (
	"fmt"
	"sync/atomic"
)

var usercounter uint32 = 1

type User struct {
	id   int
	name string
}

func NewUser(name string) *User {
	return &User{
		name: name,
		id:   int(atomic.AddUint32(&usercounter, 1)),
	}
}

func (u *User) Id() int { return u.id }

func (u *User) Name() string { return u.name }

func (u *User) String() string { return fmt.Sprintf("User (%d): %s", u.id, u.name) }
