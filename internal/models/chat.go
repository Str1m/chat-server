package models

import "time"

type Chat struct {
	ID        int64
	Usernames []string
	CreatedAt time.Time
}
