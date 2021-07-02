package model

import "time"

type Invite struct {
	UserID uint
	GroupID uint

	User User
	Group Group

	CreatedAt time.Time
}
